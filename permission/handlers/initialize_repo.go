package permissionhandlers

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/olivere/elastic"
	"github.com/samuel/go-zookeeper/zk"
	"github.com/vmihailenco/msgpack"
	"gopkg.in/mgo.v2"
	"io/ioutil"
	xconfig "konekko.me/xbasis/commons/config"
	constants "konekko.me/xbasis/commons/constants"
	"konekko.me/xbasis/commons/encrypt"
	generator "konekko.me/xbasis/commons/generator"
	"konekko.me/xbasis/commons/hashcode"
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
	Api           string   `json:"api"`
	Name          string   `json:"name"`
	AuthType      []int64  `json:"auth_type"`
	ValTokenTimes int64    `json:"val_token_times"`
	Roles         []string `json:"roles"`
	Share         bool     `json:"share"`
	Desc          string   `json:"desc"`
}

type initializeRepo struct {
	session *mgo.Session
	config  *xconfig.GosionInitializeConfig
	id      generator.IDGenerator
	bulk    *elastic.BulkService
	conn    *zk.Conn
	// data
	userRolesRelation map[int][]interface{}
	userRoles         map[int][]interface{}
	functions         map[int][]interface{}
	functionGroups    map[int][]interface{}
	functionRoles     map[int][]interface{}
	groupUsers        map[int][]interface{}
	groups            map[int][]interface{}
	adminRoles        []string
	//callback
}

func (repo *initializeRepo) AddManageApp() {
	config := repo.readFile("admin.json")
	repo.generate(repo.config.AdminAppId, config, true)
	repo.AddUGRelation(repo.config.AdminAppId)
}

func (repo *initializeRepo) AddRouteApp() {
	config := repo.readFile("route.json")
	repo.generate(repo.config.RouteAppId, config, true)
	repo.AddUGRelation(repo.config.RouteAppId)
}

func (repo *initializeRepo) AddSafeApp() {
	config := repo.readFile("safe.json")
	repo.generate(repo.config.SafeAppId, config, true)
	repo.AddUGRelation(repo.config.SafeAppId)
}

func (repo *initializeRepo) AddUserApp() {
	config := repo.readFile("user.json")
	repo.generate(repo.config.UserAppId, config, true)
	repo.AddUGRelation(repo.config.UserAppId)
}

func (repo *initializeRepo) AddUGRelation(appId string) {

	if repo.groupUsers == nil {
		repo.groupUsers = make(map[int][]interface{})
	}

	repo.groupUsers[hashcode.Equa(appId)] = append(repo.groupUsers[hashcode.Equa(appId)], &userGroupsRelation{
		AppId:       appId,
		BindGroupId: []string{constants.AppMainStructureGroup},
		CreateAt:    time.Now().UnixNano(),
		UserId:      repo.config.UserId,
	})

	repo.groupUsers[hashcode.Equa(appId)] = append(repo.groupUsers[hashcode.Equa(appId)], &userGroupsRelation{
		AppId:       appId,
		BindGroupId: []string{constants.AppUserGroup},
		CreateAt:    time.Now().UnixNano(),
		UserId:      repo.config.UserId,
	})
}

func (repo *initializeRepo) SaveAndClose() {
	defer repo.session.Close()
	if repo.bulk != nil && repo.bulk.NumberOfActions() > 0 {
		db := repo.session.DB(dbName)
		if repo.userRolesRelation != nil && len(repo.userRolesRelation) > 0 {
			for k, v := range repo.userRolesRelation {
				check(db.C(fmt.Sprintf("%s_%d", userRoleRelationCollection, k)).Insert(v...))
			}
		}

		if len(repo.userRoles) > 0 {
			for k, v := range repo.userRoles {
				check(db.C(fmt.Sprintf("%s_%d", roleCollection, k)).Insert(v...))
			}
		}

		if len(repo.functionRoles) > 0 {
			for k, v := range repo.functionRoles {
				check(db.C(fmt.Sprintf("%s_%d", functionRoleRelationCollection, k)).Insert(v...))
			}
		}

		if len(repo.functions) > 0 {
			for k, v := range repo.functions {
				check(db.C(fmt.Sprintf("%s_%d", functionCollection, k)).Insert(v...))
			}
		}

		if len(repo.functionGroups) > 0 {
			for k, v := range repo.functionGroups {
				check(db.C(fmt.Sprintf("%s_%d", functionGroupCollection, k)).Insert(v...))
			}
		}

		if len(repo.groupUsers) > 0 {
			for k, v := range repo.groupUsers {
				check(db.C(fmt.Sprintf("%s_%d", groupUsersCollection, k)).Insert(v...))
			}
		}

		index := fmt.Sprintf("xbs-index.users.%d", hashcode.Equa(repo.config.UserId))
		repo.bulk.Add(elastic.NewBulkIndexRequest().Index(index).Type("_doc").Id(encrypt.Md5(repo.config.UserId + index)).Routing(repo.config.UserId).Doc(getRolesRelation(repo.config.UserId, repo.adminRoles)))

		ok, err := repo.bulk.Do(context.Background())
		check(err)
		if ok.Errors {
			panic("init failed.")
		}
	}
	if c != nil {
		b, err := msgpack.Marshal(c)
		if err != nil {
			panic(err)
		}
		acl := zk.WorldACL(zk.PermAll)
		_, err = repo.conn.Create(constants.ZKAutonomyRegister, b, 0, acl)
		if err != nil {
			//panic(err)
		}
	}
}

