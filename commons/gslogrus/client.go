package gslogrus

import (
	"fmt"
	"github.com/Sirupsen/logrus"
	"konekko.me/gosion/commons/date"
	"konekko.me/gosion/commons/indexutils"
)

type Logger struct {
	log         *logrus.Logger
	serviceName string
	client      *indexutils.Client
}

func (l *Logger) Levels() []logrus.Level {
	return logrus.AllLevels
}

func (l *Logger) Fire(data *logrus.Entry) error {
	go func() {
		d := &logdata{
			Service:   l.serviceName,
			Timestamp: data.Time.UnixNano(),
			Message:   data.Message,
			Level:     data.Level.String(),
			Data:      data.Data,
		}

		l.client.AddData(fmt.Sprintf("gs_logger.%s", gs_commons_date.FormatDate(data.Time, gs_commons_date.YYYY_MM_DD)), d)

	}()

	return nil
}

type logdata struct {
	Service   string        `json:"service"`
	Timestamp int64         `json:"@timestamp"`
	Message   string        `json:"message"`
	Data      logrus.Fields `json:"data"`
	Level     string        `json:"level"`
}

func New(serviceName string, client *indexutils.Client) *Logger {
	log := logrus.New()

	log.SetFormatter(&logrus.JSONFormatter{PrettyPrint: true})

	l := &Logger{log: log, serviceName: serviceName, client: client}

	log.Hooks.Add(l)

	return l
}

type Map map[string]interface{}

type TCI struct {
	fields logrus.Fields
	log    *logrus.Logger
}

func (t *TCI) WithAction(action string, fields logrus.Fields) *logrus.Entry {
	fields["content"] = t.fields
	t.fields["action"] = action
	return t.log.WithFields(fields)
}

func (l *Logger) WithHeaders(traceId, clientId, ip, path, userAgent, userDevice string) *TCI {
	fields := logrus.Fields{}
	fields["trace_id"] = traceId
	fields["from_client_id"] = clientId
	fields["request_path"] = path
	fields["user_ip"] = ip
	fields["user_agent"] = userAgent
	fields["user_device"] = userDevice
	return &TCI{fields: fields, log: l.log}
}
