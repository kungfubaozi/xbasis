package gs_commons_wrapper

import (
	"context"
	"errors"
	"github.com/micro/go-micro/metadata"
	"konekko.me/gosion/commons/dto"
	"konekko.me/gosion/commons/errstate"
	"reflect"
	"strconv"
	"time"
)

type WrapperUser struct {
	User         string
	AppId        string
	FromClientId string
	RefClientId  string
	IP           string
	TraceId      string
	UserAgent    string
	UserDevice   string
	Access       *DurationAccessUser
	Platform     int64
	AppType      int64
	Token        *WrapperUserToken
}

func (w *WrapperUser) GetClientId() string {
	if len(w.RefClientId) == 0 {
		return w.FromClientId
	}
	return w.RefClientId
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
	To   string
	Auth bool
}

type WrapperEvent func(auth *WrapperUser) *gs_commons_dto.State

func GetData(md metadata.Metadata) *WrapperUser {
	auth := &WrapperUser{}
	auth.User = md["transport-user"]
	auth.AppId = md["transport-app-id"]
	auth.FromClientId = md["transport-from-client-id"]
	auth.RefClientId = md["transport-ref-client-id"]
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
	dau := &DurationAccessUser{
		To: md["transport-duration-access-to"],
	}
	a, err = strconv.ParseInt(md["transport-duration-access-auth"], 10, 64)
	if err == nil && a > 0 {
		if a == 2 {
			dau.Auth = false
		} else {
			dau.Auth = true
		}
	}
	auth.Access = dau
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
			v := errstate.ErrRequest
			v.Timestamp = time.Now().Unix()
			s.Set(reflect.ValueOf(errstate.ErrRequest))
		}
	}

	if ok {
		auth := GetData(md)

		if auth.Token != nil && len(auth.Token.UserId) > 0 {
			if auth.Token.AppType == -1 || auth.Platform == -1 || auth.AppType == -1 {
				null()
				return nil
			}
		}

		v := event(auth)
		if v != nil {
			v.Timestamp = time.Now().Unix()
			s.Set(reflect.ValueOf(v))
			return nil
		}
	}

	null()
	return nil
}
