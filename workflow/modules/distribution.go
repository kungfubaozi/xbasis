package modules

import (
	"context"
	"konekko.me/gosion/workflow/types"
)

type distribution interface {
	Data() interface{}

	Do(ctx context.Context, instanceId string, node *node, ct types.ConnectType, value interface{})

	ExclusiveGateway()

	ParallelGateway()

	InclusiveGateway()

	StartEvent()

	EndEvent()

	UserTask()

	TriggerStartEvent()

	Restore()
}

func distribute(ct types.ConnectType, t distribution) {
	switch ct {
	case types.CTStartEvent:
		t.StartEvent()
		break
	case types.CTEndEvent:
		t.EndEvent()
		break
	case types.CTUserTask:
		t.UserTask()
		break
	case types.CTExclusiveGateway:
		t.ExclusiveGateway()
		break
	case types.CTParallelGateway:
		t.ParallelGateway()
		break
	case types.CTInclusiveGateway:
		t.InclusiveGateway()
		break
	}
}
