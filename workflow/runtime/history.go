package runtime

import (
	"context"
	"github.com/garyburd/redigo/redis"
	"gopkg.in/mgo.v2"
	"konekko.me/xbasis/analysis/client"
	generator "konekko.me/xbasis/commons/generator"
	"konekko.me/xbasis/commons/indexutils"
	"konekko.me/xbasis/workflow/flowerr"
	"konekko.me/xbasis/workflow/models"
)

type history struct {
	session *mgo.Session
	pool    *redis.Pool
	client  *indexutils.Client
	id      generator.IDGenerator
	log     analysisclient.LogClient
}

func (h *history) Record(ctx context.Context, data *models.History) *flowerr.Error {
	panic("implement me")
}

func (h *history) GetInstanceNodeHistory(instanceId, nodeId string) {
	panic("implement me")
}

func (h *history) GetInstanceOperateHistory(instanceId string) {
	panic("implement me")
}
