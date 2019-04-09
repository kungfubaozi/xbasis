package safetyservice

import (
	"konekko.me/gosion/commons/config"
	"konekko.me/gosion/commons/constants"
	"konekko.me/gosion/commons/dao"
	"konekko.me/gosion/commons/microservice"
	"konekko.me/gosion/safety/handers"
	"konekko.me/gosion/safety/pb"
	"konekko.me/gosion/safety/pb/ext"
)

func StartService() {

	errc := make(chan error, 2)

	pool, err := gs_commons_dao.CreatePool("192.168.2.60:6379")
	if err != nil {
		panic(err)
	}
	defer pool.Close()

	session, err := gs_commons_dao.CreateSession("192.168.2.60:27017")
	if err != nil {
		panic(err)
	}
	defer session.Close()

	configuration := &gs_commons_config.GosionConfiguration{}

	go func() {
		s := microservice.NewService(gs_commons_constants.SafetyService)
		s.Init()

		gs_service_safety.RegisterBlacklistHandler(s.Server(), safetyhanders.NewBlacklistService(session, pool))

		gs_service_safety.RegisterFrozenHandler(s.Server(), safetyhanders.NewFrozenService())

		gs_service_safety.RegisterLockingHandler(s.Server(), safetyhanders.NewLockingService())

		errc <- s.Run()
	}()

	go func() {
		gs_commons_config.WatchGosionConfig(func(config *gs_commons_config.GosionConfiguration) {
			configuration = config
		})
	}()

	go func() {
		s := microservice.NewService(gs_commons_constants.ExtSafetyService)
		s.Init()
		gs_ext_service_safety.RegisterSecurityHandler(s.Server(), safetyhanders.NewSecurityService(session))

		errc <- s.Run()
	}()

	if err := <-errc; err != nil {
		panic(err)
	}

}
