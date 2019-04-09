package gs_commons_wrapper

import (
	"context"
	"fmt"
	"github.com/micro/go-micro/client"
	"github.com/micro/go-micro/metadata"
	"github.com/micro/go-micro/server"
	"github.com/pkg/errors"
	"konekko.me/gosion/commons/dto"
	"konekko.me/gosion/commons/errstate"
	"konekko.me/gosion/permission/client"
	"konekko.me/gosion/permission/pb"
	"reflect"
)

func set(rsp interface{}, state *gs_commons_dto.State) error {
	s := reflect.ValueOf(rsp).Elem().FieldByName("State")
	if s.CanSet() {
		s.Set(reflect.ValueOf(state))
		return nil
	}
	return errors.New("err")
}

func AuthWrapper(c client.Client, fn server.HandlerFunc) server.HandlerFunc {
	verificationClient := permissioncli.NewVerificationClient()
	return func(ctx context.Context, req server.Request, rsp interface{}) error {

		status, err := verificationClient.Check(ctx, &gs_service_permission.HasPermissionRequest{})
		if err != nil {
			fmt.Println("verification error", err)
			return set(rsp, errstate.ErrRequest)
		}

		if !status.State.Ok {
			fmt.Println("verification state error", rsp)
			return set(rsp, status.State)
		}

		fmt.Println("verification clear")

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

		return fn(ctx, req, rsp)
	}
}
