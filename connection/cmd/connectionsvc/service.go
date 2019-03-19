package connectionsvc

import (
	"fmt"
	"github.com/gorilla/websocket"
	"github.com/streadway/amqp"
	"github.com/vmihailenco/msgpack"
	"konekko.me/gosion/commons/constants"
	"konekko.me/gosion/commons/generator"
	"konekko.me/gosion/connection"
	"konekko.me/gosion/connection/cmd/connectioncli"
	"net/http"
)

func StartService() {
	m := &connection.ConnectManager{
		Clients:    make(map[*connection.UserClient]bool),
		Message:    make(chan *connection.Message),
		Register:   make(chan *connection.UserClient),
		Unregister: make(chan *connection.UserClient),
	}

	conn, err := amqp.Dial("amqp://root:123456@192.168.2.60:5672/")
	if err != nil {
		fmt.Println("connect to message queue error.")
		panic(err)
	}
	defer conn.Close()
	ch, err := conn.Channel()
	defer conn.Close()

	err = ch.ExchangeDeclare(
		gs_commons_constants.ConnectionFanoutChannel,
		"fanout",
		true,
		false,
		false,
		false,
		nil,
	)

	if err != nil {
		panic(err)
	}

	q, err := ch.QueueDeclare(
		"",
		false,
		false,
		false,
		false,
		nil,
	)

	if err != nil {
		panic(err)
	}

	err = ch.QueueBind(
		q.Name,
		"",
		gs_commons_constants.ConnectionFanoutChannel,
		false,
		nil,
	)

	if err != nil {
		panic(err)
	}

	messages, err := ch.Consume(
		q.Name,
		"",
		true,
		false,
		false,
		false,
		nil,
	)

	if err != nil {
		panic(err)
	}

	go m.Process()

	http.HandleFunc("/connection", func(writer http.ResponseWriter, request *http.Request) {
		ws(m, writer, request)
	})

	errc := make(chan error, 1)

	go func() {
		errc <- http.ListenAndServe(":9000", nil)
	}()

	go func() {
		for msg := range messages {
			t := msg.Type
			var message connectioncli.Message
			err := msgpack.Unmarshal(msg.Body, &message)
			if err == nil {
				println("onMessage", message.Content)
				switch t {
				case "otau":
					go m.OfflineToAppUser(&message)
					break
				case "otu":
					go m.OfflineToUser(&message)
					break
				case "bta":
					go m.BroadcastToApp(&message)
					break
				case "baa":
					go m.BroadcastAll(&message)
					break
				case "stau":
					go m.SendMessageToAppUser(&message)
					break
				case "stu":
					go m.SendMessageToUser(&message)
					break
				case "otacu":
					go m.OfflineToAppClientUser(&message)
					break
				case "btac":
					go m.BroadcastToAppClient(&message)
					break
				case "smtcu":
					go m.SendMessageToAppClientUser(&message)
					break
				}
			}
		}
	}()

	//start
	if err := <-errc; err != nil {
		panic(err)
	}
}

func verify(r *http.Request) bool {
	auth := r.Header.Get("Authorization")
	println("receiver auth", auth)
	return true
}

func ws(m *connection.ConnectManager, w http.ResponseWriter, r *http.Request) {

	println("new connection")

	c, err := (&websocket.Upgrader{CheckOrigin: verify}).Upgrade(w, r, nil)
	if err != nil {
		http.NotFound(w, r)
		return
	}

	c.SetPingHandler(func(appData string) error {
		println(appData)
		return nil
	})

	uc := &connection.UserClient{
		Id:   gs_commons_generator.ID().Generate().String(),
		Send: make(chan map[string]interface{}),
		Conn: c,
	}

	//register user client
	m.Register <- uc

	go uc.Read(m)
	go uc.Write()
}
