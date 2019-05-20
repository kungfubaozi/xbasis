package modules

import (
	"context"
	"konekko.me/gosion/workflow/flowerr"
	"konekko.me/gosion/workflow/models"
	"konekko.me/gosion/workflow/types"
)

type distribution interface {
	Data() interface{}

	Do(ctx context.Context, instance *models.Instance, node *models.Node, ct types.ConnectType, value ...interface{}) (context.Context, *flowerr.Error)

	ExclusiveGateway() (context.Context, *flowerr.Error)

	ParallelGateway() (context.Context, *flowerr.Error)

	InclusiveGateway() (context.Context, *flowerr.Error)

	StartEvent() (context.Context, *flowerr.Error)

	EndEvent() (context.Context, *flowerr.Error)

	ApiStartEvent() (context.Context, *flowerr.Error)

	UserTask() (context.Context, *flowerr.Error)

	NotifyTask() (context.Context, *flowerr.Error)

	TriggerStartEvent() (context.Context, *flowerr.Error)

	Restore()
}

func distribute(ctx context.Context, ct types.ConnectType, t distribution) (context.Context, *flowerr.Error) {
	switch ct {
	case types.CTStartEvent:
		return t.StartEvent()
	case types.CTEndEvent:
		return t.EndEvent()
	case types.CTUserTask:
		return t.UserTask()
	case types.CTInclusiveGateway:
		return t.InclusiveGateway()
	}
	return ctx, flowerr.ErrUnsupportedConnectType
}
