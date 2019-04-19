package applicationhanderls

import (
	"context"
	"github.com/olivere/elastic"
	"gopkg.in/mgo.v2"
	"konekko.me/gosion/commons/config"
	"konekko.me/gosion/commons/constants"
	"konekko.me/gosion/commons/generator"
	"time"
)

type initializeRepo struct {
	session *mgo.Session
	bulk    *elastic.BulkService
	config  *gs_commons_config.GosionInitializeConfig
	id      gs_commons_generator.IDGenerator

	//data
	apps []interface{}
}

func (repo *initializeRepo) AddUserApp() {
	app := repo.getApp(repo.config.UserAppId, "Gsuser", gs_commons_constants.AppTypeUser)
	repo.setStructure(repo.config.UserAppFSId, repo.config.UserAppUSId, app)
	repo.bulk.Add(elastic.NewBulkIndexRequest().Index("gs-applications").Type("_doc").Doc(app))
	repo.apps = append(repo.apps, app)
}

func (repo *initializeRepo) AddManageApp() {
	app := repo.getApp(repo.config.ManageAppId, "Gsmanage", gs_commons_constants.AppTypeManage)
	repo.setStructure(repo.config.ManageFSId, repo.config.ManageUSId, app)
	repo.bulk.Add(elastic.NewBulkIndexRequest().Index("gs-applications").Type("_doc").Doc(app))
	repo.apps = append(repo.apps, app)
}

func (repo *initializeRepo) AddRouteApp() {
	app := repo.getApp(repo.config.RouteAppId, "Gsroute", gs_commons_constants.AppTypeRoute)
	repo.setStructure(repo.config.RouteAppFSId, repo.config.RouteAppUSId, app)
	repo.bulk.Add(elastic.NewBulkIndexRequest().Index("gs-applications").Type("_doc").Doc(app))
	repo.apps = append(repo.apps, app)
}

func (repo *initializeRepo) AddSafeApp() {
	app := repo.getApp(repo.config.SafeAppId, "Gssafe", gs_commons_constants.AppTypeSafe)
	repo.setStructure(repo.config.SafeAppFSId, repo.config.SafeAppUSId, app)
	repo.bulk.Add(elastic.NewBulkIndexRequest().Index("gs-applications").Type("_doc").Doc(app))
	repo.apps = append(repo.apps, app)
}

func (repo *initializeRepo) SaveAndClose() {
	defer repo.session.Close()
	if len(repo.apps) > 0 {
		check(repo.session.DB("gs_applications").C("applications").Insert(repo.apps...))
		ok, err := repo.bulk.Do(context.Background())
		check(err)
		if ok.Errors {
			panic("application init failed.")
		}
	}
}

func (repo *initializeRepo) setStructure(funcS, userS string, info *appInfo) {
	info.FunctionS = &appStructure{
		Id:           funcS,
		LastUpdateBy: repo.config.UserId,
		LastUpdateAt: time.Now().UnixNano(),
	}
	info.UserS = &appStructure{
		Id:           userS,
		LastUpdateAt: time.Now().UnixNano(),
		LastUpdateBy: repo.config.UserId,
	}
}

func (repo *initializeRepo) getApp(id, name string, appType int64) *appInfo {
	return &appInfo{
		Name:         name,
		Id:           id,
		CreateAt:     time.Now().UnixNano(),
		CreateUserId: repo.config.UserId,
		Type:         appType,
		Settings: &appSetting{
			Enabled: gs_commons_constants.Enabled,
		},
		Clients: []*appClient{
			{
				Id:       repo.id.Short(),
				Platform: gs_commons_constants.PlatformOfWeb,
				Enabled:  gs_commons_constants.Enabled,
			},
			{
				Id:       repo.id.Short(),
				Platform: gs_commons_constants.PlatfromOfMacOS,
				Enabled:  gs_commons_constants.Closed,
			},
			{
				Id:       repo.id.Short(),
				Platform: gs_commons_constants.PlatformOfWindows,
				Enabled:  gs_commons_constants.Closed,
			},
			{
				Id:       repo.id.Short(),
				Platform: gs_commons_constants.PlatformOfIOS,
				Enabled:  gs_commons_constants.Closed,
			},
			{
				Id:       repo.id.Short(),
				Platform: gs_commons_constants.PlatformOfAndroid,
				Enabled:  gs_commons_constants.Closed,
			},
			{
				Id:       repo.id.Short(),
				Platform: gs_commons_constants.PlatformOfLinux,
				Enabled:  gs_commons_constants.Closed,
			},
			{
				Id:       repo.id.Short(),
				Platform: gs_commons_constants.PlatformOfFuchsia,
				Enabled:  gs_commons_constants.Closed,
			},
		},
	}
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}
