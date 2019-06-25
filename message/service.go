package message

import (
	"fmt"
	"github.com/streadway/amqp"
	"io/ioutil"
	"net/http"
	"strings"
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
	//appId := "EUCP-EMY-SMS0-JBZOQ"
	//secretKey := "8553947376603211"
	//timestamp := utils.FormatDate(time.Now(), utils.YYYYMMDDHHMMSS)
	//sign := utils.Md5(appId + secretKey + timestamp)
	//values := url.Values{}
	//values.Add("appId", appId)
	//values.Add("timestamp", timestamp)
	//values.Add("sign", sign)
	//values.Add("mobiles", mobile)
	//values.Add("content", c)
	//r, err := http.PostForm("http://shmtn.b2m.cn:80/simpleinter/sendSMS", values)
	//defer r.Body.Close()
	//
	//resp, err := ioutil.ReadAll(r.Body)
	//if err != nil {
	//	fmt.Println("err", err)
	//	return err
	//}
	//
	//fmt.Println("data", string(resp))

	t := utils.FormatDate(time.Now(), utils.YYYYMMDDHHMMSS)
	sig := "6635c476181c448d82057cf06bae92fafe6a975822fd4fc2afbd83ba282a951d" + t
	m := utils.Md5(sig)
	str := fmt.Sprintf("accountSid=6635c476181c448d82057cf06bae92fa&smsContent=【Gosion】%s&to=%s&timestamp=%s&sig=%s&respDataType=JSON",
		c,
		mobile,
		t,
		m)
	v, err := http.Post("https://api.miaodiyun.com/20150822/industrySMS/sendSMS",
		"application/x-www-form-urlencoded", strings.NewReader(str))
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("Send validate message to %s success.\n", mobile)
		d, err := ioutil.ReadAll(v.Body)
		if err != nil {
			fmt.Println("err", err)
			return err
		}
		fmt.Println("input", c)
		fmt.Println("data", string(d))
	}

	return err
}
