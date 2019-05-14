package applicationservice

import (
	"konekko.me/gosion/application/handerls"
	"konekko.me/gosion/application/pb"
	"konekko.me/gosion/application/pb/ext"
	"konekko.me/gosion/commons/config"
	"konekko.me/gosion/commons/config/call"
	"konekko.me/gosion/commons/constants"
	"konekko.me/gosion/commons/dao"
	"konekko.me/gosion/commons/gslogrus"
	"konekko.me/gosion/commons/indexutils"
	"konekko.me/gosion/commons/microservice"
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

	client, err := indexutils.NewClient("http://192.168.2.62:9200/")
	if err != nil {
		panic(err)
	}

	go func() {
		gs_commons_config.WatchInitializeConfig(gs_commons_constants.ApplicationService,
			applicationhanderls.Initialize(session.Clone(), client))
	}()

	go func() {
		gs_commons_config.WatchGosionConfig(serviceconfiguration.Configuration())
	}()

	go func() {
		s := microservice.NewService(gs_commons_constants.ExtApplicationService, true)
		s.Init()

		log := gslogrus.New(gs_commons_constants.ExtApplicationService, client)

		gs_ext_service_application.RegisterApplicationStatusHandler(s.Server(), applicationhanderls.NewApplicationStatusService(client, pool, log))

		gs_ext_service_application.RegisterUsersyncHandler(s.Server(), applicationhanderls.NewSyncService(client, session))

		errc <- s.Run()
	}()

	go func() {
		s := microservice.NewService(gs_commons_constants.ApplicationService, true)
		s.Init()

		gs_service_application.RegisterApplicationHandler(s.Server(), applicationhanderls.NewApplicationService(session, client))

		gs_service_application.RegisterSettingsHandler(s.Server(), applicationhanderls.NewSettingsService(session, client))

		errc <- s.Run()
	}()

	if err := <-errc; err != nil {
		panic(err)
	}
}
