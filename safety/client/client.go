package safetyclient

import (
	mircogrpc "github.com/micro/go-grpc"
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/registry"
	"github.com/micro/go-micro/registry/consul"
	"konekko.me/gosion/commons/constants"
	"konekko.me/gosion/safety/pb"
	"konekko.me/gosion/safety/pb/ext"
)

func client() micro.Service {
	s := mircogrpc.NewService(micro.Registry(consul.NewRegistry(registry.Addrs("192.168.80.67:8500"),
		registry.Secure(false))))

	s.Init()
	return s
}

func NewSecurityClient() gs_ext_service_safety.SecurityService {
	return gs_ext_service_safety.NewSecurityService(gs_commons_constants.ExtSafetyService, client().Client())
}

func NewBlacklistClient() gs_service_safety.BlacklistService {
	return gs_service_safety.NewBlacklistService(gs_commons_constants.SafetyService, client().Client())
}
