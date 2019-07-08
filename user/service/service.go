package usersvc

import (
	"konekko.me/xbasis/analysis/client"
	"konekko.me/xbasis/application/client"
	"konekko.me/xbasis/authentication/client"
	xconfig "konekko.me/xbasis/commons/config"
	"konekko.me/xbasis/commons/config/call"
	constants "konekko.me/xbasis/commons/constants"
	dao "konekko.me/xbasis/commons/dao"
	"konekko.me/xbasis/commons/indexutils"
	"konekko.me/xbasis/commons/microservice"
	"konekko.me/xbasis/message/cmd/messagecli"
	"konekko.me/xbasis/permission/client"
	"konekko.me/xbasis/safety/client"
	"konekko.me/xbasis/user/client"
	"konekko.me/xbasis/user/handlers"
	"konekko.me/xbasis/user/pb"
	"konekko.me/xbasis/user/pb/inner"
)

func StartService() {

	errc := make(chan error, 2)

	session, err := dao.CreateSession("192.168.2.60:27017")
	if err != nil {
		panic(err)
	}
	defer session.Close()

	ms, err := messagecli.NewClient()
	if err != nil {
		panic(err)
	}
	defer ms.Close()

	client, err := indexutils.NewClient("http://192.168.2.62:9200/")
	if err != nil {
		panic(err)
	}

	zk := xconfig.NewConnect("192.168.2.57:2181")
	defer zk.Close()

	logger := analysisclient.NewLoggerClient()

	go func() {
		s := microservice.NewService(constants.InternalUserService, true)

		xbasissvc_internal_user.RegisterMessageHandler(s.Server(), userhandlers.NewMessageService(ms, session))

		xbasissvc_internal_user.RegisterUserHandler(s.Server(), userhandlers.NewInnerUserService(session, logger))

		errc <- s.Run()
	}()

	go func() {
		s := microservice.NewService(constants.UserService, true)

		ss := safetyclient.NewSecurityClient(s.Client())
		ts := authenticationcli.NewTokenClient(s.Client())

		xbasissvc_external_user.RegisterLoginHandler(s.Server(), userhandlers.NewLoginService(session, ss, ts, client, logger))

		xbasissvc_external_user.RegisterRegisterHandler(s.Server(), userhandlers.NewRegisterService(logger, session, userclient.NewInviteClient(s.Client()),
			client, permissioncli.NewBindingClient(s.Client()), permissioncli.NewGroupClient(s.Client()), applicationclient.NewStatusClient(s.Client())))

		xbasissvc_external_user.RegisterSafetyHandler(s.Server(), userhandlers.NewSafetyService(session))

		xbasissvc_external_user.RegisterUpdateHandler(s.Server(), userhandlers.NewUpdateService(session))

		xbasissvc_external_user.RegisterInviteHandler(s.Server(), userhandlers.NewInviteService(session, logger, userclient.NewExtUserClient(s.Client())))

		xbasissvc_external_user.RegisterUserInfoHandler(s.Server(), userhandlers.NewUserInfoService(session, logger))

		xbasissvc_external_user.RegisterUserHandler(s.Server(), userhandlers.NewUserService(client))

		xbasissvc_external_user.RegisterOAuthHandler(s.Server(), userhandlers.NewOAuthService(session, client, logger))

		errc <- s.Run()
	}()

	go func() {

		//watch config change to init def data
		xconfig.WatchInitializeConfig(constants.UserService, userhandlers.Initialize(session.Clone(), client))

	}()

	go func() {
		xconfig.WatchGosionConfig(serviceconfiguration.Configuration())
	}()

	if err := <-errc; err != nil {
		panic(err)
	}

}
