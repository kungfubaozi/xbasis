package xbasisconfig

import (
	"fmt"
	"github.com/samuel/go-zookeeper/zk"
	"github.com/vmihailenco/msgpack"
	constants "konekko.me/xbasis/commons/constants"
	"konekko.me/xbasis/commons/encrypt"
	"time"
)

type GosionInitializeConfig struct {
	UserId     string
	Phone      string
	AppName    string
	Desc       string
	Username   string
	Email      string
	Account    string
	SecretKey  string
	UserAppId  string
	RouteAppId string
	SafeAppId  string
	AdminAppId string

	UserAppRoleId string
	SafeAppRoleId string
}

type OnNodeDataChanged func(data []byte, version int32) bool

type OnConfigNodeChanged func(config *GosionInitializeConfig)

func WatchInitializeConfig(serviceName string, event OnConfigNodeChanged) {
	c := NewConnect("192.168.2.57:2181")

	path := "/_gosion.init.locking-" + encrypt.SHA1(serviceName)

	//one same service process
	acl := zk.WorldACL(zk.PermAll)
	_, err := c.Create(path, nil, 1, acl)
	if err != nil { //delete and recreate
		//fmt.Println("end")
		//return
	}

	fmt.Println("start watch:", serviceName)

	//c.Delete(gs_commons_constants.ZKWatchInitializeConfigPath, 0)

	watch(c, constants.ZKWatchInitializeConfigPath, func(data []byte, version int32) bool {
		var config GosionInitializeConfig
		err := msgpack.Unmarshal(data, &config)
		if err != nil {
			//continue monitoring
			return false
		}
		event(&config)
		c.Delete(path, 0)
		//tryAgain(serviceName, c, 0)
		return true //stop monitoring
	})
}

func NewConnect(url string) *zk.Conn {
	c, _, err := zk.Connect([]string{url}, time.Second)
	if err != nil {
		panic(err)
	}
	return c
}

func watch(c *zk.Conn, path string, event OnNodeDataChanged) {
	if event == nil {
		panic("invalid config listener event")
	}

	acl := zk.WorldACL(zk.PermAll)
	_, err := c.Create(path, nil, int32(0), acl)
	if err != nil {
		fmt.Printf("create: %+v\n", err)
	}
	fmt.Println("listen start:", path)
	_, _, ch, err := c.GetW(path)
	if err != nil {
		panic(err)
	}
	for {
		select {
		case e := <-ch:
			if e.Type == zk.EventNodeDataChanged {
				fmt.Println("changed")
				v, s, err := c.Get(path)
				if err != nil {
					fmt.Println("err", err)
				} else {
					if event(v, s.Version) {
						return
					}
				}
			}
		}
	}
}
