package usersvc

import (
	"konekko.me/gosion/commons/config"
	"konekko.me/gosion/commons/constants"
	"konekko.me/gosion/commons/microservice"
	"konekko.me/gosion/user/handlers"
	"konekko.me/gosion/user/pb"
)

func StartService() {

	s := microservice.NewService(gs_commons_constants.UserService)
	s.Init()

	gs_service_user.RegisterLoginHandler(s.Server(), user_handlers.NewLoginService())

	gs_service_user.RegisterRegisterHandler(s.Server(), user_handlers.NewRegisterService())

	gs_service_user.RegisterSafetyHandler(s.Server(), user_handlers.NewSafetyService())

	gs_service_user.RegisterUpdateHandler(s.Server(), user_handlers.NewUpdateService())

	gs_service_user.RegisterUserHandler(s.Server(), user_handlers.NewUserService())

	go func() {

		//watch config change to init def data
		gs_commons_config.ListenInitializeConfig(gs_commons_constants.UserService, user_handlers.Initialize())

	}()

	if err := s.Run(); err != nil {
		panic(err)
	}

}
