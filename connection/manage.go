package connection

import (
	"konekko.me/xbasis/connection/cmd/connectioncli"
)

type Message struct {
	Client  *UserClient
	Content map[string]interface{}
}

type ConnectManager struct {
	Clients    map[*UserClient]bool
	Message    chan *Message
	Register   chan *UserClient
	Unregister chan *UserClient
}

//manage connection
func (cm *ConnectManager) Process() {
	for {
		select {
		case conn := <-cm.Register:
			cm.Clients[conn] = true
			println("register", conn.Id)
			break
		case conn := <-cm.Unregister:
			if _, ok := cm.Clients[conn]; ok {
				close(conn.Send)
				delete(cm.Clients, conn)
				println("unregister", conn.Id)
			}
			break
		case msg := <-cm.Message:
			select {
			case msg.Client.Send <- msg.Content:
			default:
				close(msg.Client.Send)
				delete(cm.Clients, msg.Client)
			}
			break
		}
	}
}

type condition func(k *UserClient) bool

type clientEvent func(c *UserClient)

func (cm *ConnectManager) client(c condition, single bool, event clientEvent) {
	for k, v := range cm.Clients {
		if v && c(k) {
			event(k)
			if single {
				return
			}
		}
	}
}

func (cm *ConnectManager) OfflineToAppUser(msg *connectioncli.Message) {
	cm.client(func(k *UserClient) bool {
		return msg.TargetId == k.AppId && k.UserId == msg.UserId
	}, true, func(c *UserClient) {
		cm.Unregister <- c
		c.Conn.Close()
	})
}

func (cm *ConnectManager) OfflineToAppClientUser(msg *connectioncli.Message) {
	cm.client(func(k *UserClient) bool {
		return k.ClientId == msg.TargetId && k.UserId == msg.UserId
	}, true, func(c *UserClient) {
		cm.Unregister <- c
		c.Conn.Close()
	})
}

func (cm *ConnectManager) OfflineToUser(msg *connectioncli.Message) {
	cm.client(func(k *UserClient) bool {
		return k.UserId == msg.UserId
	}, true, func(c *UserClient) {
		cm.Unregister <- c
		c.Conn.Close()
	})
}

func (cm *ConnectManager) BroadcastToApp(msg *connectioncli.Message) {
	cm.client(func(k *UserClient) bool {
		return k.AppId == msg.TargetId
	}, false, func(c *UserClient) {
		cm.Message <- &Message{Client: c, Content: msg.Content}
	})
}

func (cm *ConnectManager) BroadcastToAppClient(msg *connectioncli.Message) {
	cm.client(func(k *UserClient) bool {
		return k.ClientId == msg.TargetId
	}, false, func(c *UserClient) {
		cm.Message <- &Message{Client: c, Content: msg.Content}
	})
}

func (cm *ConnectManager) BroadcastAll(msg *connectioncli.Message) {
	cm.client(func(k *UserClient) bool {
		return true
	}, false, func(c *UserClient) {
		cm.Message <- &Message{Client: c, Content: msg.Content}
	})
}

func (cm *ConnectManager) SendMessageToAppUser(msg *connectioncli.Message) {
	cm.client(func(k *UserClient) bool {
		return k.AppId == msg.TargetId && k.UserId == msg.UserId
	}, false, func(c *UserClient) {
		cm.Message <- &Message{Client: c, Content: msg.Content}
	})
}

func (cm *ConnectManager) SendMessageToAppClientUser(msg *connectioncli.Message) {
	cm.client(func(k *UserClient) bool {
		return k.ClientId == msg.TargetId && k.UserId == msg.UserId
	}, false, func(c *UserClient) {
		cm.Message <- &Message{Client: c, Content: msg.Content}
	})
}

func (cm *ConnectManager) SendMessageToUser(msg *connectioncli.Message) {
	cm.client(func(k *UserClient) bool {
		return k.UserId == msg.UserId
	}, false, func(c *UserClient) {
		cm.Message <- &Message{Client: c, Content: msg.Content}
	})
}
