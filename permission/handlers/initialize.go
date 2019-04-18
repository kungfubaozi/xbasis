package permissionhandlers

import (
	"fmt"
	"gopkg.in/mgo.v2"
	"konekko.me/gosion/commons/config"
	"konekko.me/gosion/commons/indexutils"
	"time"
)

func Initialize(session *mgo.Session, client *indexutils.Client) gs_commons_config.OnConfigNodeChanged {
	return func(config *gs_commons_config.GosionInitializeConfig) {
		db := session.DB(dbName)
		c, err := db.C(structureCollection).Count()
		if err != nil {
			return
		}

		if c == 0 {
			structureRepo := &structureRepo{session: session, Client: client}
			repo := initializeRepo{session: session, bulk: client.GetElasticClient().Bulk(), structure: structureRepo, userOrientate: &userOrientate{
				UserId:   config.UserId,
				CreateAt: time.Now().UnixNano(),
			}}
			repo.AddManageApp()
			repo.AddRouteApp()
			repo.AddSafeApp()
			repo.AddUserApp()
			repo.SaveAndClose()

		} else {
			fmt.Println("receiver init config, bug service already initialized.")
		}
	}
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}
