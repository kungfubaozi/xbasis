package safety_handers

import (
	"context"
	"github.com/garyburd/redigo/redis"
	"gopkg.in/mgo.v2"
	"konekko.me/gosion/commons/dto"
	"konekko.me/gosion/commons/errstate"
	"konekko.me/gosion/commons/wrapper"
	"konekko.me/gosion/safety/pb"
	"konekko.me/gosion/safety/repositories"
)

type blacklistService struct {
	session *mgo.Session
	pool    *redis.Pool
}

func (svc *blacklistService) GetRepo() safety_repositories.BlacklistRepo {
	return safety_repositories.BlacklistRepo{Session: svc.session.Clone(), Conn: svc.pool.Get()}
}

func (svc *blacklistService) AddBlacklist(ctx context.Context, in *gs_service_safety.AddRequest, out *gs_commons_dto.Status) error {
	return gs_commons_wrapper.ContextToAuthorize(ctx, out, func(auth *gs_commons_dto.Authorize) *gs_commons_dto.State {

		repo := svc.GetRepo()
		defer repo.Close()

		if repo.CacheExists(in.Type, in.Content) {

			err := repo.Save(in.Type, in.Content, auth.UserId)
			if err != nil {

				return errstate.ErrRequest
			}

			return errstate.Success
		}

		return errstate.ErrBlacklistAlreadyExists
	})
}

func (svc *blacklistService) RemoveBlacklist(ctx context.Context, in *gs_service_safety.RemoveBlacklistRequest, out *gs_commons_dto.Status) error {
	return gs_commons_wrapper.ContextToAuthorize(ctx, out, func(auth *gs_commons_dto.Authorize) *gs_commons_dto.State {

		repo := svc.GetRepo()
		defer repo.Close()

		err := repo.Remove(in.Id)
		if err == nil {
			return errstate.Success
		}

		return nil
	})
}

func NewBlakclistService(session *mgo.Session, pool *redis.Pool) gs_service_safety.BlacklistHandler {
	return &blacklistService{}
}
