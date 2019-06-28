package modules

import (
	"konekko.me/xbasis/workflow/flowerr"
	"konekko.me/xbasis/workflow/models"
)

type IProcesses interface {
	AddProcess(p *models.Process)

	//重新分配某个节点的操作人
	Reassignment()

	FindNode(instanceId, nodeId string) (interface{}, *flowerr.Error)

	GetFlowDataArray(processId string) (*models.FlowDataArray, error)

	SaveFlowDataArrays(data *models.FlowDataArray) error

	Search(appId, name string, pageIndex, pageSize int64) ([]*models.SearchFlowItem, error)

	SaveImage(processId string, base64 string) error

	GetImage(processId string) (string, error)
}
