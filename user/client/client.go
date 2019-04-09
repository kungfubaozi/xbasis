package userclient

import (
	mircogrpc "github.com/micro/go-grpc"
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/registry"
	"github.com/micro/go-micro/registry/consul"
	"konekko.me/gosion/commons/constants"
	"konekko.me/gosion/user/pb/ext"
)

func client() micro.Service {
	s := mircogrpc.NewService(micro.Registry(consul.NewRegistry(registry.Addrs("192.168.80.67:8500"),
		registry.Secure(false))))

	s.Init()
	return s
}

func NewExtUserClient() gs_ext_service_user.UserService {
	return gs_ext_service_user.NewUserService(gs_commons_constants.ExtUserService, client().Client())
}

func NewExtMessageClient() gs_ext_service_user.MessageService {
	return gs_ext_service_user.NewMessageService(gs_commons_constants.ExtUserService, client().Client())
}
