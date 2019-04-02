package permissionhandlers

import (
	"github.com/garyburd/redigo/redis"
	"github.com/vmihailenco/msgpack"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"konekko.me/gosion/commons/encrypt"
)

type functionRepo struct {
	session *mgo.Session
	conn    redis.Conn
}

func (repo *functionRepo) AddFunction(function *function) error {
	b, err := msgpack.Marshal(function)
	if err != nil {
		return err
	}
	_, err = repo.conn.Do("set", gs_commons_encrypt.SHA1(function.Api+function.AppId), b)
	if err != nil {
		return err
	}
	return repo.functionCollection().Insert(function)
}

func (repo *functionRepo) AddGroup(group *functionGroup) error {
	return repo.groupCollection().Insert(group)
}

func (repo *functionRepo) FindGroup(appId, name string) error {
	return nil
}

func (repo *functionRepo) FindChildGroups(parentId string) ([]*functionGroup, error) {
	var groups []*functionGroup
	err := repo.groupCollection().Find(bson.M{"bind_group_id": parentId}).All(&groups)
	return groups, err
}

func (repo *functionRepo) FindChildFunctions(parentId string) ([]*function, error) {
	var functions []*function
	err := repo.functionCollection().Find(bson.M{"bind_group_id": parentId}).All(&functions)
	return functions, err
}

func (repo *functionRepo) FindApi(appId, api string) (*function, error) {
	var f function
	err := repo.functionCollection().Find(bson.M{"app_id": appId, "api": api}).One(&f)
	return &f, err
}

func (repo *functionRepo) FindGroupExists(groupId string) bool {
	c, err := repo.groupCollection().Find(bson.M{"_id": groupId}).Count()
	if err != nil {
		return true
	}
	return c > 0
}

func (repo *functionRepo) FindApiInCache(appId, api string) (*function, error) {
	b, err := redis.Bytes(repo.conn.Do("get", gs_commons_encrypt.SHA1(api+appId)))
	if err != nil {
		return nil, err
	}
	var f function
	err = msgpack.Unmarshal(b, &f)
	return &f, err
}

func (repo *functionRepo) groupCollection() *mgo.Collection {
	return repo.session.DB("gosion").C("function_groups")
}

func (repo *functionRepo) functionCollection() *mgo.Collection {
	return repo.session.DB("gosion").C("functions")
}

func (repo *functionRepo) Close() {
	repo.conn.Close()
	repo.session.Close()
}
