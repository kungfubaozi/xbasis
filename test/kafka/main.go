package main

import (
	"encoding/json"
	"fmt"
	"github.com/Shopify/sarama"
	"konekko.me/gosion/commons/generator"
	"time"
)

func main() {
	//log := loggerclient.NewClient()
	////id := gs_commons_generator.NewIDG()
	//
	//log.Info(&loggerclient.LogContent{
	//	Headers: &loggerclient.Headers{
	//		ServiceName: "",
	//		Ip:          "192.168.80.67",
	//		UserAgent:   "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_13_6) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/74.0.3729.169 Safari/537.36",
	//		TraceId:     "wACIHN7ukUEqKY2whzmPyZw0x/whT8XBq+ojnYiauJdkU4fZh0tYtQ0hgJJsB",
	//	},
	//	Action:    "LoginSuccess",
	//	Message:   "start verification",
	//	StateCode: 0,
	//	Fields: &loggerclient.Fields{
	//		"appId":    "5135597a5a69",
	//		"clientId": "597a5957566d",
	//		"userId":   "MDk1YThiYTdlMjMxMmU0MjBhYzY5YmYzZjhjN2E0ZjQ2OTc5ZTA2Yw==",
	//	},
	//})

}

func test() {
	config := sarama.NewConfig()
	config.Producer.RequiredAcks = sarama.WaitForAll
	config.Producer.Partitioner = sarama.NewRandomPartitioner
	config.Producer.Return.Successes = true
	config.Version = sarama.V0_10_2_1

	producer, err := sarama.NewAsyncProducer([]string{"192.168.2.62:9092"}, config)
	if err != nil {
		panic(err)
	}

	defer producer.AsyncClose()

	id := gs_commons_generator.NewIDG()

	for i := 0; i < 20; i++ {

		//构建发送的消息，
		msg := &sarama.ProducerMessage{
			Partition: 1,
			Key:       sarama.StringEncoder(id.String()),
			Topic:     "gs-kafka-analysis-topic-3",
		}

		rule := &InputData{
			Action:    "TestSendRuleAction",
			Timestamp: time.Now().UnixNano() / 1e6,
			Bindings: []string{
				"TestTag1",
			},
		}

		b, err := json.Marshal(rule)
		if err != nil {
			panic(err)
		}

		fmt.Println(string(b))

		msg.Value = sarama.StringEncoder(b)

		producer.Input() <- msg

		select {
		case suc := <-producer.Successes():
			fmt.Printf("offset: %d,  timestamp: %s \n", suc.Offset, suc.Timestamp.String())
		case fail := <-producer.Errors():
			fmt.Printf("err: %s\n", fail.Err.Error())
		}
	}

}

type InputData struct {
	Action    string   `json:"action"`
	Bindings  []string `json:"bindings"`
	Timestamp int64    `json:"timestamp"`
	Increment int64    `json:"increment"`
}

type inputRule struct {
	Action        string          `json:"action"`
	Type          string          `json:"type"`
	Restrict      string          `json:"restrict"`
	RuleType      string          `json:"ruleType"` //count, inc
	RuleValue     int64           `json:"ruleValue"`
	DivisionRules []*divisionRule `json:"divisionRules"` //hour, minute, max, eq, min
	Timestamp     int64           `json:"timestamp"`
}

type divisionRule struct {
	RuleType  string `json:"ruleType"`
	RuleValue int64  `json:"ruleValue"` //第二次提交时不会生效， 会引用上一个此类型的ruleValue
}
