package distribute

import (
	"context"
	"konekko.me/gosion/commons/gslogrus"
	"konekko.me/gosion/workflow/flowerr"
	"konekko.me/gosion/workflow/models"
	"konekko.me/gosion/workflow/modules"
	"konekko.me/gosion/workflow/types"
)

type dataGetter struct {
	modules  modules.Modules
	finished map[string]bool
	log      *gslogrus.Logger
	values   []interface{}
	node     *models.Node
	ctx      context.Context
	instance *models.Instance
	pipe     modules.Pipeline
}

func (f *dataGetter) Data() interface{} {
	panic("implement me")
}

func (f *dataGetter) Do(ctx context.Context, instance *models.Instance, node *models.Node, ct types.ConnectType, value ...interface{}) (context.Context, *flowerr.Error) {
	f.ctx = ctx
	f.node = node
	f.instance = instance
	f.pipe = value[0].(modules.Pipeline)
	return handler(ctx, ct, f)
}

func (f *dataGetter) exclusiveGateway() *flowerr.Error {
	panic("implement me")
}

func (f *dataGetter) parallelGateway() *flowerr.Error {
	panic("implement me")
}

func (f *dataGetter) inclusiveGateway() *flowerr.Error {
	panic("implement me")
}

func (f *dataGetter) startEvent() *flowerr.Error {
	panic("implement me")
}

func (f *dataGetter) endEvent() *flowerr.Error {
	panic("implement me")
}

func (f *dataGetter) apiStartEvent() *flowerr.Error {
	panic("implement me")
}

func (f *dataGetter) userTask() *flowerr.Error {
	panic("implement me")
}

func (f *dataGetter) notifyTask() *flowerr.Error {
	panic("implement me")
}

func (f *dataGetter) triggerStartEvent() *flowerr.Error {
	panic("implement me")
}

func (f *dataGetter) context(ctx context.Context) context.Context {
	panic("implement me")
}

func (f *dataGetter) metadata(key, data interface{}) {
	panic("implement me")
}

func (f *dataGetter) Restore() {
	panic("implement me")
}

//作用是查找节点提交的数据
func NewDataGetter(modules modules.Modules, log *gslogrus.Logger) Handler {
	return &dataGetter{modules: modules, log: log}
}
