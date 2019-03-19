package application_repositories

import (
	"github.com/garyburd/redigo/redis"
	"github.com/globalsign/mgo/bson"
	"github.com/vmihailenco/msgpack"
	"gopkg.in/mgo.v2"
	"konekko.me/gosion/commons/encrypt"
)

type ApplicationRepo struct {
	Session *mgo.Session
	Conn    redis.Conn
}

func (repo *ApplicationRepo) FindAll() ([]*AppInfo, error) {
	var list []*AppInfo
	err := repo.collection().Find(bson.M{}).All(&list)
	return list, err
}

func (repo *ApplicationRepo) FindByApplicationId(appId string) (*AppInfo, error) {
	var app AppInfo
	err := repo.collection().Find(bson.M{"_id": appId}).One(&app)
	return &app, err
}

func (repo *ApplicationRepo) GetApplicationInCache(appId string) (*AppInfo, error) {
	b, err := redis.Bytes(repo.Conn.Do("get", gs_commons_encrypt.SHA1("app-"+appId)))
	if err != nil {
		return nil, err
	}
	var info AppInfo
	err = msgpack.Unmarshal(b, &info)
	return &info, err
}

func (repo *ApplicationRepo) ApplicationExists(name string) bool {
	c, err := repo.collection().Find(bson.M{"name": name}).Count()
	if err != nil {
		return true
	}
	return c > 0
}

func (repo *ApplicationRepo) Upsert(info *AppInfo) error {
	_, err := repo.collection().Upsert(bson.M{"_id": info.Id}, bson.M{"$set": info})
	if err != nil {
		return err
	}
	b, err := msgpack.Marshal(info)
	if err != nil {
		return err
	}
	_, err = repo.Conn.Do("set", gs_commons_encrypt.SHA1("app-"+info.Id), b)
	return err
}

func (repo *ApplicationRepo) FindByClientId(clientId string) (*AppInfo, error) {

	var app AppInfo
	err := repo.collection().Find(bson.M{"clients": bson.M{"$elemMatch": bson.M{"id": clientId}}}).One(&app)

	return &app, err
}

func (repo *ApplicationRepo) collection() *mgo.Collection {
	return repo.Session.DB("gosion").C("applications")
}

func (repo *ApplicationRepo) Close() {
	repo.Conn.Close()
	repo.Session.Close()
}
