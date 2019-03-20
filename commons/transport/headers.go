package gs_commons_transport

import (
	"context"
	"github.com/go-kit/kit/transport/grpc"
	"google.golang.org/grpc/metadata"
	"strconv"
	"zskparker.com/foundation/base/pb"
	"zskparker.com/foundation/pkg/constants"
)

type ApiRequestHeaders struct {
	UserId   string
	AppId    string
	ClientId string
	IP       string
}

func GRPCToContext() grpc.ServerRequestFunc {
	return func(ctx context.Context, mds metadata.MD) context.Context {
		header, ok := mds["authorization"]
		if !ok {
			ctx.Err()
			return ctx
		}
		meta := &fs_base.Metadata{}
		meta.ClientId = header[0]
		meta.Ip = header[1]
		meta.UserAgent = header[2]
		meta.Api = header[3]
		meta.Token = header[4]
		meta.Device = header[5]
		meta.UserId = header[6]
		i, e := strconv.ParseInt(header[7], 10, 64)
		if e != nil || i == 0 {
			i = fs_constants.LEVEL_TOURIST
		}
		meta.Level = i
		meta.Session = header[8]
		meta.InitSession = header[9]

		return context.WithValue(ctx, "meta-info", meta)
	}
}
