package gslogrus

import (
	"fmt"
	"github.com/Sirupsen/logrus"
	"konekko.me/gosion/commons/date"
	"konekko.me/gosion/commons/indexutils"
	"time"
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
	d := &logdata{
		Service:   l.serviceName,
		Timestamp: data.Time.UnixNano(),
		Message:   data.Message,
		Level:     data.Level.String(),
		Data:      data.Data,
	}

	_, err := l.client.AddData(fmt.Sprintf("gs_logger.%s", gs_commons_date.FormatDate(time.Now(), gs_commons_date.YYYY_MM_DD)), d)
	return err
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

	l := &Logger{log: log, serviceName: serviceName, client: client}

	log.AddHook(l)

	return l
}

type Map map[string]interface{}

type TCI struct {
	fields logrus.Fields
}

func (t *TCI) WithAction(action string, fields logrus.Fields) *logrus.Entry {
	t.fields["action"] = action
	return logrus.WithFields(fields)
}

func (l *Logger) WithHeaders(traceId, clientId, ip, path, userAgent, userDevice string) *TCI {
	fields := logrus.Fields{}
	fields["trace_id"] = traceId
	fields["from_client_id"] = clientId
	fields["request_path"] = path
	fields["user_ip"] = ip
	fields["user_agent"] = userAgent
	fields["user_device"] = userDevice
	return &TCI{fields: fields}
}
