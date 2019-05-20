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
	ctx      context.Context
	instance *models.Instance
	rn       *models.NodeRelation
	data     map[string]interface{}
}

func (f *dataGetter) SetCommandFunc(call types.CommandDataGetter) {
	panic("implement me")
}

func (f *dataGetter) eventGateway() *flowerr.Error {
	panic("implement me")
}

func (f *dataGetter) timerStartEvent() *flowerr.Error {
	panic("implement me")
}

func (f *dataGetter) messageStartEvent() *flowerr.Error {
	panic("implement me")
}

func (f *dataGetter) cancelEndEvent() *flowerr.Error {
	panic("implement me")
}

func (f *dataGetter) terminateEndEvent() *flowerr.Error {
	panic("implement me")
}

func (f *dataGetter) Data() interface{} {
	panic("implement me")
}

func (f *dataGetter) Do(ctx context.Context, instance *models.Instance, node *models.Node, ct types.ConnectType, values ...interface{}) (context.Context, *flowerr.Error) {
	f.ctx = ctx
	f.instance = instance
	f.values = values
	return handler(ctx, ct, f)
}

func (f *dataGetter) relation() *models.NodeRelation {
	return f.values[0].(*models.NodeRelation)
}

func (f *dataGetter) inclusiveGateway() *flowerr.Error {
	return nil
}

func (f *dataGetter) startEvent() *flowerr.Error {
	return f.loadData()
}

func (f *dataGetter) endEvent() *flowerr.Error {
	return nil
}

func (f *dataGetter) apiStartEvent() *flowerr.Error {
	return f.loadData()
}

func (f *dataGetter) userTask() *flowerr.Error {
	return f.loadData()
}

func (f *dataGetter) notifyTask() *flowerr.Error {
	return nil
}

func (f *dataGetter) triggerStartEvent() *flowerr.Error {
	return nil
}

func (f *dataGetter) context(ctx context.Context) context.Context {
	return f.ctx
}

func (f *dataGetter) metadata(key string, data interface{}) {
	if f.data == nil {
		f.data = make(map[string]interface{})
	}
	f.data[key] = data
}

func (f *dataGetter) loadData() *flowerr.Error {
	d := f.relation()
	data, err := f.modules.Form().LoadNodeDataToStore(f.ctx, f.instance.Id, d.Id)
	if err != nil {
		return err
	}
	f.metadata(d.Key, data)
	return nil
}

func (f *dataGetter) Restore() {
	f.data = nil
}

//作用是查找节点提交的数据
func NewDataGetter(modules modules.Modules, log *gslogrus.Logger) Handler {
	return &dataGetter{modules: modules, log: log}
}
