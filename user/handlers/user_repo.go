package userhandlers

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/olivere/elastic"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"konekko.me/xbasis/commons/encrypt"
	"konekko.me/xbasis/commons/hashcode"
	"konekko.me/xbasis/commons/indexutils"
)

type userRepo struct {
	session *mgo.Session
	*indexutils.Client
}

func (repo *userRepo) FindById(id string) (*userModel, error) {
	info := &userModel{}
	err := repo.userCollection(id).Find(bson.M{"_id": id}).One(info)
	return info, err
}

func (repo *userRepo) AddUser(user *userModel) error {
	return repo.userCollection(user.Id).Insert(user)
}

func (repo *userRepo) AddUserInfo(info *userInfo) error {
	return repo.infoCollection(info.UserId).Insert(info)
}

func (repo *userRepo) userCollection(userId string) *mgo.Collection {
	return repo.session.DB(dbName).C(fmt.Sprintf("%s_%d", userCollection, hashcode.Equa(userId)))
}

func (repo *userRepo) oauthCollection(openId string) *mgo.Collection {
	return repo.session.DB(dbName).C(fmt.Sprintf("%s_%d", userOAuthCollection, hashcode.Equa(openId)))
}

func (repo *userRepo) FindIndexTable(key string, content string) (string, error) {
	var m map[string]interface{}
	ok, err := repo.QueryFirst(typeUserIndex, map[string]interface{}{key: content, "invite": false}, &m)
	if err != nil {
		return "", nil
	}
	if ok {
		return m["id"].(string), nil
	}
	return "", indexutils.ErrNotFound
}

func (repo *userRepo) index(c string) string {
	return encrypt.SHA1(c)
}

func (repo *userRepo) infoCollection(userId string) *mgo.Collection {
	return repo.session.DB(dbName).C(fmt.Sprintf("%s_%d", userInfoCollection, hashcode.Equa(userId)))
}

func (repo *userRepo) FindUserInfo(userId string) (*userIndexFields, error) {
	query := elastic.NewBoolQuery()
	query.Must(elastic.NewMatchPhraseQuery("user_id", userId))
	v, err := repo.GetElasticClient().Search(typeUserIndex).Query(query).Type("_doc").Do(context.Background())
	if err != nil {
		return nil, err
	}
	if v.Hits.TotalHits == 1 {
		i := &userIndexFields{}
		err := json.Unmarshal(*v.Hits.Hits[0].Source, i)
		if err == nil {
			return i, nil
		}
	}
	return nil, indexutils.ErrNotFound
}

func (repo *userRepo) Close() {
	repo.session.Close()
}
