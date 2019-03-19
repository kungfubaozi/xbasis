package user_handlers

import (
	"gopkg.in/mgo.v2"
	"konekko.me/gosion/commons/config"
)

func Initialize(session *mgo.Session) gs_commons_config.OnConfigNodeChanged {
	return func(config *gs_commons_config.GosionInitializeConfig) {

	}
}
