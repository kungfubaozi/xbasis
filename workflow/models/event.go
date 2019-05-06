package models

import "konekko.me/gosion/workflow/types"

type NodeEvent struct {
	Id                 string   `bson:"id" json:"id"`
	Name               string   `bson:"name" json:"name"`
	Desc               string   `bson:"desc" json:"desc"`
	CreateAt           int64    `bson:"create_at" json:"create_at"`
	ExecutionListeners []string `bson:"execution_listeners" json:"execution_listeners"`
}

type StartEvent struct {
	*NodeEvent
	FormRef string `bson:"form_ref" json:"form_ref"`
}

type EndEvent struct {
	*NodeEvent
}

type CancelEndEvent struct {
	*NodeEvent
}

type MessageStartEvent struct {
	*NodeEvent
}

type TerminateEndEvent struct {
	*NodeEvent
}

type ApiStartEvent struct {
	*NodeEvent
	Method  types.HttpRequestMethod `bson:"method" json:"method"`
	FormRef string                  `bson:"form_ref" json:"form_ref"`
}

type TimerStartEvent struct {
	*NodeEvent
	Corn string `bson:"corn" json:"corn"`
}

type TriggerStartEvent struct {
	*NodeEvent
	Trigger string `bson:"trigger" json:"trigger"`
}
