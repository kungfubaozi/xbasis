package connection

import "github.com/gorilla/websocket"

type UserClient struct {
	Id       string
	UserId   string
	AppId    string
	ClientId string
	TokenKey string
	Send     chan map[string]interface{}
	Conn     *websocket.Conn
}

func (uc *UserClient) Read(m *ConnectManager) {
	defer func() {
		m.Unregister <- uc
		uc.Conn.Close()
	}()

	for {
		_, _, err := uc.Conn.ReadMessage()
		if err != nil {
			m.Unregister <- uc
			uc.Conn.Close()
			break
		}
	}
}

func (uc *UserClient) Write() {
	defer func() {
		uc.Conn.Close()
	}()

	for {
		select {
		case message, ok := <-uc.Send:
			if !ok {
				uc.Conn.WriteMessage(websocket.CloseMessage, []byte{})
				return
			}

			uc.Conn.WriteJSON(message)
		}
	}
}
