package applicationclient

import (
	"github.com/micro/go-micro/client"
	_ "github.com/micro/go-micro/registry/consul"
	"konekko.me/gosion/application/pb"
	"konekko.me/gosion/application/pb/ext"
	"konekko.me/gosion/commons/constants"
)

//func client() micro.Service {
//	s := micro.NewService(micro.Registry(consul.NewRegistry(registry.Addrs("192.168.80.67:8500"),
//		registry.Secure(false))))
//
//	s.Init()
//	return s
//}

func NewClient(client client.Client) gs_service_application.ApplicationService {
	return gs_service_application.NewApplicationService(gs_commons_constants.ApplicationService, client)
}

func NewStatusClient(client client.Client) gs_ext_service_application.ApplicationStatusService {
	return gs_ext_service_application.NewApplicationStatusService(gs_commons_constants.ExtApplicationService, client)
}
