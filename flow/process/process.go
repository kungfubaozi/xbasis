package process

import (
	"konekko.me/gosion/flow/flow"
	"konekko.me/gosion/flow/gateway"
	"konekko.me/gosion/flow/task"
	"konekko.me/gosion/flow/types"
)

type Process struct {
	Id           string               `bson:"_id" json:"id"`
	Name         string               `bson:"name" json:"name"`
	CreateAt     int64                `bson:"create_at" json:"create_at"`
	CreateUserId string               `bson:"create_user_id" json:"create_user_id"`
	Gateways     *Gateways            `bson:"gateways" json:"gateways"`
	Tasks        *Tasks               `bson:"tasks" json:"tasks"`
	Flows        []*flow.SequenceFlow `bson:"flows" json:"flows"`
	StartEvent   *TypeEvent           `bson:"start_event" json:"start_event"`
	EndEvents    []*TypeEvent         `bson:"end_events" json:"end_events"`
}

type TypeEvent struct {
	Id    string            `bson:"id" json:"id"`
	Key   string            `bson:"key" json:"key"`
	Event interface{}       `bson:"event" json:"event"`
	Type  types.ConnectType `bson:"type" json:"type"`
}

type Tasks struct {
	UserTasks     []*task.UserTask     `bson:"user_tasks" json:"user_tasks"`
	ApiTasks      []*task.ApiTask      `bson:"api_tasks" json:"api_tasks"`
	DecisionTasks []*task.DecisionTask `bson:"decision_tasks" json:"decision_tasks"`
	GRPCTasks     []*task.GRPCTask     `bson:"grpc_tasks" json:"grpc_tasks"`
	HttpTasks     []*task.HttpTask     `bson:"http_tasks" json:"http_tasks"`
	MailTasks     []*task.MailTask     `bson:"mail_tasks" json:"mail_tasks"`
	SendTasks     []*task.SendTask     `bson:"send_tasks" json:"send_tasks"`
	StorageTasks  []*task.StorageTask  `bson:"storage_tasks" json:"storage_tasks"`
}

type Gateways struct {
	Exclusives []*gateway.ExclusiveGateway `bson:"exclusive" json:"exclusive"`
	Inclusive  []*gateway.InclusiveGateway `bson:"inclusive" json:"inclusive"`
	Parallels  []*gateway.ParallelGateway  `bson:"parallel" json:"parallel"`
}
