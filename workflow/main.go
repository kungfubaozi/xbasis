package main

import (
	"fmt"
	"github.com/davecgh/go-spew/spew"
	"github.com/vmihailenco/msgpack"
	"konekko.me/gosion/commons/dao"
	"konekko.me/gosion/commons/encrypt"
	"konekko.me/gosion/commons/gslogrus"
	"konekko.me/gosion/commons/indexutils"
	"konekko.me/gosion/workflow/runtime"
	"time"
)

func main() {
	session, err := gs_commons_dao.CreateSession("192.168.2.60:27017")
	if err != nil {
		panic(err)
	}
	defer session.Close()

	pool, err := gs_commons_dao.CreatePool("192.168.2.60:6379")
	if err != nil {
		panic(err)
	}
	defer pool.Close()

	client, err := indexutils.NewClient("http://192.168.2.62:9200/")
	if err != nil {
		panic(err)
	}

	log := gslogrus.New("gs.service.flow", client)

	flow := runtime.NewService(session, pool, client, log)

	if err := flow.Run(); err != nil {
		panic(err)
	}

	//test()

}

func test() {

	s := time.Now().UnixNano()

	key := "6333614dc0c7452eb3b29bed26a8580a"

	//v := map[string]interface{}{
	//	"userId":       "askdfjakl;sdjfasdf",
	//	"username":     "1233454634",
	//	"type":         343,
	//	"createUserId": "293483098523094584329058",
	//	"createAt":     time.Now().UnixNano(),
	//}
	//
	//b, err := msgpack.Marshal(v)
	//if err != nil {
	//	panic(err)
	//}
	//
	//c, err := encrypt.AESEncrypt(b, []byte(key))
	//if err != nil {
	//	panic(err)
	//}
	//
	//fmt.Println(c)

	var data map[string]interface{}
	a, err := encrypt.AESDecrypt("s/kym+2gARNcksJ2azZDwCUFVA04FxAHIHJbaQie3oTpIf3Jm6F8IYFfa4YJLLsRfE1eiJAhMK9F3VivMhK/IItaYdH4DqcuT3HhF3o8sznGVonKMVTYmIWrtFlYKAELAZ9elYgDY2lFdYrJ6tDih1lD6f6V3+pyK+SuF3HHjhfaQ/iDgyOS4oW1VTuxI8h9", []byte(key))
	if err != nil {
		panic(err)
	}

	err = msgpack.Unmarshal([]byte(a), &data)
	if err != nil {
		panic(err)
	}

	spew.Dump(data)

	fmt.Println((time.Now().UnixNano() - s) / 1e6)
}
