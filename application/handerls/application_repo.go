package applicationhanderls

import (
	"errors"
	"fmt"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"konekko.me/gosion/commons/indexutils"
)

type applicationRepo struct {
	session *mgo.Session
	*indexutils.Client
}

func (repo *applicationRepo) FindAll() ([]*appInfo, error) {
	var list []*appInfo
	err := repo.collection().Find(bson.M{}).All(&list)
	return list, err
}

func (repo *applicationRepo) findAppInfo(key, value string) (*appInfo, error) {
	fmt.Println("find app info")

	var info appInfo
	ok, err := repo.QueryFirst("gs_applications", map[string]interface{}{key: value}, &info)
	if err != nil {
		return nil, err
	}
	if ok {
		return &info, nil
	}
	return nil, errors.New("not found")
}

func (repo *applicationRepo) FindByApplicationId(appId string) (*appInfo, error) {
	return repo.findAppInfo("id", appId)
}

func (repo *applicationRepo) ApplicationExists(name string) bool {
	c, err := repo.collection().Find(bson.M{"name": name}).Count()
	if err != nil {
		return true
	}
	return c > 0
}

func (repo *applicationRepo) Add(info *appInfo) error {
	id, err := repo.AddData("gs_applications", info)
	if err != nil {
		return err
	}
	if len(id) > 0 {
		info.SID = id
		return repo.collection().Insert(info)
	}
	return indexutils.ErrNotFound
}

func (repo *applicationRepo) FindByClientId(clientId string) (*appInfo, error) {
	return repo.findAppInfo("clients.id", clientId)
}

func (repo *applicationRepo) collection() *mgo.Collection {
	return repo.session.DB(dbName).C(applicationCollection)
}

func (repo *applicationRepo) Close() {
	if repo.session != nil {
		repo.session.Clone()
	}
}
