package permissionhandlers

import (
	"fmt"
	"gopkg.in/mgo.v2"
	"konekko.me/gosion/commons/config"
	"konekko.me/gosion/commons/constants/api"
	"konekko.me/gosion/commons/generator"
	"konekko.me/gosion/commons/indexutils"
	"konekko.me/gosion/permission/utils"
	"time"
)

func Initialize(session *mgo.Session, client *indexutils.Client) gs_commons_config.OnConfigNodeChanged {
	return func(config *gs_commons_config.GosionInitializeConfig) {
		db := session.DB(dbName)
		c, err := db.C(structureCollection).Count()
		if err != nil {
			return
		}
		if c == 0 {
			id := gs_commons_generator.NewIDG()
			structureRepo := structureRepo{session: session, Client: client}
			defer structureRepo.Close()

			functionStructureId := config.FuncS
			userStructureId := config.UserS

			defStructure := &structure{
				Id:           functionStructureId,
				CreateUserId: config.UserId,
				CreateAt:     time.Now().UnixNano(),
				AppId:        config.AppId,
				Name:         "ROOT",
				Type:         permissionutils.TypeFunctionStructure,
			}

			//set current structure and open it
			err := structureRepo.Add(defStructure)
			if err != nil {
				fmt.Println("init def function structure err.", err)
				panic(err)
			}

			defStructure.Type = permissionutils.TypeUserStructure
			defStructure.Id = userStructureId
			defStructure.SID = ""
			err = structureRepo.Add(defStructure)
			if err != nil {
				fmt.Println("init def user structure err.", err)
				panic(err)
			}

			//add def functions
			functionRepo := functionRepo{session: session, Client: client}
			defer functionRepo.Close()

			userGroupId := id.UUID()
			err = functionRepo.AddGroup(&functionGroup{
				Id:           userGroupId,
				Name:         "User",
				CreateUserId: config.UserId,
				CreateAt:     time.Now().UnixNano(),
				StructureId:  functionStructureId,
			})

			if err != nil {
				panic(err)
			}

			f := &function{
				Name:         "LoginWithAccount",
				Id:           id.UUID(),
				CreateAt:     time.Now().UnixNano(),
				CreateUserId: config.UserId,
				BindGroupId:  userGroupId,
				Api:          gosionapis.LoginWithAccount,
				StructureId:  functionStructureId,
				AuthTypes:    []int64{},
			}

			err = functionRepo.AddFunction(f)
			if err != nil {
				panic(err)
			}

			fmt.Println("permission config init ok.")

		} else {
			fmt.Println("receiver init config, bug service already initialized.")
		}
	}
}
