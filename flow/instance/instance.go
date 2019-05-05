package instance

import (
	"github.com/garyburd/redigo/redis"
	"gopkg.in/mgo.v2"
	"konekko.me/gosion/commons/indexutils"
)

type Instance struct {
	Id           string   `bson:"_id" json:"id"`
	Name         string   `bson:"name" json:"name"`
	Originator   string   `bson:"originator" json:"originator"` //发起人
	ProcessId    string   `bson:"process_id" json:"process_id"` //对应的流程
	CreateAt     int64    `bson:"create_at" json:"create_at"`
	CurrentNodes []string `bson:"current_nodes" json:"current_nodes"` //当前节点
}

type Interface interface {
	FindById(instanceId string) (*Instance, error)

	//current status
	Status(instanceId string) (int64, error)

	//is finished
	IsFinished(instanceId string, nodeId string) (bool, error)

	//nodes running the current instance
	CurrentProcess(instanceId string)

	//new instance
	New(ins *Instance) error

	//更新实例当前进行的节点
	UpdateInstanceCurrentNodes(instanceId string, nodeIds ...string) error

	FindRequireUserProcessingInstances(userId string, pageIndex, pageSize int64)

	IsUserRequireProcessingThatNode(userId string, nodeId string)
}

type Instances struct {
	Session *mgo.Session
	Pool    *redis.Pool
	Client  *indexutils.Client
}

func (i *Instances) UpdateInstanceCurrentNodes(instanceId string, nodeIds ...string) error {
	panic("implement me")
}

func (i *Instances) IsUserRequireProcessingThatNode(userId string, nodeId string) {
	panic("implement me")
}

func (i *Instances) FindRequireUserProcessingInstances(userId string, pageIndex, pageSize int64) {
	panic("implement me")
}

func (i *Instances) FindById(instanceId string) (*Instance, error) {
	panic("implement me")
}

func (i *Instances) Status(instanceId string) (int64, error) {
	panic("implement me")
}

func (i *Instances) IsFinished(instanceId string, nodeId string) (bool, error) {
	panic("implement me")
}

func (i *Instances) CurrentProcess(instanceId string) {
	panic("implement me")
}

func (i *Instances) New(ins *Instance) error {
	panic("implement me")
}
