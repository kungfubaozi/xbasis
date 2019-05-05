package gateway

import "context"

type Gateway interface {
	Do(ctx context.Context, value interface{})
}

func NewExclusiveGateway() Gateway {
	return &exclusiveGateway{}
}

func NewparallerGateway() Gateway {
	return &parallelGateway{}
}

func NewInclusiveGateway() Gateway {
	return &inclusiveGateway{}
}
