package modules

import (
	"github.com/garyburd/redigo/redis"
	"gopkg.in/mgo.v2"
	"konekko.me/gosion/commons/generator"
	"konekko.me/gosion/commons/gslogrus"
	"konekko.me/gosion/commons/indexutils"
	"konekko.me/gosion/workflow/flowerr"
	"konekko.me/gosion/workflow/models"
)

type IInstance interface {
	FindById(instanceId string) (*models.Instance, *flowerr.Error)

	//current status
	Status(instanceId string) (int64, *flowerr.Error)

	IsStarted(instanceId string) (bool, *flowerr.Error)

	//is finished
	IsFinished(instanceId string, nodeId string) (bool, *flowerr.Error)

	//nodes running the current instance
	CurrentProcess(instanceId string)

	//开始新的实例
	New(ins *models.Instance) *flowerr.Error

	HasPermission() *flowerr.Error

	//更新实例当前进行的节点
	UpdateInstanceCurrentNodes(instanceId string, nodeIds ...string) *flowerr.Error

	FindRequireUserProcessingInstances(userId string, pageIndex, pageSize int64)

	IsUserRequireProcessingThatNode(userId string, nodeId string)
}

type instances struct {
	session *mgo.Session
	pool    *redis.Pool
	client  *indexutils.Client
	log     *gslogrus.Logger
	id      gs_commons_generator.IDGenerator
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

func (i *instances) IsFinished(instanceId string, nodeId string) (bool, *flowerr.Error) {
	panic("implement me")
}

func (i *instances) CurrentProcess(instanceId string) {
	panic("implement me")
}

func (i *instances) New(ins *models.Instance) *flowerr.Error {
	panic("implement me")
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
