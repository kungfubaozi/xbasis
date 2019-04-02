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
)

type roleService struct {
	session *mgo.Session
	pool    *redis.Pool
}

func (svc *roleService) GetRepo() *roleRepo {
	return &roleRepo{Session: svc.session.Clone(),
		id: gs_commons_generator.NewIDG(), Conn: svc.pool.Get()}
}

//add new role if not exists
func (svc *roleService) Add(ctx context.Context, in *gs_service_permission.RoleRequest, out *gs_commons_dto.Status) error {
	return gs_commons_wrapper.ContextToAuthorize(ctx, out, func(auth *gs_commons_wrapper.WrapperUser) *gs_commons_dto.State {
		repo := svc.GetRepo()
		defer repo.Close()

		_, err := repo.FindByName(in.Name, in.AppId)
		if err != nil && err == mgo.ErrNotFound {
			err = repo.Save(in.Name, auth.User, in.AppId)
			if err != nil {
				return nil
			}
			return errstate.Success
		}

		if err == nil {
			return errstate.ErrRoleAlreadyExists
		}

		return nil
	})
}

//remove role
func (svc *roleService) Remove(ctx context.Context, in *gs_service_permission.RoleRequest, out *gs_commons_dto.Status) error {
	return gs_commons_wrapper.ContextToAuthorize(ctx, out, func(auth *gs_commons_wrapper.WrapperUser) *gs_commons_dto.State {
		return nil
	})
}

func (svc *roleService) Rename(ctx context.Context, in *gs_service_permission.RoleRequest, out *gs_commons_dto.Status) error {
	return gs_commons_wrapper.ContextToAuthorize(ctx, out, func(auth *gs_commons_wrapper.WrapperUser) *gs_commons_dto.State {
		return nil
	})
}

func NewRoleService(session *mgo.Session, pool *redis.Pool) gs_service_permission.RoleHandler {
	return &roleService{session: session, pool: pool}
}
