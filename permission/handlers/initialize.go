package permissionhandlers

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/olivere/elastic"
	"gopkg.in/mgo.v2"
	"io/ioutil"
	"konekko.me/gosion/commons/config"
	"konekko.me/gosion/commons/generator"
	"konekko.me/gosion/commons/hashcode"
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
			structureRepo := structureRepo{session: session, Client: client}
			defer structureRepo.Close()

			functionStructureId := config.FuncS
			userStructureId := config.UserS

			defStructure := &structure{
				Id:           functionStructureId,
				CreateUserId: config.UserId,
				CreateAt:     time.Now().UnixNano(),
				AppId:        config.AppId,
				Name:         "Manage-UserS",
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
			defStructure.Name = "Manage-FuncS"
			err = structureRepo.Add(defStructure)
			if err != nil {
				fmt.Println("init def user structure err.", err)
				panic(err)
			}

			defRouteStructure := &structure{
				Id:           config.RouteFuncS,
				CreateAt:     time.Now().UnixNano(),
				CreateUserId: config.UserId,
				AppId:        config.RouteAppId,
				Name:         "WebRoute-FuncS",
				Type:         permissionutils.TypeFunctionStructure,
			}

			err = structureRepo.Add(defRouteStructure)
			if err != nil {
				fmt.Println("init def route function structure err.", err)
				panic(err)
			}

			defRouteStructure.Type = permissionutils.TypeUserStructure
			defRouteStructure.Id = config.RouteUserS
			defRouteStructure.Name = "WebRoute-UserS"
			err = structureRepo.Add(defRouteStructure)
			if err != nil {
				fmt.Println("init def route user structure err.", err)
				panic(err)
			}

			fuinit(functionStructureId, userStructureId, config.UserId, db, client)

		} else {
			fmt.Println("receiver init config, bug service already initialized.")
		}
	}
}

func fuinit(funcs string, userss, userId string, db *mgo.Database, client *indexutils.Client) {
	id := gs_commons_generator.NewIDG()

	buffer, err := ioutil.ReadFile("init.json")
	if err != nil {
		panic(err)
	}

	var fc *functionsConfig

	check(json.Unmarshal(buffer, &fc))

	b := client.GetElasticClient().Bulk()

	ug := &userGroup{
		Name:         "Manager",
		Id:           id.UUID(),
		CreateAt:     time.Now().UnixNano(),
		CreateUserId: userId,
		LinkStructureGroups: []*linkStructureGroup{
			{
				StructureId: userss,
			},
		},
	}

	check(db.C("user_groups").Insert(ug))
	b.Add(elastic.NewBulkIndexRequest().Index("gs_user_groups").Type("v").Doc(ug))

	uo := &userOrientate{
		UserId:   userId,
		CreateAt: time.Now().UnixNano(),
		LinkStructureGroups: []*linkStructureGroup{
			{
				StructureId: userss,
				BindGroupId: ug.Id,
			},
		},
		LinkStructureRoles: []*linkStructureRole{
			{
				StructureId: funcs,
			},
		},
	}

	var roles []interface{}
	roleMap := make(map[string]string)

	adminId := id.UUID()

	for _, v := range fc.Roles {

		role := &role{
			Name:         v,
			Id:           id.UUID(),
			CreateAt:     time.Now().UnixNano(),
			StructureId:  funcs,
			CreateUserId: userId,
		}

		if v == "Administrator" {
			role.Id = adminId
		}

		b.Add(elastic.NewBulkIndexRequest().Index("gs_roles").Type("v").Doc(role))

		roles = append(roles, role)

		roleMap[role.Name] = role.Id
	}

	if roles != nil && len(roles) > 0 {
		uo.LinkStructureRoles[0].Roles = []string{adminId}
		b.Add(elastic.NewBulkIndexRequest().Index("gs_user_ort").Type("v").Doc(uo))
		check(db.C("user_roles").Insert(roles...))
		check(db.C(fmt.Sprintf("user_ort_%d", hashcode.Get(uo.UserId))).Insert(uo))
	}

	//init functions
	for _, v := range fc.Data {

		g := &functionGroup{
			Id:           id.UUID(),
			Name:         v.GroupName,
			CreateAt:     time.Now().UnixNano(),
			StructureId:  funcs,
			CreateUserId: userId,
		}

		prefix := v.Prefix

		b.Add(elastic.NewBulkIndexRequest().Index("gs_function_groups").Type("v").Doc(g))

		functionsi := make([]interface{}, 1)
		for _, v := range v.Functions {
			f := &function{
				Name:         v.Name,
				Api:          prefix + v.Api,
				AuthTypes:    v.AuthType,
				StructureId:  funcs,
				BindGroupId:  g.Id,
				CreateAt:     time.Now().UnixNano(),
				CreateUserId: userId,
				Id:           id.UUID(),
			}
			functionsi = append(functionsi, f)

			if v.Roles != nil && len(v.Roles) > 0 {
				var nr []string
				for _, r := range v.Roles {
					nr = append(nr, roleMap[r])
				}
				f.Roles = nr
			}

			b.Add(elastic.NewBulkIndexRequest().Index("gs_functions").Type("v").Doc(f))
		}
		//add to database
		check(db.C("function_groups").Insert(g))
		check(db.C("functions").Insert(functionsi...))
	}

	ok, err := b.Do(context.Background())
	check(err)
	if ok.Errors {
		panic("add functions to es err.")
	}

}

func check(err error) {
	if err != nil {
		panic(err)
	}
}
