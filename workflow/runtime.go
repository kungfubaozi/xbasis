package workflow

import (
	"github.com/garyburd/redigo/redis"
	"gopkg.in/mgo.v2"
	"konekko.me/gosion/commons/dto"
	"konekko.me/gosion/commons/indexutils"
)

type runtime struct {
	script    *luaScript
	processes *processes
	history   *history
	form      *form
	instances *instances
	session   *mgo.Session
	pool      *redis.Pool
	client    *indexutils.Client
}

func newRuntime(session *mgo.Session, pool *redis.Pool, client *indexutils.Client) *runtime {
	r := &runtime{
		processes: &processes{
			pipelines: make(map[string]*pipeline),
			session:   session.Clone(),
		},
		instances: &instances{
			session: session.Clone(),
			conn:    pool.Get(),
			client:  client,
		},
	}
	return r
}

func (rt *runtime) add(p *process) (*gs_commons_dto.State, error) {
	if len(p.Name) < 2 {
		return ErrProcessName, nil
	}
	if len(p.Key) == 0 {
		return ErrProcessKey, nil
	}
	return rt.processes.add(p)
}

//load all process to processes pipeline
func (rt *runtime) load() {

}

func (rt *runtime) open(processId string, open bool) (*gs_commons_dto.State, error) {
	return rt.processes.open(processId, open)
}

//to instanceId
func (rt *runtime) startNewProcess(processId string) (string, error) {

}

func (rt *runtime) submit(instanceId string, nodeId string) {

}

func (rt *runtime) next(i *operate) (*gs_commons_dto.State, error) {
	ins, err := rt.instances.get(i.instanceId)
	if err != nil {
		return nil, err
	}

	if ins.Status == INSFinished {
		return ErrInstanceAlreadyFinished, nil
	}

	var pip *pipeline
	load := func() {
		pip = rt.processes.pipelines[ins.ProcessId]
	}
	load()
	if pip == nil {
		a, err := rt.processes.reloadPipelineFromDB(ins.ProcessId)
		if err != nil {
			return nil, err
		}
		if !a.Ok {
			return a, nil
		}
		load()
	}

	//load node
	var nodes []*node
	for _, v := range ins.CurrentNodeIds {
		n := pip.nodes[v]
		//check node is completed
		a, err := rt.history.isCompleted(ins.Id, v)
		if err != nil {
			return nil, err
		}
		if !a.Ok {
			return a, nil
		}
		nodes = append(nodes, n)
	}

	//process
	pip.run(ins, nodes, &nextFlowControl{ins: ins, pip: pip, end: false, records: make(map[string]bool)})

}
