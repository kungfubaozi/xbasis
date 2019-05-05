package types

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

	CTEventGateway
)
