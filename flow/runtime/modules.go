package runtime

import (
	"github.com/garyburd/redigo/redis"
	"gopkg.in/mgo.v2"
	"konekko.me/gosion/commons/gslogrus"
	"konekko.me/gosion/commons/indexutils"
	"konekko.me/gosion/flow/history"
	"konekko.me/gosion/flow/instance"
	"konekko.me/gosion/flow/process"
)

type Modules interface {
	//Operational History (Instances, Nodes)
	History() history.Interface

	//Processes
	Processes() process.Interface

	//Instances (running instances)
	Instances() instance.Interface

	//Flow controller
	Runtime() Runtime

	Run() error
}

func New(session *mgo.Session, pool *redis.Pool, client *indexutils.Client, log *gslogrus.Logger) Modules {
	return &core{session: session,
		pool:      pool,
		client:    client,
		log:       log,
		pipelines: make(map[string]process.Pipeline)}
}
