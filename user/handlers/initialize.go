package userhandlers

import (
	"fmt"
	"github.com/olivere/elastic"
	"gopkg.in/mgo.v2"
	"konekko.me/gosion/commons/config"
	"time"
)

func Initialize(session *mgo.Session, client *elastic.Client) gs_commons_config.OnConfigNodeChanged {
	return func(config *gs_commons_config.GosionInitializeConfig) {
		coll := session.DB(dbName).C(userCollection)
		c, err := coll.Count()
		if err != nil {
			panic(err)
		}
		fmt.Println("receiver initialize config.")
		if c == 0 {
			//insert user

			userRepo := userRepo{session: session, elastic: client}
			defer userRepo.Close()

			u := &userInfo{
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
