package permission_handers

import (
	"context"
	"encoding/base64"
	"github.com/garyburd/redigo/redis"
	"github.com/micro/go-micro/metadata"
	"github.com/twinj/uuid"
	"github.com/vmihailenco/msgpack"
	"gopkg.in/mgo.v2"
	"konekko.me/gosion/application/pb"
	"konekko.me/gosion/commons/config"
	"konekko.me/gosion/commons/constants"
	"konekko.me/gosion/commons/dto"
	"konekko.me/gosion/commons/encrypt"
	"konekko.me/gosion/commons/errstate"
	"konekko.me/gosion/commons/generator"
	"konekko.me/gosion/commons/wrapper"
	"konekko.me/gosion/permission/pb"
	"konekko.me/gosion/permission/repositories"
	"konekko.me/gosion/safety/pb"
	"sync"
	"time"
)

type verificationService struct {
	pool               *redis.Pool
	session            *mgo.Session
	configuration      *gs_commons_config.GosionConfiguration
	applicationService gs_service_application.ApplicationService
	blacklistService   gs_service_safety.BlacklistService
	functionService    gs_service_permission.FunctionService
}

type requestHeaders struct {
	authorization string
	userAgent     string
	userDevice    string
	ip            string
	clientId      string
	path          string
	dat           string
}

func (svc *verificationService) GetRepo() permission_repositories.FunctionRepo {
	return permission_repositories.FunctionRepo{Session: svc.session.Clone(), Conn: svc.pool.Get()}
}

//application verify
//ip, userDevice blacklist verify
//api exists and authType verify
func (svc *verificationService) Test(ctx context.Context, in *gs_service_permission.HasPermissionRequest, out *gs_service_permission.HasPermissionResponse) error {
	return gs_commons_wrapper.ContextToAuthorize(ctx, out, func(auth *gs_commons_wrapper.WrapperUser) *gs_commons_dto.State {
		md, ok := metadata.FromContext(ctx)
		if ok {
			//side service contract
			//prevent loop
			traceId := md["X-Gosion-SCT"]
			if len(traceId) > 0 {
				_, err := gs_commons_encrypt.AESDecrypt(traceId, []byte(svc.configuration.ServiceContractSecretKey))
				if err != nil {
					return nil
				}
				return errstate.Success
			}

			//new request
			rh := &requestHeaders{
				authorization: md["Authorization"],
				userAgent:     md["User-Agent"],
				ip:            md["X-Real-Ip"],
				clientId:      md["GS-Client-Id"],
				userDevice:    md["GS-User-Device"],
				path:          md["GS-Request-Path"],
				dat:           md["GS-Duration-Token"],
			}

			//check
			if len(rh.authorization) == 0 || len(rh.userDevice) == 0 ||
				len(rh.clientId) == 0 || len(rh.userAgent) == 0 || len(rh.ip) == 0 ||
				len(rh.path) == 0 {
				return nil
			}

			var wg sync.WaitGroup
			wg.Add(3)

			state := errstate.Success

			id := gs_commons_generator.ID()

			traceId = base64.StdEncoding.EncodeToString([]byte(id.Generate().Base64() + uuid.NewV4().String()))

			traceId, err := gs_commons_encrypt.AESEncrypt([]byte(traceId), []byte(svc.configuration.ServiceContractSecretKey))
			if err != nil {
				return nil
			}

			resp := func(s *gs_commons_dto.State) {
				if state.Ok {
					state = s
				}
				wg.Done()
			}

			ctx = metadata.NewContext(context.Background(), map[string]string{
				"X-Gosion-SCT": traceId,
			})

			//blacklist(ip)
			go func() {
				s, err := svc.blacklistService.Check(ctx,
					&gs_service_safety.CheckRequest{
						Type: gs_commons_constants.BlacklistOfIP,
					})
				if err != nil {
					resp(errstate.ErrRequest)
					return
				}
				resp(s.State)
			}()

			//blacklist(userDevice)
			go func() {
				s, err := svc.blacklistService.Check(ctx,
					&gs_service_safety.CheckRequest{
						Type: gs_commons_constants.BlacklistOfUserDevice,
					})
				if err != nil {
					resp(errstate.ErrRequest)
					return
				}
				resp(s.State)
			}()

			var appResp *gs_service_application.StatusResponse

			//application
			go func() {
				s, err := svc.applicationService.Status(ctx,
					&gs_service_application.FindRequest{
						Content: rh.clientId,
					})
				if err != nil {
					resp(errstate.ErrRequest)
					return
				}
				resp(s.State)
				appResp = s
			}()

			wg.Wait()

			if !state.Ok {
				out.State = state
				return nil
			}

			if appResp != nil {

				repo := svc.GetRepo()
				defer repo.Close()

				a, err := repo.FindApiInCache(appResp.AppId, rh.path)
				if err != nil {
					return nil
				}
				dat := &permission_repositories.DurationAccess{}
				datFix, userId := "", ""
				wg.Add(len(a.AuthTypes))
				for _, v := range a.AuthTypes {
					go func() {
						switch v {
						case gs_commons_constants.AuthTypeOfValcode:
							if len(rh.dat) == 0 {
								resp(errstate.ErrRequest)
								return
							}
							conn := svc.pool.Get()

							b, err := redis.Bytes(conn.Do("get", rh.dat))
							if err != nil {
								resp(errstate.ErrRequest)
								return
							}

							err = msgpack.Unmarshal(b, &dat)
							if err != nil {
								resp(errstate.ErrSystem)
								return
							}
							if time.Now().UnixNano()-dat.ExpiredAt >= 0 {
								resp(errstate.ErrDurationAccessExpired)
								return
							}
							if dat.Path != rh.path || dat.ClientId != rh.clientId {
								resp(errstate.ErrDurationAccess)
								return
							}
							break
						case gs_commons_constants.AuthTypeOfToken:
							break
						case gs_commons_constants.AuthTypeOfFace:
							break
						case gs_commons_constants.AuthTypeOfPassword:
							break
						case gs_commons_constants.AuthTypeOfMobileConfirm:
							break
						}
					}()
				}

				wg.Wait()

				if len(dat.Path) > 0 {
					//the user is ip
					if len(userId) == 0 {
						userId = rh.ip
					}
					if dat.UserId != gs_commons_encrypt.SHA1(userId+rh.userAgent+rh.userDevice+rh.ip+datFix+rh.clientId) {
						return errstate.ErrDurationAccess
					}
				}

				if !state.Ok {
					out.State = state
					return nil
				}

				return errstate.Success
			}

		}

		return nil
	})
}

func NewVerificationService(pool *redis.Pool, session *mgo.Session) gs_service_permission.VerificationHandler {
	return &verificationService{pool: pool, session: session}
}
