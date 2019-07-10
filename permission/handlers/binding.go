package permissionhandlers

import (
	"context"
	"gopkg.in/mgo.v2"
	"konekko.me/xbasis/analysis/client"
	constants "konekko.me/xbasis/commons/constants"
	commons "konekko.me/xbasis/commons/dto"
	"konekko.me/xbasis/commons/errstate"
	generator "konekko.me/xbasis/commons/generator"
	"konekko.me/xbasis/commons/indexutils"
	wrapper "konekko.me/xbasis/commons/wrapper"
	external "konekko.me/xbasis/permission/pb"
	"konekko.me/xbasis/user/pb/inner"
	"sync"
	"time"
)

type bindingService struct {
	*indexutils.Client
	session          *mgo.Session
	innerUserService xbasissvc_internal_user.UserService
	roleService      external.RoleService
	log              analysisclient.LogClient
}

func (svc *bindingService) GetRepo() *bindingRepo {
	return &bindingRepo{Client: svc.Client, session: svc.session.Clone(), id: generator.NewIDG()}
}

func (svc *bindingService) GetRoleRepo() *roleRepo {
	return &roleRepo{session: svc.session.Clone()}
}

func (svc *bindingService) UserRole(ctx context.Context, in *external.BindingRolesRequest, out *commons.Status) error {
	return wrapper.ContextToAuthorize(ctx, out, func(auth *wrapper.WrapperUser) *commons.State {
		if len(in.AppId) > 0 && len(in.Id) > 0 && len(in.Roles) > 0 {
			//check roles
			repo := svc.GetRepo()
			defer repo.Close()

			headers := &analysisclient.LogHeaders{
				TraceId:     auth.TraceId,
				ServiceName: constants.PermissionService,
				ModuleName:  "BindingUserRole",
			}

			roleRepo := svc.GetRoleRepo()
			defer roleRepo.Close()

			s := errstate.Success
			resp := func(s1 *commons.State) {
				if s.Ok {
					s = s1
				}
			}

			var wg sync.WaitGroup
			wg.Add(len(in.Roles))

			for _, v := range in.Roles {

				go func() {
					defer wg.Done()
					role, err := roleRepo.FindRoleById(v, in.AppId)

					if err != nil {
						resp(errstate.ErrSystem)
						return
					}

					if len(role.Id) == 0 {
						resp(errstate.ErrRequest)
						return
					}

					if role.AppId != in.AppId {
						resp(errstate.ErrRequest)
						return
					}
				}()

			}

			wg.Wait()
			if !s.Ok {
				return s
			}

			//去重
			role, err := repo.FindRelationUserById(in.Id, in.AppId)
			if err != nil && err == mgo.ErrNotFound {
				role = &userRolesRelation{
					CreateAt: time.Now().UnixNano(),
					AppId:    in.AppId,
					UserId:   in.Id,
				}
				err = nil
			}
			if err != nil {
				return nil
			}

			var roles []string
			for _, v := range in.Roles {
				ok := true
				for _, v1 := range role.Roles {
					if v == v1 {
						ok = false
						break
					}
				}
				if ok {
					roles = append(roles, v)
				}
			}

			if len(role.Roles) > 0 {
				roles = append(roles, role.Roles...)
			}

			role.Roles = roles

			//update database
			err = repo.SetUserRole(in.Id, in.AppId, role)
			if err != nil {
				return errstate.ErrRequest
			}

			svc.log.Info(&analysisclient.LogContent{
				Headers: headers,
				Action:  "UpdateUserRole",
				Fields: &analysisclient.LogFields{
					"userId": in.Id,
					"appId":  in.AppId,
					"roles":  in.Roles,
				},
				Index: &analysisclient.LogIndex{
					Name:     "users",
					Id:       in.Id,
					Relation: true,
					Fields: &analysisclient.LogFields{
						"roles": role.Roles,
					},
				},
			})

			return errstate.Success
		}

		return nil
	})
}

func (svc *bindingService) FunctionRole(ctx context.Context, in *external.BindingRolesRequest, out *commons.Status) error {
	return wrapper.ContextToAuthorize(ctx, out, func(auth *wrapper.WrapperUser) *commons.State {
		if len(in.AppId) > 0 && len(in.Id) > 0 && len(in.Roles) > 0 {

			headers := &analysisclient.LogHeaders{
				TraceId:     auth.TraceId,
				ServiceName: constants.PermissionService,
				ModuleName:  "BindingFunctionRole",
			}

			repo := svc.GetRepo()
			defer repo.Close()

			roleRepo := svc.GetRoleRepo()
			defer roleRepo.Close()

			s := errstate.Success
			resp := func(s1 *commons.State) {
				if s.Ok {
					s = s1
				}
			}

			var wg sync.WaitGroup
			wg.Add(len(in.Roles))

			for _, v := range in.Roles {
				go func() {
					defer wg.Done()
					role, err := roleRepo.FindRoleById(v, in.AppId)
					if err != nil {
						resp(errstate.ErrSystem)
						return
					}

					if len(role.Id) == 0 {
						resp(errstate.ErrRequest)
						return
					}

					//不在同一个应用内
					if role.AppId != in.AppId {
						resp(errstate.ErrRequest)
						return
					}
				}()
			}

			wg.Wait()

			if !s.Ok {
				return s
			}

			//先查找当前功能是否已经绑定角色
			role, err := repo.FindRelationFunctionById(in.Id, in.AppId)
			if err != nil {
				return errstate.ErrRequest
			}

			var roles []string
			for _, v := range in.Roles {
				ok := true
				for _, v1 := range role.Roles {
					if v == v1 {
						ok = false
						break
					}
				}
				if ok {
					roles = append(roles, v)
				}
			}

			if len(role.Roles) > 0 {
				roles = append(roles, role.Roles...)
			}

			role.Roles = roles

			//更新数据库
			err = repo.SetFunctionRole(in.Id, in.AppId, role)
			if err == nil {

				svc.log.Info(&analysisclient.LogContent{
					Headers: headers,
					Action:  "UpdateFunctionRole",
					Fields: &analysisclient.LogFields{
						"functionId": in.Id,
						"appId":      in.AppId,
						"roles":      in.Roles,
					},
					Index: &analysisclient.LogIndex{
						Name:     functionIndex,
						Id:       in.Id,
						Relation: true,
						Fields: &analysisclient.LogFields{
							"roles": role.Roles,
						},
					},
				})

				return errstate.Success
			}
		}

		return nil
	})
}

