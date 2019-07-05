package main

import (
	"context"
	"fmt"
	"github.com/Sirupsen/logrus"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/micro/go-micro/metadata"
	"github.com/vmihailenco/msgpack"
	"io/ioutil"
	"konekko.me/xbasis/analysis/client"
	"konekko.me/xbasis/application/client"
	"konekko.me/xbasis/application/pb/inner"
	"konekko.me/xbasis/authentication/client"
	"konekko.me/xbasis/authentication/pb/inner"
	"konekko.me/xbasis/commons/actions"
	"konekko.me/xbasis/commons/config"
	"konekko.me/xbasis/commons/config/call"
	constants "konekko.me/xbasis/commons/constants"
	"konekko.me/xbasis/commons/dto"
	commons "konekko.me/xbasis/commons/dto"
	"konekko.me/xbasis/commons/encrypt"
	"konekko.me/xbasis/commons/errstate"
	"konekko.me/xbasis/commons/generator"
	"konekko.me/xbasis/commons/microservice"
	"konekko.me/xbasis/permission/client"
	"konekko.me/xbasis/permission/handlers"
	"konekko.me/xbasis/permission/pb/inner"
	"konekko.me/xbasis/safety/client"
	"konekko.me/xbasis/safety/pb"
	"net/http"
	"sync"
	"time"
)

type request struct {
	c        *gin.Context
	services *services
	toAppId  string
	headers  map[string]string
	path     string
	startAt  int64
	ctx      context.Context
	dat      *xbasissvc_internal_permission.FunctionDat
	userId   string
	rh       *requestHeaders
	auth     bool
	funcId   string
	cv       string
	host     string
}

