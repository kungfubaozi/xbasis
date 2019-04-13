package applicationhanderls

import (
	"fmt"
	"gopkg.in/mgo.v2"
	"konekko.me/gosion/commons/config"
	"konekko.me/gosion/commons/constants"
	"konekko.me/gosion/commons/generator"
	"konekko.me/gosion/commons/indexutils"
	"time"
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
			repo := &applicationRepo{session: session, Client: client}
			defer repo.Close()
			id := gs_commons_generator.NewIDG()
			info := &appInfo{
				Name:         config.AppName,
				Desc:         config.Desc,
				Id:           config.AppId,
				CreateUserId: config.UserId,
				CreateAt:     time.Now().UnixNano(),
				Main:         101,
				UserS: &appStructure{
					Id:           config.UserS,
					LastUpdateAt: time.Now().UnixNano(),
					LastUpdateBy: config.UserId,
				},
				FunctionS: &appStructure{
					Id:           config.FuncS,
					LastUpdateAt: time.Now().UnixNano(),
					LastUpdateBy: config.UserId,
				},
				Settings: &appSetting{
					Enabled:  gs_commons_constants.Enabled,
					OpenMode: gs_commons_constants.OpenModeOfSelfOrganization,
				},
				Clients: []*appClient{
					{
						Id:       config.WebClientId,
						Platform: gs_commons_constants.PlatformOfWeb,
						Enabled:  gs_commons_constants.Enabled,
					},
					{
						Id:       id.Short(),
						Platform: gs_commons_constants.PlatfromOfMacOS,
						Enabled:  gs_commons_constants.Closed,
					},
					{
						Id:       id.Short(),
						Platform: gs_commons_constants.PlatformOfWindows,
						Enabled:  gs_commons_constants.Closed,
					},
					{
						Id:       id.Short(),
						Platform: gs_commons_constants.PlatformOfIOS,
						Enabled:  gs_commons_constants.Enabled,
					},
					{
						Id:       id.Short(),
						Platform: gs_commons_constants.PlatformOfAndroid,
						Enabled:  gs_commons_constants.Enabled,
					},
				},
			}
			err = repo.Add(info)
			if err != nil {
				fmt.Println("application initialize failed.")
			}
		}
	}
}