func (svc *bindingService) UnbindUserRole(ctx context.Context, in *external.BindingRoleRequest, out *commons.Status) error {
	return wrapper.ContextToAuthorize(ctx, out, func(auth *wrapper.WrapperUser) *commons.State {

		if len(in.AppId) > 0 && len(in.RoleId) > 0 && len(in.Id) > 0 {

			headers := &analysisclient.LogHeaders{
				TraceId:     auth.TraceId,
				ServiceName: constants.PermissionService,
				ModuleName:  "UnbindUserRole",
			}

			repo := svc.GetRepo()
			defer repo.Close()

			roleRepo := svc.GetRoleRepo()
			defer roleRepo.Close()

			role, err := roleRepo.FindRoleById(in.RoleId, in.AppId)

			if err != nil {
				return errstate.ErrSystem
			}

			if len(role.Id) == 0 {
				return errstate.ErrRequest
			}

			user, err := repo.FindRelationUserById(in.Id, in.AppId)
			if err != nil {
				return nil
			}

			if len(user.Roles) == 0 {
				return nil
			}

			var roles []string
			for _, v := range user.Roles {
				ok := true
				if v == in.RoleId {
					ok = false
				}
				if ok {
					roles = append(roles, v)
				}
			}

			user.Roles = roles

			err = repo.SetUserRole(in.Id, in.AppId, user)

			if err != nil {
				return nil
			}

			svc.log.Info(&analysisclient.LogContent{
				Headers: headers,
				Action:  "UnbindUserRole",
				Fields: &analysisclient.LogFields{
					"userId": in.Id,
					"appId":  in.AppId,
					"roleId": in.RoleId,
				},
				Index: &analysisclient.LogIndex{
					Name:     "users",
					Id:       in.Id,
					Relation: true,
					Fields: &analysisclient.LogFields{
						"roles": roles,
					},
				},
			})

		}

		return nil
	})
}

func (svc *bindingService) UnbindFunctionRole(ctx context.Context, in *external.BindingRoleRequest, out *commons.Status) error {
	return wrapper.ContextToAuthorize(ctx, out, func(auth *wrapper.WrapperUser) *commons.State {

		if len(in.AppId) > 0 && len(in.RoleId) > 0 && len(in.Id) > 0 {

			headers := &analysisclient.LogHeaders{
				TraceId:     auth.TraceId,
				ServiceName: constants.PermissionService,
				ModuleName:  "UnbindFunctionRole",
			}

			repo := svc.GetRepo()
			defer repo.Close()

			roleRepo := svc.GetRoleRepo()
			defer roleRepo.Close()

			role, err := roleRepo.FindRoleById(in.RoleId, in.AppId)

			if err != nil {
				return errstate.ErrSystem
			}

			if len(role.Id) == 0 {
				return errstate.ErrRequest
			}

			function, err := repo.FindRelationFunctionById(in.Id, in.AppId)
			if err != nil {
				return nil
			}

			if len(function.Roles) == 0 {
				return nil
			}

			var roles []string
			for _, v := range function.Roles {
				ok := true
				if v == in.RoleId {
					ok = false
				}
				if ok {
					roles = append(roles, v)
				}
			}

			function.Roles = roles

			err = repo.SetFunctionRole(in.Id, in.AppId, function)

			if err != nil {
				return nil
			}

			svc.log.Info(&analysisclient.LogContent{
				Headers: headers,
				Action:  "UnbindFunctionRole",
				Fields: &analysisclient.LogFields{
					"functionId": in.Id,
					"appId":      in.AppId,
					"roleId":     in.RoleId,
				},
				Index: &analysisclient.LogIndex{
					Name:     functionIndex,
					Id:       in.Id,
					Relation: true,
					Fields: &analysisclient.LogFields{
						"roles": roles,
					},
				},
			})
		}

		return nil
	})
}

func NewBindingService(client *indexutils.Client, session *mgo.Session,
	innerUserService xbasissvc_internal_user.UserService, log analysisclient.LogClient) external.BindingHandler {
	return &bindingService{Client: client, session: session, innerUserService: innerUserService, log: log}
}
