package flowtask

type UserTask struct {
	Id                      string   `bson:"id" json:"id"`
	Name                    string   `bson:"name" json:"name"`
	Desc                    string   `bson:"desc" json:"desc"`
	CreateAt                int64    `bson:"create_at" json:"create_at"`
	BindFormId              string   `bson:"bind_form_id" json:"bind_form_id"`
	Priority                int64    `json:"priority" bson:"priority"`
	Assignments             []string `json:"assignments" bson:"assignments"`
	AllowOriginatorFinished bool     `bson:"allow_originator_finished" json:"allow_originator_finished"`
	TaskListeners           []string `bson:"task_listeners" json:"task_listeners"`
	ExecutionListeners      []string `bson:"execution_listeners" json:"execution_listeners"`
}

type ServiceTask struct {
	Id                 string   `bson:"id" json:"id"`
	Name               string   `bson:"name" json:"name"`
	Desc               string   `bson:"desc" json:"desc"`
	CreateAt           int64    `bson:"create_at" json:"create_at"`
	TaskListeners      []string `bson:"task_listeners" json:"task_listeners"`
	ExecutionListeners []string `bson:"execution_listeners" json:"execution_listeners"`
}

type HttpTask struct {
	Id                      string   `bson:"id" json:"id"`
	Name                    string   `bson:"name" json:"name"`
	Desc                    string   `bson:"desc" json:"desc"`
	CreateAt                int64    `bson:"create_at" json:"create_at"`
	RequestMethod           int64    `json:"request_method" bson:"request_method"`     //请求方法
	RequestURL              string   `bson:"request_url" json:"request_url"`           //请求地址
	RequestHeaders          []string `json:"request_headers"`                          //请求头
	RequestTimeout          int64    `json:"request_timeout" bson:"request_timeout"`   //请求超时时间
	FailStatusCode          int64    `bson:"fail_status_code" json:"fail_status_code"` //请求失败状态码
	BindFormIdAsRequestBody string   `bson:"bind_form_id" json:"bind_form_id_as_request_body"`
	TaskListeners           []string `bson:"task_listeners" json:"task_listeners"`
	ExecutionListeners      []string `bson:"execution_listeners" json:"execution_listeners"`
}

type ManualTask struct {
	Id                 string   `bson:"id" json:"id"`
	Name               string   `bson:"name" json:"name"`
	Desc               string   `bson:"desc" json:"desc"`
	CreateAt           int64    `bson:"create_at" json:"create_at"`
	TaskListeners      []string `bson:"task_listeners" json:"task_listeners"`
	ExecutionListeners []string `bson:"execution_listeners" json:"execution_listeners"`
}

type DecisionTask struct {
	Id                 string   `bson:"id" json:"id"`
	Name               string   `bson:"name" json:"name"`
	Desc               string   `bson:"desc" json:"desc"`
	CreateAt           int64    `bson:"create_at" json:"create_at"`
	TaskListeners      []string `bson:"task_listeners" json:"task_listeners"`
	ExecutionListeners []string `bson:"execution_listeners" json:"execution_listeners"`
}
