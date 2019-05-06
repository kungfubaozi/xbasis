package main

import (
	"konekko.me/gosion/commons/dao"
	"konekko.me/gosion/commons/gslogrus"
	"konekko.me/gosion/commons/indexutils"
	"konekko.me/gosion/workflow/modules"
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

	flow := modules.NewService(session, pool, client, log)

	if err := flow.Run(); err != nil {
		panic(err)
	}
}
