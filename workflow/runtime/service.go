package runtime

import (
	"fmt"
	"github.com/garyburd/redigo/redis"
	"gopkg.in/mgo.v2"
	"konekko.me/gosion/commons/config"
	"konekko.me/gosion/commons/generator"
	"konekko.me/gosion/commons/gslogrus"
	"konekko.me/gosion/commons/indexutils"
	"konekko.me/gosion/workflow/distribute"
	"konekko.me/gosion/workflow/modules"
	"konekko.me/gosion/workflow/script"
)

type Workflow struct {
	modules *workflow
}

type workflow struct {
	shutdown chan error
	session  *mgo.Session
	pool     *redis.Pool
	client   *indexutils.Client
	log      *gslogrus.Logger
	//
	ri modules.IRuntime
	hi modules.IHistory
	ii modules.IInstance
	pi modules.IProcesses
	fi modules.IForm
	ui modules.IUser
}

func (m *workflow) User() modules.IUser {
	return m.ui
}

func (m *workflow) History() modules.IHistory {
	return m.hi
}

func (m *workflow) Instance() modules.IInstance {
	return m.ii
}

func (m *workflow) Process() modules.IProcesses {
	return m.pi
}

func (m *workflow) Runtime() modules.IRuntime {
	return m.ri
}

func (m *workflow) Form() modules.IForm {
	return m.fi
}

func NewService(session *mgo.Session, pool *redis.Pool, client *indexutils.Client, log *gslogrus.Logger) *Workflow {
	return &Workflow{&workflow{
		shutdown: make(chan error),
		pool:     pool,
		session:  session,
		client:   client,
		log:      log,
	}}
}

func (w *Workflow) Run(zookeeperURL string) error {
	fmt.Println("starting...")
	id := gs_commons_generator.NewIDG()
	m := w.modules
	p := &processes{
		session:  m.session.Clone(),
		pool:     m.pool,
		id:       id,
		log:      m.log,
		relation: distribute.NewRelation(),
		client:   m.client,
	}
	h := &history{
		session: m.session.Clone(),
		pool:    m.pool,
		log:     m.log,
		id:      id,
		client:  m.client,
	}
	i := &instances{
		session: m.session.Clone(),
		pool:    m.pool,
		log:     m.log,
		id:      id,
		client:  m.client,
	}
	f := &form{
		session:   m.session.Clone(),
		pool:      m.pool,
		log:       m.log,
		id:        id,
		secretKey: "6333614dc0c7452eb3b29bed26a8580a",
		client:    m.client,
	}
	u := &user{
		log:    m.log,
		client: m.client,
	}
	r := &runtime{
		log:      m.log,
		shutdown: m.shutdown,
		conn:     gs_commons_config.NewConnect(zookeeperURL),
	}
	m.ui = u
	m.ii = i
	m.hi = h
	m.pi = p
	m.fi = f
	r.pipelines = newPipelines(m.session.Clone(), m.log, m.pool)
	r.modules = m
	r.dataGetter = distribute.NewDataGetter(m, m.log)
	r.processing = distribute.NewProcessing(m, m.log)
	r.next = distribute.NewNextflow(m, m.log, script.NewScript())
	fmt.Println("initialize ok...")
	return <-m.shutdown
}

func (w *Workflow) Modules() modules.Modules {
	return w.modules
}
