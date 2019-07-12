package xbsgatewayclient

import (
	"github.com/Shopify/sarama"
	"konekko.me/xbasis/commons/transport"
)

type GatewayClient interface {
	SendFunctionChanged(af *xbasistransport.AppFunction)
}

type gatewayClient struct {
	producer sarama.SyncProducer
}

func (svc *gatewayClient) SendFunctionChanged(af *xbasistransport.AppFunction) {
	panic("implement me")
}

func NewClient(addr string) GatewayClient {
	config := sarama.NewConfig()
	config.Net.MaxOpenRequests = 300
	config.Producer.RequiredAcks = sarama.NoResponse
	config.Producer.Partitioner = sarama.NewRandomPartitioner
	config.Producer.Return.Successes = true
	config.Producer.Return.Errors = true
	config.Version = sarama.V0_10_2_1

	producer, err := sarama.NewSyncProducer([]string{addr}, config)
	if err != nil {
		panic(err)
	}

	return &gatewayClient{
		producer: producer,
	}
}
