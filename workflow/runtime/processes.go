package runtime

import (
	"fmt"
	"github.com/garyburd/redigo/redis"
	"gopkg.in/mgo.v2"
	"konekko.me/gosion/commons/generator"
	"konekko.me/gosion/commons/gslogrus"
	"konekko.me/gosion/commons/indexutils"
	"konekko.me/gosion/workflow/distribute"
	"konekko.me/gosion/workflow/flowerr"
	"konekko.me/gosion/workflow/models"
	"konekko.me/gosion/workflow/types"
)

type processes struct {
	session  *mgo.Session
	pool     *redis.Pool
	log      *gslogrus.Logger
	client   *indexutils.Client
	relation distribute.Handler
	id       gs_commons_generator.IDGenerator
}

func (pro *processes) FindNode(instanceId, nodeId string) (interface{}, *flowerr.Error) {
	panic("implement me")
}

func (pro *processes) Reassignment() {
	panic("implement me")
}

//TODO 修复反向连接的问题
func (pro *processes) AddProcess(p *models.Process) {
	pip := &pipeline{
		ProcessId:         p.Id,
		Name:              p.Name,
		EndEvents:         make(map[string]*models.TypeEvent),
		SequenceFlows:     make(map[string][]*models.SequenceFlow),
		Nodes:             make(map[string]*models.Node),
		Parallels:         make(map[string][]string),
		BackwardRelations: make(map[string][]*models.NodeBackwardRelation),
	}
	if p.StartEvent != nil {
		pip.StartEvent = &models.Node{
			CT:   p.StartEvent.Type,
			Id:   p.StartEvent.Id,
			Key:  p.StartEvent.Key,
			Data: p.StartEvent.Event,
		}
		pip.append(pip.StartEvent)
	} else {
		panic("no start event")
	}
	if p.EndEvents != nil && len(p.EndEvents) > 0 {
		for _, v := range p.EndEvents {
			if len(v.Id) == 0 {
				panic("err id")
			}
			pip.EndEvents[v.Id] = v
			pip.append(&models.Node{
				CT:  v.Type,
				Id:  v.Id,
				Key: v.Key,
			})
		}
	} else {
		panic("no end event")
	}

	backflows := make(map[string][]*temporary)

	if p.Tasks != nil {
		if len(p.Tasks.StorageTasks) > 0 {
			for _, v := range p.Tasks.StorageTasks {
				n := &models.Node{
					Id:   v.Id,
					Key:  v.Key,
					Data: v,
					CT:   types.CTStorageTask,
				}
				pip.append(n)
			}
		}
		if len(p.Tasks.NotifyTasks) > 0 {
			for _, v := range p.Tasks.NotifyTasks {
				n := &models.Node{
					Id:   v.Id,
					Key:  v.Key,
					Data: v,
					CT:   types.CTSendTask,
				}
				pip.append(n)
			}
		}
		if len(p.Tasks.MailTasks) > 0 {
			for _, v := range p.Tasks.MailTasks {
				n := &models.Node{
					Id: v.Id, Key: v.Key,
					Data: v,
					CT:   types.CTMailTask,
				}
				pip.append(n)
			}
		}
		if len(p.Tasks.HttpTasks) > 0 {
			for _, v := range p.Tasks.HttpTasks {
				n := &models.Node{
					Id: v.Id, Key: v.Key,
					Data: v,
					CT:   types.CTHttpTask,
				}
				pip.append(n)
			}
		}
		if len(p.Tasks.DecisionTasks) > 0 {
			for _, v := range p.Tasks.DecisionTasks {
				n := &models.Node{
					Id: v.Id, Key: v.Key,
					Data: v,
					CT:   types.CTDecisionTask,
				}
				pip.append(n)
			}
		}
		if len(p.Tasks.UserTasks) > 0 {
			for _, v := range p.Tasks.UserTasks {
				n := &models.Node{
					Id: v.Id, Key: v.Key,
					Data: v,
					CT:   types.CTUserTask,
				}
				pip.append(n)
			}
		}
		if len(p.Tasks.ApiTasks) > 0 {
			for _, v := range p.Tasks.ApiTasks {
				n := &models.Node{
					Id: v.Id, Key: v.Key,
					Data: v,
					CT:   types.CTApiTask,
				}
				pip.append(n)
			}
		}
	}

	if p.Gateways != nil {
		if len(p.Gateways.Inclusive) > 0 {
			for _, v := range p.Gateways.Inclusive {
				n := &models.Node{Key: v.Key,
					CT:   types.CTInclusiveGateway,
					Id:   v.Id,
					Data: v,
				}
				pip.append(n)
			}
		}
	}

	if len(p.Flows) > 0 {
		for _, v := range p.Flows {
			//正向
			f := pip.SequenceFlows[v.Start]
			v.Key = pip.Nodes[v.Start].Key
			pip.SequenceFlows[v.Start] = append(f, v)

			//把流程转换方向
			bf := backflows[v.End]

			t := &temporary{
				start:     v.End,
				key:       v.Key,
				end:       v.Start,
				startType: v.EndType,
				endType:   v.StartType,
			}

			backflows[v.End] = append(bf, t)
		}
	}

	if len(backflows) > 0 {
		pro.backRelation(pip, backflows)
	}

	pro.condition(pip)

	fmt.Println("load process", pip.ProcessId)

}

type temporary struct {
	start     string
	end       string
	key       string
	startType types.ConnectType
	endType   types.ConnectType
}

func (pro *processes) condition(pip *pipeline) {
	for _, v := range pip.Nodes {
		if v.CT == types.CTInclusiveGateway {
			//get flow
			flows, err := pip.Flows(v.Id)
			if err == nil && len(flows) > 0 {
				size := 0
				for _, v1 := range flows {
					if len(v1.Script) > 0 {
						size++
					}
				}
				//update target node
				gateway, ok := v.Data.(*models.InclusiveGateway)
				if ok {
					gateway.ScriptFlows = size
					pip.Nodes[v.Id] = &models.Node{
						CT:   v.CT,
						Key:  v.Key,
						Data: gateway,
						Id:   v.Id,
					}
				}
			}
		}
	}
}

func (pro *processes) backRelation(pip *pipeline, backflows map[string][]*temporary) *flowerr.Error {
	relations := make(map[string][]*models.NodeBackwardRelation)
	fmt.Println("size", len(pip.Nodes))
	for _, v := range pip.Nodes {

		var nrs []*models.NodeBackwardRelation

		rs, err := pro.loopFind(v.Id, backflows)
		if err != nil {
			return err
		}
		nrs = append(nrs, rs...)

		relations[v.Id] = nrs
	}
	pip.BackwardRelations = relations

	return nil
}

func (pro *processes) loopFind(nodeId string, backflows map[string][]*temporary) ([]*models.NodeBackwardRelation, *flowerr.Error) {
	var res []*models.NodeBackwardRelation
	v := backflows[nodeId]
	for _, r := range v {
		//在部分节点上和节点连接的数量有限制
		_, err := pro.relation.Do(nil, nil, nil, r.endType, len(v))
		if err == nil {
			res = append(res, &models.NodeBackwardRelation{
				Id:  r.end,
				Key: r.key,
				CT:  r.endType,
			})
		} else if err == flowerr.NextFlow {
			s, err := pro.loopFind(r.end, backflows)
			if err != nil {
				return nil, err
			}
			res = append(res, s...)
		}
	}
	return res, nil
}
