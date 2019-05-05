package process

import (
	"context"
	"github.com/davecgh/go-spew/spew"
	"konekko.me/gosion/flow/flow"
	"konekko.me/gosion/flow/types"
)

type Pipeline interface {
	Do(ctx context.Context, value interface{})

	Id() string

	Dump()

	NextNode()
}

type pipeline struct {
	id         string
	name       string
	flows      map[string]*flow.SequenceFlow
	startEvent interface{}
	startType  types.ConnectType
	endEvents  map[string]interface{}
	expireAt   int64
}

func (p *pipeline) Do(ctx context.Context, value interface{}) {
	panic("implement me")
}

func (p *pipeline) Id() string {
	return p.id
}

func (p *pipeline) Dump() {
	spew.Dump(p)
}

func (p *pipeline) NextNode() {
	panic("implement me")
}
