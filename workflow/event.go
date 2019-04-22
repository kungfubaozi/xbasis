package workflow

type typeEvent struct {
	Id                 string   `bson:"id" json:"id"`
	Name               string   `bson:"name" json:"name"`
	Desc               string   `bson:"desc" json:"desc"`
	CreateAt           int64    `bson:"create_at" json:"create_at"`
	ExecutionListeners []string `bson:"execution_listeners" json:"execution_listeners"`
}

type startEvent struct {
	typeEvent
	FormRef string `bson:"form_ref" json:"form_ref"`
}

//定时启动事件
type timerStartEvent struct {
	typeEvent
	TimeCron string `bson:"time_cron" json:"time_cron"`
}

//消息事件
type messageStartEvent struct {
	typeEvent
	Reference string `bson:"reference" json:"reference"`
}

//结束事件
type endEvent struct {
	typeEvent
}

type endCancelEvent struct {
	typeEvent
}

type endTerminateEvent struct {
	typeEvent
}
