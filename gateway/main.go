package main

import (
	"fmt"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"konekko.me/xbasis/commons/config"
	"konekko.me/xbasis/commons/config/call"
	"konekko.me/xbasis/gateway/xbasisgateway"

	"time"
)

func main() {
	gin.SetMode(gin.ReleaseMode)
	g := gin.New()

	g.Use(cors.New(cors.Config{
		AllowOriginFunc:  func(origin string) bool { return true },
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "PATCH"},
		AllowHeaders:     []string{"Origin", "Content-Length", "Content-Type", "authorization", "x-real-ip", "xbs-client-id", "xbs-user-device"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	gate := xbasisgateway.InitService()

	g.Any("/*action", func(c *gin.Context) {

		gate.NewRequest(c)

	})

	errc := make(chan error, 2)

	go func() {
		xbasisconfig.WatchGosionConfig(serviceconfiguration.Configuration())
	}()

	go func() {
		errc <- gate.Run()
	}()

	go func() {
		errc <- g.Run(":9081")
	}()

	fmt.Println("gateway service", <-errc)
}
