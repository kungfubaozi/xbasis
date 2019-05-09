package modules

import (
	"konekko.me/gosion/workflow/flowerr"
	"konekko.me/gosion/workflow/models"
)

type Pipeline interface {
	GetNode(nodeId string) (*models.Node, *flowerr.Error)

	Id() string

	Dump()

	Flows(nodeId string) ([]*models.SequenceFlow, *flowerr.Error)

	FindParallelNodes(id string) []string

	GetNodeBackwardRelations(id string) []*models.NodeBackwardRelation
}
