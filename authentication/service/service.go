package authenticationsvc

import (
	"konekko.me/gosion/authentication/handlers"
	"konekko.me/gosion/authentication/pb/ext"
	"konekko.me/gosion/commons/config"
	"konekko.me/gosion/commons/config/call"
	"konekko.me/gosion/commons/constants"
	"konekko.me/gosion/commons/dao"
	"konekko.me/gosion/commons/indexutils"
	"konekko.me/gosion/commons/microservice"
	"konekko.me/gosion/connection/cmd/connectioncli"
	"konekko.me/gosion/safety/client"
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

	go func() {
		gs_commons_config.WatchGosionConfig(serviceconfiguration.Configuration())
	}()

	go func() {
		s := microservice.NewService(gs_commons_constants.ExtAuthenticationService, true)
		s.Init()

		gs_ext_service_authentication.RegisterAuthHandler(s.Server(),
			authenticationhandlers.NewAuthService(pool, safetyclient.NewSecurityClient(s.Client()), conn, client))

		gs_ext_service_authentication.RegisterTokenHandler(s.Server(), authenticationhandlers.NewTokenService(pool, conn))

		errc <- s.Run()

	}()

	if err := <-errc; err != nil {
		panic(err)
	}

}
