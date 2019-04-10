package userhandlers

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/olivere/elastic"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"konekko.me/gosion/commons/encrypt"
	"konekko.me/gosion/commons/hashcode"
)

type userRepo struct {
	session *mgo.Session
	elastic *elastic.Client
}

func (repo *userRepo) FindById(id string) (*userInfo, error) {
	info := &userInfo{}
	err := repo.userCollection(id).Find(bson.M{"_id": id}).One(info)
	return info, err
}

func (repo *userRepo) AddUser(user *userInfo) error {
	index := &userIndex{TargetId: user.Id}

	addIndex := func(index *userIndex, t int) error {
		p, err := repo.elastic.Index().Index(fmt.Sprintf("%s_it_%d", typeUserIndex, t)).Type("v").BodyJson(index).Do(context.Background())
		if err != nil {
			return err
		}
		fmt.Println("create index id:", p.Id)
		return nil
	}

	//insert sql index
	if len(user.Email) > 0 {
		index.Content = repo.index(user.Email)
		err := addIndex(index, emailIndexType)
		if err != nil {
			return err
		}
	}

	if len(user.Phone) > 0 {
		index.Content = repo.index(user.Phone)
		err := addIndex(index, phoneIndexType)
		if err != nil {
			return err
		}
	}

	if len(user.Account) > 0 {
		index.Content = repo.index(user.Account)
		err := addIndex(index, accountIndexType)
		if err != nil {
			return err
		}
	}

	//add collection

	return repo.userCollection(user.Id).Insert(user)
}

func (repo *userRepo) userCollection(userId string) *mgo.Collection {
	return repo.session.DB(dbName).C(fmt.Sprintf("%s_%d", userCollection, hashcode.Get(userId)))
}

func (repo *userRepo) oauthCollection(openId string) *mgo.Collection {
	return repo.session.DB(dbName).C(fmt.Sprintf("%s_%d", userOAuthCollection, hashcode.Get(openId)))
}

func (repo *userRepo) FindIndexTable(t int, content string) (string, error) {
	content = repo.index(content)
	userIndex := &userIndex{}

	q := elastic.NewMatchQuery("content", content)

	s, err := repo.elastic.Search(fmt.Sprintf("%s_it_%d", typeUserIndex, t)).Type("v").Query(q).Do(context.Background())
	if err != nil {
		return "", nil
	}
	fmt.Println("hits", s.Hits.TotalHits)
	if s.Hits.TotalHits > 0 {
		v := s.Hits.Hits[0]
		err = json.Unmarshal(*v.Source, userIndex)
		if err == nil {
			return userIndex.TargetId, nil
		}
	}

	return "", nil
}

func (repo *userRepo) index(c string) string {
	return gs_commons_encrypt.SHA1(c)
}

func (repo *userRepo) infoCollection(userId string) *mgo.Collection {
	return repo.session.DB(dbName).C(fmt.Sprintf("%s_%d", userInfoCollection, hashcode.Get(userId)))
}

func (repo *userRepo) Close() {
	repo.session.Close()
}
