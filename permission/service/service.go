package permissionsvc

import (
	"konekko.me/gosion/analysis/client"
	"konekko.me/gosion/application/client"
	"konekko.me/gosion/authentication/client"
	"konekko.me/gosion/commons/config"
	"konekko.me/gosion/commons/config/call"
	"konekko.me/gosion/commons/constants"
	"konekko.me/gosion/commons/dao"
	"konekko.me/gosion/commons/indexutils"
	"konekko.me/gosion/commons/microservice"
	"konekko.me/gosion/permission/client"
	"konekko.me/gosion/permission/handlers"
	"konekko.me/gosion/permission/pb"
	"konekko.me/gosion/permission/pb/inner"
	"konekko.me/gosion/safety/client"
	"konekko.me/gosion/user/client"
)

func StartService() {

	errc := make(chan error, 3)

	pool, err := gs_commons_dao.CreatePool("192.168.2.60:6379")
	if err != nil {
		panic(err)
	}
	defer pool.Close()

	session, err := gs_commons_dao.CreateSession("192.168.2.60:27017")
	if err != nil {
		panic(err)
	}

	client, err := indexutils.NewClient("http://192.168.2.62:9200/")
	if err != nil {
		panic(err)
	}

	logger := analysisclient.NewLoggerClient()

	go func() {
		m := microservice.NewService(gs_commons_constants.InternalPermission, false)

		gosionsvc_internal_permission.RegisterVerificationHandler(m.Server(), permissionhandlers.NewVerificationService(pool,
			session, applicationclient.NewStatusClient(m.Client()),
			safetyclient.NewBlacklistClient(m.Client()),
			authenticationcli.NewAuthClient(m.Client()), client, logger))

		gosionsvc_internal_permission.RegisterAccessibleHandler(m.Server(), permissionhandlers.NewAccessibleService(client, logger))

		errc <- m.Run()
	}()

	go func() {

		m := microservice.NewService(gs_commons_constants.PermissionService, true)

		us := userclient.NewExtUserClient(m.Client())

		mc := userclient.NewExtMessageClient(m.Client())

		gosionsvc_external_permission.RegisterBindingHandler(m.Server(), permissionhandlers.NewBindingService(client, session, us, logger))

		gosionsvc_external_permission.RegisterDurationAccessHandler(m.Server(), permissionhandlers.NewDurationAccessService(pool, session, mc, client, logger))

		gosionsvc_external_permission.RegisterFunctionHandler(m.Server(), permissionhandlers.NewFunctionService(client, session, logger))

		gosionsvc_external_permission.RegisterUserGroupHandler(m.Server(), permissionhandlers.NewGroupService(pool, session))

		gosionsvc_external_permission.RegisterRoleHandler(m.Server(), permissionhandlers.NewRoleService(session, pool,
			permissioncli.NewBindingClient(m.Client())))

		gosionsvc_external_permission.RegisterStructureHandler(m.Server(), permissionhandlers.NewStructureService(session, client, applicationclient.NewClient(m.Client())))

		errc <- m.Run()

	}()

	go func() {
		gs_commons_config.WatchGosionConfig(serviceconfiguration.Configuration())
	}()

	go func() {
		gs_commons_config.WatchInitializeConfig(gs_commons_constants.PermissionService, permissionhandlers.Initialize(session.Clone(), client))
	}()

	if err := <-errc; err != nil {
		panic(err)
	}

}
