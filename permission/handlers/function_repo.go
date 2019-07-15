package permissionhandlers

import (
	"fmt"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"konekko.me/xbasis/commons/hashcode"
	"konekko.me/xbasis/commons/indexutils"
)

type functionRepo struct {
	session *mgo.Session
	*indexutils.Client
}

func (repo *functionRepo) AddFunction(function *function) error {
	id, err := repo.AddDataById(function.Id, functionIndex, &SimplifiedFunction{
		Id:        function.Id,
		Name:      function.Name,
		AppId:     function.AppId,
		Path:      function.Api,
		Share:     function.Share,
		JoinField: "relation",
	})
	if err != nil {
		return err
	}
	if len(id) > 0 {
		function.SID = id
		return repo.functionCollection(function.AppId).Insert(function)
	}
	return indexutils.ErrNotFound
}

func (repo *functionRepo) AddGroup(group *functionGroup) error {
	//id, err := repo.AddData(functionGroupRelationIndex, group)
	//if err != nil {
	//	return err
	//}
	//if len(id) > 0 {
	//	group.SID = id
	//	return
	//}
	return repo.groupCollection(group.AppId).Insert(group)
}

func (repo *functionRepo) FindChildGroups(appId, parentId string) ([]*functionGroup, error) {
	var groups []*functionGroup
	err := repo.groupCollection(appId).Find(bson.M{"bind_group_id": parentId, "app_id": appId}).All(&groups)
	return groups, err
}

func (repo *functionRepo) FindChildFunctions(appId, parentId string) ([]*function, error) {
	var functions []*function
	err := repo.functionCollection(appId).Find(bson.M{"bind_group_id": parentId, "app_id": appId}).All(&functions)
	return functions, err
}

func (repo *functionRepo) FindApi(appId, api string) (*function, error) {
	var f function
	err := repo.functionCollection(appId).Find(bson.M{"app_id": appId, "api": api}).One(&f)
	return &f, err
}

func (repo *functionRepo) FindApiById(appId, id string) (*function, error) {
	var f function
	err := repo.functionCollection(appId).Find(bson.M{"app_id": appId, "_id": id}).One(&f)
	return &f, err
}

func (repo *functionRepo) FindApiByPrimaryId(id, appId string) (*function, error) {
	var f function
	err := repo.functionCollection(appId).Find(bson.M{"_id": id}).One(&f)
	return &f, err
}

func (repo *functionRepo) FindGroupExists(groupId, appId string) bool {
	c, err := repo.groupCollection(appId).Find(bson.M{"_id": groupId}).Count()
	if err != nil {
		return false
	}
	return c > 0
}

func (repo *functionRepo) UpdateFunction(appId, id string, f *function) error {
	return repo.functionCollection(appId).Update(bson.M{"_id": id}, f)
}

//
//func (repo *functionRepo) SimplifiedLookupApi(appId, path string) (*SimplifiedFunction, error) {
//	var sf SimplifiedFunction
//
//	ok, err := repo.QueryFirst(functionIndex, map[string]interface{}{"app_id": appId, "path": path, "join_field": "relation"}, &sf)
//	if err != nil {
//		return nil, err
//	}
//	if ok {
//		return &sf, nil
//	}
//	return nil, indexutils.ErrNotFound
//}

func (repo *functionRepo) groupCollection(appId string) *mgo.Collection {
	return repo.session.DB(dbName).C(fmt.Sprintf("%s_%d", functionGroupCollection, hashcode.Equa(appId)))
}

func (repo *functionRepo) functionCollection(appId string) *mgo.Collection {
	return repo.session.DB(dbName).C(fmt.Sprintf("%s_%d", functionCollection, hashcode.Equa(appId)))
}

func (repo *functionRepo) Close() {
	if repo.session != nil {
		repo.session.Close()
	}
}
