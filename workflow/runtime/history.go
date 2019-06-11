package runtime

import (
	"context"
	"github.com/garyburd/redigo/redis"
	"gopkg.in/mgo.v2"
	"konekko.me/gosion/analysis/client"
	"konekko.me/gosion/commons/generator"
	"konekko.me/gosion/commons/indexutils"
	"konekko.me/gosion/workflow/flowerr"
	"konekko.me/gosion/workflow/models"
)

type history struct {
	session *mgo.Session
	pool    *redis.Pool
	client  *indexutils.Client
	id      gs_commons_generator.IDGenerator
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
