package authenticationcli

import (
	"github.com/micro/go-micro/client"
	"konekko.me/gosion/authentication/pb/inner"
	"konekko.me/gosion/commons/constants"
)

//func client() micro.Service {
//	s := micro.NewService(micro.Registry(consul.NewRegistry(registry.Addrs("192.168.80.67:8500"),
//		registry.Secure(false))))
//
//	s.Init()
//	return s
//}

func NewTokenClient(client client.Client) gosionsvc_internal_authentication.TokenService {
	return gosionsvc_internal_authentication.NewTokenService(gs_commons_constants.InternalAuthenticationService, client)
}

func NewAuthClient(client client.Client) gosionsvc_internal_authentication.AuthService {
	return gosionsvc_internal_authentication.NewAuthService(gs_commons_constants.InternalAuthenticationService, client)
}
