package usercli

import (
	mircogrpc "github.com/micro/go-grpc"
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/registry"
	"github.com/micro/go-micro/registry/consul"
	"konekko.me/gosion/commons/constants"
	"konekko.me/gosion/user/pb"
)

func NewUserClient() gs_service_user.UserService {

	s := mircogrpc.NewService(micro.Registry(consul.NewRegistry(registry.Addrs("localhost:8500"),
		registry.Secure(false))))

	s.Init()

	cl := gs_service_user.NewUserService(gs_commons_constants.UserService, s.Client())

	return cl
}
