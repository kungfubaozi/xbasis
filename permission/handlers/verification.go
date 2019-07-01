package permissionhandlers

import (
	"context"
	"fmt"
	"github.com/Sirupsen/logrus"
	"github.com/garyburd/redigo/redis"
	"github.com/micro/go-micro/metadata"
	"github.com/vmihailenco/msgpack"
	"gopkg.in/mgo.v2"
	"konekko.me/xbasis/analysis/client"
	"konekko.me/xbasis/application/pb/inner"
	"konekko.me/xbasis/authentication/pb/inner"
	"konekko.me/xbasis/commons/actions"
	xconfig "konekko.me/xbasis/commons/config"
	"konekko.me/xbasis/commons/config/call"
	constants "konekko.me/xbasis/commons/constants"
	commons "konekko.me/xbasis/commons/dto"
	"konekko.me/xbasis/commons/encrypt"
	"konekko.me/xbasis/commons/errstate"
	generator "konekko.me/xbasis/commons/generator"
	"konekko.me/xbasis/commons/indexutils"
	"konekko.me/xbasis/commons/wrapper"
	inner "konekko.me/xbasis/permission/pb/inner"
	"konekko.me/xbasis/safety/pb"
	"sync"
	"time"
)

type verificationService struct {
	pool                          *redis.Pool
	session                       *mgo.Session
	configuration                 *xconfig.GosionConfiguration
	innerApplicationStatusService xbasissvc_internal_application.ApplicationStatusService
	blacklistService              xbasissvc_external_safety.BlacklistService
	innerAuthService              xbasissvc_internal_authentication.AuthService
	*indexutils.Client
	log  analysisclient.LogClient
	_log *logrus.Logger
	id   generator.IDGenerator
}

type requestHeaders struct {
	authorization string
	userAgent     string
	userDevice    string
	ip            string
	refClientId   string //作为API share
	path          string
	dat           string
	fromClientId  string
}

func (svc *verificationService) GetRepo() *functionRepo {
	return &functionRepo{session: svc.session.Clone(), Client: svc.Client}
}

var whiteApiList = []string{"/authentication/router/refresh", "/authentication/router/logout"}

