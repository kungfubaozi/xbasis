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

type groupService struct {
	pool    *redis.Pool
	session *mgo.Session
}

func (svc *groupService) GetRepo() *groupRepo {
	return &groupRepo{session: svc.session.Clone(), id: gs_commons_generator.NewIDG()}
}

func (svc *groupService) Add(ctx context.Context, in *gs_service_permission.SimpleGroup, out *gs_commons_dto.Status) error {
	return gs_commons_wrapper.ContextToAuthorize(ctx, out, func(auth *gs_commons_wrapper.WrapperUser) *gs_commons_dto.State {

		repo := svc.GetRepo()
		defer repo.Close()

		if isStructureExists(repo.session, in.StructureId) == 0 {
			return errstate.ErrInvalidStructure
		}

		_, err := repo.FindByName(in.StructureId, in.Name)
		if err != nil && err == mgo.ErrNotFound {

			err = repo.Save(in.StructureId, auth.User, in.Name)

			if err != nil {
				return errstate.ErrRequest
			}

			return errstate.Success
		}

		if err == nil {
			return errstate.ErrGroupAlreadyExists
		}

		return nil
	})
}

//You can link to other application groups, or to this application group.
func (svc *groupService) LinkTo(ctx context.Context, in *gs_service_permission.SimpleGroup, out *gs_commons_dto.Status) error {
	return gs_commons_wrapper.ContextToAuthorize(ctx, out, func(auth *gs_commons_wrapper.WrapperUser) *gs_commons_dto.State {
		return nil
	})
}

func (svc *groupService) Unlink(ctx context.Context, in *gs_service_permission.SimpleGroup, out *gs_commons_dto.Status) error {
	return gs_commons_wrapper.ContextToAuthorize(ctx, out, func(auth *gs_commons_wrapper.WrapperUser) *gs_commons_dto.State {
		return nil
	})
}

func (svc *groupService) Rename(ctx context.Context, in *gs_service_permission.SimpleGroup, out *gs_commons_dto.Status) error {
	return gs_commons_wrapper.ContextToAuthorize(ctx, out, func(auth *gs_commons_wrapper.WrapperUser) *gs_commons_dto.State {
		return nil
	})
}

//User cannot be in the same group under the same application
func (svc *groupService) AddUser(ctx context.Context, in *gs_service_permission.SimpleUserNode, out *gs_commons_dto.Status) error {
	return gs_commons_wrapper.ContextToAuthorize(ctx, out, func(auth *gs_commons_wrapper.WrapperUser) *gs_commons_dto.State {
		return nil
	})
}

func (svc *groupService) MoveUser(ctx context.Context, in *gs_service_permission.SimpleUserNode, out *gs_commons_dto.Status) error {
	return gs_commons_wrapper.ContextToAuthorize(ctx, out, func(auth *gs_commons_wrapper.WrapperUser) *gs_commons_dto.State {
		return nil
	})
}

//If there are users under the current group, they cannot be deleted
func (svc *groupService) Remove(ctx context.Context, in *gs_service_permission.SimpleGroup, out *gs_commons_dto.Status) error {
	return gs_commons_wrapper.ContextToAuthorize(ctx, out, func(auth *gs_commons_wrapper.WrapperUser) *gs_commons_dto.State {
		return nil
	})
}

func NewGroupService(pool *redis.Pool, session *mgo.Session) gs_service_permission.UserGroupHandler {
	return &groupService{pool: pool, session: session}
}
