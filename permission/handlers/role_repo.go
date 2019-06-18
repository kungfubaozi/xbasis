package permissionhandlers

import (
	"github.com/garyburd/redigo/redis"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"konekko.me/gosion/commons/generator"
	"konekko.me/gosion/commons/indexutils"
	"konekko.me/gosion/permission/utils"
	"sync"
	"time"
)

type roleRepo struct {
	session *mgo.Session
	id      gs_commons_generator.IDGenerator
	conn    redis.Conn
	*indexutils.Client
}

func (repo *roleRepo) FindByName(appId, name string) (*role, error) {
	var r role
	err := repo.collection().Find(bson.M{"name": name, "app_id": appId}).One(&r)
	return &r, err
}

func (repo *roleRepo) Remove(appId, roleId string) error {
	names, err := repo.session.DB(dbName).CollectionNames()
	if err != nil {
		return err
	}
	var relations []string
	for _, v := range names {
		if len(v) > 19 && v[:19] == userRoleRelationCollection {
			relations = append(relations, v)
		}
	}
	var wg sync.WaitGroup
	wg.Add(len(relations) + 2)
	b := bson.M{"app_id": appId}
	u := bson.M{"$pull": bson.M{"roles": roleId}}
	resp := func(e error) {
		if err == nil {
			err = e
		}
	}

	for _, v := range relations {
		go func() {
			defer wg.Done()
			resp(repo.session.DB(dbName).C(v).Update(b, u))
		}()
	}

	go func() {
		defer wg.Done()

	}()

	go func() {
		defer wg.Done()
		resp(repo.collection().Remove(bson.M{"_id": roleId}))
	}()

	wg.Wait()

	return err

}

func (repo *roleRepo) Save(name, appId, userId string) error {
	r := &role{
		Id:           repo.id.Get(),
		CreateAt:     time.Now().UnixNano(),
		Name:         name,
		AppId:        appId,
		CreateUserId: userId,
	}
	err := repo.collection().Insert(r)
	if err != nil {
		return err
	}
	_, err = repo.conn.Do("hmset", permissionutils.GetAppRoleKey(appId), r.Id, r.Name)
	return err
}

func (repo *roleRepo) FindRoleById(roleId string) (*role, error) {
	var role *role
	err := repo.collection().Find(bson.M{"_id": roleId}).One(&role)
	return role, err
}

func (repo *roleRepo) FindRolesByAppId(appId string, page, size int64) ([]*role, error) {
	var roles []*role
	err := repo.collection().Find(bson.M{"app_id": appId}).All(&roles)
	return roles, err
}

func (repo *roleRepo) collection() *mgo.Collection {
	return repo.session.DB(dbName).C(roleCollection)
}

func (repo *roleRepo) Close() {
	repo.session.Close()
	if repo.conn != nil {
		repo.conn.Close()
	}
}
