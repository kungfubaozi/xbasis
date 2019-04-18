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
	"konekko.me/gosion/permission/utils"
	"time"
)

type functionsConfig struct {
	Version float32              `json:"version"`
	Desc    string               `json:"desc"`
	Data    []*functionGroupData `json:"api"`
	Roles   []string             `json:"roles"`
}

type functionGroupData struct {
	GroupName string          `json:"group_name"`
	Prefix    string          `json:"prefix"`
	Functions []*functionData `json:"functions"`
}

type functionData struct {
	Api      string   `json:"api"`
	Name     string   `json:"name"`
	AuthType []int64  `json:"auth_type"`
	Roles    []string `json:"roles"`
	Share    bool     `json:"share"`
}

type initializeRepo struct {
	session   *mgo.Session
	config    *gs_commons_config.GosionInitializeConfig
	structure *structureRepo
	id        gs_commons_generator.IDGenerator
	bulk      *elastic.BulkService
	// data
	userOrientate  *userOrientate
	userRoles      []interface{}
	functions      []interface{}
	functionGroups []interface{}
	structures     []interface{}

	//callback
}

func (repo *initializeRepo) AddManageApp() {
	config := repo.readFile("manage.json")
	mus := repo.buildStructure(repo.config.ManageAppId, permissionutils.TypeUserStructure)
	mfs := repo.buildStructure(repo.config.ManageAppId, permissionutils.TypeFunctionStructure)
	repo.generate(mfs.Id, config)
}

func (repo *initializeRepo) AddRouteApp() {
	config := repo.readFile("route.json")
	rus := repo.buildStructure(repo.config.RouteAppId, permissionutils.TypeUserStructure)
	rfs := repo.buildStructure(repo.config.RouteAppId, permissionutils.TypeFunctionStructure)
	repo.generate(rfs.Id, config)
}

func (repo *initializeRepo) AddSafeApp() {
	config := repo.readFile("safe.json")
	sus := repo.buildStructure(repo.config.SafeAppId, permissionutils.TypeUserStructure)
	sfs := repo.buildStructure(repo.config.SafeAppId, permissionutils.TypeFunctionStructure)
	repo.generate(sfs.Id, config)
}

func (repo *initializeRepo) AddUserApp() {
	config := repo.readFile("user.json")
	uus := repo.buildStructure(repo.config.UserAppId, permissionutils.TypeUserStructure)
	ufs := repo.buildStructure(repo.config.UserAppId, permissionutils.TypeFunctionStructure)
	repo.generate(ufs.Id, config)
}

func (repo *initializeRepo) SaveAndClose() {
	defer repo.session.Close()
	if repo.bulk != nil && repo.bulk.NumberOfActions() > 0 {
		db := repo.session.DB("gs_permission")
		if repo.userOrientate != nil && len(repo.userOrientate.LinkStructureRoles) > 0 {
			check(db.C(fmt.Sprintf("user_ort_%d", hashcode.Get(repo.config.UserId))).Insert(repo.userOrientate))
			repo.bulk.Add(elastic.NewBulkIndexRequest().Index("gs_user_ort").Type("v").Doc(repo.userOrientate))
		}

		if len(repo.userRoles) > 0 {
			check(db.C("user_roles").Insert(repo.userRoles...))
		}

		if len(repo.functions) > 0 {
			check(db.C("functions").Insert(repo.functions...))
		}

		if len(repo.structures) > 0 {
			check(db.C("structures").Insert(repo.structures...))
		}

		if len(repo.functionGroups) > 0 {
			check(db.C("function_groups").Insert(repo.functionGroups...))
		}

		ok, err := repo.bulk.Do(context.Background())
		check(err)
		if ok.Errors {
			panic("init failed.")
		}
	}
}

func (repo *initializeRepo) generate(functionStructureId string, config *functionsConfig) {
	roleMap := make(map[string]string)

	var adminRoles []string
	for _, v := range config.Roles {

		role := &role{
			Name:         v,
			Id:           repo.id.UUID(),
			CreateAt:     time.Now().UnixNano(),
			StructureId:  functionStructureId,
			CreateUserId: repo.config.UserId,
		}

		if v == "Administrator" || v == "User" {
			adminRoles = append(adminRoles, role.Id)
		}

		repo.bulk.Add(elastic.NewBulkIndexRequest().Index("gs_roles").Type("v").Doc(role))

		repo.userRoles = append(repo.userRoles, role)

		roleMap[role.Name] = role.Id
	}

	if len(adminRoles) > 0 {
		repo.userOrientate.LinkStructureRoles = append(repo.userOrientate.LinkStructureRoles, &linkStructureRole{
			StructureId: functionStructureId,
			Roles:       adminRoles,
		})
	}

	for _, v := range config.Data {

		g := &functionGroup{
			Id:           repo.id.UUID(),
			Name:         v.GroupName,
			CreateAt:     time.Now().UnixNano(),
			StructureId:  functionStructureId,
			CreateUserId: repo.config.UserId,
		}

		prefix := v.Prefix

		repo.functionGroups = append(repo.functionGroups, g)

		repo.bulk.Add(elastic.NewBulkIndexRequest().Index("gs_function_groups").Type("v").Doc(g))

		for _, v := range v.Functions {
			f := &function{
				Name:         v.Name,
				Api:          prefix + v.Api,
				AuthTypes:    v.AuthType,
				StructureId:  functionStructureId,
				BindGroupId:  g.Id,
				CreateAt:     time.Now().UnixNano(),
				CreateUserId: repo.config.UserId,
				Id:           repo.id.UUID(),
			}

			repo.functions = append(repo.functions, f)

			if v.Roles != nil && len(v.Roles) > 0 {
				var nr []string
				for _, r := range v.Roles {
					nr = append(nr, roleMap[r])
				}
				f.Roles = nr
			}

			repo.bulk.Add(elastic.NewBulkIndexRequest().Index("gs_functions").Type("v").Doc(f))
		}
	}
}

func (repo *initializeRepo) buildStructure(appId string, st int64) *structure {
	var name string
	if st == permissionutils.TypeUserStructure {
		name = "Users"
	} else {
		name = "Functions"
	}
	s := &structure{
		Id:           repo.id.UUID(),
		Type:         st,
		Name:         name,
		AppId:        appId,
		CreateUserId: repo.config.UserId,
		CreateAt:     time.Now().UnixNano(),
	}
	repo.structures = append(repo.structures, s)
	repo.bulk.Add(elastic.NewBulkIndexRequest().Index("gs_structures").Type("v").Doc(s))
	return s
}

func (repo *initializeRepo) readFile(file string) *functionsConfig {
	buffer, err := ioutil.ReadFile(file)
	if err != nil {
		panic(err)
	}

	var fc *functionsConfig

	check(json.Unmarshal(buffer, &fc))

	return fc
}
