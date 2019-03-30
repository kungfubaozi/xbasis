package permission_repositories

import (
	"github.com/garyburd/redigo/redis"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"konekko.me/gosion/commons/generator"
	"konekko.me/gosion/permission/uitls"
	"time"
)

type RoleRepo struct {
	Session *mgo.Session
	ID      gs_commons_generator.IDGenerator
	Conn    redis.Conn
}

func (repo *RoleRepo) FindByName(appId, name string) (*Role, error) {
	var r Role
	err := repo.collection().Find(bson.M{"name": name, "app_id": appId}).One(&r)
	return &r, err
}

func (repo *RoleRepo) Exists(appId, roleId string) (bool, error) {
	return redis.Bool(repo.Conn.Do("hexists", permission_uitls.GetAppRoleKey(appId), roleId))
}

func (repo *RoleRepo) GetUserRoleMembers(appId, userId string) ([]interface{}, error) {
	return redis.Values(repo.Conn.Do("SMEMBERS", permission_uitls.GetAppUserRoleKey(appId, userId)))
}

func (repo *RoleRepo) Remove(appId, roleId string) error {
	_, err := repo.Conn.Do("hdel", permission_uitls.GetAppRoleKey(appId), roleId)
	if err != nil && err == redis.ErrNil {
		err = nil
	}
	if err != nil {
		return err
	}
	return repo.collection().Remove(bson.M{"app_id": appId, "_id": roleId})
}

func (repo *RoleRepo) Save(name, appId, userId string) error {
	r := &Role{
		Id:           repo.ID.Get(),
		CreateAt:     time.Now().UnixNano(),
		Name:         name,
		CreateUserId: userId,
	}
	err := repo.collection().Insert(r)
	if err != nil {
		return err
	}
	_, err = repo.Conn.Do("hmset", permission_uitls.GetAppRoleKey(appId), r.Id, r.Name)
	return err
}

func (repo *RoleRepo) collection() *mgo.Collection {
	return repo.Session.DB("gosion").C("user_roles")
}

func (repo *RoleRepo) Close() {
	repo.Session.Close()
}
