package applicationclient

import (
	mircogrpc "github.com/micro/go-grpc"
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/registry"
	"github.com/micro/go-micro/registry/consul"
	"konekko.me/gosion/application/pb"
	"konekko.me/gosion/application/pb/ext"
	"konekko.me/gosion/commons/constants"
)

func client() micro.Service {
	s := mircogrpc.NewService(micro.Registry(consul.NewRegistry(registry.Addrs("192.168.80.67:8500"),
		registry.Secure(false))))

	s.Init()
	return s
}

func NewClient() gs_service_application.ApplicationService {
	return gs_service_application.NewApplicationService(gs_commons_constants.ApplicationService, client().Client())
}

func NewStatusClient() gs_ext_service_application.ApplicationStatusService {
	return gs_ext_service_application.NewApplicationStatusService(gs_commons_constants.ExtApplicationService, client().Client())
}
