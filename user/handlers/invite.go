package userhandlers

import (
	"context"
	"github.com/olivere/elastic"
	"gopkg.in/mgo.v2"
	"konekko.me/xbasis/analysis/client"
	"konekko.me/xbasis/commons/config/call"
	constants "konekko.me/xbasis/commons/constants"
	commons "konekko.me/xbasis/commons/dto"
	"konekko.me/xbasis/commons/errstate"
	generator "konekko.me/xbasis/commons/generator"
	"konekko.me/xbasis/commons/indexutils"
	regx "konekko.me/xbasis/commons/regx"
	wrapper "konekko.me/xbasis/commons/wrapper"
	external "konekko.me/xbasis/user/pb"
	"konekko.me/xbasis/user/pb/inner"
	"time"
)

type inviteService struct {
	session          *mgo.Session
	log              analysisclient.LogClient
	id               generator.IDGenerator
	client           *indexutils.Client
	innerUserService xbasissvc_internal_user.UserService
}

func (svc *inviteService) SetState(ctx context.Context, in *external.SetStateRequest, out *commons.Status) error {
	return wrapper.ContextToAuthorize(ctx, out, func(auth *wrapper.WrapperUser) *commons.State {

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
	return wrapper.ContextToAuthorize(ctx, out, func(auth *wrapper.WrapperUser) *commons.State {

		if len(in.UserId) < 10 {
			return nil
		}

		repo := svc.GetRepo()
		defer repo.Close()

		m, err := repo.FindByKey("user_id", in.UserId)
		if err != nil {
			return nil
		}

		out.Username = m.Username
		out.RealName = m.RealName
		out.Email = m.Email
		out.Phone = m.Phone

		var items []*external.InviteItem
		for _, v := range m.Items {
			if len(in.AppId) > 8 && v.AppId != in.AppId {
				continue
			}
			items = append(items, &external.InviteItem{
				AppId:        v.AppId,
				BindGroupIds: v.BingGroupIds,
				Roles:        v.Roles,
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
	return wrapper.ContextToAuthorize(ctx, out, func(auth *wrapper.WrapperUser) *commons.State {

		return nil
	})
}

func (svc *inviteService) HasInvited(ctx context.Context, in *external.HasInvitedRequest, out *external.HasInvitedResponse) error {
	return wrapper.ContextToAuthorize(ctx, out, func(auth *wrapper.WrapperUser) *commons.State {

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
如果被邀请用户已经注册
*/
func (svc *inviteService) User(ctx context.Context, in *external.InviteUserRequest, out *commons.Status) error {
	return wrapper.ContextToAuthorize(ctx, out, func(auth *wrapper.WrapperUser) *commons.State {
		configuration := serviceconfiguration.Get()

		header := &analysisclient.LogHeaders{
			TraceId:     auth.TraceId,
			ServiceName: constants.UserService,
			ModuleName:  "InviteUser",
		}

		key := ""
		value := ""

		if len(in.Phone) > 0 && !regx.Phone(in.Phone) {
			return errstate.ErrFormatPhone
		}
		if len(in.Email) > 0 && !regx.Email(in.Email) {
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

		if len(in.Account) > 0 {

			v, err := svc.client.GetElasticClient().Search(typeUserIndex).Type("_doc").
				Query(elastic.NewBoolQuery().Must(elastic.NewMatchPhraseQuery("fields.account", in.Account))).
				Do(context.Background())

			if err != nil {
				return nil
			}

			if v.Hits.TotalHits > 0 {
				return errstate.ErrAccountAlreadyExists
			}
		}

		if len(in.Email) > 0 {

			v, err := svc.client.GetElasticClient().Search(typeUserIndex).Type("_doc").
				Query(elastic.NewBoolQuery().Must(elastic.NewMatchPhraseQuery("fields.email", in.Email))).
				Do(context.Background())

			if err != nil {
				return nil
			}

			if v.Hits.TotalHits > 0 {
				return errstate.ErrEmailAlreadyExists
			}
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
				Account:      in.Account,
				Side:         false, //side的作用是判断user是内部还是外部新的
				State:        constants.InviteStateOfRegister,
			}

			var items []*inviteItem

			for _, v := range in.Items {
				items = append(items, &inviteItem{
					AppId:        v.AppId,
					Roles:        v.Roles,
					BingGroupIds: v.BindGroupIds,
				})
			}

			m.Items = items

			err = repo.Add(m)
			if err != nil {
				return errstate.ErrRequest
			}

			svc.log.Info(&analysisclient.LogContent{
				Headers: header,
				Action:  "NewInviteUser",
				Fields: &analysisclient.LogFields{
					"username":  in.Username,
					"user_id":   m.UserId,
					"timestamp": time.Now().Unix(),
				},
				Index: &analysisclient.LogIndex{
					Name: "users",
					Id:   m.UserId,
					Fields: &analysisclient.LogFields{
						"username":    in.Username,
						"real_name":   in.RealName,
						"phone":       in.Phone,
						"email":       in.Email,
						"user_id":     m.UserId,
						"invite":      true,
						"from_invite": true,
						"timestamp":   time.Now().Unix(),
						"state":       m.State,
					},
				},
			})

			return errstate.Success
		}

		if err != nil {
			return nil
		}

		return errstate.ErrHasInvited
	})
}

/**
Append是在邀请用户或已经注册用户中添加邀请信息，不同与User接口
*/
func (svc *inviteService) Append(ctx context.Context, in *external.AppendRequest, out *commons.Status) error {
	return wrapper.ContextToAuthorize(ctx, out, func(auth *wrapper.WrapperUser) *commons.State {
		if len(in.UserId) < 10 || in.Item == nil {
			return nil
		}

		header := &analysisclient.LogHeaders{
			TraceId:     auth.TraceId,
			ServiceName: constants.UserService,
			ModuleName:  "InviteUser",
		}

		//检查用户是否被注册/被邀请
		//如果是邀请用户则直接在原有内容上继续添加
		//如果是注册用户，则新建invite，保存

		repo := svc.GetRepo()
		defer repo.Close()

		v, err := svc.innerUserService.GetUserInfoById(ctx, &xbasissvc_internal_user.GetUserInfoByIdRequest{
			UserId: in.UserId,
		})

		if err != nil {
			return nil
		}

		if !v.State.Ok {
			return v.State
		}

		m, err := repo.FindByKey("user_id", in.UserId)
		e := false
		if err != nil && err == mgo.ErrNotFound {
			err = nil
			e = true
		}

		if err != nil {
			return nil
		}

		if e {
			s, err := svc.innerUserService.IsExists(ctx, &xbasissvc_internal_user.ExistsRequest{
				UserId: in.UserId,
			})

			if err != nil {
				return nil
			}

			if !s.State.Ok {
				return s.State
			}

			u := &inviteModel{
				UserId: in.UserId,
				Side:   true,
				Items: []*inviteItem{
					{
						BingGroupIds: in.Item.BindGroupIds,
						AppId:        in.Item.AppId,
						Roles:        in.Item.Roles,
					},
				},
				CreateAt:     time.Now().UnixNano(),
				CreateUserId: auth.Token.UserId,
			}

			//path search

			err = repo.Add(u)

			if err == nil {

				svc.log.Info(&analysisclient.LogContent{
					Headers: header,
					Action:  "InviteSideUserToApp",
					Fields: &analysisclient.LogFields{
						"app_id":    in.Item.AppId,
						"user_id":   m.UserId,
						"timestamp": time.Now().Unix(),
					},
					Index: &analysisclient.LogIndex{
						Name: "users",
						Id:   m.UserId,
						Fields: &analysisclient.LogFields{
							"user_id":                         m.UserId,
							"invite":                          true,
							"side":                            true,
							"from_invite":                     false,
							"authorize_apps." + in.Item.AppId: true,
						},
					},
				})

			}

		} else {
			//已经邀请但未注册
			if m == nil {
				return nil
			}

			//去除原有的
			for k, v := range m.Items {
				if v.AppId == in.Item.AppId {
					i := k - 1
					if i < 0 {
						i = 0
					}
					if len(in.Item.Roles) > 0 {
						v.Roles = append(v.Roles, in.Item.Roles...)
					}
					if len(in.Item.BindGroupIds) > 0 {
						v.BingGroupIds = append(v.BingGroupIds, in.Item.BindGroupIds...)
					}
					break
				}
			}

			err = repo.UpdateItems(in.UserId, m.Items)

			if err == nil {
				svc.log.Info(&analysisclient.LogContent{
					Headers: header,
					Action:  "AppendInviteUserAccess",
					Fields: &analysisclient.LogFields{
						"app_id":    in.Item.AppId,
						"user_id":   m.UserId,
						"timestamp": time.Now().Unix(),
					},
					Index: &analysisclient.LogIndex{
						Name: "users",
						Id:   m.UserId,
						Fields: &analysisclient.LogFields{
							"apps." + in.Item.AppId: true,
						},
					},
				})
			}
		}

		if err != nil {
			return nil
		}

		return errstate.Success
	})
}

func NewInviteService(session *mgo.Session, log analysisclient.LogClient, innerUserService xbasissvc_internal_user.UserService) external.InviteHandler {
	return &inviteService{session: session, log: log, id: generator.NewIDG(), innerUserService: innerUserService}
}
