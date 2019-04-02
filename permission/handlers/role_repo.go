package permissionhandlers

import (
	"github.com/garyburd/redigo/redis"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"konekko.me/gosion/commons/generator"
	"konekko.me/gosion/permission/utils"
	"time"
)

type roleRepo struct {
	session *mgo.Session
	id      gs_commons_generator.IDGenerator
	conn    redis.Conn
}

func (repo *roleRepo) FindByName(appId, name string) (*role, error) {
	var r role
	err := repo.collection().Find(bson.M{"name": name, "app_id": appId}).One(&r)
	return &r, err
}

func (repo *roleRepo) Exists(appId, roleId string) (bool, error) {
	return redis.Bool(repo.conn.Do("hexists", permissionutils.GetStructureRoleKey(appId), roleId))
}

func (repo *roleRepo) GetUserRoleMembers(appId, userId string) ([]interface{}, error) {
	return redis.Values(repo.conn.Do("SMEMBERS", permissionutils.GetStructureUserRoleKey(appId, userId)))
}

func (repo *roleRepo) Remove(appId, roleId string) error {
	_, err := repo.conn.Do("hdel", permissionutils.GetStructureRoleKey(appId), roleId)
	if err != nil && err == redis.ErrNil {
		err = nil
	}
	if err != nil {
		return err
	}
	return repo.collection().Remove(bson.M{"app_id": appId, "_id": roleId})
}

func (repo *roleRepo) Save(name, appId, userId string) error {
	r := &role{
		Id:           repo.id.Get(),
		CreateAt:     time.Now().UnixNano(),
		Name:         name,
		CreateUserId: userId,
	}
	err := repo.collection().Insert(r)
	if err != nil {
		return err
	}
	_, err = repo.conn.Do("hmset", permissionutils.GetStructureRoleKey(appId), r.Id, r.Name)
	return err
}

func (repo *roleRepo) collection() *mgo.Collection {
	return repo.session.DB("gosion").C("user_roles")
}

func (repo *roleRepo) Close() {
	repo.session.Close()
	repo.conn.Close()
}
