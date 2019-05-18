package runtime

import (
	"context"
	"github.com/olivere/elastic"
	"konekko.me/gosion/commons/gslogrus"
	"konekko.me/gosion/commons/indexutils"
	"konekko.me/gosion/workflow/flowerr"
	"konekko.me/gosion/workflow/models"
)

type store struct {
	client *indexutils.Client
	log    *gslogrus.Logger
}

var storeIndex = "gs-flow-store"

func (s *store) IsFinished(nodeId string, instanceId string) (bool, *flowerr.Error) {
	panic("implement me")
}

func (s *store) ClearParentNodesStatus(nodeId string, instanceId string) (bool, *flowerr.Error) {
	panic("implement me")
}

//流程实例化时保存状态, 初始状态
func (s *store) Finished(store *models.Holder) *flowerr.Error {

}
