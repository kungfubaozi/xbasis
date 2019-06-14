package permissionhandlers

import (
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"konekko.me/gosion/commons/indexutils"
)

type functionRepo struct {
	session *mgo.Session
	*indexutils.Client
}

func (repo *functionRepo) AddFunction(function *function) error {
	id, err := repo.AddData("gs-functions", function)
	if err != nil {
		return err
	}
	if len(id) > 0 {
		function.SID = id
		return repo.functionCollection().Insert(function)
	}
	return indexutils.ErrNotFound
}

func (repo *functionRepo) AddGroup(group *functionGroup) error {
	id, err := repo.AddData("gs-functions-groups", group)
	if err != nil {
		return err
	}
	if len(id) > 0 {
		group.SID = id
		return repo.groupCollection().Insert(group)
	}
	return indexutils.ErrNotFound
}

func (repo *functionRepo) FindChildGroups(appId, parentId string) ([]*functionGroup, error) {
	var groups []*functionGroup
	err := repo.groupCollection().Find(bson.M{"bind_group_id": parentId, "app_id": appId}).All(&groups)
	return groups, err
}

func (repo *functionRepo) FindChildFunctions(appId, parentId string) ([]*function, error) {
	var functions []*function
	err := repo.functionCollection().Find(bson.M{"bind_group_id": parentId, "app_id": appId}).All(&functions)
	return functions, err
}

func (repo *functionRepo) FindApi(appId, api string) (*function, error) {
	var f function
	err := repo.functionCollection().Find(bson.M{"app_id": appId, "api": api}).One(&f)
	return &f, err
}

func (repo *functionRepo) FindApiById(appId, id string) (*function, error) {
	var f function
	err := repo.functionCollection().Find(bson.M{"app_id": appId, "_id": id}).One(&f)
	return &f, err
}

func (repo *functionRepo) FindApiByPrimaryId(id string) (*function, error) {
	var f function
	err := repo.functionCollection().Find(bson.M{"_id": id}).One(&f)
	return &f, err
}

func (repo *functionRepo) FindGroupExists(groupId string) bool {
	c, err := repo.groupCollection().Find(bson.M{"_id": groupId}).Count()
	if err != nil {
		return false
	}
	return c > 0
}

func (repo *functionRepo) SimplifiedLookupApi(appId, api string) (*simplifiedFunction, error) {
	var sf simplifiedFunction

	ok, err := repo.QueryFirst("gs-functions", map[string]interface{}{"app_id": appId, "path": api}, &sf)
	if err != nil {
		return nil, err
	}
	if ok {
		return &sf, nil
	}
	return nil, indexutils.ErrNotFound
}

func (repo *functionRepo) groupCollection() *mgo.Collection {
	return repo.session.DB(dbName).C(functionGroupCollection)
}

func (repo *functionRepo) functionCollection() *mgo.Collection {
	return repo.session.DB(dbName).C(functionCollection)
}

func (repo *functionRepo) Close() {
	if repo.session != nil {
		repo.session.Close()
	}
}
