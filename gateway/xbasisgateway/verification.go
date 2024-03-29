package xbasisgateway

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/Sirupsen/logrus"
	"github.com/micro/go-micro/metadata"
	"github.com/vmihailenco/msgpack"
	"io/ioutil"
	"konekko.me/xbasis/analysis/client"
	"konekko.me/xbasis/application/pb/inner"
	"konekko.me/xbasis/authentication/pb/inner"
	"konekko.me/xbasis/commons/actions"
	"konekko.me/xbasis/commons/config/call"
	constants "konekko.me/xbasis/commons/constants"
	"konekko.me/xbasis/commons/date"
	commons "konekko.me/xbasis/commons/dto"
	"konekko.me/xbasis/commons/encrypt"
	"konekko.me/xbasis/commons/errstate"
	"konekko.me/xbasis/commons/transport"
	"konekko.me/xbasis/permission/handlers"
	"konekko.me/xbasis/permission/pb/inner"
	"konekko.me/xbasis/safety/pb"
	"reflect"
	"sync"
	"time"
)

type requestHeaders struct {
	authorization string
	userAgent     string
	userDevice    string
	ip            string
	refClientId   string
	path          string
	dat           string
	fromClientId  string
}

var whiteApiList = []string{"/authentication/router/logout"}

func (r *request) verification() bool {

	r.startAt = time.Now().UnixNano()

	clientId := r.c.Request.Header.Get("xbs-client-id")
	if len(clientId) != 12 {
		r.json(errstate.ErrInvalidClientId)
		return false
	}
	r.c.Request.Header.Del("xbs-client-id")
	ip := r.c.ClientIP()
	if len(ip) < 8 {
		r.json(errstate.ErrRequest)
		return false
	}
	userDevice := r.c.Request.Header.Get("xbs-user-device")
	if len(userDevice) < 8 {
		r.json(errstate.ErrRequest)
		return false
	}
	r.c.Request.Header.Del("xbs-user-device")

	rh := &requestHeaders{
		authorization: r.c.Request.Header.Get("authorization"),
		userAgent:     r.c.Request.UserAgent(),
		ip:            ip,
		fromClientId:  clientId,
		userDevice:    userDevice,
		path:          r.c.Request.URL.Path,
		dat:           r.c.Request.FormValue("_g.dat"),
		refClientId:   r.c.Request.FormValue("_g.sci"),
	}

	state := errstate.Success

	id := r.services.id.String()

	traceId, err := encrypt.AESEncrypt([]byte(id), []byte(serviceconfiguration.Get().CurrencySecretKey))
	if err != nil {

		return false
	}

	resp := func(s *commons.State) {
		if state.Ok {
			state = s
		}
	}

	headers := &analysisclient.LogHeaders{
		ServiceName:      constants.GatewayService,
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

	r.logId = r.services.id.String()
	r.logIndex = "xbs-logger." + xbasisdate.FormatDate(time.Now(), xbasisdate.YYYY_I_MM_I_DD)

	r.services.log.Info(&analysisclient.LogContent{
		Headers:   headers,
		Action:    "PermissionVerification",
		Message:   "Start verification",
		StateCode: 0,
		LogIndex:  r.logIndex,
		Id:        r.logId,
	})

	r.c.Request.Header.Del("authorization")

	var wg sync.WaitGroup

	wg.Add(3)

	ctx := metadata.NewContext(context.Background(), map[string]string{
		"transport-trace-id":  traceId,
		"transport-log-id":    r.logId,
		"transport-log-index": r.logIndex,
	})

	r.ctx = ctx

	callClientId := rh.fromClientId
	if len(rh.refClientId) > 0 {
		callClientId = rh.refClientId
	}

	//blacklist(ip)
	go func() {
		defer wg.Done()
		s, err := r.services.blacklistService.Check(ctx,
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
		s, err := r.services.blacklistService.Check(ctx,
			&xbasissvc_external_safety.CheckRequest{
				Type:    constants.BlacklistOfDevice,
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
		s, err := r.services.innerApplicationStatusService.GetAppClientStatus(ctx, &xbasissvc_internal_application.GetAppClientStatusRequest{
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
		r.services.log.Info(&analysisclient.LogContent{
			Headers: &analysisclient.LogHeaders{
				ServiceName: constants.GatewayService,
				ModuleName:  "Verification",
				TraceId:     traceId,
			},
			Action:    "ThreeBasicValidations",
			Message:   "CheckFailed",
			StateCode: state.Code,
			Index: &analysisclient.LogIndex{
				Id:   r.logId,
				Name: r.logIndex,
				Fields: &analysisclient.LogFields{
					"basic_validation": true,
				},
			},
		})
		r.json(state)
		return false
	}

	if appResp != nil && len(appResp.AppId) > 0 {

		r.services.log.Info(&analysisclient.LogContent{
			Headers: &analysisclient.LogHeaders{
				TraceId:     traceId,
				ServiceName: constants.GatewayService,
				ModuleName:  "Verification",
			},
			Action:  "ThreeBasicValidations",
			Message: "Check blacklist(ip, userDevice) and clientId to passed",
			Index: &analysisclient.LogIndex{
				Id:   r.logId,
				Name: r.logIndex,
				Fields: &analysisclient.LogFields{
					"app_id": appResp.AppId,
				},
			},
		})

		if appResp.ClientEnabled != constants.Enabled {
			r.json(errstate.ErrClientClosed)
			return false
		}

		var f *xbasistransport.AppFunction

		wl := false
		for _, v := range whiteApiList {
			if v == rh.path {
				f = &xbasistransport.AppFunction{
					AuthTypes: []int64{},
				}
				wl = true
				break
			}
		}
		if f == nil && !wl {

			//in cache
			f = r.services.functions.find(appResp.AppId, rh.path)

			fmt.Println("path", rh.path)

			if f == nil {

				//from database
				f1, err := r.services.accessibleService.LookupApi(ctx, &xbasissvc_internal_permission.LookupApiRequest{
					AppId: appResp.AppId,
					Path:  rh.path,
				})

				if err != nil {
					r.services.log.Info(&analysisclient.LogContent{
						Headers: &analysisclient.LogHeaders{
							TraceId:     traceId,
							ServiceName: constants.GatewayService,
							ModuleName:  "Verification",
						},
						Action:  "InvalidApi",
						Message: fmt.Sprintf("Not found api %s in %s", rh.path, appResp.AppId),
						Index: &analysisclient.LogIndex{
							Id:   r.logId,
							Name: r.logIndex,
							Fields: &analysisclient.LogFields{
								"invalid_api": true,
							},
						},
					})
					r.json(errstate.ErrRequest)
					return false
				}

				if !f1.State.Ok {
					r.json(f1.State)
					return false
				}

				af := &xbasistransport.AppFunction{
					Id:               f1.Id,
					Name:             f1.Name,
					AuthTypes:        f1.AuthTypes,
					ValTokenTimes:    f1.ValTokenTimes,
					NoGrantPlatforms: f1.NoGrantPlatforms,
					Share:            f1.Share,
					AppId:            f1.AppId,
					Path:             f1.Path,
					Version:          1,
				}

				r.services.functions.update(af)

				f = af
			} else {
				r.services.log.Info(&analysisclient.LogContent{
					Headers: &analysisclient.LogHeaders{
						TraceId:     traceId,
						ServiceName: constants.GatewayService,
						ModuleName:  "Verification",
					},
					Action:  "LookupApiFromCache",
					Message: "Found application function",
					Fields: &analysisclient.LogFields{
						"function_id":   f.Id,
						"function_name": f.Name,
						"auth_types":    f.AuthTypes,
						"app_id":        f.AppId,
					},
					Index: &analysisclient.LogIndex{
						Id:   r.logId,
						Name: r.logIndex,
						Fields: &analysisclient.LogFields{
							"function": f.Name,
						},
					},
				})
			}

			//grant platform check
			if f.NoGrantPlatforms != nil && len(f.NoGrantPlatforms) > 0 {
				for _, v := range f.NoGrantPlatforms {
					if v == appResp.ClientPlatform {
						r.services.log.Info(&analysisclient.LogContent{
							Headers: headers,
							Action:  "ApiPlatformAccessDenied",
							Index: &analysisclient.LogIndex{
								Id:   r.logId,
								Name: r.logIndex,
								Fields: &analysisclient.LogFields{
									"denied_api_client": true,
								},
							},
						})
						r.json(errstate.ErrRequest)
						return false
					}
				}
			}
		}

		userId := ""
		cv := ""
		wg.Add(len(f.AuthTypes))

		var credential *permissionhandlers.DurationAccessCredential
		var dat *xbasissvc_internal_permission.FunctionDat

		//清除凭证(可能会遇到其他认证不通过的情况，发送验证码的前提是: 其余的验证条件全部满足)
		cclear := false

		auth := false

		cm := make(map[string]string)

		for _, v := range f.AuthTypes {
			go func() {
				defer wg.Done()
				switch v {
				case constants.AuthTypeOfValcode:
					if len(rh.dat) == 0 {
						//生成生成验证码凭证
						//(先请求原API，生成凭证，然后调用dat获取验证码，再请求API)
						resp(errstate.Success)

						credential = &permissionhandlers.DurationAccessCredential{
							FromClientId: rh.fromClientId,
							RefClientId:  rh.refClientId,
							FuncId:       f.Id,
							AppId:        f.AppId,
							Timestamp:    time.Now().UnixNano(),
						}

						return
					}

					v, err := encrypt.AESDecrypt(rh.dat, []byte(serviceconfiguration.Get().CurrencySecretKey))
					if err != nil {
						resp(errstate.ErrNotFoundDurationAccessToken)
						return
					}

					cv = v

					d, err := r.services.accessibleService.GetDat(ctx, &xbasissvc_internal_permission.GetDatRequest{
						FuncId: f.Id,
						Key:    v,
					})

					if err != nil {
						resp(errstate.ErrRequest)
						return
					}

					if !d.State.Ok {
						resp(d.State)
						return
					}

					dat = d.Data

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

					status, err := r.services.innerAuthService.Verify(ac, &xbasissvc_internal_authentication.VerifyRequest{
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

						cm["transport-token-user-id"] = status.UserId
						cm["transport-token-app-id"] = status.AppId
						cm["transport-token-client-platform"] = fmt.Sprintf("%d", status.ClientPlatform)
						cm["transport-token-client-id"] = status.ClientId
						cm["transport-token-user-relation"] = status.Relation
						cm["transport-token-app-type"] = fmt.Sprintf("%d", status.AppType)

						userId = status.UserId

					} else {
						r.services.log.Info(&analysisclient.LogContent{
							Headers: &analysisclient.LogHeaders{
								TraceId:     traceId,
								ServiceName: constants.GatewayService,
								ModuleName:  "Verification",
							},
							Action:  "ExtAuthVerify",
							Message: "VerificationFailed",
						})
					}

					cclear = true
					resp(status.State)

					break
				case constants.AuthTypeOfFace: //需要刷脸认证时，从body中取出对应信息

					if r.c.Request.Method != "POST" {
						resp(errstate.ErrRequest)
						return
					}

					b, err := ioutil.ReadAll(r.c.Request.Body)
					if err != nil {
						return
					}
					var face map[string]interface{}
					err = json.Unmarshal(b, &face)
					if err != nil {
						return
					}

					v, ok := face[fmt.Sprintf("_xbs_face_auth")]
					if !ok {
						return
					}

					if reflect.TypeOf(v).Kind() != reflect.String {
						return
					}

					break
				case constants.AuthTypeOfMobileConfirm:
					break
				}
			}()
		}

		wg.Wait()

		if !state.Ok {
			r.services.log.Info(&analysisclient.LogContent{
				Headers: &analysisclient.LogHeaders{
					TraceId:     traceId,
					ServiceName: constants.GatewayService,
				},
				Action:    "BasicApplicationInfoCheck",
				Message:   "CheckFailed",
				StateCode: state.Code,
				Index: &analysisclient.LogIndex{
					Id:   r.logId,
					Name: r.logIndex,
					Fields: &analysisclient.LogFields{
						"passed": false,
					},
				},
			})
			r.json(state)
			return false
		}

		if dat != nil {

			r.dat = dat

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
				r.json(errstate.ErrSystem)
				return false
			}

			c, err := encrypt.AESEncrypt(b, []byte(serviceconfiguration.Get().CurrencySecretKey))
			if err != nil {
				r.json(errstate.ErrSystem)
				return false
			}

			state = errstate.ErrDurationAccessCredential
			state.Credential = c

			r.json(state)

			return true
		}

		if len(userId) == 0 {
			userId = rh.ip
		}
		var v1 int64
		if dat != nil {
			if dat.Auth {
				v1 = 1
			} else {
				v1 = 2
			}
		}

		r.services.log.Info(&analysisclient.LogContent{
			Headers: &analysisclient.LogHeaders{
				TraceId:     traceId,
				UserId:      userId,
				ServiceName: constants.GatewayService,
			},
			Action:    loggeractions.UserRequestApi,
			Message:   "verification passed",
			StateCode: 0,
			Fields: &analysisclient.LogFields{
				"func_id":   f.Id,
				"user_id":   userId,
				"client_id": callClientId,
				"app_id":    appResp.AppId,
				"platform":  appResp.ClientPlatform,
			},
			Index: &analysisclient.LogIndex{
				Id:   r.logId,
				Name: r.logIndex,
				Fields: &analysisclient.LogFields{
					"passed": true,
				},
			},
		})

		r.services._log.WithFields(logrus.Fields{
			"taking": fmt.Sprintf("%dms", (time.Now().UnixNano()-r.startAt)/1e6),
		}).Info("auth time consuming")

		cm["transport-log-id"] = r.logId
		cm["transport-log-index"] = r.logIndex
		cm["transport-user"] = userId
		cm["transport-app-id"] = appResp.AppId
		cm["transport-from-client-id"] = rh.fromClientId
		cm["transport-ref-client-id"] = rh.refClientId
		cm["transport-trace-id"] = traceId
		cm["transport-ip"] = rh.ip
		cm["transport-function-id"] = f.Id
		cm["transport-user-device"] = rh.userDevice
		cm["transport-user-agent"] = rh.userAgent
		cm["transport-app-type"] = fmt.Sprintf("%d", appResp.Type)
		cm["transport-client-platform"] = fmt.Sprintf("%d", appResp.ClientPlatform)
		if dat != nil {
			cm["transport-duration-access-to"] = dat.User
			cm["transport-duration-access-auth"] = fmt.Sprintf("%d", v1)
			r.dat = dat
			r.dat.Cv = cv
		}

		r.funcAppId = f.AppId
		r.funcName = f.Name
		r.traceId = traceId
		r.headers = cm
		r.toAppId = appResp.AppId
		r.userId = userId
		r.rh = rh
		r.funcId = f.Id
		r.auth = auth
		r.cv = cv
		r.serviceName = appResp.ServiceName

	}

	return true
}
