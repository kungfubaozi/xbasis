package modules

import (
	"github.com/garyburd/redigo/redis"
	"gopkg.in/mgo.v2"
	"konekko.me/gosion/commons/gslogrus"
	"konekko.me/gosion/commons/indexutils"
	"konekko.me/gosion/workflow/models"
)

type IInstance interface {
	FindById(instanceId string) (*models.Instance, error)

	//current status
	Status(instanceId string) (int64, error)

	//is finished
	IsFinished(instanceId string, nodeId string) (bool, error)

	//nodes running the current instance
	CurrentProcess(instanceId string)

	//new instance
	New(ins *models.Instance) error

	//更新实例当前进行的节点
	UpdateInstanceCurrentNodes(instanceId string, nodeIds ...string) error

	FindRequireUserProcessingInstances(userId string, pageIndex, pageSize int64)

	IsUserRequireProcessingThatNode(userId string, nodeId string)
}

type instances struct {
	session *mgo.Session
	pool    *redis.Pool
	client  *indexutils.Client
	log     *gslogrus.Logger
}

func (i *instances) FindById(instanceId string) (*models.Instance, error) {
	panic("implement me")
}

func (i *instances) Status(instanceId string) (int64, error) {
	panic("implement me")
}

func (i *instances) IsFinished(instanceId string, nodeId string) (bool, error) {
	panic("implement me")
}

func (i *instances) CurrentProcess(instanceId string) {
	panic("implement me")
}

func (i *instances) New(ins *models.Instance) error {
	panic("implement me")
}

func (i *instances) UpdateInstanceCurrentNodes(instanceId string, nodeIds ...string) error {
	panic("implement me")
}

func (i *instances) FindRequireUserProcessingInstances(userId string, pageIndex, pageSize int64) {
	panic("implement me")
}

func (i *instances) IsUserRequireProcessingThatNode(userId string, nodeId string) {
	panic("implement me")
}

