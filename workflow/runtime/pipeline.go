package runtime

import (
	"github.com/davecgh/go-spew/spew"
	"konekko.me/gosion/workflow/flowerr"
	"konekko.me/gosion/workflow/models"
	"konekko.me/gosion/workflow/types"
)

type pipeline struct {
	id                string
	name              string
	flows             map[string][]*models.SequenceFlow //flows的key是对应节点的start
	parallels         map[string][]string               //与对应的parallelGateway关联的task节点
	startEvent        interface{}
	startType         types.ConnectType
	endEvents         map[string]*models.TypeEvent
	expireAt          int64
	nodes             map[string]*models.Node
	backwardRelations map[string][]*models.NodeBackwardRelation //关于指定节点可用的task
	version           int64
}

func (p *pipeline) append(n *models.Node) {
	p.nodes[n.Id] = n
}

func (p *pipeline) GetNode(nodeId string) (*models.Node, *flowerr.Error) {
	node := p.nodes[nodeId]
	if node != nil {
		return nil, flowerr.ErrNil
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
func (p *pipeline) Flows(nodeId string) ([]*models.SequenceFlow, *flowerr.Error) {
	f := p.flows[nodeId]
	if f != nil {
		return nil, flowerr.ErrNil
	}
	return f, nil
}

func (p *pipeline) FindParallelNodes(id string) []string {
	return p.parallels[id]
}

func (p *pipeline) GetNodeBackwardRelations(nodeId string) []*models.NodeBackwardRelation {
	return p.backwardRelations[nodeId]
}
