package permissionhandlers

import (
	"fmt"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	generator "konekko.me/xbasis/commons/generator"
	"konekko.me/xbasis/commons/hashcode"
	"konekko.me/xbasis/commons/indexutils"
)

type bindingRepo struct {
	session *mgo.Session
	*indexutils.Client
	id generator.IDGenerator
}

func (repo *bindingRepo) Close() {
	repo.session.Close()
}

func (repo *bindingRepo) functionCollection(functionId string) *mgo.Collection {
	return repo.session.DB(dbName).C(fmt.Sprintf("%s_%d", functionRoleRelationCollection, hashcode.Equa(functionId)))
}

func (repo *bindingRepo) userRelationCollection(userId string) *mgo.Collection {
	return repo.session.DB(dbName).C(fmt.Sprintf("%s_%d", userRoleRelationCollection, hashcode.Equa(userId)))
}

func (repo *bindingRepo) FindRelationFunctionById(id, appId string) (*functionRolesRelation, error) {
	f := &functionRolesRelation{}
	err := repo.functionCollection(id).Find(bson.M{"function_id": id, "app_id": appId}).One(f)
	return f, err
}

func (repo *bindingRepo) FindRelationUserById(userId, appId string) (*userRolesRelation, error) {
	f := &userRolesRelation{}
	err := repo.userRelationCollection(userId).Find(bson.M{"user_id": userId, "app_id": appId}).One(f)
	return f, err
}

func (repo *bindingRepo) UpdateFunctionRole(id, appId string, roles []string) error {
	return repo.functionCollection(id).Update(bson.M{"function_id": id, "app_id": appId}, bson.M{"$pushAll": bson.M{"roles": roles}})
}

func (repo *bindingRepo) UpdateUserRole(id, appId string, role []string) error {
	return repo.userRelationCollection(id).Update(bson.M{"user_id": id, "app_id": appId}, bson.M{"$pushAll": bson.M{"roles": role}})
}

func (repo *bindingRepo) RemoveRoleFromFunctions(id, appId string, role string) error {
	return repo.functionCollection(id).Update(bson.M{"function_id": id, "app_id": appId}, bson.M{"$pull": bson.M{"roles": role}})
}

func (repo *bindingRepo) RemoveRoleFromUserRelation(userId, role string) error {
	return repo.userRelationCollection(userId).Update(bson.M{"user_id": userId}, bson.M{"$pull": bson.M{"roles": role}})
}
