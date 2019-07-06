package main

import (
	"context"
	"fmt"
	"github.com/Sirupsen/logrus"
	"github.com/micro/go-micro/metadata"
	"github.com/vmihailenco/msgpack"
	"konekko.me/xbasis/analysis/client"
	"konekko.me/xbasis/application/pb/inner"
	"konekko.me/xbasis/authentication/pb/inner"
	"konekko.me/xbasis/commons/actions"
	"konekko.me/xbasis/commons/config/call"
	constants "konekko.me/xbasis/commons/constants"
	commons "konekko.me/xbasis/commons/dto"
	"konekko.me/xbasis/commons/encrypt"
	"konekko.me/xbasis/commons/errstate"
	"konekko.me/xbasis/permission/handlers"
	"konekko.me/xbasis/permission/pb/inner"
	"konekko.me/xbasis/safety/pb"
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

var whiteApiList = []string{"/authentication/router/refresh",
	"/authentication/router/logout", "/apps/settings/getSetting", "/permission/userGroup/move"}

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
		ModuleName:       "Route",
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

	r.services.log.Info(&analysisclient.LogContent{
		Headers:   headers,
		Action:    "PermissionVerification",
		Message:   "start verification",
		StateCode: 0,
	})

	r.c.Request.Header.Del("authorization")

	var wg sync.WaitGroup

	wg.Add(3)

	ctx := metadata.NewContext(context.Background(), map[string]string{
		"transport-trace-id": traceId,
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
				TraceId: traceId,
			},
			Action:    "ThreeBasicValidations",
			Message:   "CheckFailed",
			StateCode: state.Code,
		})
		r.json(state)
		return false
	}

	if appResp != nil && len(appResp.AppId) > 0 {

		if appResp.ClientEnabled != constants.Enabled {
			r.json(errstate.ErrClientClosed)
			return false
		}

		var f *xbasissvc_internal_permission.LookupApiResponse

		for _, v := range whiteApiList {
			if v == rh.path {
				f = &xbasissvc_internal_permission.LookupApiResponse{
					AuthTypes: []int64{},
				}
				break
			}
		}
		if f == nil {

			f1, err := r.services.accessibleService.LookupApi(ctx, &xbasissvc_internal_permission.LookupApiRequest{
				AppId: appResp.AppId,
				Path:  encrypt.SHA1(rh.path),
			})

			if err != nil {
				r.services.log.Info(&analysisclient.LogContent{
					Headers: &analysisclient.LogHeaders{
						TraceId: traceId,
					},
					Action: "InvalidApi",
				})
				r.json(errstate.ErrRequest)
				return false
			}

			if !f1.State.Ok {
				r.json(f1.State)
				return false
			}

			f = f1

			//grant platform
			if f.GrantPlatforms != nil && len(f.GrantPlatforms) > 0 {
				for _, v := range f.GrantPlatforms {
					if v == appResp.ClientPlatform {
						r.services.log.Info(&analysisclient.LogContent{
							Headers: &analysisclient.LogHeaders{
								TraceId: traceId,
							},
							Action: "ApiPlatformAccessDenied",
						})
						r.json(errstate.ErrRequest)
						return false
					}
				}
			}
		}

		r.services.log.Info(&analysisclient.LogContent{
			Headers: headers,
			Action:  "RequestApi",
			Message: headers.Path,
			Fields: &analysisclient.LogFields{
				"id":    f.Id,
				"appId": appResp.AppId,
			},
		})

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
			r.services.log.Info(&analysisclient.LogContent{
				Headers: &analysisclient.LogHeaders{
					TraceId: traceId,
				},
				Action:    "BasicApplicationInfoCheck",
				Message:   "CheckFailed",
				StateCode: state.Code,
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

		r.services._log.WithFields(logrus.Fields{
			"taking": fmt.Sprintf("%dms", (time.Now().UnixNano()-r.startAt)/1e6),
		}).Info("auth time consuming")

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
