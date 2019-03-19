package application_handerls

import (
	"fmt"
	"gopkg.in/mgo.v2"
	"konekko.me/gosion/application/repositories"
	"konekko.me/gosion/commons/config"
	"konekko.me/gosion/commons/constants"
	"konekko.me/gosion/commons/generator"
)

func Initialize(session *mgo.Session) gs_commons_config.OnConfigNodeChanged {
	return func(config *gs_commons_config.GosionInitializeConfig) {
		defer session.Close()
		c := session.DB("gosion").C("applications")
		count, err := c.Count()
		if err != nil {
			return
		}
		//init
		if count == 0 {
			id := gs_commons_generator.ID()
			info := &application_repositories.AppInfo{
				Name: config.AppName,
				Desc: config.Desc,
				Id:   id.Generate().String(),
				Settings: &application_repositories.AppSetting{
					Enabled:  gs_commons_constants.Enabled,
					OpenMode: gs_commons_constants.OpenModeOfSelfOrganization,
				},
				Clients: []*application_repositories.AppClient{
					{
						Id:       id.Generate().String(),
						Platform: gs_commons_constants.PlatformOfWeb,
						Enabled:  gs_commons_constants.Enabled,
					},
				},
			}
			err = c.Insert(info)
			if err != nil {
				fmt.Println("application initialize failed.")
			}
		}
	}
}
