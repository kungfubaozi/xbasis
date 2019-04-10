package main

import (
	"context"
	"fmt"
	_ "github.com/micro/go-micro"
	"github.com/micro/go-micro/metadata"
	_ "github.com/micro/go-micro/registry"
	_ "github.com/micro/go-micro/registry/consul"
	"konekko.me/gosion/permission/pb"
	"konekko.me/gosion/permission/service"
)

type sayService struct {
	hiService gs_service_permission.HiService
}

type hiService struct {
}

func (svc *hiService) SayHello(ctx context.Context, request *gs_service_permission.HiRequest, out *gs_service_permission.HiResponse) error {
	fmt.Println("456wer")

	md, ok := metadata.FromContext(ctx)
	if ok {
		fmt.Println("value is ", md)
	}

	out.Msg = "asdfasdf"
	return nil
}

func (svc *sayService) Hello(ctx context.Context, request *gs_service_permission.Request, out *gs_service_permission.Response) error {
	fmt.Println("456werwer")

	md, ok := metadata.FromContext(ctx)
	if ok {
		fmt.Println("value is ", md)
	}

	svc.hiService.SayHello(context.Background(), &gs_service_permission.HiRequest{})

	out.Msg = "ikasdfasdf"
	return nil
}

//func main1() {
//	errc := make(chan error, 2)
//
//	go func() {
//		//s := mircogrpc.NewService(micro.Registry(consul.NewRegistry(registry.Addrs("192.168.80.67:8500"),
//		//	registry.Secure(false))))
//		s1 := microservice.NewService(gs_commons_constants.PermissionService, false)
//		gs_service_permission.RegisterSayHandler(s1.Server(), &sayService{
//			gs_service_permission.NewHiService(gs_commons_constants.ExtPermissionVerificationService, s1.Client()),
//		})
//
//		errc <- s1.Run()
//
//	}()
//
//	go func() {
//		s2 := microservice.NewService(gs_commons_constants.ExtPermissionVerificationService, false)
//		gs_service_permission.RegisterHiHandler(s2.Server(), &hiService{})
//
//		errc <- s2.Run()
//	}()
//
//	fmt.Println(<-errc)
//
//	//permissionsvc.StartService()
//
//}

func main() {
	permissionsvc.StartService()
	//main1()
}
