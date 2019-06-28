package safetyhanders

import (
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"konekko.me/xbasis/commons/encrypt"
	"konekko.me/xbasis/commons/indexutils"
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
	id, err := repo.AddData("gs-safety-blacklist", b)
	if err != nil {
		return err
	}
	if len(id) > 0 {
		return repo.collection().Insert(b)
	}

	return indexutils.ErrNotFound
}

func (repo *blacklistRepo) Remove(id string) error {
	ok, err := repo.Delete("gs-safety-blacklist", map[string]interface{}{"content": encrypt.SHA1(id)})
	if err != nil {
		return err
	}
	if ok {
		return repo.collection().Remove(bson.M{"content": id})
	}
	return indexutils.ErrNotFound
}

func (repo *blacklistRepo) Exists(bt int64, content string) bool {
	count, err := repo.Count("gs-safety-blacklist", map[string]interface{}{"type": bt, "content": content})
	if err != nil {
		return false
	}
	return count != 0
}

func (repo *blacklistRepo) Close() {
	if repo.session != nil {
		repo.session.Close()
	}
}

func (repo *blacklistRepo) collection() *mgo.Collection {
	return repo.session.DB(dbName).C(blacklistCollection)
}
