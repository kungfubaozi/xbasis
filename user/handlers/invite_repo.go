package userhandlers

import (
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type inviteRepo struct {
	session *mgo.Session
}

func (repo *inviteRepo) Add(model *inviteModel) error {
	return repo.collection().Insert(model)
}

func (repo *inviteRepo) FindByKey(key string, value interface{}) (*inviteModel, error) {
	var model *inviteModel
	err := repo.collection().Find(bson.M{key: value}).One(model)
	return model, err
}

func (repo *inviteRepo) UpdateItem(userId string, item *inviteItem) error {
	return repo.collection().Update(bson.M{"user_id": userId}, bson.M{"$push": bson.M{"items": item}})
}

func (repo *inviteRepo) SetState(userId, appId string, state int64) error {
	return repo.collection().Update(bson.M{"user_id": userId, "app_id": appId}, bson.M{"$set": bson.M{"state": state}})
}

func (repo *inviteRepo) collection() *mgo.Collection {
	return repo.session.DB(dbName).C(inviteCollection)
}

func (repo *inviteRepo) Close() {
	repo.session.Close()
}
