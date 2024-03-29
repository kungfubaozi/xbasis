package distribute

import (
	"context"
	"konekko.me/xbasis/workflow/flowerr"
	"konekko.me/xbasis/workflow/models"
	"konekko.me/xbasis/workflow/types"
)

type relation struct {
	values []interface{}
}

func (f *relation) startEvent() *flowerr.Error {
	return nil
}

func (f *relation) timerStartEvent() *flowerr.Error {
	return nil
}

func (f *relation) apiStartEvent() *flowerr.Error {
	return nil
}

func (f *relation) messageStartEvent() *flowerr.Error {
	return nil
}

func (f *relation) triggerStartEvent() *flowerr.Error {
	return nil
}

func (f *relation) endEvent() *flowerr.Error {
	return nil
}

func (f *relation) cancelEndEvent() *flowerr.Error {
	return nil
}

func (f *relation) terminateEndEvent() *flowerr.Error {
	return nil
}

func (f *relation) userTask() *flowerr.Error {
	return nil
}

func (f *relation) notifyTask() *flowerr.Error {
	return nil
}

func (f *relation) RunActions(values ...interface{}) (interface{}, *flowerr.Error) {
	panic("implement me")
}

func (f *relation) SetCommandFunc(call types.CommandDataGetter) {
	panic("implement me")
}

func (f *relation) Data() interface{} {
	panic("implement me")
}

func (f *relation) Do(ctx context.Context, instance *models.Instance, node *models.Node, ct types.ConnectType, values ...interface{}) (context.Context, *flowerr.Error) {
	f.values = values
	return handler(ctx, ct, f)
}

func (f *relation) size() int {
	return f.values[0].(int)
}

func (f *relation) inclusiveGateway() *flowerr.Error {
	//if f.size() == 1 {
	//	return nil
	//}
	return flowerr.NextFlow
}

func (f *relation) context(ctx context.Context) context.Context {
	return ctx
}

func (f *relation) metadata(key string, data interface{}) {
	panic("implement me")
}

func (f *relation) Restore() {
	panic("implement me")
}

func NewRelation() Handler {
	return &relation{}
}