//application verify
//ip, userDevice blacklist verify
//api exists and authType verify
func (svc *verificationService) Check(ctx context.Context, in *inner.HasPermissionRequest, out *inner.HasPermissionResponse) error {
	var wg sync.WaitGroup
	return xbasiswrapper.ContextToAuthorize(ctx, out, func(auth *xbasiswrapper.WrapperUser) *commons.State {

		md, ok := metadata.FromContext(ctx)
		svc.configuration = serviceconfiguration.Get()
		if len(svc.configuration.CurrencySecretKey) == 0 {
			return errstate.ErrAuthorization
		}

		if ok {

			traceId := md["transport-trace-id"]
			if len(traceId) > 10 {

				_, err := encrypt.AESDecrypt(traceId, []byte(svc.configuration.CurrencySecretKey))
				if err != nil {
					return nil
				}

				out.FromClient = auth.FromClientId
				out.RefClientId = auth.RefClientId
				out.TraceId = traceId
				out.Ip = auth.IP
				out.UserDevice = auth.UserDevice
				out.UserAgent = auth.UserAgent
				out.User = auth.User
				out.AppId = auth.AppId
				out.AppType = auth.AppType
				out.Platform = auth.Platform
				if auth.Access != nil {
					out.DatTo = auth.Access.To
					if auth.Access.Auth {
						out.DatAuth = 2
					}
				}
				if auth.Token != nil {
					out.Token = &inner.TokenInfo{
						UserId:   auth.Token.UserId,
						ClientId: auth.Token.ClientId,
						Platform: auth.Token.ClientPlatform,
						AppType:  auth.Token.AppType,
						AppId:    auth.Token.AppId,
						Relation: auth.Token.Relation,
					}
				}

				return errstate.SuccessTraceCheck
			}

			start := time.Now().UnixNano()

			//new request
			rh := &requestHeaders{
				authorization: md["authorization"],
				userAgent:     md["x-user-agent"],
				ip:            md["x-real-ip"],
				fromClientId:  md["gs-client-id"],
				userDevice:    md["gs-user-device"],
				path:          md["gs-request-path"],
				dat:           md["gs-duration-token"],
				refClientId:   md["gs-client-sci"],
			}

			//check
			if len(rh.userDevice) == 0 ||
				len(rh.fromClientId) == 0 || len(rh.userAgent) == 0 || len(rh.ip) == 0 ||
				len(rh.path) == 0 {
				fmt.Println("basic check failed")
				return nil
			}

			if len(rh.refClientId) == 0 {
				rh.refClientId = rh.fromClientId
			}

			state := errstate.Success

			id := svc.id.String()

			traceId, err := encrypt.AESEncrypt([]byte(id), []byte(svc.configuration.CurrencySecretKey))
			if err != nil {
				return nil
			}

			resp := func(s *commons.State) {
				if state.Ok {
					state = s
				}
			}

			/**
			调用的顺序是refClientId第一
			*/
			headers := &analysisclient.LogHeaders{
				ServiceName:      constants.InternalPermission,
				ModuleName:       "Verification",
				UserAgent:        rh.userAgent,
				RefClientId:      rh.refClientId,
				Device:           rh.userDevice,
				HasAccessToken:   len(rh.authorization) != 0,
				HasDurationToken: len(rh.dat) != 0,
				FromClientId:     rh.fromClientId,
				Ip:               rh.ip,
				Path:             rh.path,
				TraceId:          traceId,
			}

			svc.log.Info(&analysisclient.LogContent{
				Headers:   headers,
				Action:    "PermissionVerification",
				Message:   "start verification",
				StateCode: 0,
			})

			ctx = metadata.NewContext(ctx, map[string]string{
				"transport-trace-id": traceId,
			})

			var conn redis.Conn

			wg.Add(3)

			callClientId := rh.fromClientId
			if len(rh.refClientId) > 0 {
				callClientId = rh.refClientId
			}

			//blacklist(ip)
			go func() {
				defer wg.Done()
				s, err := svc.blacklistService.Check(ctx,
					&xbasissvc_external_safety.CheckRequest{
						Type:    constants.BlacklistOfIP,
						Content: rh.ip,
					})
				if err != nil {
					fmt.Println("err blacklist ip", err)
					resp(errstate.ErrRequest)
					return
				}
				resp(s.State)
			}()

			//blacklist(userDevice)
			go func() {
				defer wg.Done()
				s, err := svc.blacklistService.Check(ctx,
					&xbasissvc_external_safety.CheckRequest{
						Type:    constants.BlacklistOfUserDevice,
						Content: rh.userDevice,
					})
				if err != nil {
					fmt.Println("err blacklist user device", err)
					resp(errstate.ErrRequest)
					return
				}
				resp(s.State)
			}()

			var appResp *xbasissvc_internal_application.GetAppClientStatusResponse

			////application
			go func() {
				defer wg.Done()
				s, err := svc.innerApplicationStatusService.GetAppClientStatus(ctx, &xbasissvc_internal_application.GetAppClientStatusRequest{
					ClientId: callClientId,
				})
				if err != nil {
					resp(errstate.ErrRequest)
					fmt.Println("err app", err)
					return
				}
				resp(s.State)
				appResp = s
			}()

			wg.Wait()

			if !state.Ok {
				svc.log.Info(&analysisclient.LogContent{
					Headers: &analysisclient.LogHeaders{
						TraceId: traceId,
					},
					Action:    "ThreeBasicValidations",
					Message:   "CheckFailed",
					StateCode: state.Code,
				})

				return state
			}

			if appResp != nil && len(appResp.AppId) > 0 {

				if appResp.ClientEnabled != constants.Enabled {
					return errstate.ErrClientClosed
				}

				repo := svc.GetRepo()
				defer repo.Close()

				var f *simplifiedFunction
				var err error
				for _, v := range whiteApiList {
					if v == rh.path {
						f = &simplifiedFunction{
							AuthTypes: []int64{},
						}
						break
					}
				}
				if f == nil {
					f, err = repo.SimplifiedLookupApi(appResp.AppId, encrypt.SHA1(rh.path))
					if err != nil {
						svc.log.Info(&analysisclient.LogContent{
							Headers: &analysisclient.LogHeaders{
								TraceId: traceId,
							},
							Action: "InvalidApi",
						})
						return nil
					}

					//grant platform
					if f.GrantPlatforms != nil && len(f.GrantPlatforms) > 0 {
						for _, v := range f.GrantPlatforms {
							if v == appResp.ClientPlatform {
								svc.log.Info(&analysisclient.LogContent{
									Headers: &analysisclient.LogHeaders{
										TraceId: traceId,
									},
									Action: "ApiPlatformAccessDenied",
								})
								return errstate.ErrRequest
							}
						}
					}
				}

				svc.log.Info(&analysisclient.LogContent{
					Headers: headers,
					Action:  "RequestApi",
					Message: headers.Path,
					Fields: &analysisclient.LogFields{
						"id":    f.Id,
						"appId": appResp.AppId,
					},
				})

				out.AppType = appResp.Type

				userId := ""
				cv := ""
				wg.Add(len(f.AuthTypes))

				var credential *durationAccessCredential
				var dat *durationAccessToken

				//清除凭证(可能会遇到其他认证不通过的情况，发送验证码的前提是: 其余的验证条件全部满足)
				cclear := false

				auth := false

				for _, v := range f.AuthTypes {
					go func() {
						defer wg.Done()
						switch v {
						case constants.AuthTypeOfValcode:
							if len(rh.dat) == 0 {
								//生成生成验证码凭证
								//(先请求原API，生成凭证，然后调用dat获取验证码，再请求API)
								resp(errstate.Success)

								credential = &durationAccessCredential{
									FromClientId: rh.fromClientId,
									RefClientId:  rh.refClientId,
									FuncId:       f.Id,
									AppId:        f.AppId,
									Timestamp:    time.Now().UnixNano(),
								}

								return
							}

							v, err := encrypt.AESDecrypt(rh.dat, []byte(svc.configuration.CurrencySecretKey))
							if err != nil {
								resp(errstate.ErrNotFoundDurationAccessToken)
								return
							}

							cv = v

							if conn == nil {
								conn = svc.pool.Get()
							}

							fmt.Println("v", v)

							b, err := redis.Bytes(conn.Do("hget", "dat."+f.Id, v))
							if err != nil {
								resp(errstate.ErrRequest)
								return
							}

							err = msgpack.Unmarshal(b, &dat)
							if err != nil {
								resp(errstate.ErrSystem)
								return
							}

							break
						case constants.AuthTypeOfToken:
							//1.check token

							auth = true

							if len(rh.authorization) == 0 {
								resp(errstate.ErrAccessToken)
								cclear = true
								return
							}

							ac := metadata.NewContext(context.Background(), map[string]string{
								"transport-user-agent":      rh.userAgent,
								"transport-app-id":          appResp.AppId,
								"transport-ip":              rh.ip,
								"transport-from-client-id":  rh.fromClientId,
								"transport-ref-client-id":   rh.refClientId,
								"transport-trace-id":        traceId,
								"transport-user-device":     rh.userDevice,
								"transport-client-platform": fmt.Sprintf("%d", appResp.ClientPlatform),
							})

							status, err := svc.innerAuthService.Verify(ac, &xbasissvc_internal_authentication.VerifyRequest{
								Token:    rh.authorization,
								ClientId: callClientId,
								//FunctionRoles: f.Roles,
								FunctionId: f.Id,
								Share:      f.Share,
								AppId:      appResp.AppId,
							})
							if err != nil {
								resp(errstate.ErrSystem)
								cclear = true
								return
							}

							if status.State.Ok {
								out.Token = &inner.TokenInfo{
									UserId:   status.UserId,
									ClientId: status.ClientId,
									Platform: status.ClientPlatform,
									AppId:    status.AppId,
									Relation: status.Relation,
									AppType:  status.AppType,
								}

								userId = status.UserId

							} else {
								svc.log.Info(&analysisclient.LogContent{
									Headers: &analysisclient.LogHeaders{
										TraceId: traceId,
									},
									Action:  "ExtAuthVerify",
									Message: "VerificationFailed",
								})
							}

							cclear = true
							resp(status.State)

							break
						case constants.AuthTypeOfFace:
							break
						case constants.AuthTypeOfMobileConfirm:
							break
						}
					}()
				}

				wg.Wait()

				if !state.Ok {
					out.State = state
					out.Token = nil
					svc.log.Info(&analysisclient.LogContent{
						Headers: &analysisclient.LogHeaders{
							TraceId: traceId,
						},
						Action:    "BasicApplicationInfoCheck",
						Message:   "CheckFailed",
						StateCode: out.State.Code,
					})
					return nil
				}

				if dat != nil {
					from := encrypt.SHA1(rh.ip + rh.userAgent + rh.userDevice + rh.fromClientId)
					if dat.From != from {
						return errstate.ErrDurationAccess
					}
					user := ""
					if auth {
						if len(userId) == 0 {
							return errstate.ErrRequest
						}
						user = userId
					}
					if dat.FuncId != f.Id || dat.ClientId != rh.fromClientId || (auth && user != dat.User) {
						return errstate.ErrDurationAccess
					}

					if dat.MaxTimes >= dat.Times {
						dat.Times = dat.Times + 1
						b, err := msgpack.Marshal(dat)
						if err != nil {
							return errstate.ErrSystem
						}
						_, err = conn.Do("hset", "dat."+f.Id, cv, b)
						if err != nil {
							return errstate.ErrSystem
						}
					} else {
						//delete access token
						conn.Do("del", "dat."+f.Id, cv)
						return errstate.ErrDurationAccessExpired
					}

				}

				if cclear {
					credential = nil
				}

				//发送验证码凭证
				if credential != nil {

					//从auth获取发送对象
					if auth {
						credential.FromAuth = true
					}

					b, err := msgpack.Marshal(credential)
					if err != nil {
						return errstate.ErrSystem
					}

					c, err := encrypt.AESEncrypt(b, []byte(svc.configuration.CurrencySecretKey))
					if err != nil {
						return errstate.ErrSystem
					}

					state = errstate.ErrDurationAccessCredential
					state.Credential = c

					return state
				}

				out.AppId = appResp.AppId
				out.UserAgent = rh.userAgent
				out.UserDevice = rh.userDevice
				out.FromClient = rh.fromClientId
				out.RefClientId = rh.refClientId
				out.Ip = rh.ip
				out.FunctionId = f.Id
				out.TraceId = traceId
				out.AppType = appResp.Type
				out.Platform = appResp.ClientPlatform

				if conn != nil {
					conn.Close()
				}

				if len(userId) == 0 {
					userId = rh.ip
				}
				out.User = userId
				if dat != nil {
					var v1 int64
					if dat.Auth {
						v1 = 1
					} else {
						v1 = 2
					}
					out.DatAuth = v1
					out.DatTo = dat.User
				}

				svc.log.Info(&analysisclient.LogContent{
					Headers: &analysisclient.LogHeaders{
						TraceId: traceId,
						UserId:  userId,
					},
					Action:    loggeractions.UserRequestApi,
					Message:   "verification passed",
					StateCode: 0,
					Fields: &analysisclient.LogFields{
						"id":             f.Id,
						"user_id":        userId,
						"from_client_id": headers.FromClientId,
						"ref_client_id":  headers.RefClientId,
						"app_id":         appResp.AppId,
					},
				})

				svc._log.WithFields(logrus.Fields{
					"taking": fmt.Sprintf("%dms", (time.Now().UnixNano()-start)/1e6),
				}).Info("Time consuming")

				return errstate.Success
			}

		}

		return errstate.ErrRequest
	})
}

func NewVerificationService(pool *redis.Pool, session *mgo.Session,
	innerApplicationStatusService xbasissvc_internal_application.ApplicationStatusService, blacklistService xbasissvc_external_safety.BlacklistService,
	innerAuthService xbasissvc_internal_authentication.AuthService, client *indexutils.Client, logger analysisclient.LogClient) inner.VerificationHandler {
	return &verificationService{pool: pool, session: session, innerApplicationStatusService: innerApplicationStatusService, id: generator.NewIDG(),
		blacklistService: blacklistService, innerAuthService: innerAuthService, Client: client, log: logger, _log: logrus.New()}
}
