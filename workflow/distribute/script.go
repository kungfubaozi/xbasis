package distribute

import (
	"context"
	"fmt"
	"konekko.me/gosion/commons/gslogrus"
	"konekko.me/gosion/workflow/flowerr"
	"konekko.me/gosion/workflow/models"
	"konekko.me/gosion/workflow/modules"
	"konekko.me/gosion/workflow/script"
	"konekko.me/gosion/workflow/types"
)

type flowscript struct {
	modules  modules.Modules
	log      *gslogrus.Logger
	script   *script.LuaScript
	ctx      context.Context
	values   []interface{}
	instance *models.Instance
}

func (f *flowscript) timerStartEvent() *flowerr.Error {
	return f.do()
}

func (f *flowscript) messageStartEvent() *flowerr.Error {
	return f.do()
}

func (f *flowscript) cancelEndEvent() *flowerr.Error {
	return f.do()
}

func (f *flowscript) terminateEndEvent() *flowerr.Error {
	return f.do()
}

func (f *flowscript) Data() interface{} {
	panic("implement me")
}

func (f *flowscript) Do(ctx context.Context, instance *models.Instance, node *models.Node, ct types.ConnectType, values ...interface{}) (context.Context, *flowerr.Error) {
	f.values = values
	f.ctx = ctx
	f.instance = instance
	return handler(ctx, ct, f)
}

func (f *flowscript) nextflow() *flowerr.Error {
	err := f.do()
	if err != nil && err == flowerr.ScriptTrue {
		return flowerr.NextFlow
	}
	return err
}

func (f *flowscript) eventGateway() *flowerr.Error {
	return f.nextflow()
}

func (f *flowscript) exclusiveGateway() *flowerr.Error {
	return f.nextflow()
}

func (f *flowscript) parallelGateway() *flowerr.Error {
	return f.nextflow()
}

func (f *flowscript) inclusiveGateway() *flowerr.Error {
	return f.nextflow()
}

func (f *flowscript) startEvent() *flowerr.Error {
	return f.do()
}

func (f *flowscript) endEvent() *flowerr.Error {
	return f.do()
}

func (f *flowscript) apiStartEvent() *flowerr.Error {
	return f.do()
}

func (f *flowscript) userTask() *flowerr.Error {
	return f.do()
}

func (f *flowscript) notifyTask() *flowerr.Error {
	return f.do()
}

func (f *flowscript) triggerStartEvent() *flowerr.Error {
	return f.do()
}

func (f *flowscript) context(ctx context.Context) context.Context {
	if ctx != nil {
		f.ctx = ctx
	}
	return f.ctx
}

func (f *flowscript) metadata(key, data interface{}) {
	f.ctx = context.WithValue(f.ctx, key, data)
}

func (f *flowscript) Restore() {
	panic("implement me")
}

func (f *flowscript) do() *flowerr.Error {
	flow := f.values[0].(*models.SequenceFlow)
	if len(flow.Script) > 0 { //网关上必须设置条件脚本
		node := f.values[1].(*models.Node)
		var data map[string]interface{}
		value := f.ctx.Value(node.Key)
		if value != nil {
			data = value.(map[string]interface{})
			d := make(map[string]interface{})
			for k, v := range data {
				d[fmt.Sprintf("%s.%s", node.Key, k)] = v
			}

			//in script: flow['yourNodeKey.yourNeedParam']

			data = d
		}
		if data == nil {
			return flowerr.ErrFindSubmitForm
		}
		ok, err := f.script.Run(flow.Script, data)
		if err != nil {
			return err
		}
		if !ok {
			return flowerr.ScriptFalse
		}
		f.metadata(node.Key, data)
		return flowerr.ScriptTrue
	}
	return flowerr.ErrInvalidGatewayScript
}

func newScript(modules modules.Modules, log *gslogrus.Logger, script *script.LuaScript) Handler {
	return &flowscript{modules: modules, log: log, script: script}
}
