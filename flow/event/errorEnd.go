package event

import "context"

type errorStartEvent struct {
}

func (e *errorStartEvent) Do(ctx context.Context, value interface{}) {
	panic("implement me")
}
