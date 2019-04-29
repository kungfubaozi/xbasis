package workflow

import (
	"github.com/go-redis/redis"
	"github.com/pkg/errors"
	"gopkg.in/mgo.v2"
	"konekko.me/gosion/commons/dto"
	"konekko.me/gosion/commons/errstate"
	"konekko.me/gosion/commons/generator"
)

type process struct {
	Id           string `bson:"_id" json:"id"`
	Name         string `bson:"name" json:"name"`
	CreateUserId string `bson:"create_user_id" json:"create_user_id"`
	CreateAt     int64  `bson:"create_at" json:"create_at"`
	Desc         string `bson:"desc" json:"desc"`
	//connect flows
	Flows []*sequenceFlow `bson:"flows" json:"flows"`
	//user tasks
	UserTasks []*userTask `bson:"user_tasks" json:"user_tasks"`
	//http tasks
	HttpTasks []*httpTask `bson:"http_tasks" json:"http_tasks"`
	//decision tasks
	DecisionTasks []*decisionTask `bson:"decision_tasks" json:"decision_tasks"`
	//send tasks
	SendTasks []*sendTask `bson:"send_tasks" json:"send_tasks"`
	//grpc task
	GrpcTask []*grpcTask `bson:"grpc_task" json:"grpc_task"`
	//mail task
	MailTask []*mailTask `bson:"mail_task" json:"mail_task"`
	//version control
	Version int64 `bson:"version" json:"version"`

	Gateways []interface{} `bson:"gateways" json:"gateways"`

	StartEvent     interface{} `bson:"start_event" json:"start_event"`
	StartEventType ConnectType `bson:"start_event_type" json:"start_event_type"`

	EndEvents map[string]interface{} `bson:"end_events" json:"end_events"`
}

type pipeline struct {
	id         string
	flows      map[string][]*sequenceFlow
	nodes      map[string]*node
	startEvent interface{}
	startType  ConnectType
	endEvents  map[string]*simpleEndEvent
	lua        *luaScript
}

type simpleEndEvent struct {
	id   string
	ct   ConnectType
	data interface{}
}

type node struct {
	id   string
	data interface{}
	ct   ConnectType
}

func (pip *pipeline) currents(instanceId string) []string {

}

func (pip *pipeline) aboutFlows(nodeId string) []*sequenceFlow {
	return pip.flows[nodeId]
}

func (pip *pipeline) finished(instanceId, nodeId string) {

}

//run lua script
func (pip *pipeline) script(flow *sequenceFlow, m map[string]interface{}) (bool, error) {
	ok, err := pip.lua.Run(flow.Script, m)
	if err != nil {
		return false, err
	}
	return ok, nil
}

//获取初始节点提交的数据和当前节点前一个task产生的数据
func (pip *pipeline) lastData(instanceId string) map[string]interface{} {

}

func (pip *pipeline) waiting() {

}

//检查自身节点是否有错误（比如没有提交form等）
func (pip *pipeline) checkSelf(n *node, instanceId string) (*gs_commons_dto.State, error) {

}

func (pip *pipeline) process(n *node, instanceId, nodeId string) (*gs_commons_dto.State, error) {
	switch n.ct {
	case FTEventGateway:
		break
	case FTExclusiveGateway:
		flows := pip.aboutFlows(nodeId)
		if len(flows) > 0 {
			for _, v := range flows {
				var ok bool
				var err error
				if len(v.Script) > 0 {
					ok, err = pip.script(v, pip.lastData(instanceId))
					if err != nil {
						return errstate.ErrRequest, err
					}
				}
				if ok {
					return pip.process(n, instanceId, v.End)
				}
			}
			return ErrNoOperatingConditions, nil
		}
		break
	case FTParallelGateway:
		break
	case FTInclusiveGateway:
		break
	case FTEndEvent:
		break
	case FTEndCancelEvent:
		break
	case FTEndErrorEvent:
		break
	case FTStartEvent:
		flows := pip.aboutFlows(nodeId)

		pip.checkSelf(n, instanceId)

		if len(flows) > 0 {
			for _, v := range flows {
				pip.process(n, instanceId, v.End)
			}
		}
		break
	case FTMessageStartEvent:
		break
	case FTTimerStartEvent:
		break
	case FTGRPCTask:
		break
	case FTDecisionTask:
		break
	case FTHttpTask:
		break
	case FTUserTask:
		break
	case FTMailTask:
		break
	}
}

func (pi *processes) next(i *instance) error {
	pip := pi.processes[i.processId]
	if pip == nil {
		return errors.New("not found")
	}

	currentNodes := pip.currents(i.id)

	for _, nodeId := range currentNodes {
		pip.process(pip.nodes[nodeId], i.id, nodeId)
	}

}

//流程实例
//用来控制实例的走向等
type processes struct {
	processes map[string]*pipeline
	session   *mgo.Session
	conn      redis.Conn
	id        gs_commons_generator.IDGenerator
}

func (pi *processes) generate(p *process) error {
	pip := &pipeline{
		id: p.Id,
	}
	if p.StartEvent != nil && (p.StartEventType == FTStartEvent || p.StartEventType == FTTimerStartEvent || p.StartEventType == FTMessageStartEvent) {
		pip.startEvent = p.StartEvent
		pip.startType = p.StartEventType
	} else {
		return errors.New("err start event.")
	}
	if len(p.UserTasks) > 0 {
		for _, v := range p.UserTasks {
			node := &node{
				id:   v.Basic.Id,
				data: v,
				ct:   FTUserTask,
			}
			pip.nodes[v.Basic.Id] = node
		}
	}
	if len(p.HttpTasks) > 0 {
		for _, v := range p.HttpTasks {
			node := &node{
				id:   v.Basic.Id,
				data: v,
				ct:   FTHttpTask,
			}
			pip.nodes[v.Basic.Id] = node
		}
	}
	if len(p.SendTasks) > 0 {
		for _, v := range p.HttpTasks {
			node := &node{
				id:   v.Basic.Id,
				data: v,
				ct:   FTSendTask,
			}
			pip.nodes[v.Basic.Id] = node
		}
	}
	if len(p.DecisionTasks) > 0 {
		for _, v := range p.HttpTasks {
			node := &node{
				id:   v.Basic.Id,
				data: v,
				ct:   FTDecisionTask,
			}
			pip.nodes[v.Basic.Id] = node
		}
	}
	if len(p.Flows) > 0 {
		flows := make(map[string][]*sequenceFlow)
		for _, v := range p.Flows {
			if flows[v.Start] != nil {
				a := flows[v.Start]
				a = append(a, v)
				flows[v.Start] = a
				break
			}
			a := make([]*sequenceFlow, 1)
			flows[v.Start] = append(a, v)
		}
	}
	pi.processes[pip.id] = pip
	return nil
}

func (pi *processes) close() {
	pi.session.Close()
}
