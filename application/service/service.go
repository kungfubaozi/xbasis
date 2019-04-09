package applicationservice

import (
	"konekko.me/gosion/application/handerls"
	"konekko.me/gosion/application/pb"
	"konekko.me/gosion/application/pb/ext"
	"konekko.me/gosion/commons/config"
	"konekko.me/gosion/commons/constants"
	"konekko.me/gosion/commons/dao"
	"konekko.me/gosion/commons/microservice"
)

func StartService() {

	errc := make(chan error, 2)

	configuration := &gs_commons_config.GosionConfiguration{}

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

	go func() {
		gs_commons_config.WatchInitializeConfig(gs_commons_constants.ApplicationService,
			applicationhanderls.Initialize(session.Clone()))
	}()

	go func() {
		gs_commons_config.WatchGosionConfig(func(config *gs_commons_config.GosionConfiguration) {
			configuration = config
		})
	}()

	go func() {
		s := microservice.NewService(gs_commons_constants.ExtApplicationService)
		s.Init()

		gs_ext_service_application.RegisterApplicationStatusHandler(s.Server(), applicationhanderls.NewApplicationStatusServie(session, pool))

		errc <- s.Run()
	}()

	go func() {
		s := microservice.NewService(gs_commons_constants.ApplicationService)
		s.Init()

		gs_service_application.RegisterApplicationHandler(s.Server(), applicationhanderls.NewApplicationService(session, pool))

		gs_service_application.RegisterSettingsHandler(s.Server(), applicationhanderls.NewSettingsService(session, pool))

		errc <- s.Run()
	}()

	if err := <-errc; err != nil {
		panic(err)
	}
}
