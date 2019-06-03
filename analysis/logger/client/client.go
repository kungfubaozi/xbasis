package loggerclient

import (
	"github.com/Shopify/sarama"
	"github.com/Sirupsen/logrus"
)

type Headers struct {
	UserAgent        string `json:"userAgent"`
	Ip               string `json:"ip"`
	UserId           string `json:"userId"`
	Path             string `json:"path"`
	MicroServiceName string `json:"microServiceName"`
	ClientId         string `json:"client_id"`
	HasAccessToken   bool   `json:"hasAccessToken"`
	HasDurationToken bool   `json:"hasDurationToken"`
	TraceId          string `json:"traceId"`
}

type LogFields map[string]interface{}

type LogContent struct {
	Headers   *Headers   `json:"headers"`
	Fields    *LogFields `json:"fields"`
	Action    string     `json:"action"`
	Message   string     `json:"message"`
	Timestamp int64      `json:"timestamp"`
	Level     int64      `json:"level"`
	PublishId string     `json:"publishId"`
	Content   string     `json:"content"`
	StateCode int64      `json:"stateCode"`
	Dump      string     `json:"dump"`
}

type Logger interface {
	Info(content *LogContent)

	Error(content *LogContent)

	Danger(content *LogContent)

	Warn(content *LogContent)

	Close()
}

type logClient struct {
	log    *logrus.Logger
	client sarama.AsyncProducer
}

func (log *logClient) Close() {
	log.client.Close()
}

func (log *logClient) Error(content *LogContent) {
	panic("implement me")
}

func (log *logClient) Danger(content *LogContent) {
	panic("implement me")
}

func (log *logClient) Warn(content *LogContent) {
	panic("implement me")
}

func (log *logClient) Info(content *LogContent) {
	panic("implement me")
}

/**
build basic log info and send json to kafka to analysis
*/
func NewClient() Logger {
	return &logClient{}
}
