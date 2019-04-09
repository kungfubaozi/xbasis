package usersvc

import (
	"konekko.me/gosion/authentication/client"
	"konekko.me/gosion/commons/config"
	"konekko.me/gosion/commons/constants"
	"konekko.me/gosion/commons/dao"
	"konekko.me/gosion/commons/microservice"
	"konekko.me/gosion/message/cmd/messagecli"
	"konekko.me/gosion/safety/client"
	"konekko.me/gosion/user/handlers"
	"konekko.me/gosion/user/pb"
	"konekko.me/gosion/user/pb/ext"
)

func StartService() {

	errc := make(chan error, 2)

	session, err := gs_commons_dao.CreateSession("192.168.2.60:27017")
	if err != nil {
		panic(err)
	}
	defer session.Close()

	ss := safetyclient.NewSecurityClient()
	ts := authenticationcli.NewTokenClient()

	ms, err := messagecli.NewClient()
	if err != nil {
		panic(err)
	}
	defer ms.Close()

	configuration := &gs_commons_config.GosionConfiguration{}

	go func() {
		s := microservice.NewService(gs_commons_constants.ExtUserService)

		gs_ext_service_user.RegisterMessageHandler(s.Server(), userhandlers.NewMessageService(ms))

		gs_ext_service_user.RegisterUserHandler(s.Server(), userhandlers.NewExtUserService())

		errc <- s.Run()
	}()

	go func() {
		s := microservice.NewService(gs_commons_constants.UserService)

		gs_service_user.RegisterLoginHandler(s.Server(), userhandlers.NewLoginService(session, ss, ts))

		gs_service_user.RegisterRegisterHandler(s.Server(), userhandlers.NewRegisterService(session))

		gs_service_user.RegisterSafetyHandler(s.Server(), userhandlers.NewSafetyService(session))

		gs_service_user.RegisterUpdateHandler(s.Server(), userhandlers.NewUpdateService(session))

		errc <- s.Run()
	}()

	go func() {

		//watch config change to init def data
		gs_commons_config.WatchInitializeConfig(gs_commons_constants.UserService, userhandlers.Initialize(session))

	}()

	go func() {
		gs_commons_config.WatchGosionConfig(func(config *gs_commons_config.GosionConfiguration) {
			configuration = config
		})
	}()

	if err := <-errc; err != nil {
		panic(err)
	}

}
