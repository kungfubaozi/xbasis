package applicationhanderls

import (
	"fmt"
	"gopkg.in/mgo.v2"
	"konekko.me/gosion/commons/encrypt"
	"konekko.me/gosion/commons/hashcode"
	"konekko.me/gosion/commons/indexutils"
	"time"
)

type syncRepo struct {
	*indexutils.Client
	session *mgo.Session
}

func (repo *syncRepo) Synced(userId, appId, relation string) (int64, error) {

	count, err := repo.Client.Count(repo.GetKey(relation), map[string]interface{}{"user_id": userId, "app_id": appId, "sha_relation": relation})
	if err != nil {
		return 0, err
	}
	return count, nil
}

func (repo *syncRepo) Sync(userId, appId, relation string) error {

	s := &syncLog{
		UserId:      userId,
		AppId:       appId,
		SHARelation: encrypt.SHA1(relation),
		Timestamp:   time.Now().UnixNano(),
	}

	_, err := repo.Client.AddData(repo.GetKey(relation), s)
	return err
}

func (repo *syncRepo) GetKey(relation string) string {
	return fmt.Sprintf("%s.%d", "gosion-synced.", hashcode.Get(relation))
}

func (repo *syncRepo) Close() {
	repo.session.Close()
}
