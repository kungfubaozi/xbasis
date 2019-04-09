package gs_commons_wrapper

import (
	"context"
	"errors"
	"github.com/micro/go-micro/metadata"
	"konekko.me/gosion/commons/dto"
	gserrors "konekko.me/gosion/commons/errstate"
	"reflect"
)

type WrapperUser struct {
	User       string
	AppId      string
	ClientId   string
	IP         string
	TraceId    string
	UserAgent  string
	UserDevice string
	DAU        DurationAccessUser
}

type DurationAccessUser struct {
	SendTo string
}

type WrapperEvent func(auth *WrapperUser) *gs_commons_dto.State

func ContextToAuthorize(ctx context.Context, out interface{}, event WrapperEvent) error {
	s := reflect.ValueOf(out).Elem().FieldByName("State")
	if !s.CanSet() {
		return errors.New("err return type canSet")
	}

	md, ok := metadata.FromContext(ctx)
	if ok {
		auth := &WrapperUser{}
		auth.User = md["Transport-User"]
		auth.AppId = md["Transport-AppId"]
		auth.ClientId = md["Transport-ClientId"]
		auth.IP = md["Transport-Ip"]
		auth.TraceId = md["transport-traceId"]
		auth.UserAgent = md["Transport-UserAgent"]
		auth.UserDevice = md["Transport-UserDevice"]
		v := event(auth)
		if v != nil {
			s.Set(reflect.ValueOf(v))
		}
	}
	if s.IsNil() {
		s.Set(reflect.ValueOf(gserrors.ErrRequest))
	}
	return nil
}
