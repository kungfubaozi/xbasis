package applicationhanderls

import (
	"errors"
	"github.com/garyburd/redigo/redis"
	"github.com/vmihailenco/msgpack"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"konekko.me/xbasis/commons/indexutils"
)

type applicationRepo struct {
	session *mgo.Session
	*indexutils.Client
	conn    redis.Conn
	clients map[string]string
}

var r *applicationRepo

func getApplicationRepo(session *mgo.Session, client *indexutils.Client, conn redis.Conn) *applicationRepo {

	if r != nil {
		r.session = session
		r.Client = client
		r.conn = conn
		return r
	}

	r = &applicationRepo{
		conn:    conn,
		clients: make(map[string]string),
		session: session,
		Client:  client,
	}

	return r
}

func (repo *applicationRepo) FindAll() ([]*appInfo, error) {
	var list []*appInfo
	err := repo.collection().Find(bson.M{}).All(&list)
	return list, err
}

func (repo *applicationRepo) findAppInfo(key, value string) (*appInfo, error) {
	var info appInfo
	ok, err := repo.QueryFirst(applicationIndex, map[string]interface{}{key: value}, &info)
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

func (repo *applicationRepo) RedirectUrlExists(url string) bool {
	c, err := repo.collection().Find(bson.M{"settings.redirect_url": url}).Count()
	if err != nil {
		return true
	}
	return c > 0
}

func (repo *applicationRepo) Add(info *appInfo) error {
	id, err := repo.AddData(applicationIndex, info)
	if err != nil {
		return err
	}
	if len(id) > 0 {
		info.SID = id
		err = repo.appendToCache(info)
		if err != nil {
			return err
		}
		return repo.collection().Insert(info)
	}
	return indexutils.ErrNotFound
}

func (repo *applicationRepo) FindByClientId(clientId string) (*appInfo, error) {
	appId := repo.clients[clientId]
	var info *appInfo
	if len(appId) != 0 {
		d, err := redis.Bytes(repo.conn.Do("get", appId))
		if err == nil {
			err = msgpack.Unmarshal(d, &info)
			return info, err
		}
	}
	i, err := repo.findAppInfo("clients.id", clientId)
	if err != nil {
		return nil, err
	}
	err = repo.appendToCache(i)
	return i, err
}

func (repo *applicationRepo) appendToCache(info *appInfo) error {
	b, err := msgpack.Marshal(info)
	if err != nil {
		return err
	}
	_, err = repo.conn.Do("set", info.Id, b)
	if err != nil {
		return err
	}
	for _, v := range info.Clients {
		repo.clients[v.Id] = info.Id
	}
	return nil
}

func (repo *applicationRepo) GetApplication(appId string) (*appInfo, error) {
	var appInfo *appInfo
	err := repo.collection().Find(bson.M{"_id": appId}).One(&appInfo)
	if err == nil && appInfo != nil {
		err = repo.appendToCache(appInfo)
		if err != nil {
			return nil, err
		}
	}
	return appInfo, err
}

func (repo *applicationRepo) collection() *mgo.Collection {
	return repo.session.DB(dbName).C(applicationCollection)
}

func (repo *applicationRepo) Close() {
	if repo.session != nil {
		repo.session.Clone()
	}
}
