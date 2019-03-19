package permission_repositories

import (
	"github.com/bwmarrin/snowflake"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"time"
)

type GroupRepo struct {
	Session *mgo.Session
	ID      *snowflake.Node
}

func (repo *GroupRepo) FindByName(appId, name string) (*UserGroup, error) {
	var ug UserGroup
	err := repo.collection().Find(bson.M{"name": name, "appId": appId}).One(&ug)
	return &ug, err
}

func (repo *GroupRepo) Save(appId, userId, name string) error {
	ug := &UserGroup{
		CreateUserId: userId,
		CreateAt:     time.Now().UnixNano(),
		Name:         name,
		Id:           repo.ID.Generate().String(),
	}
	return repo.collection().Insert(ug)
}

func (repo *GroupRepo) collection() *mgo.Collection {
	return repo.Session.DB("gosion").C("user_groups")
}

func (repo *GroupRepo) Close() {
	repo.Session.Close()
}
