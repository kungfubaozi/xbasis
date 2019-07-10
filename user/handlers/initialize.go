package userhandlers

import (
	"context"
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"gopkg.in/mgo.v2"
	xconfig "konekko.me/xbasis/commons/config"
	"konekko.me/xbasis/commons/constants"
	"konekko.me/xbasis/commons/hashcode"
	"konekko.me/xbasis/commons/indexutils"
	"time"
)

func Initialize(session *mgo.Session, client *indexutils.Client) xconfig.OnConfigNodeChanged {
	return func(config *xconfig.GosionInitializeConfig) {
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

			index := fmt.Sprintf("xbs-index.users.%d", hashcode.Equa(u.Id))

			_, err = client.GetElasticClient().CreateIndex(index).BodyString(xbasisconstants.IndexMapping).Do(context.Background())
			if err != nil {
				panic(err)
			}

			_, err = client.AddDataById(u.Id, index, map[string]interface{}{
				"name":                     "users",
				"join_field":               "relation",
				"username":                 info.Username,
				"real_name":                info.RealName,
				"phone":                    u.Phone,
				"email":                    u.Email,
				"user_id":                  u.Id,
				"invite":                   false,
				"account":                  u.Account,
				"state":                    xbasisconstants.StateOk,
				"app_" + config.AdminAppId: true,
				"app_" + config.SafeAppId:  true,
				"app_" + config.RouteAppId: true,
				"app_" + config.UserAppId:  true,
			})

			if err != nil {
				panic(err)
			}

			fmt.Println("user config initialize ok.")
		}
	}
}
