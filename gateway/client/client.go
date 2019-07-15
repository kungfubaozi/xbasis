package xbsgatewayclient

import (
	"github.com/Shopify/sarama"
	"github.com/vmihailenco/msgpack"
	"konekko.me/xbasis/commons/constants"
	"konekko.me/xbasis/commons/transport"
)

type GatewayClient interface {
	SendFunctionChanged(af *xbasistransport.AppFunction)
}

type gatewayClient struct {
	producer sarama.SyncProducer
}

func (svc *gatewayClient) SendFunctionChanged(af *xbasistransport.AppFunction) {
	b, err := msgpack.Marshal(af)
	if err != nil {
		return
	}
	message := &sarama.ProducerMessage{
		Partition: 1,
		Topic:     xbasisconstants.SyncFunctionTopic,
		Value:     sarama.StringEncoder(string(b)),
	}
	_, _, err = svc.producer.SendMessage(message)
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
