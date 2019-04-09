package gs_commons_wrapper

import (
	"context"
	"github.com/micro/go-micro/metadata"
	"github.com/micro/go-micro/server"
	"konekko.me/gosion/commons/dto"
	"konekko.me/gosion/commons/errstate"
	"konekko.me/gosion/permission/client"
	"konekko.me/gosion/permission/pb"
)

func AuthWrapper(fn server.HandlerFunc) server.HandlerFunc {
	verificationClient := permissioncli.NewVerificationClient()
	return func(ctx context.Context, req server.Request, rsp interface{}) error {

		status, err := verificationClient.Test(ctx, &gs_service_permission.HasPermissionRequest{})
		if err != nil {
			rsp = &gs_commons_dto.Status{State: errstate.ErrRequest}
			return nil
		}

		if !status.State.Ok {
			rsp = &gs_commons_dto.Status{State: status.State}
			return nil
		}

		//compressed volume
		ctx = metadata.NewContext(context.Background(), map[string]string{
			"Transport-User":       status.User,
			"Transport-AppId":      status.AppId,
			"Transport-ClientId":   status.ClientId,
			"transport-traceId":    status.TraceId,
			"Transport-Ip":         status.Ip,
			"Transport-UserDevice": status.UserDevice,
			"Transport-UserAgent":  status.UserAgent,
		})

		fn(ctx, req, rsp)
		return nil
	}
}
