package user_handlers

import (
	"context"
	"gopkg.in/mgo.v2"
	"konekko.me/gosion/user/pb"
)

type userService struct {
	session *mgo.Session
}

func (svc *userService) FindUserInfoById(ctx context.Context, in *gs_service_user.FindRequest, out *gs_service_user.SimpleUserInfo) error {
	panic("implement me")
}

func (svc *userService) FindUserIdByPhone(ctx context.Context, in *gs_service_user.FindRequest, out *gs_service_user.UserIdResponse) error {
	panic("implement me")
}

func (svc *userService) FindUserIdByEmail(ctx context.Context, in *gs_service_user.FindRequest, out *gs_service_user.UserIdResponse) error {
	panic("implement me")
}

func (svc *userService) FindUserIdByAccount(ctx context.Context, in *gs_service_user.FindRequest, out *gs_service_user.UserIdResponse) error {
	panic("implement me")
}

func NewUserService() gs_service_user.UserHandler {
	return &userService{}
}
