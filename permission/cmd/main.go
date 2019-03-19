package main

import (
	"context"
	"fmt"
	mircogrpc "github.com/micro/go-grpc"
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/metadata"
	"github.com/micro/go-micro/registry"
	"github.com/micro/go-micro/registry/consul"
	"konekko.me/gosion/commons/constants"
	"konekko.me/gosion/commons/microservice"
	"konekko.me/gosion/permission/pb"
)

type SayService interface {
	Hello(ctx context.Context, request *gs_service_permission.Request, out *gs_service_permission.Response) error
}

type HiService interface {
	SayHello(ctx context.Context, request *gs_service_permission.Request, out *gs_service_permission.Response) error
}

type sayService struct {
	hiService gs_service_permission.HiService
}

type hiService struct {
}

func (svc *hiService) SayHello(ctx context.Context, request *gs_service_permission.Request, out *gs_service_permission.Response) error {
	fmt.Println("456")

	md, ok := metadata.FromContext(ctx)
	if ok {
		fmt.Println("value is ", md[":path"])
	}

	out.Msg = "asdfasdf"
	return nil
}

func (svc *sayService) Hello(ctx context.Context, request *gs_service_permission.Request, out *gs_service_permission.Response) error {
	fmt.Println("456")

	md, ok := metadata.FromContext(ctx)
	if ok {
		fmt.Println("value is ", md)
	}

	svc.hiService.SayHello(context.Background(), &gs_service_permission.Request{})

	out.Msg = "ikasdfasdf"
	return nil
}

func main() {

	s := microservice.NewService(gs_commons_constants.PermissionService)

	s.Init()

	s1 := mircogrpc.NewService(micro.Registry(consul.NewRegistry(registry.Addrs("192.168.80.67:8500"),
		registry.Secure(false))),
		micro.Name(gs_commons_constants.PermissionService))

	s1.Init()

	gs_service_permission.RegisterSayHandler(s.Server(), &sayService{
		gs_service_permission.NewHiService(gs_commons_constants.PermissionService, s1.Client()),
	})

	gs_service_permission.RegisterHiHandler(s.Server(), &hiService{})

	s.Run()

}
