package usersvc

import (
	"konekko.me/gosion/analysis/client"
	"konekko.me/gosion/authentication/client"
	"konekko.me/gosion/commons/config"
	"konekko.me/gosion/commons/config/call"
	"konekko.me/gosion/commons/constants"
	"konekko.me/gosion/commons/dao"
	"konekko.me/gosion/commons/indexutils"
	"konekko.me/gosion/commons/microservice"
	"konekko.me/gosion/message/cmd/messagecli"
	"konekko.me/gosion/safety/client"
	"konekko.me/gosion/user/handlers"
	"konekko.me/gosion/user/pb"
	"konekko.me/gosion/user/pb/inner"
)

func StartService() {

	errc := make(chan error, 2)

	session, err := gs_commons_dao.CreateSession("192.168.2.60:27017")
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

	zk := gs_commons_config.NewConnect("192.168.2.57:2181")
	defer zk.Close()

	logger := analysisclient.NewLoggerClient()

	go func() {
		s := microservice.NewService(gs_commons_constants.InternalUserService, true)

		gosionsvc_internal_user.RegisterMessageHandler(s.Server(), userhandlers.NewMessageService(ms, session))

		gosionsvc_internal_user.RegisterUserHandler(s.Server(), userhandlers.NewInnerUserService(session, logger))

		errc <- s.Run()
	}()

	go func() {
		s := microservice.NewService(gs_commons_constants.UserService, true)

		ss := safetyclient.NewSecurityClient(s.Client())
		ts := authenticationcli.NewTokenClient(s.Client())

		gosionsvc_external_user.RegisterLoginHandler(s.Server(), userhandlers.NewLoginService(session, ss, ts, client, logger))

		gosionsvc_external_user.RegisterRegisterHandler(s.Server(), userhandlers.NewRegisterService(session, zk))

		gosionsvc_external_user.RegisterSafetyHandler(s.Server(), userhandlers.NewSafetyService(session))

		gosionsvc_external_user.RegisterUpdateHandler(s.Server(), userhandlers.NewUpdateService(session))

		gosionsvc_external_user.RegisterInviteHandler(s.Server(), userhandlers.NewInviteService(session, logger))

		gosionsvc_external_user.RegisterUserInfoHandler(s.Server(), userhandlers.NewUserInfoService(session, logger))

		gosionsvc_external_user.RegisterUserHandler(s.Server(), userhandlers.NewUserService(session, client))

		gosionsvc_external_user.RegisterOAuthHandler(s.Server(), userhandlers.NewOAuthService(session, client, logger))

		errc <- s.Run()
	}()

	go func() {

		//watch config change to init def data
		gs_commons_config.WatchInitializeConfig(gs_commons_constants.UserService, userhandlers.Initialize(session.Clone(), client))

	}()

	go func() {
		gs_commons_config.WatchGosionConfig(serviceconfiguration.Configuration())
	}()

	if err := <-errc; err != nil {
		panic(err)
	}

}
