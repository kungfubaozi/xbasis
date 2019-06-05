package analysisclient

import (
	"encoding/json"
	"fmt"
	"github.com/Shopify/sarama"
	"github.com/Sirupsen/logrus"
	"konekko.me/gosion/commons/generator"
	"time"
)

type LogHeaders struct {
	ServiceName      string `json:"serviceName"`
	ModuleName       string `json:"moduleName"`
	UserAgent        string `json:"userAgent"`
	Ip               string `json:"ip"`
	UserId           string `json:"userId"`
	Path             string `json:"path"`
	Api              string `json:"api"`
	ClientId         string `json:"clientId"`
	HasAccessToken   bool   `json:"hasAccessToken"`
	HasDurationToken bool   `json:"hasDurationToken"`
	TraceId          string `json:"traceId"`
	Device           string `json:"device"`
}

type LogFields map[string]interface{}

type LogContent struct {
	Id        string      `json:"id"`
	Headers   *LogHeaders `json:"headers"`
	Fields    *LogFields  `json:"fields"`
	Action    string      `json:"action"`
	Message   string      `json:"message"`
	Timestamp int64       `json:"timestamp"`
	Level     string      `json:"level"`
	PublishId string      `json:"publishId"`
	Content   string      `json:"content"`
	StateCode int64       `json:"stateCode"`
	Dump      string      `json:"dump"`
}

type LogClient interface {
	Info(content *LogContent)

	Error(content *LogContent)

	Danger(content *LogContent)

	Warn(content *LogContent)

	Close()
}

type logClient struct {
	id       gs_commons_generator.IDGenerator
	log      *logrus.Logger
	producer sarama.SyncProducer
}

func (log *logClient) Close() {
	log.producer.Close()
}

func (log *logClient) Error(content *LogContent) {
	content.Level = "error"
	log.fields(content).Error(content.Message)
}

func (log *logClient) Danger(content *LogContent) {
	content.Level = "danger"
	log.fields(content).Error(content.Message)
}

func (log *logClient) Warn(content *LogContent) {
	content.Level = "warn"
	log.fields(content).Warn(content.Message)
}

func (log *logClient) Info(content *LogContent) {
	content.Level = "info"
	log.fields(content).Info(content.Message)
}

func (log *logClient) fields(content *LogContent) *logrus.Entry {
	content.Timestamp = time.Now().UnixNano() / 1e6
	content.Id = log.id.String()
	log.buildMessage(content)
	return log.log.WithFields(logrus.Fields{
		"action":    content.Action,
		"timestamp": content.Timestamp,
		"state":     content.StateCode,
	})
}

func (log *logClient) buildMessage(content *LogContent) {
	b, err := json.Marshal(content)
	if err != nil {
		fmt.Println("log content json marshal err", err)
		return
	}
	message := &sarama.ProducerMessage{
		Partition: 1,
		Topic:     "gs-kafka-analysis-logger-topic-4",
		Value:     sarama.StringEncoder(string(b)),
	}
	log.producer.SendMessage(message)
}

/**
build basic log info and send json to kafka to analysis
*/
func NewLoggerClient() LogClient {
	log := logrus.New()
	config := sarama.NewConfig()
	config.Net.MaxOpenRequests = 300
	config.Producer.RequiredAcks = sarama.NoResponse
	config.Producer.Partitioner = sarama.NewRandomPartitioner
	config.Producer.Return.Successes = true
	config.Producer.Return.Errors = true
	config.Version = sarama.V0_10_2_1

	producer, err := sarama.NewSyncProducer([]string{"192.168.2.62:9092"}, config)
	if err != nil {
		panic(err)
	}

	return &logClient{
		log:      log,
		producer: producer,
		id:       gs_commons_generator.NewIDG(),
	}
}
