package analysisclient

import (
	"encoding/json"
	"github.com/Shopify/sarama"
	"konekko.me/gosion/commons/generator"
	"time"
)

type content struct {
	Action    string   `json:"action"`
	Bindings  []string `json:"bindings"`
	Timestamp int64    `json:"timestamp"`
	Increment int64    `json:"increment"`
}

type FrequencyClient interface {
	Increment(action string, bindings ...string)

	Plus(number int64, action string, bindings ...string)
}

type client struct {
	id       gs_commons_generator.IDGenerator
	producer sarama.AsyncProducer
}

func (c *client) Increment(action string, bindings ...string) {
	c.build(1, action, bindings...)
}

func (c *client) Plus(number int64, action string, bindings ...string) {
	c.build(number, action, bindings...)
}

func (c *client) build(number int64, action string, bindings ...string) {
	content := &content{
		Action:    action,
		Increment: number,
		Bindings:  bindings,
		Timestamp: time.Now().UnixNano() / 1e6, //mil
	}
	b, err := json.Marshal(content)
	if err != nil {
		return
	}

	c.producer.Input() <- &sarama.ProducerMessage{
		Topic: "gs-kafka-analysis-topic-3",
		Key:   sarama.StringEncoder(c.id.String()),
		Value: sarama.StringEncoder(string(b)),
	}
}

func NewFrequencyClient() FrequencyClient {
	config := sarama.NewConfig()
	config.Producer.RequiredAcks = sarama.WaitForAll
	config.Producer.Partitioner = sarama.NewRandomPartitioner
	config.Producer.Return.Successes = true
	config.Version = sarama.V0_10_2_1

	producer, err := sarama.NewAsyncProducer([]string{"192.168.2.62:9092"}, config)
	if err != nil {
		panic(err)
	}
	return &client{
		producer: producer,
		id:       gs_commons_generator.NewIDG(),
	}
}
