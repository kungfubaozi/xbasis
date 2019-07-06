package main

import (
	"fmt"
	"github.com/Sirupsen/logrus"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/micro/go-micro/registry"
	"github.com/micro/go-micro/registry/consul"
	"konekko.me/xbasis/analysis/client"
	"konekko.me/xbasis/application/client"
	"konekko.me/xbasis/authentication/client"
	"konekko.me/xbasis/commons/config"
	"konekko.me/xbasis/commons/config/call"
	constants "konekko.me/xbasis/commons/constants"
	"konekko.me/xbasis/commons/errstate"
	"konekko.me/xbasis/commons/generator"
	"konekko.me/xbasis/commons/microservice"
	"konekko.me/xbasis/permission/client"
	"konekko.me/xbasis/safety/client"
	"time"
)

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

	cr = consul.NewRegistry(registry.Addrs("192.168.80.67:8500"),
		registry.Secure(false))

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

	apps = make(map[string][]string)

	g.Any("/*action", func(c *gin.Context) {

		r := &request{
			c:        c,
			services: s,
		}

		r.services._log.WithFields(logrus.Fields{
			"requestURI": c.Request.RequestURI,
			"remoteAddr": c.ClientIP(),
		}).Warn("new request")

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
