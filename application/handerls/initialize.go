package applicationhanderls

import (
	"gopkg.in/mgo.v2"
	"konekko.me/gosion/commons/config"
	"konekko.me/gosion/commons/generator"
	"konekko.me/gosion/commons/indexutils"
)

func Initialize(session *mgo.Session, client *indexutils.Client) gs_commons_config.OnConfigNodeChanged {
	return func(config *gs_commons_config.GosionInitializeConfig) {
		defer session.Close()
		c := session.DB(dbName).C(applicationCollection)
		count, err := c.Count()
		if err != nil {
			return
		}
		//init
		if count == 0 {
			repo := &initializeRepo{session: session, id: gs_commons_generator.NewIDG(), bulk: client.GetElasticClient().Bulk(), config: config}
			repo.AddUserApp()
			repo.AddSafeApp()
			repo.AddRouteApp()
			repo.AddManageApp()
			repo.SaveAndClose()
		}
	}
}
