package userhandlers

import (
	"fmt"
	"gopkg.in/mgo.v2"
	"konekko.me/gosion/commons/config"
	"konekko.me/gosion/commons/indexutils"
	"time"
)

func Initialize(session *mgo.Session, client *indexutils.Client) gs_commons_config.OnConfigNodeChanged {
	return func(config *gs_commons_config.GosionInitializeConfig) {
		coll := session.DB(dbName).C(userCollection)
		c, err := coll.Count()
		if err != nil {
			panic(err)
		}
		fmt.Println("receiver initialize config.")
		if c == 0 {
			//insert user

			userRepo := userRepo{session: session, Client: client}
			defer userRepo.Close()

			u := &userModel{
				Id:         config.UserId,
				CreateAt:   time.Now().UnixNano(),
				Account:    config.Username,
				Password:   config.Password,
				RegisterAt: config.WebClientId,
				Phone:      config.Phone,
				Email:      config.Email,
			}
			err = userRepo.AddUser(u)
			if err != nil {
				panic(err)
			}

			fmt.Println("user config initialize ok.")
		}
	}
}
