package permissionhandlers

import (
	"fmt"
	"github.com/garyburd/redigo/redis"
	"gopkg.in/mgo.v2"
	"konekko.me/gosion/commons/config"
	"konekko.me/gosion/commons/dao"
	"konekko.me/gosion/commons/generator"
	"konekko.me/gosion/permission/utils"
	"time"
)

func Initialize(session *mgo.Session, pool *redis.Pool) gs_commons_config.OnConfigNodeChanged {
	return func(config *gs_commons_config.GosionInitializeConfig) {
		db := session.DB(gs_commons_dao.DBName)
		c, err := db.C(gs_commons_dao.StructureCollection).Count()
		if err != nil {
			return
		}
		if c == 0 {
			id := gs_commons_generator.NewIDG()
			structureRepo := structureRepo{session: session, conn: pool.Get()}
			defer structureRepo.Close()

			defStructure := &structure{
				Id:           id.UUID(),
				CreateUserId: config.UserId,
				CreateAt:     time.Now().UnixNano(),
				AppId:        config.AppId,
				Opening:      true,
				Name:         "ROOT",
				Type:         permissionutils.TypeFunctionStructure,
			}

			//set current structure and open it
			err := structureRepo.Add(defStructure)
			if err != nil {
				fmt.Println("init def function structure err.", err)
				panic(err)
			}

			err = structureRepo.OpeningCache(defStructure.Id, defStructure.AppId, defStructure.Type)
			if err != nil {
				fmt.Println("open def function structure err.", err)
				panic(err)
			}

			defStructure.Type = permissionutils.TypeUserStructure
			defStructure.Id = id.UUID()
			err = structureRepo.Add(defStructure)
			if err != nil {
				fmt.Println("init def user structure err.", err)
				panic(err)
			}

			//open structure
			err = structureRepo.OpeningCache(defStructure.Id, defStructure.AppId, defStructure.Type)
			if err != nil {
				fmt.Println("open def user structure err.", err)
				panic(err)
			}

			fmt.Println("permission config init ok.")

		} else {
			fmt.Println("receiver init config, bug service already initialized.")
		}
	}
}
