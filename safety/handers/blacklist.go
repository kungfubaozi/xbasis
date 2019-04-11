package safetyhanders

import (
	"context"
	"github.com/garyburd/redigo/redis"
	"gopkg.in/mgo.v2"
	"konekko.me/gosion/commons/constants"
	"konekko.me/gosion/commons/dto"
	"konekko.me/gosion/commons/errstate"
	"konekko.me/gosion/commons/wrapper"
	"konekko.me/gosion/safety/pb"
)

type blacklistService struct {
	session *mgo.Session
	pool    *redis.Pool
}

func (svc *blacklistService) GetRepo() blacklistRepo {
	return blacklistRepo{session: svc.session.Clone(), conn: svc.pool.Get()}
}

func (svc *blacklistService) Check(ctx context.Context, in *gs_service_safety.CheckRequest, out *gs_commons_dto.Status) error {
	return gs_commons_wrapper.ContextToAuthorize(ctx, out, func(auth *gs_commons_wrapper.WrapperUser) *gs_commons_dto.State {

		if len(in.Content) > 0 {
			return errstate.Success
		}

		if in.Type == gs_commons_constants.BlacklistOfIP || in.Type == gs_commons_constants.BlacklistOfUserDevice {

			repo := svc.GetRepo()
			defer repo.Close()

			if !repo.CacheExists(in.Type, in.Content) {
				return errstate.Success
			}

		}
		return nil
	})
}

func (svc *blacklistService) AddBlacklist(ctx context.Context, in *gs_service_safety.AddRequest, out *gs_commons_dto.Status) error {
	return gs_commons_wrapper.ContextToAuthorize(ctx, out, func(auth *gs_commons_wrapper.WrapperUser) *gs_commons_dto.State {

		repo := svc.GetRepo()
		defer repo.Close()

		if repo.CacheExists(in.Type, in.Content) {

			err := repo.Save(in.Type, in.Content, auth.User)
			if err != nil {

				return errstate.ErrRequest
			}

			return errstate.Success
		}

		return errstate.ErrBlacklistAlreadyExists
	})
}

func (svc *blacklistService) RemoveBlacklist(ctx context.Context, in *gs_service_safety.RemoveBlacklistRequest, out *gs_commons_dto.Status) error {
	return gs_commons_wrapper.ContextToAuthorize(ctx, out, func(auth *gs_commons_wrapper.WrapperUser) *gs_commons_dto.State {

		repo := svc.GetRepo()
		defer repo.Close()

		err := repo.Remove(in.Id)
		if err == nil {
			return errstate.Success
		}

		return nil
	})
}

func NewBlacklistService(session *mgo.Session, pool *redis.Pool) gs_service_safety.BlacklistHandler {
	return &blacklistService{session: session, pool: pool}
}
