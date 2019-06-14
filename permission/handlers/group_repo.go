package permissionhandlers

import (
	"fmt"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"konekko.me/gosion/commons/generator"
	"konekko.me/gosion/commons/hashcode"
	"time"
)

type groupRepo struct {
	session *mgo.Session
	id      gs_commons_generator.IDGenerator
}

func (repo *groupRepo) FindByName(appId, name string) (*userGroup, error) {
	var ug userGroup
	err := repo.collection().Find(bson.M{"name": name, "app_id": appId}).One(&ug)
	return &ug, err
}

func (repo *groupRepo) Save(appId, userId, name, bindGroupId string) error {
	ug := &userGroup{
		CreateUserId: userId,
		BindGroupId:  bindGroupId,
		CreateAt:     time.Now().UnixNano(),
		Name:         name,
		AppId:        appId,
		Id:           repo.id.Get(),
	}
	return repo.collection().Insert(ug)
}

func (repo *groupRepo) collection() *mgo.Collection {
	return repo.session.DB(dbName).C(groupCollection)
}

func (repo *groupRepo) groupUsersCollection(groupId string) *mgo.Collection {
	return repo.session.DB(dbName).C(fmt.Sprintf("%s_%d", groupUsersCollection, hashcode.Get(groupId)%5))
}

func (repo *groupRepo) Close() {
	repo.session.Close()
}

func (repo *groupRepo) FindGroupItems(appId, id string) ([]*userGroup, error) {
	var groups []*userGroup
	err := repo.collection().Find(bson.M{"app_id": appId, "bing_group_id": id}).All(&groups)
	return groups, err
}

func (repo *groupRepo) FindGroupUsers(appId, groupId string) ([]*userGroupsRelation, error) {
	var groups []*userGroupsRelation
	err := repo.groupUsersCollection(groupId).Find(bson.M{"app_id": appId, "$elemMatch": bson.M{"bind_group_id": groupId}}).All(groups)
	return groups, err
}
