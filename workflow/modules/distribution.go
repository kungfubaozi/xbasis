package modules

import (
	"context"
	"konekko.me/gosion/commons/dto"
	"konekko.me/gosion/workflow/flowstate"
	"konekko.me/gosion/workflow/models"
	"konekko.me/gosion/workflow/types"
)

type distribution interface {
	Data() interface{}

	Do(ctx context.Context, instance *models.Instance, node *node, ct types.ConnectType, value ...interface{}) (*gs_commons_dto.State, error)

	ExclusiveGateway() (*gs_commons_dto.State, error)

	ParallelGateway() (*gs_commons_dto.State, error)

	InclusiveGateway() (*gs_commons_dto.State, error)

	StartEvent() (*gs_commons_dto.State, error)

	EndEvent() (*gs_commons_dto.State, error)

	UserTask() (*gs_commons_dto.State, error)

	NotifyTask() (*gs_commons_dto.State, error)

	TriggerStartEvent() (*gs_commons_dto.State, error)

	Restore()
}

func distribute(ct types.ConnectType, t distribution) (*gs_commons_dto.State, error) {
	switch ct {
	case types.CTStartEvent:
		return t.StartEvent()
	case types.CTEndEvent:
		return t.EndEvent()
	case types.CTUserTask:
		return t.UserTask()
	case types.CTExclusiveGateway:
		return t.ExclusiveGateway()
	case types.CTParallelGateway:
		return t.ParallelGateway()
	case types.CTInclusiveGateway:
		return t.InclusiveGateway()
	}
	return flowstate.ErrUnsupportedConnectType, nil
}
