package modules

import (
	"fmt"
	"github.com/garyburd/redigo/redis"
	"gopkg.in/mgo.v2"
	"konekko.me/gosion/commons/generator"
	"konekko.me/gosion/commons/gslogrus"
	"konekko.me/gosion/commons/indexutils"
	"konekko.me/gosion/workflow/script"
)

type Modules interface {
	History() IHistory

	Instance() IInstance

	Process() IProcesses

	Runtime() IRuntime

	Form() IForm

	User() IUser
}

type Workflow struct {
	modules *modules
}

type modules struct {
	shutdown chan error
	session  *mgo.Session
	pool     *redis.Pool
	client   *indexutils.Client
	log      *gslogrus.Logger
	//
	ri IRuntime
	hi IHistory
	ii IInstance
	pi IProcesses
	fi IForm
	ui IUser
}

func (m *modules) User() IUser {
	panic("implement me")
}

func (m *modules) History() IHistory {
	return m.hi
}

func (m *modules) Instance() IInstance {
	return m.ii
}

func (m *modules) Process() IProcesses {
	return m.pi
}

func (m *modules) Runtime() IRuntime {
	return m.ri
}

func (m *modules) Form() IForm {
	return m.fi
}

func NewService(session *mgo.Session, pool *redis.Pool, client *indexutils.Client, log *gslogrus.Logger) *Workflow {
	return &Workflow{&modules{
		shutdown: make(chan error),
		pool:     pool,
		session:  session,
		client:   client,
		log:      log,
	}}
}

func (w *Workflow) Run() error {
	id := gs_commons_generator.NewIDG()
	m := w.modules
	callback, r := createRuntime(m.shutdown, m.log)
	p := &processes{
		session: m.session.Clone(),
		pool:    m.pool,
		id:      id,
		log:     m.log,
		client:  m.client,
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
		session: m.session.Clone(),
		pool:    m.pool,
		log:     m.log,
		id:      id,
		client:  m.client,
	}
	u := &user{
		log:    m.log,
		client: m.client,
	}
	m.ui = u
	m.ii = i
	m.hi = h
	m.ri = r
	m.pi = p
	m.fi = f
	r.modules = m
	r.processing = newProcessing(m, m.log)
	r.next = newNextflow(m, m.log, script.NewScript())
	m.pi.SetCallback(callback)
	fmt.Println("Goflow initialize ok...")
	//加载所有流程
	m.pi.LoadAll()
	return <-m.shutdown
}

func (w *Workflow) Modules() Modules {
	return w.modules
}
