package permissionhandlers

import (
	"context"
	"fmt"
	"github.com/olivere/elastic"
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
			fmt.Println("err", err)
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

			for _, v := range in.Roles {
				role, err := roleRepo.FindRoleById(v, in.AppId)
				if err != nil {
					return errstate.ErrSystem
				}

				if len(role.Id) == 0 {
					return errstate.ErrRequest
				}

				//不在同一个应用内
				if role.AppId != in.AppId {
					return errstate.ErrRequest
				}
			}

			//先查找当前功能是否已经绑定角色
			f, err := repo.FindRelationFunctionById(in.Id, in.AppId)
			if err != nil {
				return errstate.ErrRequest
			}

			if in.AppId != f.AppId {
				return errstate.ErrRequest
			}

			for _, v := range f.Roles {
				for _, v1 := range in.Roles {
					if v == v1 {
						return errstate.ErrFunctionAlreadyBindRole
					}
				}
			}

			//更新数据库
			err = repo.UpdateFunctionRole(in.Id, in.AppId, in.Roles)
			if err == nil {

				svc.log.Info(&analysisclient.LogContent{
					Headers: headers,
					Action:  "UpdateUserRole",
					Fields: &analysisclient.LogFields{
						"userId": in.Id,
						"appId":  in.AppId,
						"roles":  in.Roles,
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

			//所有包含此角色的索引全都删除
			query := elastic.NewBoolQuery()
			query.Must(elastic.NewMatchPhraseQuery("userId", in.Id))
			query.Must(elastic.NewMatchPhraseQuery("appId", in.AppId))
			query.Must(elastic.NewMatchQuery("roles", in.RoleId))
			r, err := svc.GetElasticClient().DeleteByQuery(getFunctionAuthorizeIndex(in.Id)).Query(query).Type("_doc").Do(context.Background())
			if err != nil {
				return errstate.ErrRequest
			}

			if r.Total >= 0 {

				err = repo.RemoveRoleFromUserRelation(in.Id, in.RoleId)
				if err == nil {
					return errstate.Success
				}
			}

		}

		return nil
	})
}

func (svc *bindingService) UnbindFunctionRole(ctx context.Context, in *external.BindingRoleRequest, out *commons.Status) error {
	return wrapper.ContextToAuthorize(ctx, out, func(auth *wrapper.WrapperUser) *commons.State {

		if len(in.AppId) > 0 && len(in.RoleId) > 0 && len(in.Id) > 0 {

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

			//所有包含此角色的索引全都删除
			query := elastic.NewBoolQuery()
			query.Must(elastic.NewMatchPhraseQuery("functionId", in.Id))
			query.Must(elastic.NewMatchPhraseQuery("appId", in.AppId))
			query.Must(elastic.NewMatchQuery("roles", in.RoleId))
			r, err := svc.GetElasticClient().DeleteByQuery(getFunctionAuthorizeIndex(in.Id)).Query(query).Type("_doc").Do(context.Background())
			if err != nil {
				return errstate.ErrRequest
			}

			if r.Total >= 0 {

				err = repo.RemoveRoleFromFunctions(in.Id, in.AppId, in.RoleId)
				if err == nil {
					return errstate.Success
				}

			}
		}

		return nil
	})
}

func NewBindingService(client *indexutils.Client, session *mgo.Session,
	innerUserService xbasissvc_internal_user.UserService, log analysisclient.LogClient) external.BindingHandler {
	return &bindingService{Client: client, session: session, innerUserService: innerUserService, log: log}
}
