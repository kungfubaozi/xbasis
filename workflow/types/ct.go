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

	CTStartEvent

	CTTimerStartEvent

	CTMessageStartEvent

	CTApiStartEvent

	CTEndEvent

	CTEndErrorEvent

	CTEndCancelEvent

	CTTerminateEvent

	CTInclusiveGateway
)

var RollbackKey = "_#rollback"
