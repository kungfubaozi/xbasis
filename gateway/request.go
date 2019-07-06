package main

import (
	"context"
	"github.com/Sirupsen/logrus"
	"github.com/gin-gonic/gin"
	"konekko.me/xbasis/analysis/client"
	"konekko.me/xbasis/application/pb/inner"
	"konekko.me/xbasis/authentication/pb/inner"
	"konekko.me/xbasis/commons/dto"
	"konekko.me/xbasis/commons/generator"
	"konekko.me/xbasis/permission/pb/inner"
	"konekko.me/xbasis/safety/pb"
	"net/http"
)

type request struct {
	c           *gin.Context
	services    *services
	toAppId     string
	headers     map[string]string
	path        string
	startAt     int64
	ctx         context.Context
	dat         *xbasissvc_internal_permission.FunctionDat
	userId      string
	rh          *requestHeaders
	auth        bool
	funcId      string
	cv          string
	host        string
	serviceName string
	secure      bool
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
