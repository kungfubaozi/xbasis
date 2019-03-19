package gs_commons_registration

import (
	"context"
	"fmt"
	"github.com/go-kit/kit/log"
	"github.com/hashicorp/consul/api"
	stdconsul "github.com/hashicorp/consul/api"
	"google.golang.org/grpc"
	"google.golang.org/grpc/health/grpc_health_v1"
	"os"
	"zskparker.com/foundation/pkg/osenv"
)

// HealthImpl 健康检查实现
type healthImpl struct{}

// Check 实现健康检查接口，这里直接返回健康状态，这里也可以有更复杂的健康检查策略，比如根据服务器负载来返回
func (h *healthImpl) Check(ctx context.Context, req *grpc_health_v1.HealthCheckRequest) (*grpc_health_v1.HealthCheckResponse, error) {
	return &grpc_health_v1.HealthCheckResponse{
		Status: grpc_health_v1.HealthCheckResponse_SERVING,
	}, nil
}

func (h *healthImpl) Watch(in *grpc_health_v1.HealthCheckRequest, s grpc_health_v1.Health_WatchServer) error {
	return nil
}

func NewRegistrar(gs *grpc.Server, name, consulAddr string) error {

	checkPort := osenv.GetMicroPort()

	var logger log.Logger
	{
		logger = log.NewLogfmtLogger(os.Stderr)
		logger = log.With(logger, "ts", log.DefaultTimestampUTC)
		logger = log.With(logger, "caller", log.DefaultCaller)
	}

	consulConfig := api.DefaultConfig()
	consulConfig.Address = consulAddr
	consulClient, err := api.NewClient(consulConfig)
	if err != nil {
		os.Exit(1)
	}

	agent := consulClient.Agent()

	ip := osenv.GetHostIp()
	id := fmt.Sprintf("%v-%v", name, os.Getenv("NODE_NAME"))

	//HTTP:                           fmt.Sprintf("http://%s:%s%s", ip, checkPort, "/check"),

	check := api.AgentServiceCheck{
		GRPC:                           fmt.Sprintf("%v:%v/%v", ip, checkPort, name),
		Interval:                       "30s",
		DeregisterCriticalServiceAfter: "30s",
	}

	fmt.Println(check.GRPC)

	reg := &stdconsul.AgentServiceRegistration{
		Name:    name,
		Address: ip,
		ID:      id,
		Tags:    []string{"Gosion Service"},
		Port:    int(checkPort),
		Check:   &check,
	}

	grpc_health_v1.RegisterHealthServer(gs, &healthImpl{})

	return agent.ServiceRegister(reg)
}
