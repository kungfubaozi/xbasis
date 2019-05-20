package distribute

import (
	"context"
	"github.com/garyburd/redigo/redis"
	"konekko.me/gosion/commons/gslogrus"
	"konekko.me/gosion/workflow/flowerr"
	"konekko.me/gosion/workflow/models"
	"konekko.me/gosion/workflow/modules"
	"konekko.me/gosion/workflow/script"
	"konekko.me/gosion/workflow/types"
	"sync"
)

type nextflow struct {
	modules  modules.Modules
	store    modules.IStore
	finished map[string]bool
	log      *gslogrus.Logger
	values   []interface{}
	status   *models.NextStatus
	node     *models.Node
	ctx      context.Context
	instance *models.Instance
	script   Handler
	gon      bool
	pool     *redis.Pool
	call     types.CommandDataGetter
}

func (f *nextflow) SetCommandFunc(call types.CommandDataGetter) {
	f.call = call
}

func (f *nextflow) timerStartEvent() *flowerr.Error {
	panic("implement me")
}

func (f *nextflow) messageStartEvent() *flowerr.Error {
	panic("implement me")
}

func (f *nextflow) cancelEndEvent() *flowerr.Error {
	panic("implement me")
}

func (f *nextflow) terminateEndEvent() *flowerr.Error {
	panic("implement me")
}

func (f *nextflow) Data() interface{} {
	panic("implement me")
}

/**
当前是进行下一步流程的规定和辨识, 并不设计具体的操作，如 Submit
*/
func (f *nextflow) Do(ctx context.Context, instance *models.Instance, node *models.Node, ct types.ConnectType, value ...interface{}) (context.Context, *flowerr.Error) {
	f.values = value
	f.node = node
	f.ctx = ctx
	f.instance = instance
	f.status = &models.NextStatus{}
	return handler(ctx, ct, f)
}

func (f *nextflow) flows(callback func([]*models.SequenceFlow) *flowerr.Error) *flowerr.Error {
	var flows []*models.SequenceFlow
	v := f.values[1]
	if v != nil {
		flows = v.([]*models.SequenceFlow)
	}
	if flows == nil || len(flows) == 0 {
		return flowerr.ErrNode
	}
	return callback(flows)
}

