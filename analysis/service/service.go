package analysisservice

import (
	"fmt"
	"github.com/Shopify/sarama"
	"konekko.me/gosion/analysis/client"
	"konekko.me/gosion/analysis/handlers"
	"konekko.me/gosion/analysis/pb"
	"konekko.me/gosion/analysis/report"
	"konekko.me/gosion/commons/constants"
	"konekko.me/gosion/commons/indexutils"
	"konekko.me/gosion/commons/microservice"
)

func StartService() {
	errc := make(chan error, 1)

	client, err := indexutils.NewClient("http://192.168.2.62:9200/")
	if err != nil {
		panic(err)
	}

	logger := analysisclient.NewLoggerClient()

	go func() {
		s := microservice.NewService(gs_commons_constants.AnalysisService, true)
		s.Init()

		gosionsvc_external_analysis.RegisterLoggerHandler(s.Server(), analysishandlers.NewLoggerService(logger, client))

		errc <- s.Run()
	}()

	go func() {
		config := sarama.NewConfig()
		config.Consumer.Return.Errors = true
		config.Version = sarama.V0_10_2_1

		// consumer
		consumer, err := sarama.NewConsumer([]string{"192.168.2.62:9092"}, config)
		if err != nil {
			panic(err.Error())
		}

		defer consumer.Close()

		partitionConsumer, err := consumer.ConsumePartition("gs-analysis-logger-resp-1", 0, sarama.OffsetOldest)
		if err != nil {
			panic(err.Error())
		}
		defer partitionConsumer.Close()

		r := analysisreport.NewAction()

		for {
			select {
			case msg := <-partitionConsumer.Messages():
				fmt.Printf("msg offset: %d, partition: %d, timestamp: %s, value: %s\n",
					msg.Offset, msg.Partition, msg.Timestamp.String(), string(msg.Value))
				r.Action(msg.Value)
			case err := <-partitionConsumer.Errors():
				fmt.Printf("err :%s\n", err.Error())
			}
		}
	}()

	if err := <-errc; err != nil {
		panic(err)
	}
}
