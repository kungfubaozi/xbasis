package permissionsvc

import (
	"konekko.me/gosion/application/client"
	"konekko.me/gosion/authentication/client"
	"konekko.me/gosion/commons/config"
	"konekko.me/gosion/commons/config/call"
	"konekko.me/gosion/commons/constants"
	"konekko.me/gosion/commons/dao"
	"konekko.me/gosion/commons/microservice"
	"konekko.me/gosion/permission/handlers"
	"konekko.me/gosion/permission/pb"
	"konekko.me/gosion/safety/client"
	"konekko.me/gosion/user/client"
)

func StartService() {

	errc := make(chan error, 2)

	configuration := &gs_commons_config.GosionConfiguration{}

	pool, err := gs_commons_dao.CreatePool("192.168.2.60:6379")
	if err != nil {
		panic(err)
	}
	defer pool.Close()

	session, err := gs_commons_dao.CreateSession("192.168.2.60:27017")
	if err != nil {
		panic(err)
	}

	go func() {
		m := microservice.NewService(gs_commons_constants.ExtPermissionVerification, false)

		gs_service_permission.RegisterVerificationHandler(m.Server(), permissionhandlers.NewVerificationService(pool,
			session, applicationclient.NewStatusClient(m.Client()),
			safetyclient.NewBlacklistClient(m.Client()),
			authenticationcli.NewAuthClient(m.Client())))

		errc <- m.Run()
	}()

	go func() {

		m := microservice.NewService(gs_commons_constants.PermissionService, true)

		us := userclient.NewExtUserClient(m.Client())

		mc := userclient.NewExtMessageClient(m.Client())

		gs_service_permission.RegisterBindingHandler(m.Server(), permissionhandlers.NewBindingService(pool, session, us))

		gs_service_permission.RegisterDurationAccessHandler(m.Server(), permissionhandlers.NewDurationAccessService(pool, session, configuration, mc))

		gs_service_permission.RegisterFunctionHandler(m.Server(), permissionhandlers.NewFunctionService(pool, session))

		gs_service_permission.RegisterGroupStructureHandler(m.Server(), permissionhandlers.NewGroupService(pool, session))

		gs_service_permission.RegisterRoleHandler(m.Server(), permissionhandlers.NewRoleService(session, pool))

		errc <- m.Run()

	}()

	go func() {
		gs_commons_config.WatchGosionConfig(serviceconfiguration.Configuration())
	}()

	go func() {
		gs_commons_config.WatchInitializeConfig(gs_commons_constants.PermissionService, permissionhandlers.Initialize(session.Clone(), pool))
	}()

	if err := <-errc; err != nil {
		panic(err)
	}

}
