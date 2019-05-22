package distribute

import (
	"context"
	"konekko.me/gosion/workflow/flowerr"
	"konekko.me/gosion/workflow/models"
	"konekko.me/gosion/workflow/types"
)

type Handler interface {
	SetCommandFunc(call types.CommandDataGetter)

	Data() interface{}

	Do(ctx context.Context, instance *models.Instance, node *models.Node, ct types.ConnectType, values ...interface{}) (context.Context, *flowerr.Error)

	inclusiveGateway() *flowerr.Error

	startEvent() *flowerr.Error

	timerStartEvent() *flowerr.Error

	apiStartEvent() *flowerr.Error

	messageStartEvent() *flowerr.Error

	triggerStartEvent() *flowerr.Error

	endEvent() *flowerr.Error

	cancelEndEvent() *flowerr.Error

	terminateEndEvent() *flowerr.Error

	userTask() *flowerr.Error

	notifyTask() *flowerr.Error

	context(ctx context.Context) context.Context

	metadata(key string, data interface{})

	RunActions(values ...interface{}) (interface{}, *flowerr.Error)

	Restore()
}

func handler(ctx context.Context, ct types.ConnectType, t Handler) (context.Context, *flowerr.Error) {

	switch ct {
	case types.CTStartEvent:
		return t.context(nil), t.startEvent()
	case types.CTEndEvent:
		return t.context(nil), t.endEvent()
	case types.CTUserTask:
		return t.context(nil), t.userTask()
	case types.CTInclusiveGateway:
		return t.context(nil), t.inclusiveGateway()
	}
	return ctx, flowerr.ErrUnsupportedConnectType
}
