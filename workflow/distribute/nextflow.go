package distribute

import (
	"context"
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
	finished map[string]bool
	log      *gslogrus.Logger
	values   []interface{}
	status   *models.NextStatus
	node     *models.Node
	ctx      context.Context
	instance *models.Instance
	script   Handler
	gon      bool
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
	return handler(ctx, ct, f)
}

//如果到此网关会获取与当前关联的节点
func (f *nextflow) exclusiveGateway() *flowerr.Error {
	flow := f.values[0].(*models.SequenceFlow)
	var flows []*models.SequenceFlow
	v := f.values[1]
	if v != nil {
		flows = v.([]*models.SequenceFlow)
	}
	if flows == nil || len(flows) == 0 {
		return flowerr.ErrNode
	}
	var defNode string
	for _, f1 := range flows {
		if f1.DefaultFlow && len(defNode) == 0 { //默认节点
			defNode = f1.End
		}
		ctx, err := f.script.Do(f.ctx, f.instance, f.node, flow.StartType, f1)
		f.context(ctx)
		if err != nil {
			return err
		}
		if err == flowerr.NextFlow {
			f.again(flow.End)
			return nil
		}
	}
	if len(defNode) > 0 {
		f.again(defNode)
		return nil
	}
	return flowerr.ErrNoDownwardProcess
}

func (f *nextflow) parallelGateway() *flowerr.Error {
	var tasks []string
	v := f.values[1]
	if v != nil {
		tasks = v.([]string)
	}
	if tasks == nil {
		return flowerr.ErrNode
	}
	size := 0
	var err *flowerr.Error
	e := func(er *flowerr.Error) {
		if err == nil {
			err = er
		}
	}
	var wg sync.WaitGroup
	wg.Add(len(tasks))
	for _, v := range tasks {
		go func() {
			defer wg.Done()
			ok, err := f.modules.Instance().IsFinished(f.instance.Id, v)
			if err != nil {
				e(err)
				return
			}
			if ok {
				size++
			}
		}()
	}
	wg.Wait()
	if err != nil {
		return err
	}
	if size == len(tasks) {
		f.again(f.flow().End) //如果关联的task都完成了，那么继续下一步流程
	} else {
		f.next(f.node.Id) //没有完成当前任务，设置当前gateway为currentNode等待完成
	}
	return nil
}

func (f *nextflow) inclusiveGateway() *flowerr.Error {
	flow := f.flow()
	ctx, err := f.script.Do(f.ctx, f.instance, f.node, flow.StartType, f.values[0])
	f.context(ctx)
	if err != nil {
		return err
	}
	if err == flowerr.NextFlow {
		f.again(flow.End)
	} else if err == flowerr.ScriptTrue {
		f.next(f.node.Id)
	} else if err == flowerr.ScriptFalse {

	}
	return nil
}

func (f *nextflow) startEvent() *flowerr.Error {
	panic("implement me")
}

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
}

func NewNextflow(modules modules.Modules, log *gslogrus.Logger, script *script.LuaScript) Handler {
	return &nextflow{modules: modules, log: log, script: newScript(modules, log, script)}
}
