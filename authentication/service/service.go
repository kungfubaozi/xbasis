package authenticationsvc

import (
	"konekko.me/gosion/analysis/client"
	"konekko.me/gosion/application/client"
	"konekko.me/gosion/authentication/client"
	"konekko.me/gosion/authentication/handlers"
	"konekko.me/gosion/authentication/pb"
	"konekko.me/gosion/authentication/pb/inner"
	"konekko.me/gosion/commons/config"
	"konekko.me/gosion/commons/config/call"
	"konekko.me/gosion/commons/constants"
	"konekko.me/gosion/commons/dao"
	"konekko.me/gosion/commons/indexutils"
	"konekko.me/gosion/commons/microservice"
	"konekko.me/gosion/connection/cmd/connectioncli"
	"konekko.me/gosion/permission/client"
	"konekko.me/gosion/safety/client"
	"konekko.me/gosion/user/client"
)

func StartService() {

	errc := make(chan error, 2)

	pool, err := gs_commons_dao.CreatePool("192.168.2.60:6379")
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
		gs_commons_config.WatchGosionConfig(serviceconfiguration.Configuration())
	}()

	go func() {
		s := microservice.NewService(gs_commons_constants.InternalAuthenticationService, true)
		s.Init()

		gosionsvc_internal_authentication.RegisterAuthHandler(s.Server(),
			authenticationhandlers.NewAuthService(pool, safetyclient.NewSecurityClient(s.Client()), conn, client,
				applicationclient.NewStatusClient(s.Client()),
				safetyclient.NewBlacklistClient(s.Client()), permissioncli.NewAccessibleClient(s.Client()), logger))

		gosionsvc_internal_authentication.RegisterTokenHandler(s.Server(), authenticationhandlers.NewTokenService(pool, conn, logger))

		errc <- s.Run()

	}()

	go func() {
		s := microservice.NewService(gs_commons_constants.AuthenticationService, true)
		s.Init()

		gosionsvc_external_authentication.RegisterRouterHandler(s.Server(), authenticationhandlers.NewRouteService(client, pool, applicationclient.NewStatusClient(s.Client()),
			applicationclient.NewSyncClient(s.Client()), authenticationcli.NewTokenClient(s.Client()), conn, userclient.NewExtUserClient(s.Client()), permissioncli.NewAccessibleClient(s.Client())))

		errc <- s.Run()
	}()

	if err := <-errc; err != nil {
		panic(err)
	}

}
