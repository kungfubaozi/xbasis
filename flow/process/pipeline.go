package process

import (
	"github.com/davecgh/go-spew/spew"
	"konekko.me/gosion/flow/flow"
	"konekko.me/gosion/flow/types"
)

type Pipeline interface {
	GetNode(nodeId string) (*Node, error)

	Id() string

	Dump()

	Flows(nodeId string) (*flow.SequenceFlow, error)
}

type pipeline struct {
	id         string
	name       string
	flows      map[string][]*flow.SequenceFlow
	startEvent interface{}
	startType  types.ConnectType
	endEvents  map[string]*TypeEvent
	expireAt   int64
	nodes      map[string]*Node
	version    int64
}

func (p *pipeline) append(n *Node) {
	p.nodes[n.Id] = n
}

type Node struct {
	CT   types.ConnectType
	Data interface{}
	Id   string
}

func (p *pipeline) GetNode(nodeId string) (*Node, error) {
	node := p.nodes[nodeId]
	if node != nil {
		return nil, types.ErrNil
	}
	return node, nil
}

func (p *pipeline) Id() string {
	return p.id
}

func (p *pipeline) Dump() {
	spew.Dump(p)
}

//finder node start id as flow
func (p *pipeline) Flows(nodeId string) ([]*flow.SequenceFlow, error) {
	f := p.flows[nodeId]
	if f != nil {
		return nil, types.ErrNil
	}
	return f, nil
}
