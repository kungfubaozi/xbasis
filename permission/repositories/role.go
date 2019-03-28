package permission_repositories

import (
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"konekko.me/gosion/commons/generator"
	"time"
)

type RoleRepo struct {
	Session *mgo.Session
	ID      gs_commons_generator.IDGenerator
}

func (repo *RoleRepo) FindByName(appId, name string) (*Role, error) {
	var r Role
	err := repo.collection().Find(bson.M{"name": name, "app_id": appId}).One(&r)
	return &r, err
}

func (repo *RoleRepo) Remove(appId, roleId string) error {
	return repo.collection().Remove(bson.M{"app_id": appId, "_id": roleId})
}

func (repo *RoleRepo) Save(name, appId, userId string) error {
	r := &Role{
		Id:           repo.ID.Get(),
		CreateAt:     time.Now().UnixNano(),
		Name:         name,
		CreateUserId: userId,
	}
	return repo.collection().Insert(r)
}

func (repo *RoleRepo) collection() *mgo.Collection {
	return repo.Session.DB("gosion").C("user_roles")
}

func (repo *RoleRepo) Close() {
	repo.Session.Close()
}
