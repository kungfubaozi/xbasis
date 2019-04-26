package workflow

type taskListener struct {
	Id       string `bson:"id" json:"id"`
	Name     string `bson:"name" json:"name"`
	Desc     string `bson:"desc" json:"desc"`
	CreateAt int64  `bson:"create_at" json:"create_at"`
	Listener string `bson:"listener" json:"listener"`
}

type executionListener struct {
	Id       string `bson:"id" json:"id"`
	Name     string `bson:"name" json:"name"`
	Desc     string `bson:"desc" json:"desc"`
	CreateAt int64  `bson:"create_at" json:"create_at"`
	Listener string `bson:"listener" json:"listener"`
}

type TEListener struct {
	ExecutionListeners []string `bson:"execution_listeners" json:"execution_listeners"`
	TaskListeners      []string `bson:"task_listeners" json:"task_listeners"`
}

type FlowTaskFunc interface {
}

type FlowExecutionFunc interface {
}