var c *xconfig.AutonomyRouteConfig

func (repo *initializeRepo) generate(appId string, config *functionsConfig, sync bool) {
	roleMap := make(map[string]*roleIndexModel)

	var adminRoles []string
	for _, v := range config.Roles {

		role := &role{
			Name:         v,
			Id:           repo.id.Get(),
			CreateAt:     time.Now().UnixNano(),
			AppId:        appId,
			CreateUserId: repo.config.UserId,
		}

		var ru int64 = 0

		if v == "Administrator" || v == "User" {
			adminRoles = append(adminRoles, role.Id)
			ru = 1
		}

		if v == "User" && sync {
			if c == nil {
				c = &xconfig.AutonomyRouteConfig{
					AppId:  appId,
					RoleId: v,
				}
			}
		}

		if v == "User" && appId == repo.config.UserAppId {
			role.Id = repo.config.UserAppRoleId
		}

		if v == "User" && appId == repo.config.SafeAppId {
			role.Id = repo.config.SafeAppRoleId
		}

		if repo.userRoles == nil {
			repo.userRoles = make(map[int][]interface{})
		}

		v1 := repo.userRoles[hashcode.Equa(appId)]

		repo.userRoles[hashcode.Equa(appId)] = append(v1, role)

		roleMap[role.Name] = &roleIndexModel{
			Name:          v,
			Id:            role.Id,
			CreateAt:      role.CreateAt,
			AppId:         appId,
			CreateUserId:  repo.config.UserId,
			RelationUsers: ru,
		}
	}

	if len(adminRoles) > 0 {

		u := &userRolesRelation{
			UserId:   repo.config.UserId,
			AppId:    appId,
			Roles:    adminRoles,
			CreateAt: time.Now().UnixNano(),
		}

		if repo.userRolesRelation == nil {
			repo.userRolesRelation = make(map[int][]interface{})
		}

		v1 := repo.userRolesRelation[hashcode.Equa(repo.config.UserId)]

		repo.userRolesRelation[hashcode.Equa(repo.config.UserId)] = append(v1, u)

		repo.adminRoles = append(repo.adminRoles, adminRoles...)
	}

	for _, v := range config.Data {

		g := &functionGroup{
			Id:           repo.id.Get(),
			Name:         v.GroupName,
			CreateAt:     time.Now().UnixNano(),
			AppId:        appId,
			CreateUserId: repo.config.UserId,
		}

		prefix := v.Prefix

		if repo.functionGroups == nil {
			repo.functionGroups = make(map[int][]interface{})
		}

		v1 := repo.functionGroups[hashcode.Equa(appId)]

		repo.functionGroups[hashcode.Equa(appId)] = append(v1, g)

		for _, v := range v.Functions {
			f := &function{
				Name:         v.Name,
				Api:          prefix + v.Api,
				AuthTypes:    v.AuthType,
				AppId:        appId,
				BindGroupId:  g.Id,
				CreateAt:     time.Now().UnixNano(),
				CreateUserId: repo.config.UserId,
				Id:           repo.id.Get(),
				Share:        v.Share,
			}

			for _, v1 := range v.AuthType {
				if v1 == constants.AuthTypeOfValcode {
					f.ValTokenTimes = 1
				}
			}

			var roles []string

			if v.Roles != nil && len(v.Roles) > 0 {
				var nr []string

				for _, r := range v.Roles {
					r := roleMap[r]
					nr = append(nr, r.Id)
					r.RelationFunctions = r.RelationFunctions + 1
				}
				if repo.functionRoles == nil {
					repo.functionRoles = make(map[int][]interface{})
				}
				v := repo.functionRoles[hashcode.Equa(f.Id)]

				fr := &functionRolesRelation{
					FunctionId: f.Id,
					CreateAt:   time.Now().UnixNano(),
					Roles:      nr,
					AppId:      appId,
				}

				repo.functionRoles[hashcode.Equa(f.Id)] = append(v, fr)

				roles = nr
			}

			if repo.functions == nil {
				repo.functions = make(map[int][]interface{})
			}

			v1 := repo.functions[hashcode.Equa(f.AppId)]

			repo.functions[hashcode.Equa(f.AppId)] = append(v1, f)

			sf := &SimplifiedFunction{
				JoinField: "relation",
				Id:        f.Id,
				Share:     f.Share,
				AppId:     appId,
				Name:      f.Name,
				Path:      f.Api,
			}

			repo.bulk.Add(elastic.NewBulkIndexRequest().Index(functionIndex).Id(f.Id).Type("_doc").Doc(sf))

			repo.bulk.Add(elastic.NewBulkIndexRequest().Index(functionIndex).Type("_doc").Id(encrypt.Md5(f.Id + functionIndex)).Routing(f.Id).Doc(getRolesRelation(f.Id, roles)))

		}
	}
	for _, v := range roleMap {
		repo.bulk.Add(elastic.NewBulkIndexRequest().Index(roleIndex).Type("_doc").Doc(v))
	}
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
