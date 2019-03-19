package permission_handers

import (
	"context"
	"github.com/garyburd/redigo/redis"
	"gopkg.in/mgo.v2"
	"konekko.me/gosion/commons/dto"
	"konekko.me/gosion/commons/errstate"
	"konekko.me/gosion/commons/generator"
	"konekko.me/gosion/commons/wrapper"
	"konekko.me/gosion/permission/pb"
	"konekko.me/gosion/permission/repositories"
)

type groupService struct {
	pool    *redis.Pool
	session *mgo.Session
}

func (svc *groupService) GetRepo() permission_repositories.GroupRepo {
	return permission_repositories.GroupRepo{Session: svc.session.Clone(), ID: gs_commons_generator.ID()}
}

func (svc *groupService) Add(ctx context.Context, in *gs_service_permission.SimpleGroup, out *gs_commons_dto.Status) error {
	return gs_commons_wrapper.ContextToAuthorize(ctx, out, func(auth *gs_commons_dto.Authorize) *gs_commons_dto.State {

		repo := svc.GetRepo()
		defer repo.Close()

		_, err := repo.FindByName(in.AppId, in.Name)
		if err != nil && err == mgo.ErrNotFound {

			err = repo.Save(in.AppId, auth.UserId, in.Name)

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

func (svc *groupService) LinkTo(ctx context.Context, in *gs_service_permission.SimpleGroup, out *gs_commons_dto.Status) error {
	return gs_commons_wrapper.ContextToAuthorize(ctx, out, func(auth *gs_commons_dto.Authorize) *gs_commons_dto.State {
		return nil
	})
}

func (svc *groupService) Unlink(ctx context.Context, in *gs_service_permission.SimpleGroup, out *gs_commons_dto.Status) error {
	return gs_commons_wrapper.ContextToAuthorize(ctx, out, func(auth *gs_commons_dto.Authorize) *gs_commons_dto.State {
		return nil
	})
}

func (svc *groupService) Rename(ctx context.Context, in *gs_service_permission.SimpleGroup, out *gs_commons_dto.Status) error {
	return gs_commons_wrapper.ContextToAuthorize(ctx, out, func(auth *gs_commons_dto.Authorize) *gs_commons_dto.State {
		return nil
	})
}

func (svc *groupService) AddUserToGroup(ctx context.Context, in *gs_service_permission.SimpleUserNode, out *gs_commons_dto.Status) error {
	return gs_commons_wrapper.ContextToAuthorize(ctx, out, func(auth *gs_commons_dto.Authorize) *gs_commons_dto.State {
		return nil
	})
}

func (svc *groupService) MoveUserToGroup(ctx context.Context, in *gs_service_permission.SimpleUserNode, out *gs_commons_dto.Status) error {
	return gs_commons_wrapper.ContextToAuthorize(ctx, out, func(auth *gs_commons_dto.Authorize) *gs_commons_dto.State {
		return nil
	})
}

func (svc *groupService) Remove(ctx context.Context, in *gs_service_permission.SimpleGroup, out *gs_commons_dto.Status) error {
	return gs_commons_wrapper.ContextToAuthorize(ctx, out, func(auth *gs_commons_dto.Authorize) *gs_commons_dto.State {
		return nil
	})
}

func NewGroupService(pool *redis.Pool, session *mgo.Session) gs_service_permission.GroupStructureHandler {
	return &groupService{pool: pool, session: session}
}
