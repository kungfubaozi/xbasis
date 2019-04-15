package main

import (
	"fmt"
	"konekko.me/gosion/commons/dao"
)

func main() {
	pool, err := gs_commons_dao.CreatePool("192.168.2.60:6379")
	if err != nil {
		panic(err)
	}
	defer pool.Close()

	conn := pool.Get()

	_, err = conn.Do("hset", "this", "a", "234")
	fmt.Println("err", err)
}