//包容网关 所有满足的 没有走默认
//在入口方面 包容网关只会等待将会被执行的入口顺序流
func (f *nextflow) inclusiveGateway() *flowerr.Error {
	finished, ok := f.finished[f.node.Id]
	if !ok {
		finished = false
	}
	if !finished { //检查是否完成
		gateway, ok := f.node.Data.(*models.InclusiveGateway)
		if ok {
			var connect []*models.SequenceFlow
			var brsx []*models.NodeRelation
			var err *flowerr.Error
			resp := func(e *flowerr.Error) {
				if err == nil {
					err = e
				}
			}

			var wg sync.WaitGroup
			wg.Add(2)

			//如果有条件提前拿数据
			if gateway.ScriptFlows > 0 {
				wg.Add(1)
				go func() {
					defer wg.Done()
					data, err := f.call(types.GCNodeSubmitData, f.node.Id)
					if err != nil {
						resp(err)
						return
					}
					f1, ok := data.(map[string]interface{})
					if ok {
						for k, v1 := range f1 {
							f.metadata(k, v1)
						}
					} else {
						resp(flowerr.ErrUnknow)
					}
				}()
			}

			go func() {
				defer wg.Done()
				data, err := f.call(types.GCNodeFlows, f.node.Id)
				if err != nil {
					resp(err)
					return
				}
				f, ok := data.([]*models.SequenceFlow)
				if ok {
					connect = f
					return
				}
				resp(flowerr.ErrUnknow)
			}()

			go func() {
				defer wg.Done()
				bx, err := f.call(types.GCBackwardRelations, f.node.Id)
				if err != nil {
					resp(err)
					return
				}
				data, ok := bx.([]*models.NodeRelation)
				if ok {
					brsx = data
					return
				}
				resp(flowerr.ErrUnknow)
			}()

			wg.Wait()

			if err != nil {
				return err
			}

			okCount := 0

			i, err := f.store.GetInstanceIgnoreNodes(f.instance.Id)
			if err != nil {
				return err
			}

			//判断节点是否完成(反向关联的), 主要用于等待未完成的任务
			//store的作用是保存每个实例节点的状态（0：未完成，1：完成）
			wg.Add(len(brsx))
			for _, v := range brsx {
				//check finished
				c := false
				//检查忽略的节点，不需要进行检查
				for _, v1 := range i {
					if v.Id == v1 {
						c = true
						break
					}
				}
				if c {
					go func() {
						defer wg.Done()
						ok, err := f.store.IsFinished(v.Id, f.instance.Id)
						if err != nil {
							resp(err)
							return
						}
						if ok {
							okCount++
						}
					}()
				} else {
					wg.Done()
				}
			}

			wg.Wait()

			if err != nil {
				return err
			}

			//节点之前的任务没有完成，需要等待
			if okCount != len(brsx) {
				return nil
			}

			//next flow
			var defNode *models.SequenceFlow
			size := 0
			if !gateway.Exclusive {
				size++
			} else {
				size = gateway.ScriptFlows
			}

			if len(connect) > 0 {
				//clear about node store ignore nodes
				err := f.store.ClearAboutNodeIgnoreNodes(f.node.Id, f.instance.Id)
				if err != nil {
					return err
				}
			}

			var ignoreNodes []string
			var rollback string

			for _, v := range connect {
				if v.DefaultFlow && defNode == nil {
					defNode = v
				}
				//把剩下的节点都添加进来
				if gateway.Exclusive && len(f.status.Again) > 0 {
					ignoreNodes = append(ignoreNodes, v.End)
					continue
				}
				if len(v.Script) > 0 {
					//get target node
					node, err := f.call(types.GCNode, v.End)
					if err != nil {
						return err
					}

					n, ok := node.(*models.Node)
					if ok {
						ctx, err := f.script.Do(f.ctx, f.instance, nil, v.EndType, f, n)
						f.context(ctx)
						if err != nil {
							if err == flowerr.ScriptFalse {
								//添加到忽略节点里
								ignoreNodes = append(ignoreNodes, v.End)
							} else if err == flowerr.ScriptTrue || err == flowerr.NextFlow {
								if v.Rollback {
									rollback = v.End
									break
								}
								f.again(v.End)
							} else {
								return err
							}
						}
					}
				} else {
					f.again(v.End)
				}
			}

			//这一步的操作是，如果发现了是rollback操作，则清除rollback对应节点关联的所有节点（向前）的状态（未完成）
			if len(rollback) > 0 {
				//clear store
				ok, err := f.store.ClearRelationNodesStatus(rollback, f.instance.Id)
				if err != nil {
					return err
				}
				if ok {
					f.again(rollback)
					return flowerr.ErrRollback
				}
				return flowerr.ErrSystem
			}

			ignores := &models.NodeIgnore{
				InstanceId: f.instance.Id,
				GatewayId:  f.node.Id,
			}

			//添加忽略的节点
			addIgnoresNodes := func(def bool) *flowerr.Error {
				if def {
					for i, v := range ignoreNodes {
						if v == defNode.End {
							ignoreNodes = append(ignoreNodes[:i], ignoreNodes[i+1:]...)
							break
						}
					}
				}
				ignores.Ignores = ignoreNodes
				return f.store.AddIgnoreNode(ignores)
			}

			//执行默认的节点
			if defNode != nil && len(f.status.Again) == 0 {
				f.again(defNode.End)
				return addIgnoresNodes(true)
			}

			if gateway.Exclusive {
				return addIgnoresNodes(false)
			}

			return nil
		}

		return flowerr.ErrNode
	}

	return nil
}

//如果nextflow跳转到开始节点，那么需要停止当前实例，
func (f *nextflow) startEvent() *flowerr.Error {
	panic("implement me")
}

//如果到endevent会停止所有instance的交互
func (f *nextflow) endEvent() *flowerr.Error {
	panic("implement me")
}

func (f *nextflow) apiStartEvent() *flowerr.Error {
	panic("implement me")
}

func (f *nextflow) userTask() *flowerr.Error {
	f.next(f.node.Id)
	e := f.node.Data.(*models.UserTask)
	//如果有用户，则通知TA们
	f.modules.User().Notify(f.ctx, e)
	return nil
}

func (f *nextflow) notifyTask() *flowerr.Error {
	panic("implement me")
}

func (f *nextflow) triggerStartEvent() *flowerr.Error {
	panic("implement me")
}

func (f *nextflow) context(ctx context.Context) context.Context {
	if ctx != nil {
		f.ctx = ctx
	}
	return f.ctx
}

func (f *nextflow) metadata(key string, data interface{}) {
	panic("implement me")
}

func (f *nextflow) again(id string) {
	f.status.Again = append(f.status.Again, id)
}

func (f *nextflow) next(id string) {
	f.status.CurrentNodes = append(f.status.CurrentNodes, id)
}

func (f *nextflow) flow() *models.SequenceFlow {
	return f.values[0].(*models.SequenceFlow)
}

func (f *nextflow) Restore() {
	f.status = nil
	f.finished = make(map[string]bool)
}

func NewNextflow(modules modules.Modules, log *gslogrus.Logger, script *script.LuaScript, pool *redis.Pool, store modules.IStore) Handler {
	return &nextflow{modules: modules, log: log, script: newScript(modules, log, script), pool: pool, store: store}
}
