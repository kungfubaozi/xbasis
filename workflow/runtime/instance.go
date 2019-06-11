package runtime

import (
	"github.com/garyburd/redigo/redis"
	"gopkg.in/mgo.v2"
	"konekko.me/gosion/analysis/client"
	"konekko.me/gosion/commons/generator"
	"konekko.me/gosion/commons/indexutils"
	"konekko.me/gosion/workflow/flowerr"
	"konekko.me/gosion/workflow/models"
	"konekko.me/gosion/workflow/modules"
)

type instances struct {
	store   modules.IStore
	session *mgo.Session
	pool    *redis.Pool
	client  *indexutils.Client
	log     analysisclient.LogClient
	id      gs_commons_generator.IDGenerator
}

func (i *instances) NamedStart(name string) *flowerr.Error {
	panic("implement me")
}

func (i *instances) HasPermission() *flowerr.Error {
	panic("implement me")
}

func (i *instances) FindById(instanceId string) (*models.Instance, *flowerr.Error) {
	panic("implement me")
}

func (i *instances) Status(instanceId string) (int64, *flowerr.Error) {
	panic("implement me")
}

func (i *instances) IsStarted(instanceId string) (bool, *flowerr.Error) {
	panic("implement me")
}

func (i *instances) CurrentProcess(instanceId string) {
	panic("implement me")
}

func (i *instances) New(ins *models.Instance) *flowerr.Error {
	panic("")
}

func (i *instances) UpdateInstanceCurrentNodes(instanceId string, nodeIds ...string) *flowerr.Error {
	panic("implement me")
}

func (i *instances) FindRequireUserProcessingInstances(userId string, pageIndex, pageSize int64) {
	panic("implement me")
}

func (i *instances) IsUserRequireProcessingThatNode(userId string, nodeId string) {
	panic("implement me")
}
