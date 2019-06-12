package safetyservice

import (
	"konekko.me/gosion/commons/config"
	"konekko.me/gosion/commons/config/call"
	"konekko.me/gosion/commons/constants"
	"konekko.me/gosion/commons/dao"
	"konekko.me/gosion/commons/indexutils"
	"konekko.me/gosion/commons/microservice"
	"konekko.me/gosion/safety/handers"
	"konekko.me/gosion/safety/pb"
	"konekko.me/gosion/safety/pb/inner"
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

	client, err := indexutils.NewClient("http://192.168.2.62:9200/")
	if err != nil {
		panic(err)
	}

	go func() {
		s := microservice.NewService(gs_commons_constants.SafetyService, true)
		s.Init()

		gosionsvc_external_safety.RegisterBlacklistHandler(s.Server(), safetyhanders.NewBlacklistService(session, client))

		gosionsvc_external_safety.RegisterLockingHandler(s.Server(), safetyhanders.NewLockingService())

		gosionsvc_external_safety.RegisterUserlockHandler(s.Server(), safetyhanders.NewUserlockService())

		errc <- s.Run()
	}()

	go func() {
		gs_commons_config.WatchGosionConfig(serviceconfiguration.Configuration())
	}()

	go func() {
		s := microservice.NewService(gs_commons_constants.InternalSafetyService, true)
		s.Init()
		gosionsvc_internal_safety.RegisterSecurityHandler(s.Server(), safetyhanders.NewSecurityService(session))

		errc <- s.Run()
	}()

	if err := <-errc; err != nil {
		panic(err)
	}

}
