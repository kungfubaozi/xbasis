package applicationhanderls

import (
	"context"
	"github.com/olivere/elastic"
	"gopkg.in/mgo.v2"
	config "konekko.me/xbasis/commons/config"
	constants "konekko.me/xbasis/commons/constants"
	"konekko.me/xbasis/commons/generator"
	"time"
)

type initializeRepo struct {
	session *mgo.Session
	bulk    *elastic.BulkService
	config  *config.GosionInitializeConfig
	id      xbasisgenerator.IDGenerator

	//data
	apps []interface{}
}

//user应用是用户基本操作应用，用户可以在此项目进行更改用户信息
func (repo *initializeRepo) AddUserApp() {
	app := repo.getApp(repo.config.UserAppId, "User", constants.AppTypeUser)
	app.Settings.AllowNewUsers = &allowNewUsersToEnter{
		DefaultRole:  []string{repo.config.UserAppRoleId},
		DefaultGroup: constants.AppUserGroup,
		Enabled:      true,
	}
	repo.AppServiceName(app)
	repo.bulk.Add(elastic.NewBulkIndexRequest().Index(applicationIndex).Type("_doc").Doc(app))
	repo.apps = append(repo.apps, app)
}

//不允许用户注册到此项目，邀请可忽略
func (repo *initializeRepo) AddManageApp() {
	app := repo.getApp(repo.config.AdminAppId, "Admin", constants.AppTypeManage)
	app.Settings.Quarantine = true
	app.Settings.RedirectURL = "http://192.168.80.67:9527"
	repo.AppServiceName(app)
	repo.bulk.Add(elastic.NewBulkIndexRequest().Index(applicationIndex).Type("_doc").Doc(app))
	repo.apps = append(repo.apps, app)
}

//route项目是用户默认项目，不需要同步，也不用隔离
func (repo *initializeRepo) AddRouteApp() {
	app := repo.getApp(repo.config.RouteAppId, "Router", constants.AppTypeRoute)
	app.Settings.AllowNewUsers = &allowNewUsersToEnter{
		DefaultGroup: constants.AppUserGroup,
		Enabled:      true,
	}
	repo.AppServiceName(app)
	repo.bulk.Add(elastic.NewBulkIndexRequest().Index(applicationIndex).Type("_doc").Doc(app))
	repo.apps = append(repo.apps, app)
}

//安全项目，不要同步数据，也不用隔离
func (repo *initializeRepo) AddSafeApp() {
	app := repo.getApp(repo.config.SafeAppId, "Safety", constants.AppTypeSafe)
	app.Settings.AllowNewUsers = &allowNewUsersToEnter{
		DefaultRole:  []string{repo.config.SafeAppRoleId},
		DefaultGroup: constants.AppUserGroup,
		Enabled:      true,
	}
	repo.AppServiceName(app)
	repo.bulk.Add(elastic.NewBulkIndexRequest().Index(applicationIndex).Type("_doc").Doc(app))
	repo.apps = append(repo.apps, app)
}

func (repo *initializeRepo) AppServiceName(info *appInfo) {
	info.Settings.ServiceName = "go.micro.api"
}

func (repo *initializeRepo) SaveAndClose() {
	defer repo.session.Close()
	if len(repo.apps) > 0 {
		check(repo.session.DB(dbName).C(applicationCollection).Insert(repo.apps...))
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
			Enabled:    constants.Enabled,
			Quarantine: false,
		},
		Clients: []*appClient{
			{
				Id:       repo.id.Short(),
				Platform: constants.PlatformOfWeb,
				Enabled:  constants.Enabled,
			},
			{
				Id:       repo.id.Short(),
				Platform: constants.PlatfromOfMacOS,
				Enabled:  constants.Closed,
			},
			{
				Id:       repo.id.Short(),
				Platform: constants.PlatformOfWindows,
				Enabled:  constants.Closed,
			},
			{
				Id:       repo.id.Short(),
				Platform: constants.PlatformOfIOS,
				Enabled:  constants.Closed,
			},
			{
				Id:       repo.id.Short(),
				Platform: constants.PlatformOfAndroid,
				Enabled:  constants.Closed,
			},
			{
				Id:       repo.id.Short(),
				Platform: constants.PlatformOfLinux,
				Enabled:  constants.Closed,
			},
			{
				Id:       repo.id.Short(),
				Platform: constants.PlatformOfFuchsia,
				Enabled:  constants.Closed,
			},
		},
	}
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}
