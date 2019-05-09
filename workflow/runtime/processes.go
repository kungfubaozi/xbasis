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
	"konekko.me/gosion/workflow/modules"
	"konekko.me/gosion/workflow/types"
)

type processes struct {
	callback modules.AddProcessToPipelineCallback
	session  *mgo.Session
	pool     *redis.Pool
	log      *gslogrus.Logger
	client   *indexutils.Client
	size     int64
	relation distribute.Handler
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

func (pro *processes) SetCallback(callback modules.AddProcessToPipelineCallback) {
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
			if len(p.Gateways.Exclusives) > 0 {
				for _, v := range p.Gateways.Exclusives {
					n := &models.Node{Key: v.Key,
						CT:   types.CTExclusiveGateway,
						Id:   v.Id,
						Data: v,
					}
					pip.append(n)
				}
			}
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
			if len(p.Gateways.Parallels) > 0 {
				for _, v := range p.Gateways.Parallels {
					n := &models.Node{Key: v.Key,
						CT:   types.CTParallelGateway,
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
				f := pip.flows[v.Start]
				if f == nil {
					pip.flows[v.Start] = []*models.SequenceFlow{v}
					continue
				}
				pip.flows[v.Start] = append(f, v)

				//把流程转换方向
				bf := backflows[v.End]

				t := &temporary{
					start:     v.End,
					key:       v.Key,
					end:       v.Start,
					startType: v.EndType,
					endType:   v.EndType,
				}

				if bf == nil {
					backflows[v.End] = []*temporary{t}
					continue
				}
				backflows[v.End] = append(bf, t)
			}
		}

		if len(backflows) > 0 {
			pro.backRelation(pip, backflows)
		}

		pro.size++
		fmt.Println("load process", pip.id)
		pro.callback(pip)
	}
}

type temporary struct {
	start     string
	end       string
	key       string
	startType types.ConnectType
	endType   types.ConnectType
}

func (pro *processes) backRelation(pip *pipeline, backflows map[string][]*temporary) *flowerr.Error {
	relations := make(map[string][]*models.NodeBackwardRelation)
	for _, v := range pip.nodes {
		flows := backflows[v.Id]
		var nrs []*models.NodeBackwardRelation
		for _, v1 := range flows {
			rs, err := pro.loopFind(v1.start, backflows)
			if err != nil {
				return err
			}
			nrs = append(nrs, rs...)
		}
		relations[v.Id] = nrs
	}
	pip.backwardRelations = relations
	return nil
}

func (pro *processes) loopFind(nodeId string, backflows map[string][]*temporary) ([]*models.NodeBackwardRelation, *flowerr.Error) {
	var res []*models.NodeBackwardRelation
	v := backflows[nodeId]
	for _, r := range v {
		//在部分节点上和节点连接的数量有限制
		_, err := pro.relation.Do(nil, nil, nil, r.startType, len(v))
		if err == nil {
			res = append(res, &models.NodeBackwardRelation{
				Id:  r.end,
				Key: r.key,
				CT:  r.endType,
			})
		} else if err == flowerr.NextFlow {
			s, err := pro.loopFind(r.start, backflows)
			if err != nil {
				return nil, err
			}
			res = append(res, s...)
		}
	}
	return res, nil
}
