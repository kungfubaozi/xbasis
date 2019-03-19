package permission_repositories

import (
	"github.com/bwmarrin/snowflake"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"time"
)

type RoleRepo struct {
	Session *mgo.Session
	ID      *snowflake.Node
}

func (repo *RoleRepo) FindByName(appId, name string) (*UserRole, error) {
	var r UserRole
	err := repo.collection().Find(bson.M{"name": name, "app_id": appId}).One(&r)
	return &r, err
}

func (repo *RoleRepo) Save(name, appId, userId string) error {
	r := &UserRole{
		Id:           repo.ID.Generate().String(),
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
