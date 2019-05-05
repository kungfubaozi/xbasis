package gateway

import (
	"context"
	"konekko.me/gosion/flow/base"
)

//排他网关
//寻找第一条符合条件的流向
type ExclusiveGateway struct {
	*base.Info
}

func (g *ExclusiveGateway) Do(ctx context.Context, value interface{}) {
	panic("implement me")
}
