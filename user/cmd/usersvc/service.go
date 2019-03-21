package usersvc

import (
	"konekko.me/gosion/commons/config"
	"konekko.me/gosion/commons/constants"
	"konekko.me/gosion/commons/microservice"
	"konekko.me/gosion/user/handlers"
	"konekko.me/gosion/user/handlers/nops"
	"konekko.me/gosion/user/pb"
	"konekko.me/gosion/user/pb/nops"
)

func StartService() {

	errc := make(chan error, 2)

	go func() {
		s := microservice.NewService(gs_commons_constants.NOPSPermissionService)

		gs_nops_service_message.RegisterMessageHandler(s.Server(), user_nops_handlers.NewMessageService())

		errc <- s.Run()
	}()

	go func() {
		s := microservice.NewService(gs_commons_constants.UserService)

		gs_service_user.RegisterLoginHandler(s.Server(), user_handlers.NewLoginService())

		gs_service_user.RegisterRegisterHandler(s.Server(), user_handlers.NewRegisterService())

		gs_service_user.RegisterSafetyHandler(s.Server(), user_handlers.NewSafetyService())

		gs_service_user.RegisterUpdateHandler(s.Server(), user_handlers.NewUpdateService())

		gs_service_user.RegisterUserHandler(s.Server(), user_handlers.NewUserService())

		errc <- s.Run()
	}()

	go func() {

		//watch config change to init def data
		gs_commons_config.WatchInitializeConfig(gs_commons_constants.UserService, user_handlers.Initialize())

	}()

	go func() {
		gs_commons_config.WatchGosionConfig(func(config *gs_commons_config.GosionConfiguration) {

		})
	}()

	if err := <-errc; err != nil {
		panic(err)
	}

}
