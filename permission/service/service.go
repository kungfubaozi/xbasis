package permissionsvc

import (
	"konekko.me/gosion/application/client"
	"konekko.me/gosion/authentication/client"
	"konekko.me/gosion/commons/config"
	"konekko.me/gosion/commons/constants"
	"konekko.me/gosion/commons/dao"
	"konekko.me/gosion/commons/microservice"
	"konekko.me/gosion/permission/handlers"
	"konekko.me/gosion/permission/pb"
	"konekko.me/gosion/safety/client"
	"konekko.me/gosion/user/client"
)

func StartService() {

	m := microservice.NewService(gs_commons_constants.PermissionService)

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

	us := userclient.NewExtUserClient()

	mc := userclient.NewExtMessageClient()

	gs_service_permission.RegisterBindingHandler(m.Server(), permissionhandlers.NewBindingService(pool, session, us))

	gs_service_permission.RegisterDurationAccessHandler(m.Server(), permissionhandlers.NewDurationAccessService(pool, session, configuration, mc))

	gs_service_permission.RegisterFunctionHandler(m.Server(), permissionhandlers.NewFunctionService(pool, session))

	gs_service_permission.RegisterGroupStructureHandler(m.Server(), permissionhandlers.NewGroupService(pool, session))

	gs_service_permission.RegisterRoleHandler(m.Server(), permissionhandlers.NewRoleService(session, pool))

	gs_service_permission.RegisterVerificationHandler(m.Server(), permissionhandlers.NewVerificationService(pool,
		session, configuration, applicationclient.NewStatusClient(),
		safetyclient.NewBlacklistClient(),
		authenticationcli.NewAuthClient()))

	go func() {
		gs_commons_config.WatchGosionConfig(func(config *gs_commons_config.GosionConfiguration) {
			configuration = config
		})
	}()

	go func() {
		gs_commons_config.WatchInitializeConfig(gs_commons_constants.PermissionService, permissionhandlers.Initialize(session, pool))
	}()

	if err := m.Run(); err != nil {
		panic(err)
	}

}
