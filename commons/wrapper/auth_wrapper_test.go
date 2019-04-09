package gs_commons_wrapper

import (
	"context"
	"fmt"
	mircogrpc "github.com/micro/go-grpc"
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/registry"
	"github.com/micro/go-micro/registry/consul"
	"konekko.me/gosion/commons/constants"
	"konekko.me/gosion/permission/pb"
	"testing"
)

func TestAuthWrapper(t *testing.T) {

	s := mircogrpc.NewService(micro.Registry(consul.NewRegistry(registry.Addrs("192.168.80.67:8500"),
		registry.Secure(false))))

	verificationClient := gs_service_permission.NewVerificationService(gs_commons_constants.ExtPermissionVerificationService, s.Client())

	ss, err := verificationClient.Check(context.Background(), &gs_service_permission.HasPermissionRequest{})

	if err != nil {
		fmt.Println("err", err)
		return
	}

	fmt.Println("ok", ss)

}
