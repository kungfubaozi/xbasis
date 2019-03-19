package microservice

import (
	"github.com/juju/ratelimit"
	"konekko.me/gosion/commons/wrapper"

	//"github.com/micro/go-grpc"
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/registry"
	"github.com/micro/go-micro/registry/consul"
	r "github.com/micro/go-plugins/wrapper/ratelimiter/ratelimit"
	"time"
)

func NewService(name string) micro.Service {
	c := ratelimit.NewBucketWithRate(float64(500), int64(100))

	cr := consul.NewRegistry(registry.Addrs("192.168.80.67:8500"),
		registry.Secure(false))

	cr.Watch()

	s := micro.NewService(micro.Registry(cr),
		micro.Name(name),
		micro.RegisterTTL(time.Second*30),
		micro.RegisterInterval(time.Second*15),
		micro.WrapHandler(gs_commons_wrapper.AuthWrapper),
		micro.WrapClient(r.NewClientWrapper(c, false)),
	)

	return s
}

func Watcher() {

}
