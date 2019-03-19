package permission_repositories

import (
	"github.com/garyburd/redigo/redis"
	"github.com/vmihailenco/msgpack"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"konekko.me/gosion/commons/encrypt"
)

type FunctionRepo struct {
	Session *mgo.Session
	Conn    redis.Conn
}

func (repo *FunctionRepo) AddFunction(function *Function) error {
	b, err := msgpack.Marshal(function)
	if err != nil {
		return err
	}
	_, err = repo.Conn.Do("set", gs_commons_encrypt.SHA1(function.Api+function.AppId), b)
	if err != nil {
		return err
	}
	return repo.functionCollection().Insert(function)
}

func (repo *FunctionRepo) AddGroup(group *FunctionGroup) error {
	return repo.groupCollection().Insert(group)
}

func (repo *FunctionRepo) FindGroup(appId, name string) error {

}

func (repo *FunctionRepo) FindChildGroups(parentId string) ([]*FunctionGroup, error) {
	var groups []*FunctionGroup
	err := repo.groupCollection().Find(bson.M{"bind_group_id": parentId}).All(&groups)
	return groups, err
}

func (repo *FunctionRepo) FindChildFunctions(parentId string) ([]*Function, error) {
	var functions []*Function
	err := repo.functionCollection().Find(bson.M{"bind_group_id": parentId}).All(&functions)
	return functions, err
}

func (repo *FunctionRepo) FindApi(appId, api string) (*Function, error) {
	var f Function
	err := repo.functionCollection().Find(bson.M{"app_id": appId, "api": api}).One(&f)
	return &f, err
}

func (repo *FunctionRepo) FindGroupExits(groupId string) bool {
	c, err := repo.groupCollection().Find(bson.M{"_id": groupId}).Count()
	if err != nil {
		return true
	}
	return c > 0
}

func (repo *FunctionRepo) FindApiInCache(appId, api string) (*Function, error) {
	b, err := redis.Bytes(repo.Conn.Do("get", gs_commons_encrypt.SHA1(api+appId)))
	if err != nil {
		return nil, err
	}
	var f Function
	err = msgpack.Unmarshal(b, &f)
	return &f, err
}

func (repo *FunctionRepo) groupCollection() *mgo.Collection {
	return repo.Session.DB("gosion").C("function_groups")
}

func (repo *FunctionRepo) functionCollection() *mgo.Collection {
	return repo.Session.DB("gosion").C("functions")
}

func (repo *FunctionRepo) Close() {
	repo.Conn.Close()
	repo.Session.Close()
}
