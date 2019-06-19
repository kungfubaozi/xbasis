package userhandlers

import (
	"context"
	"gopkg.in/mgo.v2"
	"konekko.me/gosion/analysis/client"
	"konekko.me/gosion/commons/config/call"
	"konekko.me/gosion/commons/dto"
	"konekko.me/gosion/commons/errstate"
	"konekko.me/gosion/commons/generator"
	"konekko.me/gosion/commons/regx"
	"konekko.me/gosion/commons/wrapper"
	external "konekko.me/gosion/user/pb"
	"konekko.me/gosion/user/pb/inner"
)

type inviteService struct {
	session     *mgo.Session
	userService gosionsvc_internal_user.UserService
	log         analysisclient.LogClient
	id          gs_commons_generator.IDGenerator
}

func (svc *inviteService) SetState(ctx context.Context, in *external.SetStateRequest, out *gs_commons_dto.Status) error {
	return gs_commons_wrapper.ContextToAuthorize(ctx, out, func(auth *gs_commons_wrapper.WrapperUser) *gs_commons_dto.State {
		return nil
	})
}

func (svc *inviteService) GetDetail(ctx context.Context, in *external.HasInvitedRequest, out *external.GetDetailResponse) error {
	return gs_commons_wrapper.ContextToAuthorize(ctx, out, func(auth *gs_commons_wrapper.WrapperUser) *gs_commons_dto.State {
		return nil
	})
}

func (svc *inviteService) Search(ctx context.Context, in *external.InviteSearchRequest, out *external.InviteSearchResponse) error {
	return gs_commons_wrapper.ContextToAuthorize(ctx, out, func(auth *gs_commons_wrapper.WrapperUser) *gs_commons_dto.State {
		return nil
	})
}

func (svc *inviteService) HasInvited(ctx context.Context, in *external.HasInvitedRequest, out *gs_commons_dto.Status) error {
	return gs_commons_wrapper.ContextToAuthorize(ctx, out, func(auth *gs_commons_wrapper.WrapperUser) *gs_commons_dto.State {
		return nil
	})
}

func (svc *inviteService) GetRepo() *inviteRepo {
	return &inviteRepo{session: svc.session.Clone()}
}

/**
必须要填写初始化配置的信息，registerType
当用户注册时会检测(按照registerType查找对应的数据匹配)是否有对应邀请用户，如果有则会合并数据，没有则进入正常流程
如果被邀请用户已经注册会不通过
*/
func (svc *inviteService) User(ctx context.Context, in *external.InviteDetail, out *gs_commons_dto.Status) error {
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

		repo := svc.GetRepo()
		defer repo.Close()

		userId := svc.id.Get()

		//repo.IsExists()

		return nil
	})
}

func NewInviteService(session *mgo.Session, log analysisclient.LogClient) external.InviteHandler {
	return &inviteService{session: session, log: log}
}
