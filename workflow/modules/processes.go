package modules

import (
	"fmt"
	"github.com/garyburd/redigo/redis"
	"gopkg.in/mgo.v2"
	"konekko.me/gosion/commons/generator"
	"konekko.me/gosion/commons/gslogrus"
	"konekko.me/gosion/commons/indexutils"
	"konekko.me/gosion/workflow/flowerr"
	"konekko.me/gosion/workflow/models"
	"konekko.me/gosion/workflow/types"
)

type AddProcessToPipelineCallback func(pip Pipeline)

type IProcesses interface {
	AddProcess(p *models.Process)

	SetCallback(callback AddProcessToPipelineCallback)

	//重新分配某个节点的操作人
	Reassignment()

	LoadAll()

	FindNode(instanceId, nodeId string) (interface{}, *flowerr.Error)
}

type processes struct {
	callback AddProcessToPipelineCallback
	session  *mgo.Session
	pool     *redis.Pool
	log      *gslogrus.Logger
	client   *indexutils.Client
	size     int64
	id       gs_commons_generator.IDGenerator
}

func (pro *processes) FindNode(instanceId, nodeId string) (interface{}, *flowerr.Error) {
	panic("implement me")
}

func (pro *processes) Reassignment() {
	panic("implement me")
}

func (pro *processes) LoadAll() {
	fmt.Println("Goflow start load all process...")
	fmt.Println("load done, process size", pro.size)
}

func (pro *processes) SetCallback(callback AddProcessToPipelineCallback) {
	if pro.callback == nil {
		pro.callback = callback
	}
}

func (pro *processes) AddProcess(p *models.Process) {
	if pro.callback != nil {
		pip := &pipeline{
			id:        p.Id,
			name:      p.Name,
			endEvents: make(map[string]*models.TypeEvent),
			flows:     make(map[string][]*models.SequenceFlow),
		}
		if p.StartEvent != nil {
			pip.startEvent = p.StartEvent.Event
			pip.startType = p.StartEvent.Type
		} else {
			panic("no start event")
		}
		if p.EndEvents != nil && len(p.EndEvents) > 0 {
			for _, v := range p.EndEvents {
				if len(v.Id) == 0 {
					panic("err id")
				}
				pip.endEvents[v.Id] = v
			}
		} else {
			panic("no end event")
		}

		if len(p.Flows) > 0 {
			for _, v := range p.Flows {
				f := pip.flows[v.Start]
				if f == nil {
					pip.flows[v.Start] = []*models.SequenceFlow{v}
					continue
				}
				f = append(f, v)
			}
		}
		if p.Tasks != nil {
			if len(p.Tasks.StorageTasks) > 0 {
				for _, v := range p.Tasks.StorageTasks {
					n := &node{
						id:   v.Id,
						key:  v.Key,
						data: v,
						ct:   types.CTStorageTask,
					}
					pip.append(n)
				}
			}
			if len(p.Tasks.NotifyTasks) > 0 {
				for _, v := range p.Tasks.NotifyTasks {
					n := &node{
						id:   v.Id,
						key:  v.Key,
						data: v,
						ct:   types.CTSendTask,
					}
					pip.append(n)
				}
			}
			if len(p.Tasks.MailTasks) > 0 {
				for _, v := range p.Tasks.MailTasks {
					n := &node{
						id: v.Id, key: v.Key,
						data: v,
						ct:   types.CTMailTask,
					}
					pip.append(n)
				}
			}
			if len(p.Tasks.HttpTasks) > 0 {
				for _, v := range p.Tasks.HttpTasks {
					n := &node{
						id: v.Id, key: v.Key,
						data: v,
						ct:   types.CTHttpTask,
					}
					pip.append(n)
				}
			}
			if len(p.Tasks.DecisionTasks) > 0 {
				for _, v := range p.Tasks.DecisionTasks {
					n := &node{
						id: v.Id, key: v.Key,
						data: v,
						ct:   types.CTDecisionTask,
					}
					pip.append(n)
				}
			}
			if len(p.Tasks.UserTasks) > 0 {
				for _, v := range p.Tasks.UserTasks {
					n := &node{
						id: v.Id, key: v.Key,
						data: v,
						ct:   types.CTUserTask,
					}
					pip.append(n)
				}
			}
			if len(p.Tasks.ApiTasks) > 0 {
				for _, v := range p.Tasks.ApiTasks {
					n := &node{
						id: v.Id, key: v.Key,
						data: v,
						ct:   types.CTApiTask,
					}
					pip.append(n)
				}
			}
		}

		if p.Gateways != nil {
			if len(p.Gateways.Exclusives) > 0 {
				for _, v := range p.Gateways.Exclusives {
					n := &node{key: v.Key,
						ct:   types.CTExclusiveGateway,
						id:   v.Id,
						data: v,
					}
					pip.append(n)
				}
			}
			if len(p.Gateways.Inclusive) > 0 {
				for _, v := range p.Gateways.Inclusive {
					n := &node{key: v.Key,
						ct:   types.CTInclusiveGateway,
						id:   v.Id,
						data: v,
					}
					pip.append(n)
				}
			}
			if len(p.Gateways.Parallels) > 0 {
				for _, v := range p.Gateways.Parallels {
					n := &node{key: v.Key,
						ct:   types.CTParallelGateway,
						id:   v.Id,
						data: v,
					}
					pip.append(n)
				}
			}
		}
		pro.size++
		fmt.Println("load process", pip.id)
		pro.callback(pip)
	}
}
