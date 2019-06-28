package analysisservice

import (
	"fmt"
	"github.com/Shopify/sarama"
	"konekko.me/xbasis/analysis/client"
	"konekko.me/xbasis/analysis/handlers"
	"konekko.me/xbasis/analysis/pb"
	"konekko.me/xbasis/analysis/report"
	"konekko.me/xbasis/commons/constants"
	"konekko.me/xbasis/commons/indexutils"
	"konekko.me/xbasis/commons/microservice"
)

func StartService() {
	errc := make(chan error, 1)

	client, err := indexutils.NewClient("http://192.168.2.62:9200/")
	if err != nil {
		panic(err)
	}

	logger := analysisclient.NewLoggerClient()

	go func() {
		s := microservice.NewService(xbasisconstants.AnalysisService, true)
		s.Init()

		xbasissvc_external_analysis.RegisterLoggerHandler(s.Server(), analysishandlers.NewLoggerService(logger, client))

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

		partitionConsumer, err := consumer.ConsumePartition("xbs-analysis-loggerresp-topic-1", 0, sarama.OffsetOldest)
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
