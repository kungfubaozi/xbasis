package event

import "context"

//定时启动事件
type timerStartEvent struct {
	typeEvent
	TimeCron string `bson:"time_cron" json:"time_cron"`
}

func (e *timerStartEvent) Do(ctx context.Context, value interface{}) {
	panic("implement me")
}

