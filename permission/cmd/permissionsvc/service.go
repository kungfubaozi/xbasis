package permissionsvc

import (
	"github.com/micro/go-micro"
	"konekko.me/gosion/commons/constants"
	"konekko.me/gosion/commons/microservice"
	"konekko.me/gosion/commons/wrapper"
)

func StartService() {

	m := microservice.NewService(gs_commons_constants.PermissionService)

	m.Init(micro.WrapHandler(gs_commons_wrapper.AuthWrapper))

}