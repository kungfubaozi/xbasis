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

func (s *store) AddIgnoreNode(ignore *models.NodeIgnore) *flowerr.Error {
	panic("implement me")
}

func (s *store) ClearAboutNodeIgnoreNodes(nodeId, instanceId string) *flowerr.Error {
	panic("implement me")
}

func (s *store) GetInstanceIgnoreNodes(instanceId string) ([]string, *flowerr.Error) {
	panic("implement me")
}

var storeIndex = "gs-flow-store"

func (s *store) IsFinished(nodeId string, instanceId string) (bool, *flowerr.Error) {
	c, err := s.client.Count(storeIndex, map[string]interface{}{"node_id": nodeId, "instance_id": instanceId})
	if err != nil {
		return false, flowerr.FromError(err)
	}
	return c > 0, nil
}

func (s *store) ClearRelationNodesStatus(nodeId string, instanceId string) (bool, *flowerr.Error) {
	bq := elastic.NewBoolQuery().Must(elastic.NewMatchPhraseQuery("instance_id", instanceId), elastic.NewMatchQuery("relation_nodes", nodeId))
	s1 := elastic.NewScript("ctx._source['status'] = params.status").Params(map[string]interface{}{"status": 0}).Lang("painless")
	result, err := s.client.GetElasticClient().UpdateByQuery().Query(bq).Script(s1).Do(context.Background())
	if err != nil {
		return false, flowerr.FromError(err)
	}
	if result.Total > 0 {
		return true, nil
	}
	return false, flowerr.ErrUnknow
}

//节点完成
func (s *store) Finished(store *models.Holder) *flowerr.Error {
	id, err := s.client.AddData(storeIndex, store)
	if err != nil {
		return flowerr.FromError(err)
	}
	if len(id) > 0 {
		return nil
	}
	return flowerr.ErrUnknow
}
