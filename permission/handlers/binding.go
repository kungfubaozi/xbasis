package permission_handlers

import (
	"context"
	"github.com/garyburd/redigo/redis"
	"gopkg.in/mgo.v2"
	"konekko.me/gosion/commons/dto"
	"konekko.me/gosion/commons/generator"
	"konekko.me/gosion/commons/wrapper"
	"konekko.me/gosion/permission/pb"
	"konekko.me/gosion/permission/repositories"
)

type bindingService struct {
	pool    *redis.Pool
	session *mgo.Session
}

func (svc *bindingService) GetRepo() *permission_repositories.RoleRepo {
	return &permission_repositories.RoleRepo{Conn: svc.pool.Get(), Session: svc.session.Clone(), ID: gs_commons_generator.NewIDG()}
}

//verify roles effectiveness
func (svc *bindingService) verifyRoleEffectiveness() {

}

func (svc *bindingService) UserRole(ctx context.Context, in *gs_service_permission.BindingRoleRequest, out *gs_commons_dto.Status) error {
	return gs_commons_wrapper.ContextToAuthorize(ctx, out, func(auth *gs_commons_wrapper.WrapperUser) *gs_commons_dto.State {
		if len(in.AppId) == 0 && len(in.Id) == 0 && len(in.RoleId) == 0 {
			//check roles
			repo := svc.GetRepo()
			defer repo.Close()

			//check for user roles
			urs := make(map[string][]string)
			for _, v := range in.Id {
				ums, err := repo.GetUserRoleMembers(in.AppId, v)

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
				}
				if err != nil {
					return nil
				}
				add(v)
			}

			//check for role exists
			for _, v := range in.RoleId {
				ok, err := repo.Exists(in.AppId, v)
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
