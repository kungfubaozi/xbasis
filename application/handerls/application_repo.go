package applicationhanderls

import (
	"github.com/garyburd/redigo/redis"
	"github.com/vmihailenco/msgpack"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"konekko.me/gosion/commons/dao"
	"konekko.me/gosion/commons/encrypt"
)

type applicationRepo struct {
	session *mgo.Session
	conn    redis.Conn
}

func (repo *applicationRepo) FindAll() ([]*appInfo, error) {
	var list []*appInfo
	err := repo.collection().Find(bson.M{}).All(&list)
	return list, err
}

func (repo *applicationRepo) FindByApplicationId(appId string) (*appInfo, error) {
	var app appInfo
	err := repo.collection().Find(bson.M{"_id": appId}).One(&app)
	return &app, err
}

func (repo *applicationRepo) GetApplicationInCache(appId string) (*appInfo, error) {
	b, err := redis.Bytes(repo.conn.Do("get", gs_commons_encrypt.SHA1("app-"+appId)))
	if err != nil {
		return nil, err
	}
	var info appInfo
	err = msgpack.Unmarshal(b, &info)
	return &info, err
}

func (repo *applicationRepo) ApplicationExists(name string) bool {
	c, err := repo.collection().Find(bson.M{"name": name}).Count()
	if err != nil {
		return true
	}
	return c > 0
}

func (repo *applicationRepo) Upsert(info *appInfo) error {
	_, err := repo.collection().Upsert(bson.M{"_id": info.Id}, bson.M{"$set": info})
	if err != nil {
		return err
	}
	b, err := msgpack.Marshal(info)
	if err != nil {
		return err
	}
	_, err = repo.conn.Do("set", gs_commons_encrypt.SHA1("app-"+info.Id), b)
	return err
}

func (repo *applicationRepo) FindByClientId(clientId string) (*appInfo, error) {
	var app appInfo
	err := repo.collection().Find(bson.M{"clients": bson.M{"$elemMatch": bson.M{"id": clientId}}}).One(&app)

	return &app, err
}

func (repo *applicationRepo) collection() *mgo.Collection {
	return repo.session.DB(gs_commons_dao.DBName).C(gs_commons_dao.ApplicationCollection)
}

func (repo *applicationRepo) Close() {
	repo.conn.Close()
	repo.session.Close()
}
