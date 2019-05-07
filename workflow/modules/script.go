package modules

import (
	"context"
	"konekko.me/gosion/commons/dto"
	"konekko.me/gosion/commons/gslogrus"
	"konekko.me/gosion/workflow/flowstate"
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

func (f *flowscript) Data() interface{} {
	panic("implement me")
}

func (f *flowscript) Do(ctx context.Context, instance *models.Instance, node *node, ct types.ConnectType, value ...interface{}) (*gs_commons_dto.State, error) {
	f.values = value
	f.ctx = ctx
	f.instance = instance
	return distribute(ct, f)
}

func (f *flowscript) ExclusiveGateway() (*gs_commons_dto.State, error) {
	return flowstate.NextFlow, nil
}

func (f *flowscript) ParallelGateway() (*gs_commons_dto.State, error) {
	return flowstate.NextFlow, nil
}

func (f *flowscript) InclusiveGateway() (*gs_commons_dto.State, error) {
	return flowstate.NextFlow, nil
}

func (f *flowscript) StartEvent() (*gs_commons_dto.State, error) {
	return f.do()
}

func (f *flowscript) EndEvent() (*gs_commons_dto.State, error) {
	return f.do()
}

func (f *flowscript) UserTask() (*gs_commons_dto.State, error) {
	return f.do()
}

func (f *flowscript) NotifyTask() (*gs_commons_dto.State, error) {
	return f.do()
}

func (f *flowscript) TriggerStartEvent() (*gs_commons_dto.State, error) {
	return f.do()
}

func (f *flowscript) do() (*gs_commons_dto.State, error) {
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
			return flowstate.FlowScriptFalse, nil
		}
	}
	return flowstate.FlowScriptTrue, nil
}

func (f *flowscript) Restore() {
	panic("implement me")
}

func newScript(modules Modules, log *gslogrus.Logger, script *script.LuaScript) distribution {
	return &flowscript{modules: modules, log: log, script: script}
}
