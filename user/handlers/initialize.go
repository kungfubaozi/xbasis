package userhandlers

import (
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"gopkg.in/mgo.v2"
	"konekko.me/gosion/commons/config"
	"konekko.me/gosion/commons/hashcode"
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

			err = userRepo.AddUser(u)
			if err != nil {
				panic(err)
			}

			err = userRepo.AddUserInfo(info)
			if err != nil {
				panic(err)
			}

			_, err = client.AddData(fmt.Sprintf("gosion-index.users.%d", hashcode.Equa(u.Id)), map[string]interface{}{
				"index": map[string]interface{}{
					"name": "users",
					"id":   u.Id,
					"fields": map[string]interface{}{
						"username":  info.Username,
						"real_name": info.RealName,
						"phone":     u.Phone,
						"email":     u.Email,
						"user_id":   u.Id,
						"invite":    false,
						"account":   u.Account,
					},
				},
			})

			if err != nil {
				panic(err)
			}

			fmt.Println("user config initialize ok.")
		}
	}
}
