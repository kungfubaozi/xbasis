package permissionhandlers

import (
	"fmt"
	"gopkg.in/mgo.v2"
	"konekko.me/gosion/commons/config"
	"konekko.me/gosion/commons/generator"
	"konekko.me/gosion/commons/indexutils"
)

func Initialize(session *mgo.Session, client *indexutils.Client) gs_commons_config.OnConfigNodeChanged {
	return func(config *gs_commons_config.GosionInitializeConfig) {
		db := session.DB(dbName)
		c, err := db.C(functionGroupCollection).Count()
		if err != nil {
			return
		}

		if c == 0 {
			repo := &initializeRepo{session: session, bulk: client.GetElasticClient().Bulk(), config: config, id: gs_commons_generator.NewIDG()}
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
