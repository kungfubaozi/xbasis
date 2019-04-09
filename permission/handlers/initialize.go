package permissionhandlers

import (
	"fmt"
	"github.com/garyburd/redigo/redis"
	"gopkg.in/mgo.v2"
	"konekko.me/gosion/commons/config"
	"konekko.me/gosion/commons/dao"
	"konekko.me/gosion/commons/generator"
)

func Initialize(session *mgo.Session, pool *redis.Pool) gs_commons_config.OnConfigNodeChanged {
	return func(config *gs_commons_config.GosionInitializeConfig) {
		c, err := session.DB(gs_commons_dao.DBName).C("functions").Count()
		if err != nil {
			fmt.Println("db err", err)
			return
		}
		if c == 0 {
			id := gs_commons_generator.NewIDG()
			functionRepo := functionRepo{session: session, conn: pool.Get()}
			defer functionRepo.Close()
			groupRepo := groupRepo{session: session, id: id}
			defer groupRepo.Close()
			roleRepo := roleRepo{session: session, id: id}
			defer roleRepo.Close()
		}
	}
}
