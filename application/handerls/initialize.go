package applicationhanderls

import (
	"gopkg.in/mgo.v2"
	config "konekko.me/xbasis/commons/config"
	"konekko.me/xbasis/commons/generator"
	"konekko.me/xbasis/commons/indexutils"
)

func Initialize(session *mgo.Session, client *indexutils.Client) config.OnConfigNodeChanged {
	return func(config *config.GosionInitializeConfig) {
		defer session.Close()
		c := session.DB(dbName).C(applicationCollection)
		count, err := c.Count()
		if err != nil {
			return
		}
		//init
		if count == 0 {
			repo := &initializeRepo{session: session, id: xbasisgenerator.NewIDG(), bulk: client.GetElasticClient().Bulk(), config: config}
			repo.AddUserApp()
			repo.AddSafeApp()
			repo.AddRouteApp()
			repo.AddManageApp()
			repo.SaveAndClose()
		}
	}
}
