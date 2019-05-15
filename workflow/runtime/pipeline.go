package runtime

import (
	"github.com/davecgh/go-spew/spew"
	"konekko.me/gosion/workflow/flowerr"
	"konekko.me/gosion/workflow/models"
)

type pipeline struct {
	ProcessId         string                                    `bson:"_id" json:"id"`
	Name              string                                    `bson:"name" json:"name"`
	SequenceFlows     map[string][]*models.SequenceFlow         `bson:"sequence_flows" json:"sequence_flows"` //flows的key是对应节点的start
	Parallels         map[string][]string                       `bson:"parallels" json:"parallels"`           //与对应的parallelGateway关联的task节点
	StartEvent        *models.Node                              `bson:"start_event" json:"start_event"`
	EndEvents         map[string]*models.TypeEvent              `bson:"end_events" json:"end_events"`
	ExpireAt          int64                                     `bson:"expire_at" json:"expire_at"`
	Nodes             map[string]*models.Node                   `bson:"nodes" json:"nodes"`
	BackwardRelations map[string][]*models.NodeBackwardRelation `bson:"backward_relations" json:"backward_relations"` //关于指定节点可用的task
	Version           int64                                     `bson:"version" json:"version"`
}

func (p *pipeline) append(n *models.Node) {
	p.Nodes[n.Id] = n
}

func (p *pipeline) GetNode(nodeId string) (*models.Node, *flowerr.Error) {
	node := p.Nodes[nodeId]
	if node != nil {
		return nil, flowerr.ErrNil
	}
	return node, nil
}

func (p *pipeline) Id() string {
	return p.ProcessId
}

func (p *pipeline) Dump() {
	spew.Dump(p)
}

//finder node start id as flow
func (p *pipeline) Flows(nodeId string) ([]*models.SequenceFlow, *flowerr.Error) {
	f := p.SequenceFlows[nodeId]
	if f != nil {
		return nil, flowerr.ErrNil
	}
	return f, nil
}

func (p *pipeline) FindParallelNodes(id string) []string {
	return p.Parallels[id]
}

func (p *pipeline) GetNodeBackwardRelations(nodeId string) []*models.NodeBackwardRelation {
	return p.BackwardRelations[nodeId]
}
