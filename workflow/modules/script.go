package modules

import (
	"context"
	"konekko.me/gosion/commons/gslogrus"
	"konekko.me/gosion/workflow/flowerr"
	"konekko.me/gosion/workflow/models"
	"konekko.me/gosion/workflow/script"
	"konekko.me/gosion/workflow/types"
)

type flowscript struct {
	modules  Modules
	log      *gslogrus.Logger
	script   *script.LuaScript
	ctx      context.Context
	values   []interface{}
	instance *models.Instance
}

func (f *flowscript) ApiStartEvent() (context.Context, *flowerr.Error) {
	panic("implement me")
}

func (f *flowscript) Data() interface{} {
	panic("implement me")
}

func (f *flowscript) Do(ctx context.Context, instance *models.Instance, node *node, ct types.ConnectType, value ...interface{}) (context.Context, *flowerr.Error) {
	f.values = value
	f.ctx = ctx
	f.instance = instance
	return distribute(ctx, ct, f)
}

func (f *flowscript) ExclusiveGateway() (context.Context, *flowerr.Error) {
	return f.ctx, flowerr.NextFlow
}

func (f *flowscript) ParallelGateway() (context.Context, *flowerr.Error) {
	return f.ctx, flowerr.NextFlow
}

func (f *flowscript) InclusiveGateway() (context.Context, *flowerr.Error) {
	return f.ctx, flowerr.NextFlow
}

func (f *flowscript) StartEvent() (context.Context, *flowerr.Error) {
	return f.do()
}

func (f *flowscript) EndEvent() (context.Context, *flowerr.Error) {
	return f.do()
}

func (f *flowscript) UserTask() (context.Context, *flowerr.Error) {
	return f.do()
}

func (f *flowscript) NotifyTask() (context.Context, *flowerr.Error) {
	return f.do()
}

func (f *flowscript) TriggerStartEvent() (context.Context, *flowerr.Error) {
	return f.do()
}

func (f *flowscript) do() (context.Context, *flowerr.Error) {
	flow := f.values[0].(*models.SequenceFlow)
	data, err := f.modules.Form().LoadNodeDataToStore(f.ctx, f.instance.Id, flow.Start)
	if err != nil {
		return nil, err
	}
	if len(flow.Script) > 0 {
		ok, err := f.script.Run(flow.Script, data)
		if err != nil {
			return nil, err
		}
		if !ok {
			return f.ctx, flowerr.ScriptFalse
		}
	}
	return f.ctx, flowerr.ScriptTrue
}

func (f *flowscript) Restore() {
	panic("implement me")
}

func newScript(modules Modules, log *gslogrus.Logger, script *script.LuaScript) distribution {
	return &flowscript{modules: modules, log: log, script: script}
}
