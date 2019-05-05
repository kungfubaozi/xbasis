package types

import "context"

type ConnectType int64

const (
	CTUserTask ConnectType = iota

	CTHttpTask

	CTDecisionTask

	CTSendTask

	CTGRPCTask

	CTMailTask

	CTApiTask

	CTStorageTask

	CTDataCache

	CTStartEvent

	CTTimerStartEvent

	CTMessageStartEvent

	CTEndEvent

	CTEndErrorEvent

	CTEndCancelEvent

	CTTerminateEvent

	CTExclusiveGateway

	CTParallelGateway

	CTInclusiveGateway
)

type TypeTasks interface {
	ExclusiveGateway()

	ParallelGateway()

	InclusiveGateway()

	StartEvent()

	EndEvent()

	UserTask()

	Do(ctx context.Context, instanceId, nodeId string, ct ConnectType, value map[string]interface{})
}
