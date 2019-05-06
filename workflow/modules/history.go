package modules

import (
	"github.com/garyburd/redigo/redis"
	"gopkg.in/mgo.v2"
	"konekko.me/gosion/commons/gslogrus"
	"konekko.me/gosion/commons/indexutils"
)

type IHistory interface {
	GetInstanceNodeHistory(instanceId, nodeId string)

	GetInstanceOperateHistory(instanceId string)

	GetInstanceStatus()
}

type history struct {
	session  *mgo.Session
	pool     *redis.Pool
	client   *indexutils.Client
	log      *gslogrus.Logger
}

func (h *history) GetInstanceNodeHistory(instanceId, nodeId string) {
	panic("implement me")
}

func (h *history) GetInstanceOperateHistory(instanceId string) {
	panic("implement me")
}

func (h *history) GetInstanceStatus() {
	panic("implement me")
}
