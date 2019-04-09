package userhandlers

import (
	"golang.org/x/crypto/bcrypt"
	"gopkg.in/mgo.v2"
	"konekko.me/gosion/commons/config"
	"konekko.me/gosion/commons/dao"
	"time"
)

func Initialize(session *mgo.Session) gs_commons_config.OnConfigNodeChanged {
	return func(config *gs_commons_config.GosionInitializeConfig) {
		b, err := bcrypt.GenerateFromPassword([]byte(config.Password), bcrypt.DefaultCost)
		if err != nil {
			panic(err)
		}
		coll := session.DB(gs_commons_dao.DBName).C(gs_commons_dao.UserCollection)
		c, err := coll.Count()
		if err != nil {
			panic(err)
		}
		if c == 0 {
			//insert user
			u := &userInfo{
				Id:         config.UserId,
				CreateAt:   time.Now().UnixNano(),
				Account:    config.Username,
				Password:   string(b),
				RegisterAt: config.AppId,
			}
			err = coll.Insert(u)
			if err != nil {
				panic(err)
			}
		}
	}
}
