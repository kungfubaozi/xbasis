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
	"konekko.me/gosion/commons/config"
	"konekko.me/gosion/commons/constants"
	"konekko.me/gosion/commons/encrypt"
	"konekko.me/gosion/commons/generator"
	"konekko.me/gosion/commons/hashcode"
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
	Desc     string   `json:"desc"`
}

type initializeRepo struct {
	session *mgo.Session
	config  *gs_commons_config.GosionInitializeConfig
	id      gs_commons_generator.IDGenerator
	bulk    *elastic.BulkService
	conn    *zk.Conn
	// data
	userRolesRelation []interface{}
	userRoles         []interface{}
	functions         []interface{}
	functionGroups    []interface{}
	structures        []interface{}
	groupUsers        []interface{}
	groups            []interface{}
	//callback
}

func (repo *initializeRepo) AddManageApp() {
	config := repo.readFile("admin.json")
	repo.generate(repo.config.AdminAppId, config, false)
}

func (repo *initializeRepo) AddRouteApp() {
	config := repo.readFile("route.json")
	repo.generate(repo.config.RouteAppId, config, true)
}

func (repo *initializeRepo) AddSafeApp() {
	config := repo.readFile("safe.json")
	repo.generate(repo.config.SafeAppId, config, false)
}

func (repo *initializeRepo) AddUserApp() {
	config := repo.readFile("user.json")
	repo.generate(repo.config.UserAppId, config, false)
}

func (repo *initializeRepo) SaveAndClose() {
	defer repo.session.Close()
	if repo.bulk != nil && repo.bulk.NumberOfActions() > 0 {
		db := repo.session.DB("gs_permission")
		if repo.userRolesRelation != nil && len(repo.userRolesRelation) > 0 {
			check(db.C(fmt.Sprintf("user_roles_relation_%d", hashcode.Get(repo.config.UserId))).Insert(repo.userRolesRelation...))
		}

		if len(repo.userRoles) > 0 {
			check(db.C(roleCollection).Insert(repo.userRoles...))
		}

		if len(repo.functions) > 0 {
			check(db.C(functionCollection).Insert(repo.functions...))
		}

		if len(repo.functionGroups) > 0 {
			check(db.C(functionGroupCollection).Insert(repo.functionGroups...))
		}

		if len(repo.groupUsers) > 0 {
			check(db.C(fmt.Sprintf("%s_%d", groupUsersCollection, hashcode.Get("")%5)).Insert(repo.groupUsers...))
		}

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
		_, err = repo.conn.Create(gs_commons_constants.ZKAutonomyRegister, b, 0, acl)
		if err != nil {
			//panic(err)
		}
	}
}

var c *gs_commons_config.AutonomyRouteConfig

func (repo *initializeRepo) generate(appId string, config *functionsConfig, sync bool) {
	roleMap := make(map[string]string)

	var adminRoles []string
	for _, v := range config.Roles {

		role := &role{
			Name:         v,
			Id:           repo.id.UUID(),
			CreateAt:     time.Now().UnixNano(),
			AppId:        appId,
			CreateUserId: repo.config.UserId,
		}

		if v == "Administrator" || v == "User" {
			adminRoles = append(adminRoles, role.Id)
		}

		if v == "User" && sync {
			if c == nil {
				c = &gs_commons_config.AutonomyRouteConfig{
					AppId:  appId,
					RoleId: v,
				}
			}
		}

		if v == "User" && appId == repo.config.UserAppId {
			role.Id = repo.config.UserAppRoleId
		}

		repo.bulk.Add(elastic.NewBulkIndexRequest().Index("gs-roles").Type("_doc").Doc(role))

		repo.userRoles = append(repo.userRoles, role)

		roleMap[role.Name] = role.Id
	}

	if len(adminRoles) > 0 {

		u := &userRolesRelation{
			UserId:   repo.config.UserId,
			AppId:    appId,
			Roles:    adminRoles,
			CreateAt: time.Now().UnixNano(),
		}

		repo.userRolesRelation = append(repo.userRolesRelation, u)

		//		repo.bulk.Add(elastic.NewBulkIndexRequest().Index("gs-user-roles-relation").Type("_doc").Doc(u))

	}

	for _, v := range config.Data {

		g := &functionGroup{
			Id:           repo.id.UUID(),
			Name:         v.GroupName,
			CreateAt:     time.Now().UnixNano(),
			AppId:        appId,
			CreateUserId: repo.config.UserId,
		}

		prefix := v.Prefix

		repo.functionGroups = append(repo.functionGroups, g)

		repo.bulk.Add(elastic.NewBulkIndexRequest().Index("gs-function-groups").Type("_doc").Doc(g))

		for _, v := range v.Functions {
			f := &function{
				Name:         v.Name,
				Api:          prefix + v.Api,
				AuthTypes:    v.AuthType,
				AppId:        appId,
				BindGroupId:  g.Id,
				CreateAt:     time.Now().UnixNano(),
				CreateUserId: repo.config.UserId,
				Id:           repo.id.UUID(),
				Share:        v.Share,
			}

			if v.Roles != nil && len(v.Roles) > 0 {
				var nr []string
				for _, r := range v.Roles {
					id := roleMap[r]
					nr = append(nr, id)
					for _, v := range adminRoles {
						if id == v {
							dr := &directrelation{
								Function:   true,
								User:       true,
								UserId:     repo.config.UserId,
								FunctionId: f.Id,
								RoleId:     v,
								Enabled:    true,
							}

							id := encrypt.Md5(dr.FunctionId + dr.UserId)

							repo.bulk.Add(elastic.NewBulkIndexRequest().Index(fmt.Sprintf("gosion-urf-relations.%d", hashcode.Get(repo.config.UserId)%5)).Id(id).Type("_doc").Doc(dr))
						}
					}
				}
				f.Roles = nr
			}

			repo.functions = append(repo.functions, f)

			sf := &simplifiedFunction{
				Id:            f.Id,
				AuthTypes:     f.AuthTypes,
				Share:         f.Share,
				AppId:         appId,
				ValTokenTimes: f.ValTokenTimes,
				Roles:         f.Roles,
				Name:          f.Name,
				Path:          encrypt.SHA1(f.Api),
			}

			repo.bulk.Add(elastic.NewBulkIndexRequest().Index("gs-functions").Type("_doc").Doc(sf))
		}
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
