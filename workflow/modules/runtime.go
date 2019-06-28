package modules

import (
	"context"
	"konekko.me/xbasis/workflow/flowerr"
)

type IRuntime interface {
	Submit(ctx context.Context, instanceId, nodeId string, value map[string]interface{}) *flowerr.Error

	RunningProcessSize() int
}
