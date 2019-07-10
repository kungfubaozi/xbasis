package permissionhandlers

import (
	"context"
	"fmt"
	"github.com/samuel/go-zookeeper/zk"
	"gopkg.in/mgo.v2"
	xconfig "konekko.me/xbasis/commons/config"
	"konekko.me/xbasis/commons/constants"
	generator "konekko.me/xbasis/commons/generator"
	"konekko.me/xbasis/commons/indexutils"
)

func Initialize(session *mgo.Session, client *indexutils.Client, zk *zk.Conn) xconfig.OnConfigNodeChanged {
	return func(config *xconfig.GosionInitializeConfig) {
		db := session.DB(dbName)
		c, err := db.C(functionGroupCollection).Count()
		if err != nil {
			return
		}

		if c == 0 {

			bulk := client.GetElasticClient().Bulk()

			_, err := client.GetElasticClient().CreateIndex(functionIndex).BodyString(xbasisconstants.IndexMapping).Do(context.Background())
			if err != nil {
				panic(err)
			}
			_, err = client.GetElasticClient().CreateIndex(roleIndex).BodyString(xbasisconstants.IndexMapping).Do(context.Background())
			if err != nil {
				panic(err)
			}
			repo := &initializeRepo{session: session, conn: zk, bulk: bulk, config: config, id: generator.NewIDG()}

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
