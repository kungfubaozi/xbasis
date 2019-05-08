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
	UserTasks     []*UserTask     `bson:"user_tasks" json:"user_tasks"`
	ApiTasks      []*ApiTask      `bson:"api_tasks" json:"api_tasks"`
	DecisionTasks []*DecisionTask `bson:"decision_tasks" json:"decision_tasks"`
	HttpTasks     []*HttpTask     `bson:"http_tasks" json:"http_tasks"`
	MailTasks     []*MailTask     `bson:"mail_tasks" json:"mail_tasks"`
	NotifyTasks   []*NotifyTask   `bson:"notify_tasks" json:"notify_tasks"`
	StorageTasks  []*StorageTask  `bson:"storage_tasks" json:"storage_tasks"`
}

type Gateways struct {
	Exclusives []*ExclusiveGateway `bson:"exclusive" json:"exclusive"`
	Inclusive  []*InclusiveGateway `bson:"inclusive" json:"inclusive"`
	Parallels  []*ParallelGateway  `bson:"parallel" json:"parallel"`
}
