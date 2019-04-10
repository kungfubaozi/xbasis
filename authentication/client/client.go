package authenticationcli

import (
	"github.com/micro/go-micro/client"
	"konekko.me/gosion/authentication/pb/ext"
	"konekko.me/gosion/commons/constants"
)

//func client() micro.Service {
//	s := micro.NewService(micro.Registry(consul.NewRegistry(registry.Addrs("192.168.80.67:8500"),
//		registry.Secure(false))))
//
//	s.Init()
//	return s
//}

func NewTokenClient(client client.Client) gs_ext_service_authentication.TokenService {
	return gs_ext_service_authentication.NewTokenService(gs_commons_constants.ExtAuthenticationService, client)
}

func NewAuthClient(client client.Client) gs_ext_service_authentication.AuthService {
	return gs_ext_service_authentication.NewAuthService(gs_commons_constants.ExtAuthenticationService, client)
}
