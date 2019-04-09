package permissioncli

import (
	mircogrpc "github.com/micro/go-grpc"
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/registry"
	"github.com/micro/go-micro/registry/consul"
	"konekko.me/gosion/commons/constants"
	"konekko.me/gosion/permission/pb"
)

func NewVerificationClient() gs_service_permission.VerificationService {
	s := mircogrpc.NewService(micro.Registry(consul.NewRegistry(registry.Addrs("192.168.80.67:8500"),
		registry.Secure(false))))

	s.Init()

	return gs_service_permission.NewVerificationService(gs_commons_constants.PermissionService, s.Client())
}
