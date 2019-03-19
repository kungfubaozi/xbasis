package applicationsvc

import (
	"konekko.me/gosion/application/handerls"
	"konekko.me/gosion/commons/config"
	"konekko.me/gosion/commons/constants"
	"konekko.me/gosion/commons/dao"
)

func StartService() {

	session, err := gs_commons_dao.CreateSession("192.168.2.60:27017")
	if err != nil {
		panic(err)
	}
	defer session.Close()

	errc := make(chan error, 1)

	go func() {
		gs_commons_config.WatchInitializeConfig(gs_commons_constants.PermissionService, application_handerls.Initialize(session.Clone()))
	}()

	<-errc
}
