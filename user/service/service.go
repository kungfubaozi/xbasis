package usersvc

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"konekko.me/gosion/authentication/client"
	"konekko.me/gosion/commons/config"
	"konekko.me/gosion/commons/constants"
	"konekko.me/gosion/commons/dao"
	"konekko.me/gosion/commons/microservice"
	"konekko.me/gosion/message/cmd/messagecli"
	"konekko.me/gosion/safety/client"
	"konekko.me/gosion/user/handlers"
	"konekko.me/gosion/user/pb"
	"konekko.me/gosion/user/pb/ext"
)

func StartService() {

	errc := make(chan error, 2)

	session, err := gs_commons_dao.CreateSession("192.168.2.60:27017")
	if err != nil {
		panic(err)
	}
	defer session.Close()

	ms, err := messagecli.NewClient()
	if err != nil {
		panic(err)
	}
	defer ms.Close()

	db, err := gorm.Open("mysql", "root:root123@(192.168.2.60:3306)/gs_index?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	db.DB().SetMaxIdleConns(30)
	db.DB().SetMaxOpenConns(35)

	userhandlers.InitializeTables(db)

	configuration := &gs_commons_config.GosionConfiguration{}

	go func() {
		s := microservice.NewService(gs_commons_constants.ExtUserService, true)

		gs_ext_service_user.RegisterMessageHandler(s.Server(), userhandlers.NewMessageService(ms))

		gs_ext_service_user.RegisterUserHandler(s.Server(), userhandlers.NewExtUserService())

		errc <- s.Run()
	}()

	go func() {
		s := microservice.NewService(gs_commons_constants.UserService, true)

		ss := safetyclient.NewSecurityClient(s.Client())
		ts := authenticationcli.NewTokenClient(s.Client())

		gs_service_user.RegisterLoginHandler(s.Server(), userhandlers.NewLoginService(session, ss, ts, db))

		gs_service_user.RegisterRegisterHandler(s.Server(), userhandlers.NewRegisterService(session))

		gs_service_user.RegisterSafetyHandler(s.Server(), userhandlers.NewSafetyService(session))

		gs_service_user.RegisterUpdateHandler(s.Server(), userhandlers.NewUpdateService(session))

		errc <- s.Run()
	}()

	go func() {

		//watch config change to init def data
		gs_commons_config.WatchInitializeConfig(gs_commons_constants.UserService, userhandlers.Initialize(session.Clone(), db))

	}()

	go func() {
		gs_commons_config.WatchGosionConfig(func(config *gs_commons_config.GosionConfiguration) {
			configuration = config
		})
	}()

	if err := <-errc; err != nil {
		panic(err)
	}

}
