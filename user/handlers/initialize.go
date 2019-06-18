package userhandlers

import (
	"fmt"
	"golang.org/x/crypto/bcrypt"
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

			b, err := bcrypt.GenerateFromPassword([]byte("admin123"), bcrypt.DefaultCost)
			if err != nil {
				panic(err)
			}

			u := &userModel{
				Id:         config.UserId,
				CreateAt:   time.Now().UnixNano(),
				Account:    config.Account,
				Password:   string(b),
				RegisterAt: "",
				Phone:      config.Phone,
				Email:      config.Email,
			}

			info := &userInfo{
				UserId:   u.Id,
				Username: config.Username,
				CreateAt: time.Now().UnixNano(),
			}

			index := &userModelIndex{
				Username: info.Username,
				Phone:    u.Phone,
				Email:    u.Email,
				UserId:   u.Id,
				Account:  u.Account,
			}

			err = userRepo.AddUserIndex(index)
			if err != nil {
				panic(err)
			}

			err = userRepo.AddUser(u)
			if err != nil {
				panic(err)
			}

			err = userRepo.AddUserInfo(info)
			if err != nil {
				panic(err)
			}

			fmt.Println("user config initialize ok.")
		}
	}
}
