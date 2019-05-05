package event

import (
	"context"
)

type startEvent struct {
	typeEvent
	FormRef string `bson:"form_ref" json:"form_ref"`
}

func (e *startEvent) Do(ctx context.Context, value interface{}) {
	panic("implement me")
}
