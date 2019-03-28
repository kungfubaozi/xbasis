package user_handlers

import (
	"golang.org/x/crypto/bcrypt"
	"gopkg.in/mgo.v2"
	"konekko.me/gosion/commons/config"
)

func Initialize(session *mgo.Session) gs_commons_config.OnConfigNodeChanged {
	return func(config *gs_commons_config.GosionInitializeConfig) {
		b, err := bcrypt.GenerateFromPassword([]byte("root123"), bcrypt.DefaultCost)
		if err != nil {
			panic(err)
		}

	}
}
