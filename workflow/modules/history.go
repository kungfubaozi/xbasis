package modules

import (
	"context"
	"konekko.me/xbasis/workflow/flowerr"
	"konekko.me/xbasis/workflow/models"
)

type IHistory interface {
	GetInstanceNodeHistory(instanceId, nodeId string)

	GetInstanceOperateHistory(instanceId string)

	//记录操作历史
	Record(ctx context.Context, data *models.History) *flowerr.Error
}
