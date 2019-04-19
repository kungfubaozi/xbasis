package permissionhandlers

import (
	"errors"
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

func (repo *functionRepo) FindApi(structureId, api string) (*function, error) {
	var f function
	err := repo.functionCollection().Find(bson.M{"structure_id": structureId, "api": api}).One(&f)
	return &f, err
}

func (repo *functionRepo) FindGroupExists(groupId string) bool {
	c, err := repo.groupCollection().Find(bson.M{"_id": groupId}).Count()
	if err != nil {
		return true
	}
	return c > 0
}

func (repo *functionRepo) FindApiInCache(structureId, api string) (*function, error) {
	var function function
	ok, err := repo.QueryFirst("gs-functions", map[string]interface{}{
		"api":          api,
		"structure_id": structureId,
	}, &function)
	if err != nil {
		return nil, err
	}
	if ok {
		return &function, nil
	}
	return nil, errors.New("not found")
}

func (repo *functionRepo) SimplifiedLookupApi(structureId, api string) (*simplifiedFunction, error) {
	var sf simplifiedFunction

	ok, err := repo.QueryFirst("gs-functions", map[string]interface{}{"structure_id": structureId, "api": api}, &sf, "id", "roles", "auth_types", "grant_platforms")
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
