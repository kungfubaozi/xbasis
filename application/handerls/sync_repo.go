package applicationhanderls

import (
	"gopkg.in/mgo.v2"
	"konekko.me/gosion/commons/date"
	"konekko.me/gosion/commons/encrypt"
	"konekko.me/gosion/commons/indexutils"
	"time"
)

type syncRepo struct {
	*indexutils.Client
	session *mgo.Session
}

func (repo *syncRepo) Synced(userId, appId, relation string) (int64, error) {

	d := gs_commons_date.FormatDate(time.Now(), gs_commons_date.YYYY_I_MM_I_DD)

	count, err := repo.Client.Count("gosion-synced."+d, map[string]interface{}{"user_id": userId, "app_id": appId, "sha_relation": relation})
	if err != nil {
		return 0, err
	}
	return count, nil
}

func (repo *syncRepo) Sync(userId, appId, relation string) error {

	d := gs_commons_date.FormatDate(time.Now(), gs_commons_date.YYYY_I_MM_I_DD)

	s := &syncLog{
		UserId:      userId,
		AppId:       appId,
		SHARelation: encrypt.SHA1(relation),
		Timestamp:   time.Now().UnixNano(),
	}

	_, err := repo.Client.AddData("gosion-synced."+d, s)
	return err
}

func (repo *syncRepo) Close() {
	repo.session.Close()
}
