package connectioncli

import (
	"fmt"
	"github.com/pkg/errors"
	"github.com/streadway/amqp"
	"github.com/vmihailenco/msgpack"
	"konekko.me/gosion/commons/constants"
)

type Message struct {
	UserId   string
	TargetId string
	Type     int64
	Content  map[string]interface{}
}

type ConnectionClient interface {
	//disconnect user in target application
	OfflineToAppUser(msg *Message) error

	//disconnect user to target app client
	OfflineToAppClientUser(msg *Message) error

	//disconnect user all socket connection
	OfflineToUser(msg *Message) error

	BroadcastToApp(msg *Message) error

	BroadcastToAppClient(msg *Message) error

	BroadcastAll(msg *Message) error

	SendMessageToAppUser(msg *Message) error

	SendMessageToAppClientUser(msg *Message) error

	SendMessageToUser(msg *Message) error

	Close()
}

type operation struct {
	ch   *amqp.Channel
	conn *amqp.Connection
}

func (o *operation) OfflineToAppClientUser(msg *Message) error {
	if len(msg.TargetId) == 0 || len(msg.UserId) == 0 {
		return errMessage
	}
	return o.send("otacu", msg)
}

func (o *operation) BroadcastToAppClient(msg *Message) error {
	if len(msg.TargetId) == 0 {
		return errMessage
	}
	return o.send("btac", msg)
}

func (o *operation) SendMessageToAppClientUser(msg *Message) error {
	if len(msg.TargetId) == 0 || len(msg.UserId) == 0 {
		return errMessage
	}
	return o.send("smtcu", msg)
}

func (o *operation) Close() {
	o.ch.Close()
	o.conn.Close()
}

var errMessage = errors.New("err message")

func (o *operation) OfflineToAppUser(msg *Message) error {
	if len(msg.TargetId) == 0 || len(msg.UserId) == 0 {
		return errMessage
	}
	return o.send("otau", msg)
}

func (o *operation) OfflineToUser(msg *Message) error {
	if len(msg.UserId) == 0 {
		return errMessage
	}
	return o.send("otu", msg)
}

func (o *operation) BroadcastToApp(msg *Message) error {
	if len(msg.TargetId) == 0 || msg.Content == nil {
		return errMessage
	}
	return o.send("bta", msg)
}

func (o *operation) BroadcastAll(msg *Message) error {
	if msg.Content == nil {
		return errMessage
	}
	return o.send("baa", msg)
}

func (o *operation) SendMessageToAppUser(msg *Message) error {
	if len(msg.TargetId) == 0 || msg.Content == nil || len(msg.UserId) == 0 {
		return errMessage
	}
	return o.send("stau", msg)
}

func (o *operation) SendMessageToUser(msg *Message) error {
	if msg.Content == nil || len(msg.UserId) == 0 {
		return errMessage
	}
	return o.send("stu", msg)
}

func (o *operation) send(t string, msg *Message) error {
	b, err := msgpack.Marshal(msg)
	if err != nil {
		return err
	}
	return o.ch.Publish(gs_commons_constants.ConnectionFanoutChannel, "",
		false, false, amqp.Publishing{
			Type: t,
			Body: b,
		})
}

func NewClient() (ConnectionClient, error) {
	conn, err := amqp.Dial("amqp://root:123456@192.168.2.60:5672/")
	if err != nil {
		fmt.Println("connect to message queue error.")
		return nil, err
	}
	ch, err := conn.Channel()
	err = ch.ExchangeDeclare(
		gs_commons_constants.ConnectionFanoutChannel, // name
		"fanout",
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		return nil, err
	}
	return &operation{ch: ch, conn: conn}, nil

}
