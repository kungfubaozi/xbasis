package gs_commons_wrapper

import (
	"context"
	"github.com/micro/go-micro/server"
)

func AuthWrapper(fn server.HandlerFunc) server.HandlerFunc {
	return func(ctx context.Context, req server.Request, rsp interface{}) error {

		//fmt.Println("header", req.Header())

		fn(ctx, req, rsp)
		return nil
	}
}
