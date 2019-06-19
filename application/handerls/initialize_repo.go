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

//user应用是用户基本操作应用，用户可以在此项目进行更改用户信息
func (repo *initializeRepo) AddUserApp() {
	app := repo.getApp(repo.config.UserAppId, "Gsuser", gs_commons_constants.AppTypeUser)
	app.Settings.Authorize = true
	repo.bulk.Add(elastic.NewBulkIndexRequest().Index("gs-applications").Type("_doc").Doc(app))
	repo.apps = append(repo.apps, app)
}

//不允许用户注册到此项目，邀请可忽略
func (repo *initializeRepo) AddManageApp() {
	app := repo.getApp(repo.config.AdminAppId, "Gsadmin", gs_commons_constants.AppTypeManage)
	app.Settings.Quarantine = true
	app.Settings.Authorize = true
	app.Settings.RedirectURL = "http://localhost:9527"
	repo.bulk.Add(elastic.NewBulkIndexRequest().Index("gs-applications").Type("_doc").Doc(app))
	repo.apps = append(repo.apps, app)
}

//route项目是用户默认项目，不需要同步，也不用隔离
func (repo *initializeRepo) AddRouteApp() {
	app := repo.getApp(repo.config.RouteAppId, "Gsrouter", gs_commons_constants.AppTypeRoute)
	repo.bulk.Add(elastic.NewBulkIndexRequest().Index("gs-applications").Type("_doc").Doc(app))
	repo.apps = append(repo.apps, app)
}

//安全项目，不要同步数据，也不用隔离
func (repo *initializeRepo) AddSafeApp() {
	app := repo.getApp(repo.config.SafeAppId, "Gssafety", gs_commons_constants.AppTypeSafe)
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

func (repo *initializeRepo) getApp(id, name string, appType int64) *appInfo {
	return &appInfo{
		Name:         name,
		Id:           id,
		CreateAt:     time.Now().UnixNano(),
		CreateUserId: repo.config.UserId,
		Type:         appType,
		Settings: &appSetting{
			Enabled:    gs_commons_constants.Enabled,
			Quarantine: false,
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
