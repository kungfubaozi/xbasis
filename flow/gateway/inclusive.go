package gateway

import (
	"context"
	"konekko.me/gosion/flow/base"
)

//包容网关
//包容网关会寻找所有符合条件的流向
//如果没有符合的，那么就去走默认的flow
type InclusiveGateway struct {
	*base.Info
}

func (g *InclusiveGateway) Do(ctx context.Context, value interface{}) {
	panic("implement me")
}
