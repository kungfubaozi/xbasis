package applicationservice

import (
	"konekko.me/xbasis/analysis/client"
	"konekko.me/xbasis/application/handerls"
	"konekko.me/xbasis/application/pb"
	"konekko.me/xbasis/application/pb/inner"
	config "konekko.me/xbasis/commons/config"
	"konekko.me/xbasis/commons/config/call"
	constants "konekko.me/xbasis/commons/constants"
	"konekko.me/xbasis/commons/dao"
	"konekko.me/xbasis/commons/indexutils"
	"konekko.me/xbasis/commons/microservice"
	"konekko.me/xbasis/permission/client"
	"konekko.me/xbasis/user/client"
)

func StartService() {

	errc := make(chan error, 2)

	session, err := xbasisdao.CreateSession("192.168.2.60:27017")
	if err != nil {
		panic(err)
	}
	defer session.Close()

	pool, err := xbasisdao.CreatePool("192.168.2.60:6379")
	if err != nil {
		panic(err)
	}
	defer pool.Close()

	c, err := indexutils.NewClient("http://192.168.2.62:9200/")
	if err != nil {
		panic(err)
	}

	logger := analysisclient.NewLoggerClient()

	go func() {
		config.WatchInitializeConfig(constants.ApplicationService,
			applicationhanderls.Initialize(session.Clone(), c))
	}()

	go func() {
		config.WatchGosionConfig(serviceconfiguration.Configuration())
	}()

	go func() {
		s := microservice.NewService(constants.InternalApplicationService, true)
		s.Init()

		xbasissvc_internal_application.RegisterApplicationStatusHandler(s.Server(), applicationhanderls.NewApplicationStatusService(c, pool, logger))

		xbasissvc_internal_application.RegisterUserSyncHandler(s.Server(),
			applicationhanderls.NewSyncService(session, userclient.NewInviteClient(s.Client()), permissioncli.NewAccessibleClient(s.Client()), permissioncli.NewBindingClient(s.Client()),
				permissioncli.NewGroupClient(s.Client()), pool))

		errc <- s.Run()
	}()

	go func() {
		s := microservice.NewService(constants.ApplicationService, true)
		s.Init()

		xbasissvc_external_application.RegisterApplicationHandler(s.Server(), applicationhanderls.NewApplicationService(session, c, logger, pool))

		xbasissvc_external_application.RegisterSettingsHandler(s.Server(), applicationhanderls.NewSettingsService(session, c))

		errc <- s.Run()
	}()

	if err := <-errc; err != nil {
		panic(err)
	}
}
