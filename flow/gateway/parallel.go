package gateway

import (
	"context"
	"konekko.me/gosion/flow/base"
)

//并行网关
//允许将流程分成多条分支
//即使顺序流中定义了条件，也会被忽略
type ParallelGateway struct {
	*base.Info
}

func (g *ParallelGateway) Do(ctx context.Context, value interface{}) {
	panic("implement me")
}
