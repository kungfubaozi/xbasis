package userhandlers

import (
	"context"
	"gopkg.in/mgo.v2"
	commons "konekko.me/xbasis/commons/dto"
	wrapper "konekko.me/xbasis/commons/wrapper"
	external "konekko.me/xbasis/user/pb"
)

type updateService struct {
	session *mgo.Session
}

func (svc *updateService) Username(ctx context.Context, in *external.UpdateRequest, out *commons.Status) error {
	return wrapper.ContextToAuthorize(ctx, out, func(auth *wrapper.WrapperUser) *commons.State {
		return nil
	})
}

func (svc *updateService) RealName(ctx context.Context, in *external.UpdateRequest, out *commons.Status) error {
	return wrapper.ContextToAuthorize(ctx, out, func(auth *wrapper.WrapperUser) *commons.State {
		return nil
	})
}

func (svc *updateService) Phone(ctx context.Context, in *external.UpdateRequest, out *commons.Status) error {
	return wrapper.ContextToAuthorize(ctx, out, func(auth *wrapper.WrapperUser) *commons.State {
		return nil
	})
}

func (svc *updateService) Email(ctx context.Context, in *external.UpdateRequest, out *commons.Status) error {
	return wrapper.ContextToAuthorize(ctx, out, func(auth *wrapper.WrapperUser) *commons.State {
		return nil
	})
}

func (svc *updateService) UserIcon(ctx context.Context, in *external.UpdateRequest, out *commons.Status) error {
	return wrapper.ContextToAuthorize(ctx, out, func(auth *wrapper.WrapperUser) *commons.State {
		return nil
	})
}

func (svc *updateService) Password(ctx context.Context, in *external.UpdatePasswordRequest, out *commons.Status) error {
	return wrapper.ContextToAuthorize(ctx, out, func(auth *wrapper.WrapperUser) *commons.State {
		return nil
	})
}

func NewUpdateService(session *mgo.Session) external.UpdateHandler {
	return &updateService{session: session}
}
