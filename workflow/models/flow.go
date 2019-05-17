package models

import (
	"konekko.me/gosion/workflow/types"
)

type SequenceFlow struct {
	*Info
	Script      string            `bson:"script" json:"script"`
	DefaultFlow bool              `bson:"default_flow" json:"default_flow"`
	Start       string            `bson:"start" json:"start"`
	StartType   types.ConnectType `bson:"start_type" json:"start_type"`
	End         string            `bson:"end" json:"end"`
	EndType     types.ConnectType `bson:"end_type" json:"end_type"`
	Rollback    bool              `bson:"rollback" json:"rollback"` //rollback肯定是个具体的task. 不可能是Gateway
	Priority    int64             `bson:"priority" json:"priority"`
}

type NextStatus struct {
	Again        string
	CurrentNodes []string
}

type NodeBackwardRelation struct {
	Id  string
	Key string
	CT  types.ConnectType
}
