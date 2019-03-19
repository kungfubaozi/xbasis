package messagecli

import (
	"fmt"
	"github.com/streadway/amqp"
	"github.com/vmihailenco/msgpack"
	"konekko.me/gosion/commons/constants"
	"konekko.me/gosion/message"
)

type messageChannelClient struct {
	channel *amqp.Channel
	conn    *amqp.Connection
}

func (cli *messageChannelClient) SendEmail(to, subject, content string) error {
	return cli.send("email", &message.Message{To: to, Content: content, Subject: subject})
}

func (cli *messageChannelClient) SendSMS(to, content string) error {
	return cli.send("sms", &message.Message{To: to, Content: content})
}

func (cli *messageChannelClient) send(k string, message *message.Message) error {
	b, err := msgpack.Marshal(message)
	if err != nil {
		return err
	}
	return cli.channel.Publish("", gs_commons_constants.MessageChannel, false, false, amqp.Publishing{
		Body: b,
		Type: k,
	})
}

func (cli *messageChannelClient) Close() {
	cli.channel.Close()
	cli.conn.Close()
}

type MessageClient interface {
	SendEmail(to, subject, content string) error

	SendSMS(to, content string) error

	Close()
}

func NewClient() (MessageClient, error) {
	conn, err := amqp.Dial("amqp://root:123456@192.168.2.60:5672/")
	if err != nil {
		fmt.Println("message connect to message queue error.")
		return nil, err
	}
	ch, err := conn.Channel()
	if err != nil {
		fmt.Println("message channel error.")
		return nil, err
	}
	return &messageChannelClient{
		channel: ch, conn: conn,
	}, nil
}
