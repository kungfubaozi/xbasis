package permissionhandlers

import (
	"context"
	"fmt"
	"github.com/olivere/elastic"
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

func (repo *bindingRepo) SetFunctionRole(id, appId string, r *functionRolesRelation) error {
	_, err := repo.functionCollection(id).Upsert(bson.M{"function_id": id, "app_id": appId}, r)
	return err
}

func (repo *bindingRepo) SetUserRole(id, appId string, r *userRolesRelation) error {
	_, err := repo.userRelationCollection(id).Upsert(bson.M{"user_id": id, "app_id": appId}, r)
	return err
}

func (repo *bindingRepo) RecheckFunctionAuthorize(functionId string) error {
	q := elastic.NewBoolQuery()
	q.Must(elastic.NewMatchPhraseQuery("functionId", functionId))
	_, err := repo.GetElasticClient().UpdateByQuery("xbs-function-authorize.*").Query(q).Script(elastic.NewScript("ctx._source.recheck = true")).Do(context.Background())
	return err
}
