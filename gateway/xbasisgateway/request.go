package xbasisgateway

import (
	"context"
	"github.com/Sirupsen/logrus"
	"github.com/gin-gonic/gin"
	"github.com/micro/go-micro"
	"konekko.me/xbasis/analysis/client"
	"konekko.me/xbasis/application/client"
	"konekko.me/xbasis/application/pb/inner"
	"konekko.me/xbasis/authentication/client"
	"konekko.me/xbasis/authentication/pb/inner"
	constants "konekko.me/xbasis/commons/constants"
	"konekko.me/xbasis/commons/dto"
	"konekko.me/xbasis/commons/errstate"
	"konekko.me/xbasis/commons/generator"
	"konekko.me/xbasis/commons/microservice"
	"konekko.me/xbasis/permission/client"
	"konekko.me/xbasis/permission/pb/inner"
	"konekko.me/xbasis/safety/client"
	"konekko.me/xbasis/safety/pb"
	"net/http"
)

type request struct {
	c             *gin.Context
	services      *services
	toAppId       string
	headers       map[string]string
	path          string
	startAt       int64
	ctx           context.Context
	dat           *xbasissvc_internal_permission.FunctionDat
	userId        string
	rh            *requestHeaders
	auth          bool
	funcId        string
	cv            string
	host          string
	serviceName   string
	secure        bool
	traceId       string
	logId         string
	logIndex      string
	funcName      string
	requestMethod string
	requestPath   string
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

func (r *request) buildHeader(request *http.Request) {
	for k, v := range r.headers {
		request.Header.Set(k, v)
	}
}
func (r *request) json(state *xbasis_commons_dto.State) {
	r.c.JSON(200, xbasis_commons_dto.Status{State: state})
}

type Gateway struct {
	s       *services
	service micro.Service
}

func InitService() *Gateway {
	svc := microservice.NewService(constants.GatewayService, false)

	g := &Gateway{
		s: &services{
			verification:                  permissioncli.NewVerificationClient(svc.Client()),
			accessibleService:             permissioncli.NewAccessibleClient(svc.Client()),
			innerApplicationStatusService: applicationclient.NewStatusClient(svc.Client()),
			blacklistService:              safetyclient.NewBlacklistClient(svc.Client()),
			innerAuthService:              authenticationcli.NewAuthClient(svc.Client()),
			id:                            xbasisgenerator.NewIDG(),
			log:                           analysisclient.NewLoggerClient(),
			_log:                          logrus.New(),
		},
		service: svc,
	}

	return g
}

func (g *Gateway) Run() error {
	return g.service.Run()
}

func (g *Gateway) NewRequest(c *gin.Context) {

	g.s._log.WithFields(logrus.Fields{
		"requestURI": c.Request.RequestURI,
		"remoteAddr": c.ClientIP(),
	}).Warn("new request")

	r := &request{
		c:        c,
		services: g.s,
	}

	if !r.verification() {
		return
	}

	if len(r.serviceName) == 0 {
		r.json(errstate.ErrInvalidApplicationServiceName)
		return
	}

	if !r.address() {
		r.json(errstate.ErrRequest)
		return
	}

	r.call(c.Request.Method)

}
