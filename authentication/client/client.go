package authenticationcli

import (
	mircogrpc "github.com/micro/go-grpc"
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/registry"
	"github.com/micro/go-micro/registry/consul"
	"konekko.me/gosion/authentication/pb/ext"
	"konekko.me/gosion/commons/constants"
)

func client() micro.Service {
	s := mircogrpc.NewService(micro.Registry(consul.NewRegistry(registry.Addrs("192.168.80.67:8500"),
		registry.Secure(false))))

	s.Init()
	return s
}

func NewTokenClient() gs_ext_service_authentication.TokenService {
	return gs_ext_service_authentication.NewTokenService(gs_commons_constants.ExtAuthenticationService, client().Client())
}

func NewAuthClient() gs_ext_service_authentication.AuthService {
	return gs_ext_service_authentication.NewAuthService(gs_commons_constants.ExtAuthenticationService, client().Client())
}
