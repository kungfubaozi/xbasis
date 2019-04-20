package flow

type Flow struct {
	Id                 string   `bson:"id" json:"id"`
	Name               string   `bson:"name" json:"name"`
	Desc               string   `bson:"desc" json:"desc"`
	CreateAt           int64    `bson:"create_at" json:"create_at"`
	Script             string   `bson:"script" json:"script"`
	DefaultFlow        bool     `bson:"default_flow" json:"default_flow"`
	Start              string   `bson:"start" json:"start"`
	StartType          int64    `bson:"start_type" bson:"start_type"`
	End                string   `bson:"end" json:"end"`
	EndType            string   `bson:"end_type" json:"end_type"`
	ExecutionListeners []string `bson:"execution_listeners" json:"execution_listeners"`
}

type StartEvent struct {
	Id                 string   `bson:"id" json:"id"`
	Name               string   `bson:"name" json:"name"`
	Desc               string   `bson:"desc" json:"desc"`
	CreateAt           int64    `bson:"create_at" json:"create_at"`
	ExecutionListeners []string `bson:"execution_listeners" json:"execution_listeners"`
}

//定时启动事件
type TimerStartEvent struct {
	StartEvent
	TimeCron string `bson:"time_cron" json:"time_cron"`
}

//消息事件
type MessageStartEvent struct {
	StartEvent
	Reference string `bson:"reference" json:"reference"`
}

//结束事件
type EndEvent struct {
	Id                 string   `bson:"id" json:"id"`
	Name               string   `bson:"name" json:"name"`
	Desc               string   `bson:"desc" json:"desc"`
	CreateAt           int64    `bson:"create_at" json:"create_at"`
	ExecutionListeners []string `bson:"execution_listeners" json:"execution_listeners"`
}
