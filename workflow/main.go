package main

import (
	"konekko.me/xbasis/analysis/client"
	constants "konekko.me/xbasis/commons/constants"
	dao "konekko.me/xbasis/commons/dao"
	generator "konekko.me/xbasis/commons/generator"
	"konekko.me/xbasis/commons/indexutils"
	"konekko.me/xbasis/commons/microservice"
	"konekko.me/xbasis/workflow/handlers"
	pb "konekko.me/xbasis/workflow/pb"
	"konekko.me/xbasis/workflow/runtime"
)

func main() {

	session, err := dao.CreateSession("192.168.2.60:27017")
	if err != nil {
		panic(err)
	}
	defer session.Close()

	pool, err := dao.CreatePool("192.168.2.60:6379")
	if err != nil {
		panic(err)
	}
	defer pool.Close()

	client, err := indexutils.NewClient("http://192.168.2.62:9200/")
	if err != nil {
		panic(err)
	}

	log := analysisclient.NewLoggerClient()

	flow := runtime.NewService(session, pool, client, log)

	id := generator.NewIDG()

	errc := make(chan error, 2)

	go func() {
		errc <- flow.Run("192.168.2.57:2181")
	}()

	go func() {
		s := microservice.NewService(constants.WorkflowService, false)

		pb.RegisterProcessHandler(s.Server(), workflowhandlers.NewProcessService(flow.Modules(), id, log))

		errc <- s.Run()
	}()

	<-errc

}
