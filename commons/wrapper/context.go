package gs_commons_wrapper

import (
	"context"
	"errors"
	"fmt"
	"github.com/micro/go-micro/metadata"
	"konekko.me/gosion/commons/dto"
	"konekko.me/gosion/commons/errstate"
	"reflect"
	"strconv"
)

type WrapperUser struct {
	User       string
	AppId      string
	ClientId   string
	IP         string
	TraceId    string
	UserAgent  string
	UserDevice string
	DAU        *DurationAccessUser
	Platform   int64
	AppType    int64
	Token      *WrapperUserToken
}

type WrapperUserToken struct {
	ClientId       string
	ClientPlatform int64
	AppId          string
	AppType        int64
	Relation       string
	UserId         string
}

type DurationAccessUser struct {
	SendTo string
}

type WrapperEvent func(auth *WrapperUser) *gs_commons_dto.State

func GetData(md metadata.Metadata) *WrapperUser {
	auth := &WrapperUser{}
	auth.User = md["transport-user"]
	auth.AppId = md["transport-app-id"]
	auth.ClientId = md["transport-client-id"]
	auth.IP = md["transport-ip"]
	auth.TraceId = md["transport-trace-id"]
	auth.UserAgent = md["transport-user-agent"]
	auth.UserDevice = md["transport-user-device"]
	auth.Platform = -1
	auth.AppType = -1
	a, err := strconv.ParseInt(md["transport-app-type"], 10, 64)
	if err == nil && a > 0 {
		auth.AppType = a
	}
	a, err = strconv.ParseInt(md["transport-client-platform"], 10, 64)
	if err == nil && a > 0 {
		auth.Platform = a
	}
	wut := &WrapperUserToken{
		ClientId: md["transport-token-client-id"],
		AppId:    md["transport-token-app-id"],
		Relation: md["transport-token-relation"],
		UserId:   md["transport-token-user-id"],
	}
	wut.ClientPlatform = -1
	a, err = strconv.ParseInt(md["transport-token-client-platform"], 10, 64)
	if err == nil && a > 0 {
		wut.ClientPlatform = a
	}
	wut.AppType = -1
	a, err = strconv.ParseInt(md["transport-token-app-type"], 10, 64)
	if err == nil && a > 0 {
		wut.AppType = a
	}
	auth.Token = wut
	return auth
}

func ContextToAuthorize(ctx context.Context, out interface{}, event WrapperEvent) error {
	s := reflect.ValueOf(out).Elem().FieldByName("State")
	if !s.CanSet() {
		return errors.New("err return type canSet")
	}

	md, ok := metadata.FromContext(ctx)

	null := func() {
		if s.IsNil() {
			s.Set(reflect.ValueOf(errstate.ErrRequest))
		}
	}

	if ok {
		auth := GetData(md)
		//if auth.Platform == -1 || auth.AppType == -1 {
		//	null()
		//	return nil
		//}

		fmt.Println("entry md ok")

		if auth.Token != nil && len(auth.Token.UserId) > 0 {
			if auth.Token.AppType == -1 || auth.Platform == -1 || auth.AppType == -1 {
				null()
				return nil
			}
		}

		v := event(auth)
		if v != nil {
			s.Set(reflect.ValueOf(v))
			return nil
		}
	}

	fmt.Println("context md null")

	null()
	return nil
}
