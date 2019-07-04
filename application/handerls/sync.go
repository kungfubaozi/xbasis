package applicationhanderls

import (
	"context"
	"fmt"
	"github.com/garyburd/redigo/redis"
	"gopkg.in/mgo.v2"
	inner "konekko.me/xbasis/application/pb/inner"
	constants "konekko.me/xbasis/commons/constants"
	commons "konekko.me/xbasis/commons/dto"
	"konekko.me/xbasis/commons/errstate"
	wrapper "konekko.me/xbasis/commons/wrapper"
	"konekko.me/xbasis/permission/pb"
	"konekko.me/xbasis/permission/pb/inner"
	"konekko.me/xbasis/user/pb"
	"sync"
)

type syncService struct {
	session           *mgo.Session
	inviteService     xbasissvc_external_user.InviteService
	accessibleService xbasissvc_internal_permission.AccessibleService
	bindingService    xbasissvc_external_permission.BindingService
	groupService      xbasissvc_external_permission.UserGroupService
	pool              *redis.Pool
}

func (svc *syncService) GetRepo() *syncRepo {
	return &syncRepo{session: svc.session.Clone()}
}

func (svc *syncService) GetAppRepo() *applicationRepo {
	return getApplicationRepo(svc.session.Clone(), nil, svc.pool.Get())
}

func (svc *syncService) Check(ctx context.Context, in *inner.CheckRequest, out *commons.Status) error {
	return wrapper.ContextToAuthorize(ctx, out, func(auth *wrapper.WrapperUser) *commons.State {

		if len(in.UserId) > 0 && len(in.AppId) > 0 && len(auth.Token.Relation) > 10 {

			repo := svc.GetRepo()
			defer repo.Close()

			//Check if the user is synchronized to the corresponding application
			c, err := repo.Synced(in.UserId, in.AppId, auth.Token.Relation)
			if err == nil && c {
				return errstate.Success
			}

			return errstate.ErrUserNotAuthorize

		}

		fmt.Println("enter 1")

		return nil
	})
}

//这里的操作主要有两个
//1.获取app的AllowNewUser设置项，设置用户在此app的默认权限
//2.同步用户信息至app（SyncUserURL）
func (svc *syncService) Update(ctx context.Context, in *inner.UserInfo, out *commons.Status) error {
	return wrapper.ContextToAuthorize(ctx, out, func(auth *wrapper.WrapperUser) *commons.State {

		if len(in.GId) > 0 && len(in.AppId) > 0 {

			appRepo := svc.GetAppRepo()
			defer appRepo.Close()

			appInfo, err := appRepo.GetApplication(in.AppId)
			if err != nil {
				return nil
			}

			next := true
			switch appInfo.Type {
			case constants.AppTypeRoute, constants.AppTypeSafe:
				next = false
			}

			if next {
				//sync to target application
				if len(appInfo.Settings.SyncUserURL) > 0 {

				}

				i, err := svc.inviteService.HasInvited(ctx, &xbasissvc_external_user.HasInvitedRequest{
					UserId: in.GId,
					AppId:  in.AppId,
				})
				if err != nil {
					return nil
				}

				if !i.State.Ok {
					return i.State
				}

				invited := len(i.UserId) > 16

				var roles []string
				var bindGroupIds []string

				if invited {
					as, err := svc.inviteService.GetDetail(ctx, &xbasissvc_external_user.HasInvitedRequest{
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
						bindGroupIds = append(bindGroupIds, item.BindGroupIds...)
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

				resp := func(s2 *commons.State) {
					if s1.Ok {
						s1 = s2
					}
				}

				go func() {
					defer wg.Done()
					if len(roles) > 0 {
						//binding
						s, err := svc.bindingService.UserRole(ctx, &xbasissvc_external_permission.BindingRolesRequest{
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
						s, err := svc.groupService.AddUser(ctx, &xbasissvc_external_permission.AddUserRequest{
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

func NewSyncService(session *mgo.Session, inviteService xbasissvc_external_user.InviteService,
	accessibleService xbasissvc_internal_permission.AccessibleService,
	bindingService xbasissvc_external_permission.BindingService,
	groupService xbasissvc_external_permission.UserGroupService, pool *redis.Pool) inner.UserSyncHandler {
	return &syncService{session: session, inviteService: inviteService, accessibleService: accessibleService, bindingService: bindingService, groupService: groupService, pool: pool}
}
