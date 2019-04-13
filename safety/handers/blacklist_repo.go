package safetyhanders

import (
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"konekko.me/gosion/commons/encrypt"
	"konekko.me/gosion/commons/indexutils"
	"time"
)

type blacklistRepo struct {
	session *mgo.Session
	*indexutils.Client
}

func (repo *blacklistRepo) Save(bt int64, content, userId string) error {
	b := &blacklist{
		Type:         bt,
		Content:      content,
		CreateAt:     time.Now().UnixNano(),
		CreateUserId: userId,
	}

	b.Content = encrypt.SHA1(b.Content)
	id, err := repo.AddData("gs_safety_blacklist", b)
	if err != nil {
		return err
	}
	if len(id) > 0 {
		return repo.collection().Insert(b)
	}

	return indexutils.ErrNotFound
}

func (repo *blacklistRepo) Remove(id string) error {

	ok, err := repo.Delete("gs_safety_blacklist", map[string]interface{}{"content": encrypt.SHA1(id)})
	if err != nil {
		return err
	}
	if ok {
		return repo.collection().Remove(bson.M{"content": id})
	}
	return indexutils.ErrNotFound
}

func (repo *blacklistRepo) Exists(bt int64, content string) bool {
	var b blacklist
	ok, err := repo.QueryFirst("gs_safety_blacklist", map[string]interface{}{"type": bt, "content": content}, &b)
	if err != nil {
		return false
	}
	return !ok
}

func (repo *blacklistRepo) Close() {
	if repo.session != nil {
		repo.session.Close()
	}
}

func (repo *blacklistRepo) collection() *mgo.Collection {
	return repo.session.DB(dbName).C(blacklistCollection)
}
