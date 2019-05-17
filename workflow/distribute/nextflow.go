package distribute

import (
	"context"
	"github.com/garyburd/redigo/redis"
	"github.com/vmihailenco/msgpack"
	"konekko.me/gosion/commons/gslogrus"
	"konekko.me/gosion/workflow/flowerr"
	"konekko.me/gosion/workflow/models"
	"konekko.me/gosion/workflow/modules"
	"konekko.me/gosion/workflow/script"
	"konekko.me/gosion/workflow/types"
	"sync"
)

type nextflow struct {
	modules            modules.Modules
	finished           map[string]bool
	log                *gslogrus.Logger
	values             []interface{}
	status             *models.NextStatus
	node               *models.Node
	ctx                context.Context
	instance           *models.Instance
	script             Handler
	gon                bool
	ignoreNodes        []string
	ignoreNodesChanged bool
	pool               *redis.Pool
	conn               redis.Conn
	call               types.CommandDataGetter
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
	if f.conn == nil {
		f.conn = f.pool.Get()
	}
	if f.ignoreNodes == nil {
		v, err := redis.Bytes(f.conn.Do("get", instance.Id))
		if err != nil && err == redis.ErrNil {
			err = nil
			f.ignoreNodes = []string{}
		}
		if err != nil {
			return ctx, flowerr.FromError(err)
		}
		err = msgpack.Unmarshal(v, f.ignoreNodes)
		if err != nil {
			return ctx, flowerr.FromError(err)
		}
		if f.ignoreNodes == nil {
			f.ignoreNodes = []string{}
		}
	}
	return handler(ctx, ct, f)
}

func (f *nextflow) updateIgnoreNodes() *flowerr.Error {
	if len(f.ignoreNodes) > 0 {
		v, err := msgpack.Marshal(f.ignoreNodes)
		if err != nil {
			return flowerr.FromError(err)
		}
		_, err = f.conn.Do("set", f.instance.Id, v)
		if err != nil {
			return flowerr.FromError(err)
		}
	}
	return nil
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
			var flows []*models.SequenceFlow
			var brsx []*models.NodeBackwardRelation
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
					flows = f
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
				data, ok := bx.([]*models.NodeBackwardRelation)
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

			//判断节点是否完成(反向关联的), 主要用于等待未完成的任务
			for _, v := range brsx {
				//check finished

			}

			//next flow
			var defNode *models.SequenceFlow
			passCount := 0
			size := 0
			if !gateway.Exclusive {
				size++
			} else {
				size = gateway.ScriptFlows
			}
			wg.Add(size)
			for _, v := range flows {
				if v.DefaultFlow && defNode == nil {
					defNode = v
				}
				if len(v.Script) > 0 {
					go func() {
						defer wg.Done()
						node, err := f.call(types.GCNode, v.End)
						if err != nil {
							resp(err)
							return
						}
						n, ok := node.(*models.Node)
						if ok {
							ctx, err := f.script.Do(f.ctx, f.instance, nil, v.EndType, f, n)
							if err != nil {
								if err == flowerr.ScriptFalse {

								} else if err == flowerr.ScriptTrue {
									passCount++
									f.next(v.End)
									if gateway.Exclusive {
										break
									}
								} else if err == flowerr.NextFlow {
									f.again(v.End)
								} else {
									return err
								}
							}
							f.context(ctx)
						} else {
							return flowerr.ErrNode
						}
					}()
				}
			}

			wg.Wait()

			if passCount == 0 {
				//run defNode

			}
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
	f.status.Again = id
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
	f.ignoreNodes = nil
	f.conn.Close()
	f.ignoreNodesChanged = false
}

func NewNextflow(modules modules.Modules, log *gslogrus.Logger, script *script.LuaScript, pool *redis.Pool) Handler {
	return &nextflow{modules: modules, log: log, script: newScript(modules, log, script), pool: pool}
}
