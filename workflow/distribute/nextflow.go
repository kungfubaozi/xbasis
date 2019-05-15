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

func (f *nextflow) eventGateway() *flowerr.Error {
	panic("implement me")
}

//排他网关
func (f *nextflow) exclusiveGateway() *flowerr.Error {
	return f.flows(func(flows []*models.SequenceFlow) *flowerr.Error {
		var defNode string
		for _, v := range flows {
			if v.DefaultFlow && len(defNode) == 0 { //默认节点
				defNode = v.End
			}
			//check script
		}
		if len(defNode) > 0 {
			f.again(defNode)
			return nil
		}
		return flowerr.ErrNoDownwardProcess
	})
}

//并行网关
func (f *nextflow) parallelGateway() *flowerr.Error {
	return f.flows(func(flows []*models.SequenceFlow) *flowerr.Error {

	})
}

//包容网关
//此为nextflow控制
//执行所有满足条件的flow
func (f *nextflow) inclusiveGateway() *flowerr.Error {
	return f.flows(func(flows []*models.SequenceFlow) *flowerr.Error {

	})
}

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

func (f *nextflow) metadata(key, data interface{}) {
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
