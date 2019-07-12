package permissionsvc

import (
	"konekko.me/xbasis/analysis/client"
	"konekko.me/xbasis/application/client"
	"konekko.me/xbasis/authentication/client"
	xconfig "konekko.me/xbasis/commons/config"
	"konekko.me/xbasis/commons/config/call"
	constants "konekko.me/xbasis/commons/constants"
	"konekko.me/xbasis/commons/dao"
	"konekko.me/xbasis/commons/indexutils"
	"konekko.me/xbasis/commons/microservice"
	"konekko.me/xbasis/gateway/client"
	"konekko.me/xbasis/permission/client"
	"konekko.me/xbasis/permission/handlers"
	"konekko.me/xbasis/permission/pb"
	"konekko.me/xbasis/permission/pb/inner"
	"konekko.me/xbasis/safety/client"
	"konekko.me/xbasis/user/client"
)

func StartService() {

	errc := make(chan error, 3)

	pool, err := xbasisdao.CreatePool("192.168.2.60:6379")
	if err != nil {
		panic(err)
	}
	defer pool.Close()

	session, err := xbasisdao.CreateSession("192.168.2.60:27017")
	if err != nil {
		panic(err)
	}

	client, err := indexutils.NewClient("http://192.168.2.62:9200/")
	if err != nil {
		panic(err)
	}

	logger := analysisclient.NewLoggerClient()

	zk := xconfig.NewConnect("192.168.2.57:2181")
	defer zk.Close()

	go func() {
		m := microservice.NewService(constants.InternalPermission, false)

		xbasissvc_internal_permission.RegisterVerificationHandler(m.Server(), permissionhandlers.NewVerificationService(pool,
			session, applicationclient.NewStatusClient(m.Client()),
			safetyclient.NewBlacklistClient(m.Client()),
			authenticationcli.NewAuthClient(m.Client()), client, logger))

		xbasissvc_internal_permission.RegisterAccessibleHandler(m.Server(), permissionhandlers.NewAccessibleService(client, session, logger))

		errc <- m.Run()
	}()

	go func() {

		gateway := xbsgatewayclient.NewClient("192.168.2.62:9092")

		m := microservice.NewService(constants.PermissionService, true)

		us := userclient.NewExtUserClient(m.Client())

		mc := userclient.NewExtMessageClient(m.Client())

		xbasissvc_external_permission.RegisterBindingHandler(m.Server(), permissionhandlers.NewBindingService(client, session, us, logger))

		xbasissvc_external_permission.RegisterDurationAccessHandler(m.Server(), permissionhandlers.NewDurationAccessService(pool, session, zk, mc, client, logger))

		xbasissvc_external_permission.RegisterFunctionHandler(m.Server(), permissionhandlers.NewFunctionService(client, session, logger, gateway))

		xbasissvc_external_permission.RegisterUserGroupHandler(m.Server(), permissionhandlers.NewGroupService(pool, session, userclient.NewExtUserClient(m.Client()), logger))

		xbasissvc_external_permission.RegisterRoleHandler(m.Server(), permissionhandlers.NewRoleService(session, pool,
			permissioncli.NewBindingClient(m.Client()), client))

		errc <- m.Run()

	}()

	go func() {
		xconfig.WatchGosionConfig(serviceconfiguration.Configuration())
	}()

	go func() {
		xconfig.WatchInitializeConfig(constants.PermissionService, permissionhandlers.Initialize(session.Clone(), client, zk))
	}()

	if err := <-errc; err != nil {
		panic(err)
	}

}
