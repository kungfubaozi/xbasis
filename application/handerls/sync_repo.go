package applicationhanderls

import (
	"fmt"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"konekko.me/xbasis/commons/encrypt"
	"konekko.me/xbasis/commons/hashcode"
	"time"
)

type syncRepo struct {
	session *mgo.Session
}

func (repo *syncRepo) Synced(userId, appId, relation string) (bool, error) {
	n, err := repo.collection(relation).Find(bson.M{"user_id": userId, "app_id": appId, "sha_relation": encrypt.SHA1(relation)}).Count()
	if err != nil {
		return false, nil
	}
	return n == 1, nil
}

func (repo *syncRepo) Sync(userId, appId, relation string) error {

	s := &syncLog{
		UserId:      userId,
		AppId:       appId,
		SHARelation: encrypt.SHA1(relation),
		Timestamp:   time.Now().UnixNano(),
	}

	err := repo.collection(relation).Insert(s)
	if err != nil {
		return err
	}

	return err
}

func (repo *syncRepo) collection(relation string) *mgo.Collection {
	return repo.session.DB(dbName).C(repo.GetKey(relation))
}

func (repo *syncRepo) GetKey(relation string) string {
	return fmt.Sprintf("%s_%d", synclogCollection, hashcode.Equa(relation))
}

func (repo *syncRepo) Close() {
	repo.session.Close()
}
