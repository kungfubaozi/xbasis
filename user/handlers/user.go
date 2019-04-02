package userhandlers

import (
	"context"
	"gopkg.in/mgo.v2"
	"konekko.me/gosion/commons/dto"
	"konekko.me/gosion/commons/wrapper"
	"konekko.me/gosion/user/pb"
)

type userService struct {
	session *mgo.Session
}

func (svc *userService) FindUserInfoById(ctx context.Context, in *gs_service_user.FindRequest, out *gs_service_user.SimpleUserInfo) error {
	return gs_commons_wrapper.ContextToAuthorize(ctx, out, func(auth *gs_commons_wrapper.WrapperUser) *gs_commons_dto.State {
		return nil
	})
}

func (svc *userService) FindUserIdByPhone(ctx context.Context, in *gs_service_user.FindRequest, out *gs_service_user.UserIdResponse) error {
	return gs_commons_wrapper.ContextToAuthorize(ctx, out, func(auth *gs_commons_wrapper.WrapperUser) *gs_commons_dto.State {
		return nil
	})
}

func (svc *userService) FindUserIdByEmail(ctx context.Context, in *gs_service_user.FindRequest, out *gs_service_user.UserIdResponse) error {
	return gs_commons_wrapper.ContextToAuthorize(ctx, out, func(auth *gs_commons_wrapper.WrapperUser) *gs_commons_dto.State {
		return nil
	})
}

func (svc *userService) FindUserIdByAccount(ctx context.Context, in *gs_service_user.FindRequest, out *gs_service_user.UserIdResponse) error {
	return gs_commons_wrapper.ContextToAuthorize(ctx, out, func(auth *gs_commons_wrapper.WrapperUser) *gs_commons_dto.State {
		return nil
	})
}

func NewUserService(session *mgo.Session) gs_service_user.UserHandler {
	return &userService{session: session}
}
