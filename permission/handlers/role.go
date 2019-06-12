package permissionhandlers

import (
	"context"
	"fmt"
	"github.com/garyburd/redigo/redis"
	"gopkg.in/mgo.v2"
	"konekko.me/gosion/commons/dto"
	"konekko.me/gosion/commons/errstate"
	"konekko.me/gosion/commons/generator"
	"konekko.me/gosion/commons/wrapper"
	external "konekko.me/gosion/permission/pb"
)

type roleService struct {
	session *mgo.Session
	pool    *redis.Pool
}

func (svc *roleService) EffectUserSize(context.Context, *external.EffectUserSizeRequest, *external.EffectUserSizeResponse) error {
	panic("implement me")
}

func (svc *roleService) GetStructureRoles(ctx context.Context, in *external.GetStructureRolesRequest, out *external.GetRoleResponse) error {
	return gs_commons_wrapper.ContextToAuthorize(ctx, out, func(auth *gs_commons_wrapper.WrapperUser) *gs_commons_dto.State {
		if len(in.StructureId) == 0 {
			return nil
		}
		repo := svc.GetRepo()
		defer repo.Close()

		roles, err := repo.FindRolesByStructure(in.StructureId, in.Page, in.Size)
		if err != nil {
			return nil
		}

		fmt.Println("data", len(roles))
		var rs []*external.SimpleRoleInfo
		for _, v := range roles {
			rs = append(rs, &external.SimpleRoleInfo{
				Id:       v.Id,
				Name:     v.Name,
				CreateAt: v.CreateAt,
			})
		}

		out.Data = rs
		return errstate.Success
	})
}

func (svc *roleService) GetRole(ctx context.Context, in *external.GetRoleRequest, out *external.GetRoleResponse) error {
	return gs_commons_wrapper.ContextToAuthorize(ctx, out, func(auth *gs_commons_wrapper.WrapperUser) *gs_commons_dto.State {
		return nil
	})
}

func (svc *roleService) GetRepo() *roleRepo {
	return &roleRepo{session: svc.session.Clone(),
		id: gs_commons_generator.NewIDG(), conn: svc.pool.Get()}
}

//add new role if not exists
func (svc *roleService) Add(ctx context.Context, in *external.RoleRequest, out *gs_commons_dto.Status) error {
	return gs_commons_wrapper.ContextToAuthorize(ctx, out, func(auth *gs_commons_wrapper.WrapperUser) *gs_commons_dto.State {
		repo := svc.GetRepo()
		defer repo.Close()

		//check
		if isStructureExists(repo.session, in.StructureId) == 0 {
			return errstate.ErrInvalidStructure
		}

		_, err := repo.FindByName(in.Name, in.StructureId)
		if err != nil && err == mgo.ErrNotFound {
			err = repo.Save(in.Name, auth.User, in.StructureId)
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
//需要删除所有关联的角色对象包括(gs-user-roles-relation)
func (svc *roleService) Remove(ctx context.Context, in *external.RoleRequest, out *gs_commons_dto.Status) error {
	return gs_commons_wrapper.ContextToAuthorize(ctx, out, func(auth *gs_commons_wrapper.WrapperUser) *gs_commons_dto.State {

		return nil
	})
}

func (svc *roleService) Rename(ctx context.Context, in *external.RoleRequest, out *gs_commons_dto.Status) error {
	return gs_commons_wrapper.ContextToAuthorize(ctx, out, func(auth *gs_commons_wrapper.WrapperUser) *gs_commons_dto.State {
		return nil
	})
}

func NewRoleService(session *mgo.Session, pool *redis.Pool) external.RoleHandler {
	return &roleService{session: session, pool: pool}
}
