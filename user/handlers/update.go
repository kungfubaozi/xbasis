package user_handlers

import (
	"context"
	"gopkg.in/mgo.v2"
	"konekko.me/gosion/commons/dto"
	"konekko.me/gosion/commons/wrapper"
	"konekko.me/gosion/user/pb"
)

type updateService struct {
	session *mgo.Session
}

func (svc *updateService) Username(ctx context.Context, in *gs_service_user.UpdateRequest, out *gs_commons_dto.Status) error {
	return gs_commons_wrapper.ContextToAuthorize(ctx, out, func(auth *gs_commons_wrapper.WrapperUser) *gs_commons_dto.State {
		return nil
	})
}

func (svc *updateService) RealName(ctx context.Context, in *gs_service_user.UpdateRequest, out *gs_commons_dto.Status) error {
	return gs_commons_wrapper.ContextToAuthorize(ctx, out, func(auth *gs_commons_wrapper.WrapperUser) *gs_commons_dto.State {
		return nil
	})
}

func (svc *updateService) Phone(ctx context.Context, in *gs_service_user.UpdateRequest, out *gs_commons_dto.Status) error {
	return gs_commons_wrapper.ContextToAuthorize(ctx, out, func(auth *gs_commons_wrapper.WrapperUser) *gs_commons_dto.State {
		return nil
	})
}

func (svc *updateService) Email(ctx context.Context, in *gs_service_user.UpdateRequest, out *gs_commons_dto.Status) error {
	return gs_commons_wrapper.ContextToAuthorize(ctx, out, func(auth *gs_commons_wrapper.WrapperUser) *gs_commons_dto.State {
		return nil
	})
}

func (svc *updateService) UserIcon(ctx context.Context, in *gs_service_user.UpdateRequest, out *gs_commons_dto.Status) error {
	return gs_commons_wrapper.ContextToAuthorize(ctx, out, func(auth *gs_commons_wrapper.WrapperUser) *gs_commons_dto.State {
		return nil
	})
}

func (svc *updateService) Password(ctx context.Context, in *gs_service_user.UpdatePasswordRequest, out *gs_commons_dto.Status) error {
	return gs_commons_wrapper.ContextToAuthorize(ctx, out, func(auth *gs_commons_wrapper.WrapperUser) *gs_commons_dto.State {
		return nil
	})
}

func NewUpdateService() gs_service_user.UpdateHandler {
	return &updateService{}
}
