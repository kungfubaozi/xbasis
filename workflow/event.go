package workflow

type TypeEvent struct {
	Id                 string   `bson:"id" json:"id"`
	Name               string   `bson:"name" json:"name"`
	Desc               string   `bson:"desc" json:"desc"`
	CreateAt           int64    `bson:"create_at" json:"create_at"`
	ExecutionListeners []string `bson:"execution_listeners" json:"execution_listeners"`
}

type StartEvent struct {
	TypeEvent
	FormRef string `bson:"form_ref" json:"form_ref"`
}

//定时启动事件
type TimerStartEvent struct {
	TypeEvent
	TimeCron string `bson:"time_cron" json:"time_cron"`
}

//消息事件
type MessageStartEvent struct {
	TypeEvent
	Reference string `bson:"reference" json:"reference"`
}

//结束事件
type EndEvent struct {
	TypeEvent
}

type EndCancelEvent struct {
	TypeEvent
}

type EndTerminateEvent struct {
	TypeEvent
}
