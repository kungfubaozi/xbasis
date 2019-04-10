package userhandlers

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"konekko.me/gosion/commons/encrypt"
	"konekko.me/gosion/commons/hashcode"
)

type userRepo struct {
	session *mgo.Session
	db      *gorm.DB
}

func (repo *userRepo) FindById(id string) (*userInfo, error) {
	info := &userInfo{}
	err := repo.userCollection(id).Find(bson.M{"_id": id}).One(info)
	return info, err
}

func (repo *userRepo) AddUser(user *userInfo) error {
	index := &userIndex{TargetId: user.Id}
	//insert sql index
	if len(user.Email) > 0 {
		index.Type = emailIndexType
		index.Content = repo.index(user.Email)
		index.Code = hashcode.Get(index.Content)
		err := repo.db.Create(index)
		if err.Error != nil {
			return err.Error
		}
	}

	if len(user.Phone) > 0 {
		index.Type = phoneIndexType
		index.Content = repo.index(user.Phone)
		index.Code = hashcode.Get(index.Content)
		err := repo.db.Create(index)
		if err.Error != nil {
			return err.Error
		}
	}

	if len(user.Account) > 0 {
		index.Type = accountIndexType
		index.Content = repo.index(user.Account)
		index.Code = hashcode.Get(index.Content)
		err := repo.db.Create(index)
		if err.Error != nil {
			return err.Error
		}
	}

	//add collection

	return repo.userCollection(user.Id).Insert(user)
}

func (repo *userRepo) userCollection(userId string) *mgo.Collection {
	return repo.session.DB(dbName).C(fmt.Sprintf("%s_%d", userCollection, hashcode.Get(userId)))
}

func (repo *userRepo) oauthCollection(openId string) *mgo.Collection {
	return repo.session.DB(dbName).C(fmt.Sprintf("%s_%d", userOAuthCollection, hashcode.Get(openId)))
}

func (repo *userRepo) FindIndexTable(t int, content string) (string, error) {
	content = repo.index(content)
	userIndex := &userIndex{Code: hashcode.Get(content), Type: t}
	err := repo.db.Where(map[string]interface{}{"type": t, "content": content}).Find(userIndex)
	if err.Error != nil {
		return "", err.Error
	}
	return userIndex.TargetId, nil
}

func (repo *userRepo) index(c string) string {
	return gs_commons_encrypt.SHA1(c)
}

func (repo *userRepo) infoCollection(userId string) *mgo.Collection {
	return repo.session.DB(dbName).C(fmt.Sprintf("%s_%d", userInfoCollection, hashcode.Get(userId)))
}

func (repo *userRepo) Close() {
	repo.session.Close()
}
