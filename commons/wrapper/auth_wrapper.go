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
	verificationClient := permissioncli.NewVerificationClient(c)
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

		//fmt.Println("verification clear data is", status)

		if status.State.Ok {

			cm := make(map[string]string)
			cm["transport-user"] = status.User
			cm["transport-app-id"] = status.AppId
			cm["transport-client-id"] = status.ClientId
			cm["transport-trace-id"] = status.TraceId
			cm["transport-ip"] = status.Ip
			cm["transport-user-device"] = status.UserDevice
			cm["transport-user-agent"] = status.UserAgent
			cm["transport-client-platform"] = fmt.Sprintf("%d", status.Platform)

			if status.Token != nil {
				cm["transport-token-user-id"] = status.Token.UserId
				cm["transport-token-client-platform"] = fmt.Sprintf("%d", status.Token.Platform)
				cm["transport-token-client-id"] = status.Token.ClientId
			}

			//compressed volume
			ctx = metadata.NewContext(ctx, cm)
		}

		return fn(ctx, req, rsp)
	}
}
