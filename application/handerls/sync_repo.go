package applicationhanderls

import (
	"gopkg.in/mgo.v2"
	"konekko.me/gosion/commons/indexutils"
)

type syncRepo struct {
	*indexutils.Client
	session *mgo.Session
}

func (repo *syncRepo) IsSynced(userId, appId string) (int64, error) {
	count, err := repo.Client.Count("gs-usersync-his", map[string]interface{}{"user_id": userId, "app_id": appId})
	if err != nil {
		return 0, err
	}
	return count, nil
}

func (repo *syncRepo) Close() {
	repo.session.Close()
}
