package userhandlers

import (
	"context"
	"gopkg.in/mgo.v2"
	"konekko.me/gosion/commons/dto"
	"konekko.me/gosion/commons/wrapper"
	external "konekko.me/gosion/user/pb"
)

type registerService struct {
	session       *mgo.Session
	inviteService external.InviteService
}

//自注册的用户只能有访问当前项目的权限
//管理员invite可以选择可以访问哪些项目
func (svc *registerService) New(ctx context.Context, in *external.NewRequest, out *gs_commons_dto.Status) error {
	return gs_commons_wrapper.ContextToAuthorize(ctx, out, func(auth *gs_commons_wrapper.WrapperUser) *gs_commons_dto.State {

		return nil
	})
}

func NewRegisterService(session *mgo.Session) external.RegisterHandler {
	return &registerService{session: session}
}
