package modules

import (
	"context"
	"konekko.me/gosion/workflow/flowerr"
)

type IRuntime interface {
	Submit(ctx context.Context, instanceId, nodeId string, value map[string]interface{}) *flowerr.Error

	RunningProcessSize() int
}
