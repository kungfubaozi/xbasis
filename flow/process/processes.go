package process

import (
	"github.com/garyburd/redigo/redis"
	"gopkg.in/mgo.v2"
	"konekko.me/gosion/commons/indexutils"
	"konekko.me/gosion/flow/flow"
	"konekko.me/gosion/flow/types"
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
		pip := &pipeline{
			id:        p.Id,
			name:      p.Name,
			endEvents: make(map[string]*TypeEvent),
			flows:     make(map[string][]*flow.SequenceFlow),
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
					pip.flows[v.Start] = []*flow.SequenceFlow{v}
					continue
				}
				f = append(f, v)
			}
		}
		if p.Tasks != nil {
			if len(p.Tasks.StorageTasks) > 0 {
				for _, v := range p.Tasks.StorageTasks {
					n := &Node{
						Id:   v.Id,
						Data: v,
						CT:   types.CTStorageTask,
					}
					pip.append(n)
				}
			}
			if len(p.Tasks.SendTasks) > 0 {
				for _, v := range p.Tasks.SendTasks {
					n := &Node{
						Id:   v.Id,
						Data: v,
						CT:   types.CTSendTask,
					}
					pip.append(n)
				}
			}
			if len(p.Tasks.MailTasks) > 0 {
				for _, v := range p.Tasks.MailTasks {
					n := &Node{
						Id:   v.Id,
						Data: v,
						CT:   types.CTMailTask,
					}
					pip.append(n)
				}
			}
			if len(p.Tasks.HttpTasks) > 0 {
				for _, v := range p.Tasks.HttpTasks {
					n := &Node{
						Id:   v.Id,
						Data: v,
						CT:   types.CTHttpTask,
					}
					pip.append(n)
				}
			}
			if len(p.Tasks.GRPCTasks) > 0 {
				for _, v := range p.Tasks.GRPCTasks {
					n := &Node{
						Id:   v.Id,
						Data: v,
						CT:   types.CTGRPCTask,
					}
					pip.append(n)
				}
			}
			if len(p.Tasks.DecisionTasks) > 0 {
				for _, v := range p.Tasks.DecisionTasks {
					n := &Node{
						Id:   v.Id,
						Data: v,
						CT:   types.CTDecisionTask,
					}
					pip.append(n)
				}
			}
			if len(p.Tasks.UserTasks) > 0 {
				for _, v := range p.Tasks.UserTasks {
					n := &Node{
						Id:   v.Id,
						Data: v,
						CT:   types.CTUserTask,
					}
					pip.append(n)
				}
			}
			if len(p.Tasks.ApiTasks) > 0 {
				for _, v := range p.Tasks.ApiTasks {
					n := &Node{
						Id:   v.Id,
						Data: v,
						CT:   types.CTApiTask,
					}
					pip.append(n)
				}
			}
		}

		if p.Gateways != nil {
			if len(p.Gateways.Exclusives) > 0 {
				for _, v := range p.Gateways.Exclusives {
					n := &Node{
						CT:   types.CTExclusiveGateway,
						Id:   v.Id,
						Data: v,
					}
					pip.append(n)
				}
			}
			if len(p.Gateways.Inclusive) > 0 {
				for _, v := range p.Gateways.Inclusive {
					n := &Node{
						CT:   types.CTInclusiveGateway,
						Id:   v.Id,
						Data: v,
					}
					pip.append(n)
				}
			}
			if len(p.Gateways.Parallels) > 0 {
				for _, v := range p.Gateways.Parallels {
					n := &Node{
						CT:   types.CTParallelGateway,
						Id:   v.Id,
						Data: v,
					}
					pip.append(n)
				}
			}
		}

		pro.Callback(pip)
	}
}
