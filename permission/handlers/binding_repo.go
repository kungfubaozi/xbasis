package permissionhandlers

import (
	"fmt"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"konekko.me/gosion/commons/generator"
	"konekko.me/gosion/commons/hashcode"
	"konekko.me/gosion/commons/indexutils"
)

type bindingRepo struct {
	session *mgo.Session
	*indexutils.Client
	id gs_commons_generator.IDGenerator
}

func (repo *bindingRepo) Close() {
	repo.session.Close()
}

func (repo *bindingRepo) functionCollection() *mgo.Collection {
	return repo.session.DB(dbName).C(functionCollection)
}

func (repo *bindingRepo) userRelationCollection(userId string) *mgo.Collection {
	return repo.session.DB(dbName).C(fmt.Sprintf("%s_%d", userRoleRelationCollection, hashcode.Get(userId)%5))
}

func (repo *bindingRepo) FindFunctionById(id string) (*function, error) {
	f := &function{}
	err := repo.functionCollection().Find(bson.M{"_id": id}).One(f)
	return f, err
}

func (repo *bindingRepo) FindUserById(id, appId string) (*userRolesRelation, error) {
	f := &userRolesRelation{}
	err := repo.functionCollection().Find(bson.M{"user_id": id, "app_id": appId}).One(f)
	return f, err
}

func (repo *bindingRepo) UpdateFunctionRole(id, role string) error {
	return repo.functionCollection().Update(bson.M{"_id": id}, bson.M{"$push": bson.M{"roles": role}})
}

func (repo *bindingRepo) UpdateUserRole(id, appId string, role []string) error {
	return repo.userRelationCollection(id).Update(bson.M{"user_id": id, "app_id": appId}, bson.M{"$push": bson.M{"roles": role}})
}

func (repo *bindingRepo) RemoveRoleFromFunctions(id, role string) error {
	return repo.functionCollection().Update(bson.M{"_id": id}, bson.M{"$pull": bson.M{"roles": role}})
}

func (repo *bindingRepo) RemoveRoleFromUserRelation(userId, role string) error {
	return repo.userRelationCollection(userId).Update(bson.M{"user_id": userId}, bson.M{"$pull": bson.M{"roles": role}})
}
