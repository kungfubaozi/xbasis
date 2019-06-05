package analysisservice

import (
	"konekko.me/gosion/analysis/client"
	"konekko.me/gosion/analysis/handlers"
	"konekko.me/gosion/analysis/pb"
	"konekko.me/gosion/commons/constants"
	"konekko.me/gosion/commons/indexutils"
	"konekko.me/gosion/commons/microservice"
)

func StartService() {
	errc := make(chan error, 1)

	client, err := indexutils.NewClient("http://192.168.2.62:9200/")
	if err != nil {
		panic(err)
	}

	logger := analysisclient.NewLoggerClient()

	go func() {
		s := microservice.NewService(gs_commons_constants.AnalysisService, true)
		s.Init()

		gs_service_analysis.RegisterLoggerHandler(s.Server(), analysishandlers.NewLoggerService(logger, client))

		errc <- s.Run()
	}()

	if err := <-errc; err != nil {
		panic(err)
	}
}
