package permissionhandlers

import (
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"konekko.me/gosion/commons/generator"
	"time"
)

type groupRepo struct {
	session *mgo.Session
	id      gs_commons_generator.IDGenerator
}

func (repo *groupRepo) FindByName(appId, name string) (*userGroup, error) {
	var ug userGroup
	err := repo.collection().Find(bson.M{"name": name, "appId": appId}).One(&ug)
	return &ug, err
}

func (repo *groupRepo) Save(appId, userId, name string) error {
	ug := &userGroup{
		CreateUserId: userId,
		CreateAt:     time.Now().UnixNano(),
		Name:         name,
		Id:           repo.id.Get(),
	}
	return repo.collection().Insert(ug)
}

func (repo *groupRepo) collection() *mgo.Collection {
	return repo.session.DB("gosion").C("user_groups")
}

func (repo *groupRepo) Close() {
	repo.session.Close()
}
