package main

import (
	"fmt"
	"github.com/gorilla/websocket"
	"net/http"
	"net/url"
)

func main() {

	u := url.URL{Scheme: "ws", Host: "localhost:9000", Path: "/connection"}
	var dialer *websocket.Dialer

	println(u.String())

	header := http.Header{}

	header.Add("Authorization", "alskdfjalksdjf")

	conn, _, err := dialer.Dial(u.String(), header)
	if err != nil {
		fmt.Println(err)
		return
	}

	for {
		_, message, err := conn.ReadMessage()
		if err != nil {
			fmt.Println("read:", err)
			return
		}

		fmt.Printf("received: %s\n", message)
	}

}
