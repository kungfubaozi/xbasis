package modules

import (
	"konekko.me/gosion/workflow/flowerr"
	"konekko.me/gosion/workflow/models"
)

type IStore interface {
	IsFinished(nodeId string, instanceId string) (bool, *flowerr.Error)

	ClearParentNodesStatus(nodeId string, instanceId string) (bool, *flowerr.Error)

	Finished(store *models.Holder) *flowerr.Error
}
