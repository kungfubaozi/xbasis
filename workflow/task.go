package workflow

type ConnectionType int64

const (
	FTUserTask ConnectionType = iota

	FTHttpTask

	FTDecisionTask

	FTStartEvent

	FTHeader

	FTTimerStartEvent

	FTMessageStartEvent

	FTEndEvent

	FTEndErrorEvent

	FTEndCancelEvent

	FTTerminateEvent

	FTExclusiveGateway

	FTParallelGateway

	FTInclusiveGateway

	FTEventGateway
)

type basicModel struct {
	Id       string `bson:"id" json:"id"`
	Name     string `bson:"name" json:"name"`
	Desc     string `bson:"desc" json:"desc"`
	CreateAt int64  `bson:"create_at" json:"create_at"`
}

type sequenceFlow struct {
	Basic       *basicModel    `bson:"basic" json:"basic"`
	Listeners   *TEListener    `bson:"listeners" json:"listeners"`
	Script      string         `bson:"script" json:"script"`
	DefaultFlow bool           `bson:"default_flow" json:"default_flow"`
	Start       string         `bson:"start" json:"start"`
	StartType   ConnectionType `bson:"start_type" json:"start_type"`
	End         string         `bson:"end" json:"end"`
	EndType     ConnectionType `bson:"end_type" json:"end_type"`
	Priority    int64          `bson:"priority" json:"priority"`
}

type userTask struct {
	Basic                   *basicModel `bson:"basic" json:"basic"`
	Listeners               *TEListener `bson:"listeners" json:"listeners"`
	FormRef                 string      `bson:"form_ref" json:"form_ref"`
	Priority                int64       `json:"priority" bson:"priority"`
	Assignments             []string    `json:"assignments" bson:"assignments"`
	AssignmentType          int64       `bson:"assignment_type" json:"assignment_type"`
	AllowOriginatorFinished bool        `bson:"allow_originator_finished" json:"allow_originator_finished"` //允许创建人取消流程
}

type httpTask struct {
	Basic              *basicModel `bson:"basic" json:"basic"`
	Listeners          *TEListener `bson:"listeners" json:"listeners"`
	RequestMethod      int64       `json:"request_method" bson:"request_method"`   //请求方法
	RequestURL         string      `bson:"request_url" json:"request_url"`         //请求地址
	RequestHeaders     []string    `json:"request_headers"`                        //请求头
	RequestTimeout     int64       `json:"request_timeout" bson:"request_timeout"` //请求超时时间
	RequestBody        interface{} `bson:"request_body" json:"request_body"`
	ResponseScript     int64       `bson:"response_script" json:"response_script"`
	SaveResponseAsJson bool        `bson:"save_response_as_json" json:"save_response_as_json"`
}

type decisionTask struct {
	Basic          *basicModel `bson:"basic" json:"basic"`
	Listeners      *TEListener `bson:"listeners" json:"listeners"`
	TableReference string      `json:"table_reference" bson:"table_reference"`
}

type sendTask struct {
	Basic     *basicModel `bson:"basic" json:"basic"`
	Listeners *TEListener `bson:"listeners" json:"listeners"`
	UserIds   []string    `bson:"user_ids" json:"user_ids"`
}

type grpcTask struct {
}

type mailTask struct {
	Basic             *basicModel `bson:"basic" json:"basic"`
	ServerHost        string      `bson:"mail_server_host" json:"mail_server_host"`
	ServerPost        int64       `bson:"mail_server_post" json:"mail_server_post"`
	ServerDefaultFrom string      `bson:"mail_server_default_from" json:"mail_server_default_from"`
	ServerUsername    string      `bson:"mail_server_username" json:"mail_server_username"`
	ServerPassword    string      `bson:"mail_server_password" json:"mail_server_password"`
}
