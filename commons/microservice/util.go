package microservice

import (
	"github.com/juju/ratelimit"
	"github.com/micro/go-grpc"
	"github.com/micro/go-micro/server"
	"konekko.me/gosion/commons/wrapper"

	//"github.com/micro/go-grpc"
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/registry"
	"github.com/micro/go-micro/registry/consul"
	r "github.com/micro/go-plugins/wrapper/ratelimiter/ratelimit"

	"time"
)

func NewService(name string, init bool) micro.Service {
	c := ratelimit.NewBucketWithRate(float64(500), int64(100))

	cr := consul.NewRegistry(registry.Addrs("192.168.80.67:8500"),
		registry.Secure(false))

	srv := grpc.NewService(
		micro.Registry(cr), micro.Name(name),
		micro.RegisterTTL(time.Second*30),
		micro.RegisterInterval(time.Second*15),
		micro.WrapClient(r.NewClientWrapper(c, false)),
	)

	if init {
		srv.Init(micro.WrapHandler(func(handlerFunc server.HandlerFunc) server.HandlerFunc {
			return gs_commons_wrapper.AuthWrapper(srv.Client(), handlerFunc)
		}))

	}

	return srv
}

func Watcher() {

}
