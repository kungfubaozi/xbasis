package userhandlers

import (
	"errors"
	"fmt"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"konekko.me/gosion/commons/encrypt"
	"konekko.me/gosion/commons/hashcode"
	"konekko.me/gosion/commons/indexutils"
)

type userRepo struct {
	session *mgo.Session
	*indexutils.Client
}

func (repo *userRepo) FindById(id string) (*userModel, error) {
	info := &userModel{}
	err := repo.userCollection(id).Find(bson.M{"_id": id}).One(info)
	return info, err
}

func (repo *userRepo) AddUser(user *userModel) error {
	return repo.userCollection(user.Id).Insert(user)
}

func (repo *userRepo) AddUserIndex(index *userModelIndex) error {
	id, err := repo.AddData(typeUserIndex, index)
	if err != nil {
		return nil
	}
	if len(id) > 0 {
		return nil
	}
	return errors.New("")
}

func (repo *userRepo) AddUserInfo(info *userInfo) error {
	return repo.infoCollection(info.UserId).Insert(info)
}

func (repo *userRepo) userCollection(userId string) *mgo.Collection {
	return repo.session.DB(dbName).C(fmt.Sprintf("%s_%d", userCollection, hashcode.Get(userId)))
}

func (repo *userRepo) oauthCollection(openId string) *mgo.Collection {
	return repo.session.DB(dbName).C(fmt.Sprintf("%s_%d", userOAuthCollection, hashcode.Get(openId)))
}

func (repo *userRepo) FindIndexTable(key string, content string) (string, error) {
	userIndex := &userModelIndex{}

	ok, err := repo.QueryFirst(typeUserIndex, map[string]interface{}{key: content}, &userIndex)
	if err != nil {
		return "", nil
	}
	if ok {
		return userIndex.UserId, nil
	}
	return "", indexutils.ErrNotFound
}

func (repo *userRepo) index(c string) string {
	return encrypt.SHA1(c)
}

func (repo *userRepo) infoCollection(userId string) *mgo.Collection {
	return repo.session.DB(dbName).C(fmt.Sprintf("%s_%d", userInfoCollection, hashcode.Get(userId)))
}

func (repo *userRepo) FindUserInfo(userId string) (*userInfo, error) {
	u := &userInfo{}
	err := repo.infoCollection(userId).Find(bson.M{"_id": userId}).One(u)
	return u, err
}

func (repo *userRepo) Close() {
	repo.session.Close()
}
