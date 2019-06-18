package applicationhanderls

import (
	"gopkg.in/mgo.v2"
	"konekko.me/gosion/commons/indexutils"
	"time"
)

type syncRepo struct {
	*indexutils.Client
	session *mgo.Session
}

func (repo *syncRepo) IsSynced(userId, appId string) (int64, error) {
	count, err := repo.Client.Count("gosion-sync-his", map[string]interface{}{"user_id": userId, "app_id": appId})
	if err != nil {
		return 0, err
	}
	return count, nil
}

func (repo *syncRepo) Sync(userId, appId string) error {
	_, err := repo.Client.AddData("gosion-sync-his", map[string]interface{}{"user_id": userId, "app_id": appId, "create_at": time.Now().UnixNano()})
	return err
}

func (repo *syncRepo) Close() {
	repo.session.Close()
}
