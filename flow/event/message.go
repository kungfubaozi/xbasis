package event

import (
	"context"
)

//消息事件
type messageStartEvent struct {
	typeEvent
	Reference string `bson:"reference" json:"reference"`
}

func (e *messageStartEvent) Do(ctx context.Context, value interface{}) {
	panic("implement me")
}
