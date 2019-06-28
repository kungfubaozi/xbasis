package modules

import (
	"konekko.me/xbasis/workflow/flowerr"
	"konekko.me/xbasis/workflow/models"
)

type IStore interface {
	//节点是否完成
	IsFinished(nodeId string, instanceId string) (bool, *flowerr.Error)

	//清除节点关联的所有状态
	ClearRelationNodesStatus(nodeId string, instanceId string) (bool, *flowerr.Error)

	//完成
	Finished(store *models.Holder) *flowerr.Error

	//添加实例忽略节点
	AddIgnoreNode(ignore *models.NodeIgnore) *flowerr.Error

	//清除当前节点所有的忽略节点
	ClearAboutNodeIgnoreNodes(nodeId, instanceId string) *flowerr.Error

	//获取节点所有的忽略节点
	GetInstanceIgnoreNodes(instanceId string) ([]string, *flowerr.Error)
}
