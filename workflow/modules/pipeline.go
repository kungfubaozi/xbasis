package modules

import (
	"github.com/davecgh/go-spew/spew"
	"konekko.me/gosion/workflow/models"
	"konekko.me/gosion/workflow/types"
)

type pipeline struct {
	id         string
	name       string
	flows      map[string][]*models.SequenceFlow
	startEvent interface{}
	startType  types.ConnectType
	endEvents  map[string]*models.TypeEvent
	expireAt   int64
	nodes      map[string]*node
	version    int64
}

type node struct {
	id   string
	ct   types.ConnectType
	data interface{}
}

type Pipeline interface {
	GetNode(nodeId string) (*node, error)

	Id() string

	Dump()

	Flows(nodeId string) ([]*models.SequenceFlow, error)
}

func (p *pipeline) append(n *node) {
	p.nodes[n.id] = n
}

func (p *pipeline) GetNode(nodeId string) (*node, error) {
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
func (p *pipeline) Flows(nodeId string) ([]*models.SequenceFlow, error) {
	f := p.flows[nodeId]
	if f != nil {
		return nil, types.ErrNil
	}
	return f, nil
}
