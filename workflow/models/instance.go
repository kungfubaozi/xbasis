package models

import (
	"github.com/garyburd/redigo/redis"
	"gopkg.in/mgo.v2"
	"konekko.me/gosion/commons/indexutils"
)

type Instance struct {
	Id           string   `bson:"_id" json:"id"`
	Name         string   `bson:"name" json:"name"`
	AppId        string   `bson:"app_id" json:"app_id"`
	Originator   string   `bson:"originator" json:"originator"` //发起人
	ProcessId    string   `bson:"process_id" json:"process_id"` //对应的流程
	CreateAt     int64    `bson:"create_at" json:"create_at"`
	CurrentNodes []string `bson:"current_nodes" json:"current_nodes"` //当前节点
}

type Instances struct {
	Session *mgo.Session
	Pool    *redis.Pool
	Client  *indexutils.Client
}
