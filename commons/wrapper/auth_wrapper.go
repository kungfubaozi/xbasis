package xbasiswrapper

import (
	"context"
	"fmt"
	"github.com/micro/go-micro/client"
	"github.com/micro/go-micro/metadata"
	"github.com/micro/go-micro/server"
	"github.com/pkg/errors"
	"konekko.me/xbasis/commons/dto"
	"konekko.me/xbasis/commons/errstate"
	"konekko.me/xbasis/permission/client"
	"konekko.me/xbasis/permission/pb/inner"
	"reflect"
)

func set(rsp interface{}, state *xbasis_commons_dto.State) error {
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

		status, err := verificationClient.Check(ctx, &xbasissvc_internal_permission.HasPermissionRequest{})
		if err != nil {
			return set(rsp, errstate.ErrRequest)
		}

		// spew.Dump(status)

		if !status.State.Ok {
			return set(rsp, status.State)
		}

		//fmt.Println("verification clear data is", status)

		if status.State.Ok {

			cm := make(map[string]string)
			cm["transport-user"] = status.User
			cm["transport-app-id"] = status.AppId
			cm["transport-from-client-id"] = status.FromClient
			cm["transport-ref-client-id"] = status.RefClientId
			cm["transport-trace-id"] = status.TraceId
			cm["transport-log-id"] = status.LogId
			cm["transport-log-index"] = status.LogIndex
			cm["transport-ip"] = status.Ip
			cm["transport-function-id"] = status.FunctionId
			cm["transport-user-device"] = status.UserDevice
			cm["transport-user-agent"] = status.UserAgent
			cm["transport-app-type"] = fmt.Sprintf("%d", status.AppType)
			cm["transport-client-platform"] = fmt.Sprintf("%d", status.Platform)
			cm["transport-duration-access-to"] = status.DatTo
			cm["transport-duration-access-auth"] = fmt.Sprintf("%d", status.DatAuth)

			if status.Token != nil {
				//fmt.Println("set token data")

				cm["transport-token-user-id"] = status.Token.UserId
				cm["transport-token-app-id"] = status.Token.AppId
				cm["transport-token-client-platform"] = fmt.Sprintf("%d", status.Token.Platform)
				cm["transport-token-client-id"] = status.Token.ClientId
				cm["transport-token-user-relation"] = status.Token.Relation
				cm["transport-token-app-type"] = fmt.Sprintf("%d", status.Token.AppType)
			}

			//compressed volume
			ctx = metadata.NewContext(ctx, cm)

			//spew.Dump(cm)
		}

		return fn(ctx, req, rsp)
	}
}
