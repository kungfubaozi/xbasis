package user_repositories

import (
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type UserRepo struct {
	Session *mgo.Session
}

func (repo *UserRepo) FindByContract(contract string) (*UserContractInfo, error) {
	info := &UserContractInfo{}
	err := repo.contractCollection().Find(bson.M{"contract": contract}).One(info)
	return info, err
}

func (repo *UserRepo) FindByAccount(account string) (*UserInfo, error) {
	info := &UserInfo{}
	err := repo.userCollection().Find(bson.M{"account": account}).One(info)
	return info, err
}

func (repo *UserRepo) FindById(id string) (*UserInfo, error) {
	info := &UserInfo{}
	err := repo.userCollection().Find(bson.M{"_id": id}).One(info)
	return info, err
}

func (repo *UserRepo) userCollection() *mgo.Collection {
	return repo.Session.DB("gosion").C("users")
}

func (repo *UserRepo) contractCollection() *mgo.Collection {
	return repo.Session.DB("gosion").C("user_contract")
}

func (repo *UserRepo) infoCollection() *mgo.Collection {
	return repo.Session.DB("gosion").C("user_info")
}

func (repo *UserRepo) Close() {
	repo.Session.Close()
}
