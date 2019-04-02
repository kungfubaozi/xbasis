package permissionhandlers

import (
	"context"
	"github.com/garyburd/redigo/redis"
	"gopkg.in/mgo.v2"
	"konekko.me/gosion/commons/dto"
	"konekko.me/gosion/commons/errstate"
	"konekko.me/gosion/commons/generator"
	"konekko.me/gosion/commons/wrapper"
	"konekko.me/gosion/permission/pb"
	"konekko.me/gosion/user/pb/ext"
)

type bindingService struct {
	pool           *redis.Pool
	session        *mgo.Session
	extUserService gs_ext_service_user.UserService
}

func (svc *bindingService) GetRepo() *roleRepo {
	return &roleRepo{conn: svc.pool.Get(), session: svc.session.Clone(), id: gs_commons_generator.NewIDG()}
}

//verify roles effectiveness
func (svc *bindingService) verifyRoleEffectiveness() {

}

func (svc *bindingService) UserRole(ctx context.Context, in *gs_service_permission.BindingRoleRequest, out *gs_commons_dto.Status) error {
	return gs_commons_wrapper.ContextToAuthorize(ctx, out, func(auth *gs_commons_wrapper.WrapperUser) *gs_commons_dto.State {
		if len(in.StructureId) > 0 && len(in.Id) > 0 && len(in.RoleId) > 0 {
			//check roles
			repo := svc.GetRepo()
			defer repo.Close()

			//check structure exists
			if isStructureExists(repo.session, in.StructureId) == 0 {
				return errstate.ErrInvalidStructure
			}

			urs := make(map[string][]string)

			//check for user roles
			for _, v := range in.Id {
				ums, err := repo.GetUserRoleMembers(in.StructureId, v)

				add := func(userId string) {
					var uml []string
					if ums != nil {
						for _, v := range ums {
							s := string(v.([]byte))
							uml = append(uml, s)
						}
					}
					if uml == nil {
						uml = make([]string, 0)
					}
					urs[userId] = uml
				}

				if err != nil && err == redis.ErrNil {
					err = nil
					//check user exists
					//....
					s, err := svc.extUserService.IsExists(ctx, &gs_ext_service_user.ExistsRequest{
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

			for user, rs := range urs {
				var nars []string //need add roles

			}

		}

		return nil
	})
}

func (svc *bindingService) FunctionRole(ctx context.Context, in *gs_service_permission.BindingRoleRequest, out *gs_commons_dto.Status) error {
	return gs_commons_wrapper.ContextToAuthorize(ctx, out, func(auth *gs_commons_wrapper.WrapperUser) *gs_commons_dto.State {
		return nil
	})
}

func (svc *bindingService) UnbindUserRole(ctx context.Context, in *gs_service_permission.BindingRoleRequest, out *gs_commons_dto.Status) error {
	return gs_commons_wrapper.ContextToAuthorize(ctx, out, func(auth *gs_commons_wrapper.WrapperUser) *gs_commons_dto.State {
		return nil
	})
}

func (svc *bindingService) UnbindFunctionRole(ctx context.Context, in *gs_service_permission.BindingRoleRequest, out *gs_commons_dto.Status) error {
	return gs_commons_wrapper.ContextToAuthorize(ctx, out, func(auth *gs_commons_wrapper.WrapperUser) *gs_commons_dto.State {
		return nil
	})
}

func NewBindingService(pool *redis.Pool, session *mgo.Session) gs_service_permission.BindingHandler {
	return &bindingService{pool: pool, session: session}
}
