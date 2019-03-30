package permission_handlers

import (
	"context"
	"encoding/base64"
	"github.com/garyburd/redigo/redis"
	"github.com/micro/go-micro/metadata"
	"github.com/twinj/uuid"
	"github.com/vmihailenco/msgpack"
	"gopkg.in/mgo.v2"
	"konekko.me/gosion/application/pb/nops"
	"konekko.me/gosion/authentication/pb/nops"
	"konekko.me/gosion/commons/config"
	"konekko.me/gosion/commons/constants"
	"konekko.me/gosion/commons/dto"
	"konekko.me/gosion/commons/encrypt"
	"konekko.me/gosion/commons/errstate"
	"konekko.me/gosion/commons/generator"
	"konekko.me/gosion/commons/wrapper"
	"konekko.me/gosion/permission/pb"
	"konekko.me/gosion/permission/repositories"
	"konekko.me/gosion/permission/uitls"
	"konekko.me/gosion/safety/pb"
	"sync"
)

type verificationService struct {
	pool                        *redis.Pool
	session                     *mgo.Session
	configuration               *gs_commons_config.GosionConfiguration
	nopApplicationStatusService gs_nops_service_application.ApplicationStatusService
	blacklistService            gs_service_safety.BlacklistService
	functionService             gs_service_permission.FunctionService
	nopAuthService              gs_nops_service_authentication.AuthService
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
			traceId := md["transport-traceId"]
			if len(traceId) > 0 {
				_, err := gs_commons_encrypt.AESDecrypt(traceId, []byte(svc.configuration.CurrencySecretKey))
				if err != nil {
					return nil
				}
				return errstate.Success
			}

			//new request
			rh := &requestHeaders{
				authorization: md["authorization"],
				userAgent:     md["user-agent"],
				ip:            md["x-real-ip"],
				clientId:      md["gs-client-id"],
				userDevice:    md["gs-user-device"],
				path:          md["gs-request-path"],
				dat:           md["gs-duration-token"],
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

			traceId, err := gs_commons_encrypt.AESEncrypt([]byte(traceId), []byte(svc.configuration.CurrencySecretKey))
			if err != nil {
				return nil
			}

			resp := func(s *gs_commons_dto.State) {
				if state.Ok {
					state = s
				}
			}

			ctx = metadata.NewContext(context.Background(), map[string]string{
				"transport-traceId": traceId,
			})

			//blacklist(ip)
			go func() {
				defer wg.Done()
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
				defer wg.Done()
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

			var appResp *gs_nops_service_application.GetAppClientStatusResponse

			//application
			go func() {
				defer wg.Done()
				s, err := svc.nopApplicationStatusService.GetAppClientStatus(ctx, &gs_nops_service_application.GetAppClientStatusRequest{
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
				out.State = state
				return nil
			}

			if appResp != nil {

				if appResp.ClientEnabled != gs_commons_constants.Enabled {
					return errstate.ErrClientClosed
				}

				repo := svc.GetRepo()
				defer repo.Close()

				a, err := repo.FindApiInCache(appResp.AppId, rh.path)
				if err != nil {
					return nil
				}

				//grant platform
				if a.GrantPlatforms != nil && len(a.GrantPlatforms) > 0 {
					for _, v := range a.GrantPlatforms {
						if v == appResp.ClientPlatform {
							return errstate.ErrRequest
						}
					}
				}

				dat := &permission_repositories.DurationAccess{}
				userId := ""
				conn := svc.pool.Get()
				wg.Add(len(a.AuthTypes))
				for _, v := range a.AuthTypes {
					go func() {
						defer wg.Done()
						switch v {
						case gs_commons_constants.AuthTypeOfValcode:
							//user must login
							if len(rh.dat) == 0 {
								resp(errstate.ErrRequest)
								return
							}

							v, err := gs_commons_encrypt.AESDecrypt(rh.dat, []byte(svc.configuration.CurrencySecretKey))
							if err != nil {
								resp(errstate.ErrNotFoundDurationAccessToken)
								return
							}

							key := gs_commons_encrypt.SHA1(string(v) + rh.clientId)
							dat.Key = key
							b, err := redis.Bytes(conn.Do("hmget", key, a.ApiTag))
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
							status, err := svc.nopAuthService.Verify(ctx, &gs_nops_service_authentication.VerifyRequest{
								Token:    rh.authorization,
								ClientId: rh.clientId,
							})
							if err != nil {
								resp(errstate.ErrSystem)
								return
							}
							if !status.State.Ok {
								resp(status.State)
								return
							}
							//2.check user roles
							//appId.userId
							var userRoles, functionRoles []interface{}
							var swg sync.WaitGroup
							swg.Add(2)
							urk := permission_uitls.GetAppUserRoleKey(appResp.AppId, status.Content)
							frk := permission_uitls.GetAppFunctionRoleKey(appResp.AppId, a.Id)
							go func() {
								defer swg.Done()
								userRoles, err = redis.Values(conn.Do("SMEMBERS", urk))
							}()
							go func() {
								defer swg.Done()
								functionRoles, err = redis.Values(conn.Do("SMEMBERS", frk))
							}()
							swg.Wait()
							if userRoles != nil && functionRoles != nil && len(userRoles) > 0 && len(functionRoles) > 0 {
								roles := make(map[string]string)
								ok := false
								for _, v := range userRoles {
									b := string(v.([]byte))
									roles[b] = "ok"
								}
								for _, v := range functionRoles {
									b := string(v.([]byte))
									//The current design is to delete the role only by deleting the corresponding data,
									//not deleting the data corresponding to the role, so we need to do a layer of dynamic deletion.
									if roles[b] != "ok" {
										//check role
										_, err := conn.Do("hmget", permission_uitls.GetAppRoleKey(appResp.AppId), b)
										if err != nil && err == redis.ErrNil { //invalid role
											//possibly due to the removal of roles
											//remove role
											conn.Do("srem", urk, b)
											conn.Do("srem", frk, b)
											break
										}
										ok = true
									}
								}
								if ok {
									resp(errstate.Success)
									return
								} else {
									resp(errstate.ErrUserPermission)
									return
								}
							} else {
								resp(errstate.ErrRequest)
								return
							}
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
