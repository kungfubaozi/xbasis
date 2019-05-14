package modules

import (
	"konekko.me/gosion/workflow/flowerr"
	"konekko.me/gosion/workflow/models"
)

type IProcesses interface {
	AddProcess(p *models.Process)

	//重新分配某个节点的操作人
	Reassignment()

	FindNode(instanceId, nodeId string) (interface{}, *flowerr.Error)
}
