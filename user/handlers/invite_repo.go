package userhandlers

import "gopkg.in/mgo.v2"

type inviteRepo struct {
	session *mgo.Session
}

func (repo *inviteRepo) Add(model *inviteModel) error {
	return repo.collection().Insert(model)
}

func (repo *inviteRepo) IsExists(key, value string) (bool, error) {
	return false, nil
}

func (repo *inviteRepo) Cancel(id string) {

}

func (repo *inviteRepo) FindExpiredList(id string) {

}

func (repo *inviteRepo) collection() *mgo.Collection {
	return repo.session.DB(dbName).C(inviteCollection)
}

func (repo *inviteRepo) Close() {
	repo.session.Close()
}
