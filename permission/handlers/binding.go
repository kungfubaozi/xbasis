package permissionhandlers

import (
	"context"
	"github.com/garyburd/redigo/redis"
	"gopkg.in/mgo.v2"
	"konekko.me/gosion/analysis/client"
	"konekko.me/gosion/commons/dto"
	"konekko.me/gosion/commons/errstate"
	"konekko.me/gosion/commons/generator"
	"konekko.me/gosion/commons/wrapper"
	external "konekko.me/gosion/permission/pb"
	"konekko.me/gosion/user/pb/inner"
)

type bindingService struct {
	pool             *redis.Pool
	session          *mgo.Session
	innerUserService gosionsvc_internal_user.UserService
	log              analysisclient.LogClient
}

func (svc *bindingService) GetRepo() *bindingRepo {
	return &bindingRepo{conn: svc.pool.Get(), session: svc.session.Clone(), id: gs_commons_generator.NewIDG()}
}

//verify roles effectiveness
func (svc *bindingService) verifyRoleEffectiveness() {

}

func (svc *bindingService) UserRole(ctx context.Context, in *external.BindingRoleRequest, out *gs_commons_dto.Status) error {
	return gs_commons_wrapper.ContextToAuthorize(ctx, out, func(auth *gs_commons_wrapper.WrapperUser) *gs_commons_dto.State {
		if len(in.StructureId) > 0 && len(in.Id) > 0 && len(in.RoleId) > 0 {
			//check roles
			repo := svc.GetRepo()
			defer repo.Close()

			//check structure exists
			if isStructureExists(repo.session, in.StructureId) == 0 {
				return errstate.ErrInvalidStructure
			}

			//all roles corresponding to valid users
			urs := make(map[string]map[string]bool)

			//check for user roles
			for _, v := range in.Id {
				ums, err := repo.GetUserRoleMembers(in.StructureId, v)

				add := func(userId string) {
					uml := make(map[string]bool)
					if ums != nil {
						for _, v := range ums {
							s := string(v.([]byte))
							uml[s] = true
						}
					}
					urs[userId] = uml
				}

				if err != nil && err == redis.ErrNil {
					err = nil
					//check user exists
					//....
					s, err := svc.innerUserService.IsExists(ctx, &gosionsvc_internal_user.ExistsRequest{
						UserId: v,
					})
					if err != nil {
						return nil
					}
					if !s.State.Ok {
						return nil
					}
				}
				if err != nil {
					return nil
				}
				add(v)
			}

			//check for role exists
			for _, v := range in.RoleId {
				ok, err := repo.Exists(in.StructureId, v)
				if err != nil {
					return nil
				}
				if !ok {
					return nil
				}
			}

			for userId, userRoles := range urs {
				var roles []string
				if len(userRoles) > 0 {
					for _, v := range in.RoleId {
						if !userRoles[v] {
							roles = append(roles, v)
						}
					}
				} else {
					roles = in.RoleId
				}
				if roles != nil && len(roles) > 0 {
					err := repo.SetUserRoleMembersInCache(userId, in.StructureId, roles)
					if err == nil {
						//add  to database
						return errstate.Success
					}
				}
			}

		}

		return nil
	})
}

func (svc *bindingService) FunctionRole(ctx context.Context, in *external.BindingRoleRequest, out *gs_commons_dto.Status) error {
	return gs_commons_wrapper.ContextToAuthorize(ctx, out, func(auth *gs_commons_wrapper.WrapperUser) *gs_commons_dto.State {
		if len(in.StructureId) > 0 && len(in.Id) > 0 && len(in.RoleId) > 0 {

			repo := svc.GetRepo()
			defer repo.Close()

			//check structure exists
			if isStructureExists(repo.session, in.StructureId) == 0 {
				return errstate.ErrInvalidStructure
			}
			//
			//for _, v := range in.Id {
			//	f, err := repo.GetFunctionRoleMembers(in.StructureId, v)
			//	if err != nil && err == redis.ErrNil {
			//		//check function exists
			//
			//	}
			//}

		}

		return nil
	})
}

func (svc *bindingService) UnbindUserRole(ctx context.Context, in *external.BindingRoleRequest, out *gs_commons_dto.Status) error {
	return gs_commons_wrapper.ContextToAuthorize(ctx, out, func(auth *gs_commons_wrapper.WrapperUser) *gs_commons_dto.State {

		if len(in.StructureId) > 0 && len(in.RoleId) > 0 && len(in.Id) > 0 {

			repo := svc.GetRepo()
			defer repo.Close()

			//check structure exists
			if isStructureExists(repo.session, in.StructureId) == 0 {
				return errstate.ErrInvalidStructure
			}

		}

		return nil
	})
}

func (svc *bindingService) UnbindFunctionRole(ctx context.Context, in *external.BindingRoleRequest, out *gs_commons_dto.Status) error {
	return gs_commons_wrapper.ContextToAuthorize(ctx, out, func(auth *gs_commons_wrapper.WrapperUser) *gs_commons_dto.State {

		if len(in.StructureId) > 0 && len(in.RoleId) > 0 && len(in.Id) > 0 {

			repo := svc.GetRepo()
			defer repo.Close()

			//check structure exists
			if isStructureExists(repo.session, in.StructureId) == 0 {
				return errstate.ErrInvalidStructure
			}

		}

		return nil
	})
}

func NewBindingService(pool *redis.Pool, session *mgo.Session, innerUserService gosionsvc_internal_user.UserService, log analysisclient.LogClient) external.BindingHandler {
	return &bindingService{pool: pool, session: session, innerUserService: innerUserService, log: log}
}
