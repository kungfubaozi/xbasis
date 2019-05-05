package runtime

import (
	"fmt"
	"github.com/garyburd/redigo/redis"
	"gopkg.in/mgo.v2"
	"konekko.me/gosion/commons/gslogrus"
	"konekko.me/gosion/commons/indexutils"
	"konekko.me/gosion/flow/history"
	"konekko.me/gosion/flow/instance"
	"konekko.me/gosion/flow/process"
	"konekko.me/gosion/flow/script"
)

type core struct {
	pool      *redis.Pool
	session   *mgo.Session
	log       *gslogrus.Logger
	script    *script.LuaScript
	client    *indexutils.Client
	pipelines map[string]process.Pipeline
}

func (c *core) Run() error {
	fmt.Println("start")
	return nil
}

func (c *core) History() history.Interface {
	panic("implement me")
}

func (c *core) Processes() process.Interface {
	return &process.Processes{Callback: func(pip process.Pipeline) {
		c.pipelines[pip.Id()] = pip
	}, Session: c.session, Pool: c.pool, Client: c.client}
}

func (c *core) Instances() instance.Interface {
	return &instance.Instances{Session: c.session, Pool: c.pool, Client: c.client}
}

func (c *core) Runtime() Runtime {
	panic("implement me")
}

type Runtime interface {
	Next()
}
