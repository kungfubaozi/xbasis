package xbasisgateway

import (
	"fmt"
	"github.com/Sirupsen/logrus"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"konekko.me/xbasis/analysis/client"
	"konekko.me/xbasis/application/client"
	"konekko.me/xbasis/authentication/client"
	"konekko.me/xbasis/commons/config"
	"konekko.me/xbasis/commons/config/call"
	constants "konekko.me/xbasis/commons/constants"
	"konekko.me/xbasis/commons/generator"
	"konekko.me/xbasis/commons/microservice"
	"konekko.me/xbasis/permission/client"
	"konekko.me/xbasis/safety/client"
	"time"
)

func StartService() {
	gin.SetMode(gin.ReleaseMode)
	g := gin.New()

	g.Use(cors.New(cors.Config{
		AllowOriginFunc:  func(origin string) bool { return true },
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "PATCH"},
		AllowHeaders:     []string{"Origin", "Content-Length", "Content-Type", "authorization", "x-real-ip", "xbs-client-id", "xbs-user-device"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	svc := microservice.NewService(constants.GatewayService, false)

	f := &functions{
		funcs: make(map[string]AppFunctions),
	}

	gate := &gateway{
		s: &services{
			verification:                  permissioncli.NewVerificationClient(svc.Client()),
			accessibleService:             permissioncli.NewAccessibleClient(svc.Client()),
			innerApplicationStatusService: applicationclient.NewStatusClient(svc.Client()),
			blacklistService:              safetyclient.NewBlacklistClient(svc.Client()),
			innerAuthService:              authenticationcli.NewAuthClient(svc.Client()),
			id:                            xbasisgenerator.NewIDG(),
			log:                           analysisclient.NewLoggerClient(),
			_log:                          logrus.New(),
			functions:                     f,
		},
		service: svc,
	}

	g.Any("/*action", func(c *gin.Context) {

		gate.request(c)

	})

	errc := make(chan error, 2)

	go func() {
		errc <- f.listen("192.168.2.62:9092")
	}()

	go func() {
		xbasisconfig.WatchGosionConfig(serviceconfiguration.Configuration())
	}()

	go func() {
		errc <- gate.run()
	}()

	go func() {
		errc <- g.Run(":9081")
	}()

	fmt.Println("gateway service", <-errc)
}
