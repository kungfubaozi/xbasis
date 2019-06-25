package messagesvc

import (
	"fmt"
	"github.com/streadway/amqp"
	"github.com/vmihailenco/msgpack"
	"konekko.me/gosion/commons/constants"
	"konekko.me/gosion/message"
)

func StartService() {

	service := message.NewService()

	conn, err := amqp.Dial("amqp://root:123456@192.168.2.60:5672/")
	if err != nil {
		fmt.Println("connect to message queue error.")
		panic(err)
	}

	ch, err := conn.Channel()
	_, err = ch.QueueDeclare(gs_commons_constants.MessageChannel, true,
		true, false, true, nil)
	if err != nil {
		panic(err)
	}

	errc := make(chan error, 1)

	messages, err := ch.Consume(gs_commons_constants.MessageChannel, "",
		false, false, false, false, nil)
	if err != nil {
		panic(err)
	}
	go func() {
		for m := range messages {
			var msg message.Message
			err := msgpack.Unmarshal(m.Body, &msg)
			println("receiver message")
			if err == nil {
				m.Acknowledger.Ack(m.DeliveryTag, false)
				switch m.Type {
				case "sms":
					service.SendSMS(msg.To, msg.Content)
					break
				case "email":
					service.SendEmail(msg.To, msg.Subject, msg.Content)
					break
				}
			}
		}
	}()

	<-errc
}
