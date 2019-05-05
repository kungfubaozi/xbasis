package gateway

import "context"

//排他网关
//寻找第一条符合条件的流向
type exclusiveGateway struct {
}

func (g *exclusiveGateway) Do(ctx context.Context, value interface{}) {
	panic("implement me")
}
