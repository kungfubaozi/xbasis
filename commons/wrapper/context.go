package gs_commons_wrapper

import (
	"context"
	"errors"
	"github.com/micro/go-micro/metadata"
	"konekko.me/gosion/commons/dto"
	gserrors "konekko.me/gosion/commons/errstate"
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
	Token      *WrapperUserToken
}

type WrapperUserToken struct {
	ClientId       string
	ClientPlatform int64
	AppId          string
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
	auth.Platform = 0
	a, err := strconv.ParseInt(md["transport-client-platform"], 10, 64)
	if err == nil {
		auth.Platform = a
	}
	wut := &WrapperUserToken{
		ClientId: md["transport-token-client-id"],
		AppId:    md["transport-token-app-id"],
	}
	a, err = strconv.ParseInt(md["transport-token-client-platform"], 10, 64)
	if err == nil {
		wut.ClientPlatform = a
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
	if ok {
		auth := GetData(md)

		//fmt.Println("the verification ctx is:", auth)

		v := event(auth)
		if v != nil {
			s.Set(reflect.ValueOf(v))
			return nil
		}
	}
	if s.IsNil() {
		s.Set(reflect.ValueOf(gserrors.ErrRequest))
	}
	return nil
}
