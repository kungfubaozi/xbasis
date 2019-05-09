package modules

import (
	"context"
	"konekko.me/gosion/workflow/flowerr"
	"konekko.me/gosion/workflow/models"
)

type IForm interface {
	FindById(id string) (*models.TypeForm, *flowerr.Error)

	Submit(ctx context.Context, instanceId, nodeId, formId string, encryption bool, value map[string]interface{}) *flowerr.Error

	//load key.form
	LoadNodeDataToStore(ctx context.Context, instanceId, nodeId string) (map[string]interface{}, *flowerr.Error)
}
