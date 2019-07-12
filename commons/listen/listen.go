package listen

import (
	"fmt"
	"github.com/Shopify/sarama"
)

type OnMessage func(msg []byte)

func CreateKafkaListener(addr, topic string, message OnMessage) error {
	config := sarama.NewConfig()
	config.Consumer.Return.Errors = true
	config.Version = sarama.V0_10_2_1

	// consumer
	consumer, err := sarama.NewConsumer([]string{addr}, config)
	if err != nil {
		panic(err.Error())
	}

	defer consumer.Close()

	partitionConsumer, err := consumer.ConsumePartition(topic, 0, sarama.OffsetOldest)
	if err != nil {
		panic(err.Error())
	}
	defer partitionConsumer.Close()

	errc := make(chan error)

	fmt.Println("function sync topic listen started")

	for {
		select {
		case msg := <-partitionConsumer.Messages():
			fmt.Printf("msg offset: %d, partition: %d, timestamp: %s, value: %s\n",
				msg.Offset, msg.Partition, msg.Timestamp.String(), string(msg.Value))
			message(msg.Value)
		case err := <-partitionConsumer.Errors():
			fmt.Printf("err :%s\n", err.Error())
		}
	}
	return <-errc
}
