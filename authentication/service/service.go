package authenticationsvc

import (
	"konekko.me/gosion/commons/config"
	"konekko.me/gosion/commons/constants"
	"konekko.me/gosion/commons/microservice"
)

func StartService() {

	s := microservice.NewService(gs_commons_constants.AuthenticationService)
	s.Init()

	go func() {
		gs_commons_config.WatchGosionConfig(func(config *gs_commons_config.GosionConfiguration) {

		})
	}()

	if err := s.Run(); err != nil {
		panic(err)
	}

}
