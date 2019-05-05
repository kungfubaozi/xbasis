package flow

import (
	"konekko.me/gosion/flow/base"
	"konekko.me/gosion/flow/types"
)

type SequenceFlow struct {
	*base.Info
	Script      string            `bson:"script" json:"script"`
	DefaultFlow bool              `bson:"default_flow" json:"default_flow"`
	Start       string            `bson:"start" json:"start"`
	StartType   types.ConnectType `bson:"start_type" json:"start_type"`
	End         string            `bson:"end" json:"end"`
	EndType     types.ConnectType `bson:"end_type" json:"end_type"`
	Priority    int64             `bson:"priority" json:"priority"`
}
