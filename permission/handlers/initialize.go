package permissionhandlers

import (
	"fmt"
	"github.com/samuel/go-zookeeper/zk"
	"gopkg.in/mgo.v2"
	xconfig "konekko.me/xbasis/commons/config"
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
			repo := &initializeRepo{session: session, conn: zk, bulk: client.GetElasticClient().Bulk(), config: config, id: generator.NewIDG()}
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
