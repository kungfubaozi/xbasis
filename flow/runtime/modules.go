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
	return &core{shutdown: make(chan error), session: session, pool: pool, client: client, log: log}
}

type core struct {
	session  *mgo.Session
	pool     *redis.Pool
	client   *indexutils.Client
	log      *gslogrus.Logger
	shutdown chan error
	ri       *runtime
	hi       history.Interface
	pi       process.Interface
	ii       instance.Interface
}

func (c *core) Run() error {
	go func() {
		c.init()
		c.start()
	}()
	return <-c.shutdown
}

func (c *core) History() history.Interface {
	return c.hi
}

func (c *core) Processes() process.Interface {
	return c.pi
}

func (c *core) Instances() instance.Interface {
	return c.ii
}

func (c *core) Runtime() Runtime {
	return c.ri
}

func (c *core) init() {
	c.ri = &runtime{
		shutdown:  c.shutdown,
		pipelines: make(map[string]process.Pipeline),
	}
	c.pi = &process.Processes{Callback: func(pip process.Pipeline) {
		c.ri.pipelines[pip.Id()] = pip
	}, Session: c.session, Pool: c.pool, Client: c.client}
	c.ii = &instance.Instances{Session: c.session, Pool: c.pool, Client: c.client}
	c.hi = &history.History{}
	c.ri.hi = c.hi
	c.ri.ii = c.ii
	c.ri.pi = c.pi
}

func (c *core) start() {

}
