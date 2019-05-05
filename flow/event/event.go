package event

import "context"

type typeEvent struct {
	Id                 string   `bson:"id" json:"id"`
	Name               string   `bson:"name" json:"name"`
	Desc               string   `bson:"desc" json:"desc"`
	CreateAt           int64    `bson:"create_at" json:"create_at"`
	ExecutionListeners []string `bson:"execution_listeners" json:"execution_listeners"`
}

type Event interface {
	Do(ctx context.Context, value interface{})
}

func NewMessageStartEvent() Event {
	return &messageStartEvent{}
}

func NewErrorStartEvent() Event {
	return &errorStartEvent{}
}

func NewTimerStartEvent() Event {
	return &timerStartEvent{}
}
