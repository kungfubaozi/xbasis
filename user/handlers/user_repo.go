package userhandlers

import (
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"konekko.me/gosion/commons/dao"
)

type userRepo struct {
	session *mgo.Session
}

func (repo *userRepo) FindByContract(contract string) (*userContractInfo, error) {
	info := &userContractInfo{}
	err := repo.contractCollection().Find(bson.M{"contract": contract}).One(info)
	return info, err
}

func (repo *userRepo) FindByAccount(account string) (*userInfo, error) {
	info := &userInfo{}
	err := repo.userCollection().Find(bson.M{"account": account}).One(info)
	return info, err
}

func (repo *userRepo) FindById(id string) (*userInfo, error) {
	info := &userInfo{}
	err := repo.userCollection().Find(bson.M{"_id": id}).One(info)
	return info, err
}

func (repo *userRepo) userCollection() *mgo.Collection {
	return repo.session.DB(gs_commons_dao.DBName).C(gs_commons_dao.UserCollection)
}

func (repo *userRepo) contractCollection() *mgo.Collection {
	return repo.session.DB(gs_commons_dao.DBName).C(gs_commons_dao.UserContractCollection)
}

func (repo *userRepo) infoCollection() *mgo.Collection {
	return repo.session.DB(gs_commons_dao.DBName).C("user_info_")
}

func (repo *userRepo) Close() {
	repo.session.Close()
}
