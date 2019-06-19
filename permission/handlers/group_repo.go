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

func (repo *groupRepo) groupUsersCollection(appId string) *mgo.Collection {
	return repo.session.DB(dbName).C(fmt.Sprintf("%s_%d", groupUsersCollection, hashcode.Get(appId)%5))
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
	err := repo.groupUsersCollection(appId).Find(bson.M{"$elemMatch": bson.M{"bind_group_id": groupId}}).All(groups)
	return groups, err
}

func (repo *groupRepo) FindUserById(userId, appId string) (*userGroupsRelation, error) {
	var groups *userGroupsRelation
	err := repo.groupUsersCollection(appId).Find(bson.M{"user_id": userId}).One(groups)
	return groups, err
}

func (repo *groupRepo) SetGroupRelation(u *userGroupsRelation) error {
	_, err := repo.groupUsersCollection(u.AppId).Upsert(bson.M{"user_id": u.UserId}, u)
	return err
}
