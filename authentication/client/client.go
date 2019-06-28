package authenticationcli

import (
	"github.com/micro/go-micro/client"
	"konekko.me/xbasis/authentication/pb/inner"
	constants "konekko.me/xbasis/commons/constants"
)

//func client() micro.Service {
//	s := micro.NewService(micro.Registry(consul.NewRegistry(registry.Addrs("192.168.80.67:8500"),
//		registry.Secure(false))))
//
//	s.Init()
//	return s
//}

func NewTokenClient(client client.Client) xbasissvc_internal_authentication.TokenService {
	return xbasissvc_internal_authentication.NewTokenService(constants.InternalAuthenticationService, client)
}

func NewAuthClient(client client.Client) xbasissvc_internal_authentication.AuthService {
	return xbasissvc_internal_authentication.NewAuthService(constants.InternalAuthenticationService, client)
}
