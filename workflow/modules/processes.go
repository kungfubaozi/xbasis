package modules

import (
	"konekko.me/gosion/workflow/flowerr"
	"konekko.me/gosion/workflow/models"
)

type AddProcessToPipelineCallback func(pip Pipeline)

type IProcesses interface {
	AddProcess(p *models.Process)

	SetCallback(callback AddProcessToPipelineCallback)

	//重新分配某个节点的操作人
	Reassignment()

	LoadAll()

	FindNode(instanceId, nodeId string) (interface{}, *flowerr.Error)
}
