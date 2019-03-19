package main

import (
	"fmt"
	"github.com/samuel/go-zookeeper/zk"
	"gopkg.in/mgo.v2"
	"time"
)

var path = "/_gosion.test-13"

func main() {

	c, _, err := zk.Connect([]string{"192.168.2.57:2181"}, time.Second) //*10)
	if err != nil {
		panic(err)
	}

	acl := zk.WorldACL(zk.PermAll)
	_, err = c.Create(path, nil, int32(0), acl)
	if err != nil {
		fmt.Printf("create: %+v\n", err)
	}

	errc := make(chan error)

	go func() {
		for {
			_, _, ch, err := c.GetW(path)
			if err != nil {
				panic(err)
			}
			select {
			case e := <-ch:
				if e.Type == zk.EventNodeDataChanged {
					_, s, err := c.Get(path)
					if err != nil {
						fmt.Println("err", err)
					} else {
						fmt.Println(".", s.Version)
						if s.Version == 100 {
							errc <- mgo.ErrCursor
						}
					}
				}
			}
		}
	}()

	fmt.Println("ff")

	for i := 0; i <= 99; i++ {
		//fmt.Println("i", i)
		go func() {
			tryAgain(c, 0)
		}()
	}

	<-errc
}

func tryAgain(conn *zk.Conn, version int32) {
	_, err := conn.Set(path, []byte("update"), version)
	if err != nil && err == zk.ErrBadVersion {
		//fmt.Println("try again", version)
		time.Sleep(100)
		tryAgain(conn, version+1)
	}
}
