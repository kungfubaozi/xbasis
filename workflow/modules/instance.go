package modules

import (
	"konekko.me/gosion/workflow/flowerr"
	"konekko.me/gosion/workflow/models"
)

type IInstance interface {
	//通过实例id查找
	FindById(instanceId string) (*models.Instance, *flowerr.Error)

	//当前状态
	Status(instanceId string) (int64, *flowerr.Error)

	//实例是否启动
	IsStarted(instanceId string) (bool, *flowerr.Error)

	//nodes running the current instance
	CurrentProcess(instanceId string)

	//开始新的实例
	New(ins *models.Instance) *flowerr.Error

	//使用名称开始
	NamedStart(name string) *flowerr.Error

	HasPermission() *flowerr.Error

	//更新实例当前进行的节点
	UpdateInstanceCurrentNodes(instanceId string, nodeIds ...string) *flowerr.Error

	//通过用户ID查找关联的实例信息
	FindRequireUserProcessingInstances(userId string, pageIndex, pageSize int64)

	//用户是否可以处理指定节点
	IsUserRequireProcessingThatNode(userId string, nodeId string)
}