type services struct {
	verification                  xbasissvc_internal_permission.VerificationService
	innerApplicationStatusService xbasissvc_internal_application.ApplicationStatusService
	blacklistService              xbasissvc_external_safety.BlacklistService
	innerAuthService              xbasissvc_internal_authentication.AuthService
	accessibleService             xbasissvc_internal_permission.AccessibleService
	id                            xbasisgenerator.IDGenerator
	log                           analysisclient.LogClient
	_log                          *logrus.Logger
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

var apps map[string]string

var whiteApiList = []string{"/authentication/router/refresh",
	"/authentication/router/logout", "/apps/settings/getSetting", "/permission/userGroup/move"}

func main() {
	//gin.SetMode(gin.ReleaseMode)
	g := gin.New()

	g.Use(cors.New(cors.Config{
		AllowOriginFunc:  func(origin string) bool { return true },
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "PATCH"},
		AllowHeaders:     []string{"Origin", "Content-Length", "Content-Type", "authorization", "x-real-ip", "xbs-client-id", "xbs-user-device"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	svc := microservice.NewService(constants.GatewayService, false)

	s := &services{
		verification:                  permissioncli.NewVerificationClient(svc.Client()),
		accessibleService:             permissioncli.NewAccessibleClient(svc.Client()),
		innerApplicationStatusService: applicationclient.NewStatusClient(svc.Client()),
		blacklistService:              safetyclient.NewBlacklistClient(svc.Client()),
		innerAuthService:              authenticationcli.NewAuthClient(svc.Client()),
		id:                            xbasisgenerator.NewIDG(),
		log:                           analysisclient.NewLoggerClient(),
		_log:                          logrus.New(),
	}

	apps = map[string]string{
		"51334e445530": "http://192.168.80.67:8080", //router
		"4d324f574e6d": "http://192.168.80.67:8080", //admin
	}

	g.Any("/*action", func(c *gin.Context) {

		r := &request{
			c:        c,
			services: s,
		}

		r.services._log.WithFields(logrus.Fields{
			"requestURI": c.Request.RequestURI,
			"remoteAddr": c.ClientIP(),
		}).Warn("new request")

		//check path
		if c.Request.Method == "OPTIONS" {
			c.Status(200)
			fmt.Println("options")
			return
		}

		if !checkRequestBasicInfo(r) {
			return
		}

		if !checkRequestPath(r) {
			return
		}

		switch c.Request.Method {
		case "GET":
			get(r)
			return
		case "POST":
			post(r)
			return
		case "PUT":
			put(r)
			return
		case "PATCH":
			patch(r)
			return
		case "HEAD":
			return
		case "OPTIONS":
			options(r)
			return
		case "DELETE":
			delete(r)
			return
		case "TRACE":
			return
		case "CONNECT":
			return
		}
	})

	errc := make(chan error, 2)

	go func() {
		xbasisconfig.WatchGosionConfig(serviceconfiguration.Configuration())
	}()

	go func() {
		errc <- svc.Run()
	}()

	go func() {
		errc <- g.Run(":9081")
	}()

	fmt.Println("gateway service", <-errc)
}

func abort(r *request, code int) {
	r.c.AbortWithStatus(code)
}

func json(r *request, state *xbasis_commons_dto.State) {
	r.c.JSON(200, xbasis_commons_dto.Status{State: state})
}

func buildHeader(request *http.Request, r *request) {
	for k, v := range r.headers {
		request.Header.Set(k, v)
	}
}

func checkRequestPath(r *request) bool {
	//path := r.c.Request.URL.Path
	url := apps[r.toAppId]
	if len(url) <= 10 {
		json(r, errstate.ErrRequest)
		return false
	}
	r.path = url
	return true
}

func checkRequestBasicInfo(r *request) bool {

	r.startAt = time.Now().UnixNano()

	clientId := r.c.Request.Header.Get("xbs-client-id")
	if len(clientId) != 12 {
		json(r, errstate.ErrInvalidClientId)
		return false
	}
	r.c.Request.Header.Del("xbs-client-id")
	ip := r.c.ClientIP()
	if len(ip) < 8 {
		json(r, errstate.ErrRequest)
		return false
	}
	userDevice := r.c.Request.Header.Get("xbs-user-device")
	if len(userDevice) < 8 {
		json(r, errstate.ErrRequest)
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
		json(r, state)
		return false
	}

	if appResp != nil && len(appResp.AppId) > 0 {

		if appResp.ClientEnabled != constants.Enabled {
			json(r, errstate.ErrClientClosed)
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
				json(r, errstate.ErrRequest)
				return false
			}

			if !f1.State.Ok {
				json(r, f1.State)
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
						json(r, errstate.ErrRequest)
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
			json(r, state)
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
				json(r, errstate.ErrSystem)
				return false
			}

			c, err := encrypt.AESEncrypt(b, []byte(serviceconfiguration.Get().CurrencySecretKey))
			if err != nil {
				json(r, errstate.ErrSystem)
				return false
			}

			state = errstate.ErrDurationAccessCredential
			state.Credential = c

			json(r, state)

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

	}

	return true
}

func do(r *request, req *http.Request) {
	buildHeader(req, r)

	req.Header.Set("Content-Type", "application/json")

	r.services._log.WithFields(logrus.Fields{
		"routeTo": r.path,
	}).Info("request redirect")

	rt := time.Now().UnixNano()

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Println("err", err)
		r.c.JSON(200, errstate.ErrRequest)
		return
	}

	if r.dat != nil {
		from := encrypt.SHA1(r.rh.ip + r.rh.userAgent + r.rh.userDevice + r.rh.fromClientId)
		if r.dat.From != from {
			json(r, errstate.ErrDurationAccess)
			return
		}
		user := ""
		if r.auth {
			if len(r.userId) == 0 {
				json(r, errstate.ErrRequest)
				return
			}
			user = r.userId
		}
		if r.dat.FuncId != r.funcId || r.dat.ClientId != r.rh.fromClientId || (r.auth && user != r.dat.User) {
			json(r, errstate.ErrDurationAccess)
			return
		}

		s, err := r.services.accessibleService.DatReduce(r.ctx, r.dat)
		if err != nil {
			json(r, errstate.ErrRequest)
			return
		}
		if !s.State.Ok {
			json(r, s.State)
			return
		}
	}

	b, err := ioutil.ReadAll(resp.Body)
	r.c.Writer.Write(b)

	r.services._log.WithFields(logrus.Fields{
		"all":        fmt.Sprintf("%dms", (time.Now().UnixNano()-r.startAt)/1e6),
		"processing": fmt.Sprintf("%dms", (time.Now().UnixNano()-rt)/1e6),
	}).Info("all time consuming")
}

func get(r *request) {

	req, err := http.NewRequest("GET", fmt.Sprintf("%s%s", r.path, r.c.Request.RequestURI), nil)
	req.Form = r.c.Request.Form

	if err != nil {
		r.c.JSON(200, errstate.ErrRequest)
		return
	}

	do(r, req)
}

func post(r *request) {

	req, err := http.NewRequest("POST", fmt.Sprintf("%s%s", r.path, r.c.Request.RequestURI), r.c.Request.Body)

	req.Form = r.c.Request.Form
	req.MultipartForm = r.c.Request.MultipartForm
	req.PostForm = r.c.Request.PostForm

	if err != nil {
		r.c.JSON(200, errstate.ErrRequest)
		return
	}

	do(r, req)
}

func put(r *request) {

}

func patch(r *request) {

}

func options(r *request) {

}

func delete(r *request) {

}
