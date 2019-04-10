package permissionhandlers

import (
	"context"
	"fmt"
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
	"konekko.me/gosion/commons/wrapper"
	"konekko.me/gosion/permission/pb"
	"konekko.me/gosion/permission/utils"
	"konekko.me/gosion/safety/pb"
	"sync"
	"time"
)

type verificationService struct {
	pool                        *redis.Pool
	session                     *mgo.Session
	configuration               *gs_commons_config.GosionConfiguration
	extApplicationStatusService gs_ext_service_application.ApplicationStatusService
	blacklistService            gs_service_safety.BlacklistService
	extAuthService              gs_ext_service_authentication.AuthService
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

func (svc *verificationService) GetRepo(conn redis.Conn) *functionRepo {
	return &functionRepo{session: svc.session.Clone(), conn: conn}
}

var openPermission = false

//application verify
//ip, userDevice blacklist verify
//api exists and authType verify
func (svc *verificationService) Check(ctx context.Context, in *gs_service_permission.HasPermissionRequest, out *gs_service_permission.HasPermissionResponse) error {
	a := time.Now().UnixNano()
	var wgc sync.WaitGroup
	wgc.Add(1)
	var wg sync.WaitGroup
	var conn redis.Conn
	go func() {
		//Maybe it's a network problem. It takes a long time to connect, so start a collaboration to establish the connection.
		defer wgc.Done()
		conn = svc.pool.Get()
	}()

	return gs_commons_wrapper.ContextToAuthorize(ctx, out, func(auth *gs_commons_wrapper.WrapperUser) *gs_commons_dto.State {

		fmt.Println("time.now-wrapper", (time.Now().UnixNano()-a)/1e6)

		md, ok := metadata.FromContext(ctx)
		svc.configuration = serviceconfiguration.Get()
		if len(svc.configuration.CurrencySecretKey) == 0 {
			return errstate.ErrAuthorization
		}

		if ok {

			fmt.Println("time.now-wrapper-1", (time.Now().UnixNano()-a)/1e6)

			//side service contract
			//prevent loop

			//fmt.Println("ctx", md)

			traceId := md["transport-trace-id"]
			if len(traceId) > 0 {
				_, err := gs_commons_encrypt.AESDecrypt(traceId, []byte(svc.configuration.CurrencySecretKey))
				if err != nil {
					return nil
				}
				//fmt.Println("entry traceId is:", traceId)
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

			fmt.Println("request headers:", rh)

			//check
			if len(rh.userDevice) == 0 ||
				len(rh.clientId) == 0 || len(rh.userAgent) == 0 || len(rh.ip) == 0 ||
				len(rh.path) == 0 {
				fmt.Println("check failed")
				return nil
			}

			state := errstate.Success

			id := gs_commons_generator.NewIDG().String()

			fmt.Println("time.now-2", (time.Now().UnixNano()-a)/1e6)

			traceId, err := gs_commons_encrypt.AESEncrypt([]byte(id), []byte(svc.configuration.CurrencySecretKey))
			if err != nil {
				return nil
			}

			resp := func(s *gs_commons_dto.State) {
				if state.Ok {
					state = s
				}
			}

			ctx = metadata.NewContext(ctx, map[string]string{
				"transport-trace-id": traceId,
			})

			fmt.Println("traceId:", traceId)

			//conn := svc.pool.Get()
			//var conn redis.Conn

			fmt.Println("time.now-3", (time.Now().UnixNano()-a)/1e6)

			wg.Add(3)

			//blacklist(ip)
			go func() {
				fmt.Println("time.func1-", (time.Now().UnixNano()-a)/1e6)
				defer wg.Done()
				s, err := svc.blacklistService.Check(ctx,
					&gs_service_safety.CheckRequest{
						Type:    gs_commons_constants.BlacklistOfIP,
						Content: rh.ip,
					})
				if err != nil {
					fmt.Println("err service", err)
					resp(errstate.ErrRequest)
					return
				}
				fmt.Println("ip clear", s.State)
				resp(s.State)
				fmt.Println("time.func1-", (time.Now().UnixNano()-a)/1e6)
			}()

			//blacklist(userDevice)
			go func() {
				fmt.Println("time.func-2", (time.Now().UnixNano()-a)/1e6)
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
				fmt.Println("userDevice clear", s.State)
				resp(s.State)
				fmt.Println("time.func-2", (time.Now().UnixNano()-a)/1e6)
			}()

			var appResp *gs_ext_service_application.GetAppClientStatusResponse
			ccs := &cacheStructure{}

			////application
			go func() {
				defer wg.Done()
				s, err := svc.extApplicationStatusService.GetAppClientStatus(ctx, &gs_ext_service_application.GetAppClientStatusRequest{
					ClientId: rh.clientId,
				})
				if err != nil {
					fmt.Println("err")
					resp(errstate.ErrRequest)
					return
				}
				fmt.Println("app", s.State)
				resp(s.State)
				appResp = s
				if s != nil && s.State.Ok && len(s.AppId) > 0 {
					wgc.Wait()
					//get current structure id
					fmt.Println("check application structure")
					get := func(t int64) (string, error) {
						return redis.String(conn.Do("get",
							permissionutils.GetTypeCurrentStructureKey(s.AppId, t)))
					}

					v, err := get(permissionutils.TypeUserStructure)
					if err != nil {
						fmt.Println("err", err)
					}
					if err == nil && len(v) > 0 {
						ccs.UserStructureId = v
						v, err = get(permissionutils.TypeFunctionStructure)
						if err == nil && len(v) > 0 {
							fmt.Println("check application structure,-> clear")
							ccs.FunctionStructureId = v
						}
					}
				}
			}()

			wg.Wait()

			if !state.Ok {
				fmt.Println("time.stop", (time.Now().UnixNano()-a)/1e6)
				fmt.Println("basic check failed:", state.Message)
				return state
			}

			fmt.Println("time.now", (time.Now().UnixNano()-a)/1e6)

			if appResp != nil && len(ccs.UserStructureId) > 0 && len(ccs.FunctionStructureId) > 0 {

				fmt.Println("entry check process")

				if appResp.ClientEnabled != gs_commons_constants.Enabled {
					return errstate.ErrClientClosed
				}

				repo := svc.GetRepo(conn)
				defer repo.Close()

				a, err := repo.FindApiInCache(ccs.FunctionStructureId, rh.path)
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

				dat := &durationAccess{}
				userId := ""
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
							b, err := redis.Bytes(conn.Do("hmget", key, a.Api))
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
							status, err := svc.extAuthService.Verify(ctx, &gs_ext_service_authentication.VerifyRequest{
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
							if !openPermission {
								resp(errstate.Success)
								return
							}
							//2.check user roles
							//appId.userId
							var userRoles, functionRoles []interface{}
							var swg sync.WaitGroup
							swg.Add(2)
							urk := permissionutils.GetStructureUserRoleKey(ccs.UserStructureId, status.Content)
							frk := permissionutils.GetStructureFunctionRoleKey(ccs.UserStructureId, a.Id)
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
										_, err := conn.Do("hmget", permissionutils.GetStructureRoleKey(ccs.UserStructureId), b)
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

		return errstate.ErrRequest
	})
}

func NewVerificationService(pool *redis.Pool, session *mgo.Session,
	extApplicationStatusService gs_ext_service_application.ApplicationStatusService, blacklistService gs_service_safety.BlacklistService,
	extAuthService gs_ext_service_authentication.AuthService) gs_service_permission.VerificationHandler {
	return &verificationService{pool: pool, session: session, extApplicationStatusService: extApplicationStatusService,
		blacklistService: blacklistService, extAuthService: extAuthService}
}
