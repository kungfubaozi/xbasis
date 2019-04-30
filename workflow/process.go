package workflow

import (
	"github.com/go-redis/redis"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"konekko.me/gosion/commons/dto"
	"konekko.me/gosion/commons/errstate"
	"konekko.me/gosion/commons/generator"
)

const (
	STProcessOpening = 1

	STProcessDev = 2
)

type process struct {
	Id           string `bson:"_id" json:"id"`
	Key          string `bson:"key" json:"key"`
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

	Status int64 `bson:"status" json:"status"`

	ExpireTime int64 `bson:"expire_time" json:"expire_time"`
}

type pipeline struct {
	id         string
	name       string
	key        string
	flows      map[string][]*sequenceFlow
	nodes      map[string]*node
	startEvent interface{}
	startType  ConnectType
	endEvents  map[string]*simpleEndEvent
	lua        *luaScript
	expireAt   int64
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

//run lua script
func (pip *pipeline) script(flow *sequenceFlow, m map[string]interface{}) (bool, error) {
	ok, err := pip.lua.Run(flow.Script, m)
	if err != nil {
		return false, err
	}
	return ok, nil
}

func (pip *pipeline) status(id string) {

}

func (pip *pipeline) run(i *instance, nodes []*node, flowNodes flowNodes) {
	//load connect flow
	for _, v := range nodes {
		flow := pip.flows[v.id]
		flowNodes.Process(flow, v)
	}
}

//流程实例
//用来控制实例的走向等
type processes struct {
	pipelines map[string]*pipeline
	session   *mgo.Session
	conn      redis.Conn
	id        gs_commons_generator.IDGenerator
}

func (p *processes) reloadPipelineFromDB(processId string) (*gs_commons_dto.State, error) {

}

func (p *processes) open(pid string, open bool) (*gs_commons_dto.State, error) {
	process := &process{}
	err := p.collection().Find(bson.M{"_id": pid}).One(process)
	if err != nil {
		return nil, err
	}
	pip := &pipeline{
		id: process.Id,
	}
	if process.StartEvent != nil && (process.StartEventType == CTStartEvent || process.StartEventType == CTTimerStartEvent || process.StartEventType == CTMessageStartEvent) {
		pip.startEvent = process.StartEvent
		pip.startType = process.StartEventType
	} else {
		return ErrNoStartEvent, nil
	}
	if len(process.UserTasks) > 0 {
		for _, v := range process.UserTasks {
			node := &node{
				id:   v.Id,
				data: v,
				ct:   CTUserTask,
			}
			pip.nodes[v.Id] = node
		}
	}
	if len(process.HttpTasks) > 0 {
		for _, v := range process.HttpTasks {
			node := &node{
				id:   v.Id,
				data: v,
				ct:   CTHttpTask,
			}
			pip.nodes[v.Id] = node
		}
	}
	if len(process.SendTasks) > 0 {
		for _, v := range process.HttpTasks {
			node := &node{
				id:   v.Id,
				data: v,
				ct:   CTSendTask,
			}
			pip.nodes[v.Id] = node
		}
	}
	if len(process.DecisionTasks) > 0 {
		for _, v := range process.HttpTasks {
			node := &node{
				id:   v.Id,
				data: v,
				ct:   CTDecisionTask,
			}
			pip.nodes[v.Id] = node
		}
	}
	if len(process.Flows) > 0 {
		flows := make(map[string][]*sequenceFlow)
		for _, v := range process.Flows {
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
	if len(process.EndEvents) > 0 {

	} else {
		return ErrNoEndEvent, nil
	}
	return errstate.Success, nil
}

func (p *processes) add(process *process) (*gs_commons_dto.State, error) {
	//check name and key exists
	c := p.collection()
	count, err := c.Find(bson.M{"$or": []bson.M{
		{"name": process.Name}, {"key": process.Key},
	}}).Count()
	if err != nil {
		return nil, err
	}
	if count > 0 {
		return ErrProcessExists, nil
	}
	err = c.Insert(process)
	if err != nil {
		return nil, err
	}
	return errstate.Success, nil
}

func (p *processes) collection() *mgo.Collection {
	return p.session.DB("gs_workflow").C("processes")
}

func (p *processes) close() {
	p.session.Close()
}
