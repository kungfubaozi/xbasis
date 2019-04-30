package workflow

import (
	"github.com/garyburd/redigo/redis"
	"gopkg.in/mgo.v2"
	"konekko.me/gosion/commons/dto"
	"konekko.me/gosion/commons/indexutils"
)

type InstanceStatus int

const (
	INSRunning = iota

	INSStop

	INSWaitingEvent

	INSFinished
)

type instances struct {
	conn    redis.Conn
	session *mgo.Session
	client  *indexutils.Client
}

//流程实例
//每启动一个流程都会产生一个对应的实例
type instance struct {
	Id             string         `bson:"_id" json:"id"`
	CurrentNodeIds []string       `bson:"current_node_ids" json:"current_node_ids"`
	StartByUserId  string         `bson:"start_by_user_id" json:"start_by_user_id"` //流程启动者
	StartAt        int64          `bson:"start_at" json:"start_at"`
	ProcessId      string         `bson:"process_id" json:"process_id"` //对应的流程
	Times          int64          `bson:"times" json:"times"`           //运行多少次
	Status         InstanceStatus `bson:"status" json:"status"`
	CreateAt       int64          `bson:"create_at" json:"create_at"`
}

type operate struct {
	instanceId string
}

func (i *instances) get(instanceId string) (*instance, error) {

}

func (i *instances) collection() {

}

func (i *instances) status() {

}

func (i *instances) isCompleted(insId, nodeId string) (*gs_commons_dto.State, error) {

}
