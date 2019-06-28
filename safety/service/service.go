package safetyservice

import (
	xconfig "konekko.me/xbasis/commons/config"
	"konekko.me/xbasis/commons/config/call"
	constants "konekko.me/xbasis/commons/constants"
	"konekko.me/xbasis/commons/dao"
	"konekko.me/xbasis/commons/indexutils"
	"konekko.me/xbasis/commons/microservice"
	"konekko.me/xbasis/safety/handers"
	"konekko.me/xbasis/safety/pb"
	"konekko.me/xbasis/safety/pb/inner"
)

func StartService() {

	errc := make(chan error, 2)

	pool, err := xbasisdao.CreatePool("192.168.2.60:6379")
	if err != nil {
		panic(err)
	}
	defer pool.Close()

	session, err := xbasisdao.CreateSession("192.168.2.60:27017")
	if err != nil {
		panic(err)
	}
	defer session.Close()

	client, err := indexutils.NewClient("http://192.168.2.62:9200/")
	if err != nil {
		panic(err)
	}

	go func() {
		s := microservice.NewService(constants.SafetyService, true)
		s.Init()

		xbasissvc_external_safety.RegisterBlacklistHandler(s.Server(), safetyhanders.NewBlacklistService(session, client))

		xbasissvc_external_safety.RegisterLockingHandler(s.Server(), safetyhanders.NewLockingService())

		xbasissvc_external_safety.RegisterUserlockHandler(s.Server(), safetyhanders.NewUserlockService())

		errc <- s.Run()
	}()

	go func() {
		xconfig.WatchGosionConfig(serviceconfiguration.Configuration())
	}()

	go func() {
		s := microservice.NewService(constants.InternalSafetyService, true)
		s.Init()
		xbasissvc_internal_safety.RegisterSecurityHandler(s.Server(), safetyhanders.NewSecurityService(session, pool))

		errc <- s.Run()
	}()

	if err := <-errc; err != nil {
		panic(err)
	}

}
