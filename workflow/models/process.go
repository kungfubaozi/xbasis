package models

import (
	"konekko.me/gosion/workflow/types"
)

type Process struct {
	Id           string          `bson:"_id" json:"id"`
	AppId        string          `bson:"app_id" json:"app_id"`
	Name         string          `bson:"name" json:"name"`
	CreateAt     int64           `bson:"create_at" json:"create_at"`
	CreateUserId string          `bson:"create_user_id" json:"create_user_id"`
	Gateways     *Gateways       `bson:"gateways" json:"gateways"`
	Tasks        *Tasks          `bson:"tasks" json:"tasks"`
	Flows        []*SequenceFlow `bson:"flows" json:"flows"`
	StartEvent   *TypeEvent      `bson:"start_event" json:"start_event"`
	EndEvents    []*TypeEvent    `bson:"end_events" json:"end_events"`
	Version      int64           `bson:"version" json:"version"`
	Status       int64           `bson:"status" json:"status"`
	Title        string          `bson:"title" json:"title"`
}

type TypeEvent struct {
	Id    string            `bson:"id" json:"id"`
	Key   string            `bson:"key" json:"key"`
	Event interface{}       `bson:"event" json:"event"`
	Type  types.ConnectType `bson:"type" json:"type"`
}

type Tasks struct {
	UserTasks     []*UserTask      `bson:"user_tasks" json:"user_tasks"`
	ApiTasks      []*ApiTask       `bson:"api_tasks" json:"api_tasks"`
	DecisionTasks []*DecisionTask  `bson:"decision_tasks" json:"decision_tasks"`
	HttpTasks     []*HttpTask      `bson:"http_tasks" json:"http_tasks"`
	MailTasks     []*MailTask      `bson:"mail_tasks" json:"mail_tasks"`
	NotifyTasks   []*NotifyTask    `bson:"notify_tasks" json:"notify_tasks"`
	StorageTasks  []*StorageTask   `bson:"storage_tasks" json:"storage_tasks"`
	AsyncCallTask []*AsyncCallTask `bson:"async_call_task" json:"async_call_task"`
	TimerTask     []*TimerTask     `bson:"timer_task" json:"timer_task"`
	EventTask     []*EventTask     `bson:"event_task" json:"event_task"`
}

type Gateways struct {
	Inclusive []*InclusiveGateway `bson:"inclusive" json:"inclusive"`
}

type FlowDataArray struct {
	ProcessId     string        `bson:"_id" json:"_id"`
	CreateAt      int64         `bson:"create_at" json:"create_at"`
	CreateUserId  string        `bson:"create_user_id" json:"create_user_id"`
	AppId         string        `bson:"app_id" json:"app_id"`
	Name          string        `bson:"name" json:"name"`
	Desc          string        `bson:"desc" json:"desc"`
	LinkDataArray []interface{} `bson:"link_data_array" json:"link_data_array"`
	NodeDataArray []interface{} `bson:"node_data_array" json:"node_data_array"`
	Image         string        `bson:"image" json:"image"`
}

type SearchFlowItem struct {
	ProcessId string `json:"process_id"`
	UpdateAt  string `json:"update_at"`
	Name      string `json:"name"`
	AppId     string `json:"app_id"`
	Desc      string `json:"desc"`
}

type SearchFlowResponse struct {
	Name      string          `json:"name"`
	Timestamp int64           `json:"timestamp"`
	Fields    *SearchFlowItem `json:"fields"`
	Id        string          `json:"id"`
}

type FlowImage struct {
	ProcessId string `bson:"process_id"`
	Image     string `bson:"image"`
}
