package applicationhanderls

import (
	"context"
	"gopkg.in/mgo.v2"
	inner "konekko.me/gosion/application/pb/inner"
	"konekko.me/gosion/commons/constants"
	"konekko.me/gosion/commons/dto"
	"konekko.me/gosion/commons/errstate"
	"konekko.me/gosion/commons/indexutils"
	"konekko.me/gosion/commons/wrapper"
	"konekko.me/gosion/permission/pb"
	"konekko.me/gosion/permission/pb/inner"
	"konekko.me/gosion/user/pb"
	"sync"
)

type syncService struct {
	*indexutils.Client
	session           *mgo.Session
	inviteService     gosionsvc_external_user.InviteService
	accessibleService gosionsvc_internal_permission.AccessibleService
	bindingService    gosionsvc_external_permission.BindingService
	groupService      gosionsvc_external_permission.UserGroupService
}

func (svc *syncService) GetRepo() *syncRepo {
	return &syncRepo{Client: svc.Client, session: svc.session.Clone()}
}

func (svc *syncService) GetAppRepo() *applicationRepo {
	return &applicationRepo{session: svc.session.Clone()}
}

func (svc *syncService) Check(ctx context.Context, in *inner.CheckRequest, out *gs_commons_dto.Status) error {
	return gs_commons_wrapper.ContextToAuthorize(ctx, out, func(auth *gs_commons_wrapper.WrapperUser) *gs_commons_dto.State {

		if len(in.UserId) > 0 && len(in.AppId) > 0 && len(auth.Token.Relation) > 16 {

			repo := svc.GetRepo()
			defer repo.Close()

			//Check if the user is synchronized to the corresponding application
			c, err := repo.Synced(in.UserId, in.AppId, auth.Token.Relation)
			if err == nil && c == 1 {
				return errstate.Success
			}

			return errstate.ErrUserNotAuthorize

		}

		return nil
	})
}

//这里的操作主要有两个
//1.获取app的AllowNewUser设置项，设置用户在此app的默认权限
//2.同步用户信息至app（SyncUserURL）
func (svc *syncService) Update(ctx context.Context, in *inner.UserInfo, out *gs_commons_dto.Status) error {
	return gs_commons_wrapper.ContextToAuthorize(ctx, out, func(auth *gs_commons_wrapper.WrapperUser) *gs_commons_dto.State {
		if len(in.GId) > 0 && len(in.AppId) > 0 {
			appRepo := svc.GetAppRepo()
			defer appRepo.Close()

			appInfo, err := appRepo.GetApplication(in.AppId)
			if err != nil {
				return nil
			}

			next := true
			switch appInfo.Type {
			case gs_commons_constants.AppTypeManage:
			case gs_commons_constants.AppTypeRoute:
			case gs_commons_constants.AppTypeSafe:
			case gs_commons_constants.AppTypeUser:
				next = false
			}

			if next {
				//sync to target application
				if len(appInfo.Settings.SyncUserURL) > 0 {

				}

				i, err := svc.inviteService.HasInvited(ctx, &gosionsvc_external_user.HasInvitedRequest{
					UserId: in.GId,
					AppId:  in.AppId,
				})
				if err != nil {
					return nil
				}

				if !i.State.Ok {
					return i.State
				}

				invited := len(i.Content) > 16

				var roles []string
				var bindGroupIds []string

				if invited {
					as, err := svc.inviteService.GetDetail(ctx, &gosionsvc_external_user.HasInvitedRequest{
						UserId: in.GId,
						AppId:  in.AppId,
					})
					if err != nil {
						return nil
					}
					if !as.State.Ok {
						return as.State
					}
					if len(as.Items) > 0 {
						item := as.Items[0]
						roles = item.Roles
						bindGroupIds = append(bindGroupIds, item.BindGroupId)
					}
				}

				if appInfo.Settings.AllowNewUsers != nil && len(appInfo.Settings.AllowNewUsers.DefaultRole) > 0 {
					allow := appInfo.Settings.AllowNewUsers
					roles = append(roles, allow.DefaultRole...)
					if len(allow.DefaultGroup) > 0 {
						bindGroupIds = append(bindGroupIds, allow.DefaultGroup)
					}
				}

				var wg sync.WaitGroup
				wg.Add(2)

				s1 := errstate.Success

				resp := func(s2 *gs_commons_dto.State) {
					if s1.Ok {
						s1 = s2
					}
				}

				go func() {
					defer wg.Done()
					if len(roles) > 0 {
						//binding
						s, err := svc.bindingService.UserRole(ctx, &gosionsvc_external_permission.BindingRolesRequest{
							Id:    in.GId,
							Roles: roles,
							AppId: in.AppId,
						})
						if err != nil {
							resp(errstate.ErrRequest)
							return
						}
						if !s.State.Ok {
							resp(s.State)
							return
						}
					}
				}()

				go func() {
					defer wg.Done()
					if len(bindGroupIds) > 0 {
						s, err := svc.groupService.AddUser(ctx, &gosionsvc_external_permission.AddUserRequest{
							GroupIds: bindGroupIds,
							UserId:   in.GId,
							AppId:    in.AppId,
						})
						if err != nil {
							resp(errstate.ErrRequest)
							return
						}
						if !s.State.Ok {
							resp(s.State)
							return
						}
					}
				}()

				wg.Wait()

				if !s1.Ok {
					return s1
				}
			}

			repo := svc.GetRepo()
			defer repo.Close()

			err = repo.Sync(in.GId, in.AppId, auth.Token.Relation)
			if err == nil {
				return errstate.Success
			}
		}
		return nil
	})
}

func NewSyncService(client *indexutils.Client, session *mgo.Session, inviteService gosionsvc_external_user.InviteService,
	accessibleService gosionsvc_internal_permission.AccessibleService,
	bindingService gosionsvc_external_permission.BindingService,
	groupService gosionsvc_external_permission.UserGroupService) inner.UserSyncHandler {
	return &syncService{Client: client, session: session, inviteService: inviteService, accessibleService: accessibleService, bindingService: bindingService, groupService: groupService}
}
