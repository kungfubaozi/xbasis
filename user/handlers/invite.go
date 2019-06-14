package userhandlers

import (
	"context"
	"gopkg.in/mgo.v2"
	"konekko.me/gosion/analysis/client"
	"konekko.me/gosion/commons/config/call"
	"konekko.me/gosion/commons/dto"
	"konekko.me/gosion/commons/errstate"
	"konekko.me/gosion/commons/regx"
	"konekko.me/gosion/commons/wrapper"
	external "konekko.me/gosion/user/pb"
	"konekko.me/gosion/user/pb/inner"
)

type inviteService struct {
	session     *mgo.Session
	userService gosionsvc_internal_user.UserService
	log         analysisclient.LogClient
}

func (svc *inviteService) List(ctx context.Context, in *external.ListRequest, out *external.ListResponse) error {
	return gs_commons_wrapper.ContextToAuthorize(ctx, out, func(auth *gs_commons_wrapper.WrapperUser) *gs_commons_dto.State {
		return nil
	})
}

func (svc *inviteService) Find(ctx context.Context, in *external.FindInviteRequest, out *external.FindInviteResponse) error {
	return gs_commons_wrapper.ContextToAuthorize(ctx, out, func(auth *gs_commons_wrapper.WrapperUser) *gs_commons_dto.State {
		return nil
	})
}

func (svc *inviteService) HasInvited(ctx context.Context, in *external.HasInvitedRequest, out *gs_commons_dto.Status) error {
	return gs_commons_wrapper.ContextToAuthorize(ctx, out, func(auth *gs_commons_wrapper.WrapperUser) *gs_commons_dto.State {
		return nil
	})
}

/**
必须要填写初始化配置的信息，registerType
当用户注册时会检测(按照registerType查找对应的数据匹配)是否有对应邀请用户，如果有则会合并数据，没有则进入正常流程
如果被邀请用户已经注册会不通过
*/
func (svc *inviteService) User(ctx context.Context, in *external.InviteRequest, out *gs_commons_dto.Status) error {
	return gs_commons_wrapper.ContextToAuthorize(ctx, out, func(auth *gs_commons_wrapper.WrapperUser) *gs_commons_dto.State {
		configuration := serviceconfiguration.Get()
		if len(in.Phone) > 0 && gs_commons_regx.Phone(in.Phone) {
			return errstate.ErrFormatPhone
		}
		if len(in.Email) > 0 && gs_commons_regx.Email(in.Email) {
			return errstate.ErrFormatEmail
		}
		if configuration.RegisterType == 1001 { //phone
			if len(in.Phone) <= 8 {
				return errstate.ErrRequest
			}
		} else if configuration.RegisterType == 1002 { //email
			if len(in.Email) <= 8 {
				return errstate.ErrRequest
			}
		}

		return nil
	})
}

func NewInviteService(session *mgo.Session, log analysisclient.LogClient) external.InviteHandler {
	return &inviteService{session: session, log: log}
}
