package message

import (
	"fmt"
	"github.com/streadway/amqp"
	"net/http"
	"net/url"
	"time"
	"zskparker.com/foundation/pkg/utils"
)

type Service interface {
	SendEmail(to, subject, content string) error

	SendSMS(to, content string)
}

type messageService struct {
	channel *amqp.Channel
	conn    *amqp.Connection
}

func (svc *messageService) SendEmail(to, subject, content string) error {
	panic("implement me")
}

func (svc *messageService) SendSMS(to, content string) {
	println("send sms to", to)
	err := sms(to, content)
	if err != nil {
		fmt.Println("send to ", to, " error")
	}
}

func NewService() Service {
	var svc Service
	{
		svc = &messageService{}
	}
	return svc
}

func sms(mobile, c string) error {
	appId := "EUCP-EMY-SMS0-JBZOQ"
	secretKey := "8553947376603211"
	timestamp := utils.FormatDate(time.Now(), utils.YYYYMMDDHHMMSS)
	sign := utils.Md5(appId + secretKey + timestamp)
	values := url.Values{}
	values.Add("appId", appId)
	values.Add("timestamp", timestamp)
	values.Add("sign", sign)
	values.Add("mobiles", mobile)
	values.Add("content", c)
	_, err := http.PostForm("http://shmtn.b2m.cn:80/simpleinter/sendSMS", values)
	return err
}
