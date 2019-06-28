package models

import "konekko.me/xbasis/workflow/types"

type History struct {
	*Info
	InstanceId string            `bson:"instance_id" json:"instance_id"`
	Operate    types.OperateType `bson:"operate" json:"operate"`
	NodeId     string            `bson:"node_id" json:"node_id"`
	Comments   string            `bson:"comments" json:"comments"`
	Status     int64             `bson:"status" json:"status"`
}
