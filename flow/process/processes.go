package process

import (
	"github.com/garyburd/redigo/redis"
	"gopkg.in/mgo.v2"
	"konekko.me/gosion/commons/indexutils"
)

type AddProcessToPipelineCallback func(pip Pipeline)

type Interface interface {
	AddProcess(p *Process)
}

type Processes struct {
	Callback AddProcessToPipelineCallback
	Session  *mgo.Session
	Pool     *redis.Pool
	Client   *indexutils.Client
}

func (pro *Processes) AddProcess(p *Process) {
	if pro.Callback != nil {
		pip := &pipeline{}
		pro.Callback(pip)
	}
}
