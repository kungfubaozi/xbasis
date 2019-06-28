package authenticationsvc

import (
	"konekko.me/xbasis/analysis/client"
	"konekko.me/xbasis/application/client"
	"konekko.me/xbasis/authentication/client"
	"konekko.me/xbasis/authentication/handlers"
	"konekko.me/xbasis/authentication/pb"
	"konekko.me/xbasis/authentication/pb/inner"
	"konekko.me/xbasis/commons/config"
	"konekko.me/xbasis/commons/config/call"
	constants "konekko.me/xbasis/commons/constants"
	"konekko.me/xbasis/commons/dao"
	"konekko.me/xbasis/commons/indexutils"
	"konekko.me/xbasis/commons/microservice"
	"konekko.me/xbasis/connection/cmd/connectioncli"
	"konekko.me/xbasis/permission/client"
	"konekko.me/xbasis/safety/client"
	"konekko.me/xbasis/user/client"
)

func StartService() {

	errc := make(chan error, 2)

	pool, err := xbasisdao.CreatePool("192.168.2.60:6379")
	if err != nil {
		panic(err)
	}
	defer pool.Close()

	conn, err := connectioncli.NewClient()
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	client, err := indexutils.NewClient("http://192.168.2.62:9200/")
	if err != nil {
		panic(err)
	}

	logger := analysisclient.NewLoggerClient()

	go func() {
		xbasisconfig.WatchGosionConfig(serviceconfiguration.Configuration())
	}()

	go func() {
		s := microservice.NewService(constants.InternalAuthenticationService, true)
		s.Init()

		xbasissvc_internal_authentication.RegisterAuthHandler(s.Server(),
			authenticationhandlers.NewAuthService(pool, safetyclient.NewSecurityClient(s.Client()), conn, client,
				applicationclient.NewStatusClient(s.Client()),
				safetyclient.NewBlacklistClient(s.Client()), permissioncli.NewAccessibleClient(s.Client()), logger))

		xbasissvc_internal_authentication.RegisterTokenHandler(s.Server(), authenticationhandlers.NewTokenService(pool, conn, logger))

		errc <- s.Run()

	}()

	go func() {
		s := microservice.NewService(constants.AuthenticationService, true)
		s.Init()

		gosionsvc_external_authentication.RegisterRouterHandler(s.Server(), authenticationhandlers.NewRouteService(client, pool, applicationclient.NewStatusClient(s.Client()),
			applicationclient.NewSyncClient(s.Client()), authenticationcli.NewTokenClient(s.Client()), conn, userclient.NewExtUserClient(s.Client()), permissioncli.NewAccessibleClient(s.Client())))

		errc <- s.Run()
	}()

	if err := <-errc; err != nil {
		panic(err)
	}

}
