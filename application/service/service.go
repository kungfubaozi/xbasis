package applicationservice

import (
	"konekko.me/gosion/analysis/client"
	"konekko.me/gosion/application/handerls"
	"konekko.me/gosion/application/pb"
	"konekko.me/gosion/application/pb/inner"
	"konekko.me/gosion/commons/config"
	"konekko.me/gosion/commons/config/call"
	"konekko.me/gosion/commons/constants"
	"konekko.me/gosion/commons/dao"
	"konekko.me/gosion/commons/indexutils"
	"konekko.me/gosion/commons/microservice"
	"konekko.me/gosion/permission/client"
	"konekko.me/gosion/user/client"
)

func StartService() {

	errc := make(chan error, 2)

	session, err := gs_commons_dao.CreateSession("192.168.2.60:27017")
	if err != nil {
		panic(err)
	}
	defer session.Close()

	pool, err := gs_commons_dao.CreatePool("192.168.2.60:6379")
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
		gs_commons_config.WatchInitializeConfig(gs_commons_constants.ApplicationService,
			applicationhanderls.Initialize(session.Clone(), c))
	}()

	go func() {
		gs_commons_config.WatchGosionConfig(serviceconfiguration.Configuration())
	}()

	go func() {
		s := microservice.NewService(gs_commons_constants.InternalApplicationService, true)
		s.Init()

		gosionsvc_internal_application.RegisterApplicationStatusHandler(s.Server(), applicationhanderls.NewApplicationStatusService(c, pool, logger))

		gosionsvc_internal_application.RegisterUserSyncHandler(s.Server(),
			applicationhanderls.NewSyncService(c, session, userclient.NewInviteClient(s.Client()), permissioncli.NewAccessibleClient(s.Client()), permissioncli.NewBindingClient(s.Client()),
				permissioncli.NewGroupClient(s.Client())))

		errc <- s.Run()
	}()

	go func() {
		s := microservice.NewService(gs_commons_constants.ApplicationService, true)
		s.Init()

		gosionsvc_external_application.RegisterApplicationHandler(s.Server(), applicationhanderls.NewApplicationService(session, c, logger))

		gosionsvc_external_application.RegisterSettingsHandler(s.Server(), applicationhanderls.NewSettingsService(session, c))

		errc <- s.Run()
	}()

	if err := <-errc; err != nil {
		panic(err)
	}
}
