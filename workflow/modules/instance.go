package modules

import (
	"konekko.me/gosion/workflow/flowerr"
	"konekko.me/gosion/workflow/models"
)

type IInstance interface {
	FindById(instanceId string) (*models.Instance, *flowerr.Error)

	//current status
	Status(instanceId string) (int64, *flowerr.Error)

	IsStarted(instanceId string) (bool, *flowerr.Error)

	//nodes running the current instance
	CurrentProcess(instanceId string)

	//开始新的实例
	New(ins *models.Instance) *flowerr.Error

	HasPermission() *flowerr.Error

	//更新实例当前进行的节点
	UpdateInstanceCurrentNodes(instanceId string, nodeIds ...string) *flowerr.Error

	FindRequireUserProcessingInstances(userId string, pageIndex, pageSize int64)

	IsUserRequireProcessingThatNode(userId string, nodeId string)
}
