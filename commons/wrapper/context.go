package gs_commons_wrapper

import (
	"context"
	"errors"
	"github.com/micro/go-micro/metadata"
	"konekko.me/gosion/commons/dto"
	gserrors "konekko.me/gosion/commons/errstate"
	"reflect"
)

type WrapperEvent func(auth *gs_commons_dto.Authorize) *gs_commons_dto.State

func ContextToAuthorize(ctx context.Context, out interface{}, event WrapperEvent) error {
	s := reflect.ValueOf(out).Elem().FieldByName("State")
	if !s.CanSet() {
		return errors.New("err return type canSet")
	}
	state := s.Interface().(*gs_commons_dto.State)
	if state == nil {
		state = new(gs_commons_dto.State)
	}
	s.Set(reflect.ValueOf(state))
	md, ok := metadata.FromContext(ctx)
	if ok {
		auth := gs_commons_dto.Authorize{}
		auth.UserId = md["TRANSPORT-USER-ID"]
		auth.AppId = md["TRANSPORT-APP-ID"]
		auth.ClientId = md["TRANSPORT-CLIENT-ID"]
		auth.UserAgent = md["TRANSPORT-USER-AGENT"]
		auth.Ip = md["X-Real-IP"]
		state = event(&auth)
	}
	if state == nil {
		state = gserrors.ErrRequest
	}
	return nil
}
