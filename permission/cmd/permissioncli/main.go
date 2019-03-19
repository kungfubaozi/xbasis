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
	"konekko.me/gosion/permission/pb"
)

func NewClient() {
	s := mircogrpc.NewService(micro.Registry(consul.NewRegistry(registry.Addrs("192.168.80.67:8500"),
		registry.Secure(false))))

	s.Init()

	cl := gs_service_permission.NewSayService(gs_commons_constants.PermissionService, s.Client())

	ctx := metadata.NewContext(context.Background(), map[string]string{
		"key": "ajsdflkajsfd;kl",
	})

	out, err := cl.Hello(ctx, &gs_service_permission.Request{})
	if err != nil {
		fmt.Println(err)
		panic(err)
	}

	fmt.Println(out.Msg)
}

func main() {
	NewClient()
}
