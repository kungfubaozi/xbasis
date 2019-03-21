package permission_handlers

import (
	"fmt"
	"github.com/garyburd/redigo/redis"
	"gopkg.in/mgo.v2"
	"konekko.me/gosion/commons/config"
	"konekko.me/gosion/commons/generator"
	"konekko.me/gosion/permission/repositories"
)

func Initialize(session *mgo.Session, pool *redis.Pool) gs_commons_config.OnConfigNodeChanged {
	return func(config *gs_commons_config.GosionInitializeConfig) {
		c, err := session.DB("gosion").C("functions").Count()
		if err != nil {
			fmt.Println("db err", err)
			return
		}
		if c == 0 {
			id := gs_commons_generator.ID()
			functionRepo := permission_repositories.FunctionRepo{Session: session, Conn: pool.Get()}
			defer functionRepo.Close()
			groupRepo := permission_repositories.GroupRepo{Session: session, ID: id}
			defer groupRepo.Close()
			roleRepo := permission_repositories.RoleRepo{Session: session, ID: id}
			defer roleRepo.Close()

		}
	}
}
