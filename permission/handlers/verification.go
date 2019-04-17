package permissionhandlers

import (
	"context"
	"fmt"
	"github.com/Sirupsen/logrus"
	"github.com/garyburd/redigo/redis"
	"github.com/micro/go-micro/metadata"
	"github.com/vmihailenco/msgpack"
	"gopkg.in/mgo.v2"
	"konekko.me/gosion/application/pb/ext"
	"konekko.me/gosion/authentication/pb/ext"
	"konekko.me/gosion/commons/config"
	"konekko.me/gosion/commons/config/call"
	"konekko.me/gosion/commons/constants"
	"konekko.me/gosion/commons/dto"
	"konekko.me/gosion/commons/encrypt"
	"konekko.me/gosion/commons/errstate"
	"konekko.me/gosion/commons/generator"
	"konekko.me/gosion/commons/gslogrus"
	"konekko.me/gosion/commons/indexutils"
	"konekko.me/gosion/commons/wrapper"
	"konekko.me/gosion/permission/pb/ext"
	"konekko.me/gosion/safety/pb"
	"sync"
)

type verificationService struct {
	pool                        *redis.Pool
	session                     *mgo.Session
	configuration               *gs_commons_config.GosionConfiguration
	extApplicationStatusService gs_ext_service_application.ApplicationStatusService
	blacklistService            gs_service_safety.BlacklistService
	extAuthService              gs_ext_service_authentication.AuthService
	*indexutils.Client
	*gslogrus.Logger
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

func (svc *verificationService) GetRepo() *functionRepo {
	return &functionRepo{session: svc.session.Clone(), Client: svc.Client}
}

//application verify
//ip, userDevice blacklist verify
//api exists and authType verify
func (svc *verificationService) Check(ctx context.Context, in *gs_ext_service_permission.HasPermissionRequest, out *gs_ext_service_permission.HasPermissionResponse) error {
	var wg sync.WaitGroup

	return gs_commons_wrapper.ContextToAuthorize(ctx, out, func(auth *gs_commons_wrapper.WrapperUser) *gs_commons_dto.State {

		md, ok := metadata.FromContext(ctx)
		svc.configuration = serviceconfiguration.Get()
		if len(svc.configuration.CurrencySecretKey) == 0 {
			return errstate.ErrAuthorization
		}

		if ok {

			traceId := md["transport-trace-id"]
			if len(traceId) > 0 {
				_, err := encrypt.AESDecrypt(traceId, []byte(svc.configuration.CurrencySecretKey))
				if err != nil {
					return nil
				}

				out.ClientId = auth.ClientId
				out.TraceId = traceId
				out.Ip = auth.IP
				out.UserDevice = auth.UserDevice
				out.UserAgent = auth.UserAgent
				out.User = auth.User
				out.AppId = auth.AppId

				return errstate.SuccessTraceCheck
			}

			//new request
			rh := &requestHeaders{
				authorization: md["authorization"],
				userAgent:     md["x-user-agent"],
				ip:            md["x-real-ip"],
				clientId:      md["gs-client-id"],
				userDevice:    md["gs-user-device"],
				path:          md["gs-request-path"],
				dat:           md["gs-duration-token"],
			}

			//check
			if len(rh.userDevice) == 0 ||
				len(rh.clientId) == 0 || len(rh.userAgent) == 0 || len(rh.ip) == 0 ||
				len(rh.path) == 0 {
				return nil
			}

			state := errstate.Success

			id := gs_commons_generator.NewIDG().String()

			traceId, err := encrypt.AESEncrypt([]byte(id), []byte(svc.configuration.CurrencySecretKey))
			if err != nil {
				return nil
			}

			resp := func(s *gs_commons_dto.State) {
				if state.Ok {
					state = s
				}
			}

			log := svc.WithHeaders(traceId, rh.clientId, rh.ip, rh.path, rh.userAgent, rh.userDevice)

			ctx = metadata.NewContext(ctx, map[string]string{
				"transport-trace-id": traceId,
			})

			var conn redis.Conn

			wg.Add(3)

			//blacklist(ip)
			go func() {
				defer wg.Done()
				s, err := svc.blacklistService.Check(ctx,
					&gs_service_safety.CheckRequest{
						Type:    gs_commons_constants.BlacklistOfIP,
						Content: rh.ip,
					})
				if err != nil {
					resp(errstate.ErrRequest)
					return
				}
				resp(s.State)
			}()

			//blacklist(userDevice)
			go func() {
				defer wg.Done()
				s, err := svc.blacklistService.Check(ctx,
					&gs_service_safety.CheckRequest{
						Type:    gs_commons_constants.BlacklistOfUserDevice,
						Content: rh.userDevice,
					})
				if err != nil {
					resp(errstate.ErrRequest)
					return
				}
				resp(s.State)
			}()

			var appResp *gs_ext_service_application.GetAppClientStatusResponse

			////application
			go func() {
				defer wg.Done()
				s, err := svc.extApplicationStatusService.GetAppClientStatus(ctx, &gs_ext_service_application.GetAppClientStatusRequest{
					ClientId: rh.clientId,
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
				log.WithAction("ThreeBasicValidations", logrus.Fields{
					"state": state.Code,
				}).Info("check failed.")
				return state
			}

			if appResp != nil && len(appResp.UserStructure) > 0 && len(appResp.FunctionStructure) > 0 {

				if appResp.ClientEnabled != gs_commons_constants.Enabled {
					return errstate.ErrClientClosed
				}

				repo := svc.GetRepo()
				defer repo.Close()

				//fmt.Println("function structure id", ccs.FunctionStructureId)
				f, err := repo.SimplifiedLookupApi(appResp.FunctionStructure, rh.path)
				if err != nil {
					fmt.Println("invalid api", rh.path)
					return nil
				}

				//grant platform
				if f.GrantPlatforms != nil && len(f.GrantPlatforms) > 0 {
					for _, v := range f.GrantPlatforms {
						if v == appResp.ClientPlatform {
							return errstate.ErrRequest
						}
					}
				}

				out.CurrentMain = appResp.Main

				dat := &durationAccess{}
				userId := ""
				wg.Add(len(f.AuthTypes))
				for _, v := range f.AuthTypes {
					go func() {
						defer wg.Done()
						switch v {
						case gs_commons_constants.AuthTypeOfValcode:
							//user must login
							if len(rh.dat) == 0 {
								resp(errstate.ErrRequest)
								return
							}

							v, err := encrypt.AESDecrypt(rh.dat, []byte(svc.configuration.CurrencySecretKey))
							if err != nil {
								resp(errstate.ErrNotFoundDurationAccessToken)
								return
							}

							key := encrypt.SHA1(string(v) + rh.clientId)
							dat.Key = key
							b, err := redis.Bytes(conn.Do("hget", key, f.Id))
							if err != nil && err == redis.ErrNil {
								resp(errstate.ErrRequest)
								return
							}

							err = msgpack.Unmarshal(b, &dat)
							if err != nil {
								resp(errstate.ErrSystem)
								return
							}
							if dat.Stat != rh.dat || dat.Path != rh.path || dat.ClientId != rh.clientId {
								resp(errstate.ErrDurationAccess)
								return
							}
							//if time.Now().UnixNano()-dat.ExpiredAt >= 0 {
							//	resp(errstate.ErrDurationAccessExpired)
							//	return
							//}
							resp(errstate.Success)
							break
						case gs_commons_constants.AuthTypeOfToken:
							//1.check token

							ac := metadata.NewContext(context.Background(), map[string]string{
								"transport-user-agent":      rh.userAgent,
								"transport-app-id":          appResp.AppId,
								"transport-ip":              rh.ip,
								"transport-client-id":       rh.clientId,
								"transport-trace-id":        traceId,
								"transport-user-device":     rh.userDevice,
								"transport-client-platform": fmt.Sprintf("%d", appResp.ClientPlatform),
							})

							status, err := svc.extAuthService.Verify(ac, &gs_ext_service_authentication.VerifyRequest{
								Token:         rh.authorization,
								ClientId:      rh.clientId,
								FunctionRoles: f.Roles,
								Funcs:         appResp.FunctionStructure,
							})
							if err != nil {
								resp(errstate.ErrSystem)
								return
							}

							if status.State.Ok {
								out.Token = &gs_ext_service_permission.TokenInfo{
									UserId:   status.UserId,
									ClientId: status.ClientId,
									Platform: status.ClientPlatform,
									AppId:    status.AppId,
									Relation: status.Relation,
								}

								userId = status.UserId
							} else {
								log.WithAction("ExtAuthVerify", logrus.Fields{
									"state": status.State.Code,
									"who":   status.UserId,
								}).Info("verification failed.")
							}

							resp(status.State)

							break
						case gs_commons_constants.AuthTypeOfFace:
							break
						case gs_commons_constants.AuthTypeOfMobileConfirm:
							break
						}
					}()
				}

				wg.Wait()

				if len(dat.Path) > 0 {
					if len(userId) == 0 {
						userId = rh.ip
					}
					if dat.User != userId {
						out.Token = nil
						return errstate.ErrDurationAccess
					}

				}

				if !state.Ok {
					out.State = state
					out.Token = nil
					log.WithAction("BasicApplicationInfoCheck", logrus.Fields{
						"state": state.Code,
					}).Info("check failed.")
					return nil
				}

				out.AppId = appResp.AppId
				out.UserAgent = rh.userAgent
				out.UserDevice = rh.userDevice
				out.ClientId = rh.clientId
				out.Ip = rh.ip
				out.TraceId = traceId
				out.Platform = appResp.ClientPlatform
				if len(userId) == 0 {
					userId = rh.ip
				}
				out.User = userId

				log.WithAction("Verification", logrus.Fields{
					"state": errstate.Success.Code,
					"who":   userId,
				}).Info("verification success.")

				return errstate.Success
			}

		}

		return errstate.ErrRequest
	})
}

func NewVerificationService(pool *redis.Pool, session *mgo.Session,
	extApplicationStatusService gs_ext_service_application.ApplicationStatusService, blacklistService gs_service_safety.BlacklistService,
	extAuthService gs_ext_service_authentication.AuthService, client *indexutils.Client, log *gslogrus.Logger) gs_ext_service_permission.VerificationHandler {
	return &verificationService{pool: pool, session: session, extApplicationStatusService: extApplicationStatusService,
		blacklistService: blacklistService, extAuthService: extAuthService, Client: client, Logger: log}
}
