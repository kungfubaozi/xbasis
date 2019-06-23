package userhandlers

import (
	"context"
	"gopkg.in/mgo.v2"
	"konekko.me/gosion/analysis/client"
	"konekko.me/gosion/commons/config/call"
	"konekko.me/gosion/commons/constants"
	"konekko.me/gosion/commons/dto"
	"konekko.me/gosion/commons/errstate"
	"konekko.me/gosion/commons/generator"
	"konekko.me/gosion/commons/regx"
	"konekko.me/gosion/commons/wrapper"
	external "konekko.me/gosion/user/pb"
	"time"
)

type inviteService struct {
	session *mgo.Session
	log     analysisclient.LogClient
	id      gs_commons_generator.IDGenerator
}

func (svc *inviteService) SetState(ctx context.Context, in *external.SetStateRequest, out *gs_commons_dto.Status) error {
	return gs_commons_wrapper.ContextToAuthorize(ctx, out, func(auth *gs_commons_wrapper.WrapperUser) *gs_commons_dto.State {

		if len(in.UserId) > 10 && len(in.AppId) > 8 && in.State > 0 {
			repo := svc.GetRepo()
			defer repo.Close()

			err := repo.SetState(in.UserId, in.AppId, in.State)
			if err == nil {
				return errstate.Success
			}
		}

		return nil
	})
}

func (svc *inviteService) GetDetail(ctx context.Context, in *external.HasInvitedRequest, out *external.GetDetailResponse) error {
	return gs_commons_wrapper.ContextToAuthorize(ctx, out, func(auth *gs_commons_wrapper.WrapperUser) *gs_commons_dto.State {

		if len(in.UserId) < 10 {
			return nil
		}

		repo := svc.GetRepo()
		defer repo.Close()

		m, err := repo.FindByKey("user_id", in.UserId)
		if err != nil {
			return nil
		}

		var items []*external.InviteItem
		for _, v := range m.Items {
			if len(in.AppId) > 8 && v.AppId != in.AppId {
				continue
			}
			items = append(items, &external.InviteItem{
				AppId:       v.AppId,
				BindGroupId: v.BingGroupId,
				Roles:       v.Roles,
			})
			if len(in.AppId) > 8 {
				break
			}
		}

		out.Items = items

		return errstate.Success
	})
}

func (svc *inviteService) Search(ctx context.Context, in *external.InviteSearchRequest, out *external.InviteSearchResponse) error {
	return gs_commons_wrapper.ContextToAuthorize(ctx, out, func(auth *gs_commons_wrapper.WrapperUser) *gs_commons_dto.State {
		return nil
	})
}

func (svc *inviteService) HasInvited(ctx context.Context, in *external.HasInvitedRequest, out *external.HasInvitedResponse) error {
	return gs_commons_wrapper.ContextToAuthorize(ctx, out, func(auth *gs_commons_wrapper.WrapperUser) *gs_commons_dto.State {

		key := ""
		var value interface{}
		if len(in.UserId) > 10 {
			key = "user_id"
			value = in.UserId
		} else if len(in.Email) > 10 {
			key = "email"
			value = in.Email
		} else if len(in.Phone) > 10 {
			key = "phone"
			value = in.Phone
		}

		if len(key) == 0 || value == nil {
			return nil
		}

		repo := svc.GetRepo()
		defer repo.Close()

		m, err := repo.FindByKey(key, value)
		if err != nil && err == mgo.ErrNotFound {
			return errstate.Success
		}
		if err != nil {
			return nil
		}

		out.UserId = m.UserId
		out.Status = m.State

		return errstate.Success
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

		key := ""
		value := ""
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
			key = "phone"
			value = in.Phone
		} else if configuration.RegisterType == 1002 { //email
			if len(in.Email) <= 8 {
				return errstate.ErrRequest
			}
			key = "email"
			value = in.Email
		}

		if len(key) == 0 || len(value) < 6 {
			return nil
		}

		repo := svc.GetRepo()
		defer repo.Close()

		_, err := repo.FindByKey(key, value)
		if err != nil && err == mgo.ErrNotFound {
			m := &inviteModel{
				Phone:        in.Phone,
				Email:        in.Email,
				CreateAt:     time.Now().UnixNano(),
				CreateUserId: auth.Token.UserId,
				UserId:       svc.id.Get(),
				Username:     in.Username,
				RealName:     in.RealName,
				State:        gs_commons_constants.InviteStateOfWaiting,
			}

			var items []*inviteItem

			for _, v := range in.Items {
				items = append(items, &inviteItem{
					AppId:       v.AppId,
					Roles:       v.Roles,
					BingGroupId: v.BindGroupId,
				})
			}

			m.Items = items

			err = repo.Add(m)
			if err != nil {
				return errstate.ErrRequest
			}

			return errstate.Success
		}

		if err != nil {
			return nil
		}

		return errstate.ErrHasInvited
	})
}

func (svc *inviteService) Append(ctx context.Context, in *external.AppendRequest, out *gs_commons_dto.Status) error {
	return gs_commons_wrapper.ContextToAuthorize(ctx, out, func(auth *gs_commons_wrapper.WrapperUser) *gs_commons_dto.State {

		return nil
	})
}

func NewInviteService(session *mgo.Session, log analysisclient.LogClient) external.InviteHandler {
	return &inviteService{session: session, log: log, id: gs_commons_generator.NewIDG()}
}
