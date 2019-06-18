package safetyhanders

import (
	"context"
	"fmt"
	"github.com/garyburd/redigo/redis"
	"gopkg.in/mgo.v2"
	"konekko.me/gosion/analysis/client"
	"konekko.me/gosion/commons/dto"
	"konekko.me/gosion/commons/errstate"
	"konekko.me/gosion/commons/wrapper"
	external "konekko.me/gosion/safety/pb"
	"konekko.me/gosion/user/pb/inner"
	"time"
)

type lockingService struct {
	session     *mgo.Session
	log         analysisclient.LogClient
	pool        *redis.Pool
	userService gosionsvc_internal_user.UserService
}

func (svc *lockingService) Search(context.Context, *external.SearchRequest, *external.SearchResponse) error {
	panic("implement me")
}

func (svc *lockingService) GetRepo() *lockingRepo {
	return &lockingRepo{session: svc.session.Clone(), conn: svc.pool.Get()}
}

func (svc *lockingService) Lock(ctx context.Context, in *external.LockRequest, out *gs_commons_dto.Status) error {
	return gs_commons_wrapper.ContextToAuthorize(ctx, out, func(auth *gs_commons_wrapper.WrapperUser) *gs_commons_dto.State {
		fmt.Println("lock user", in.UserId)
		if len(in.UserId) < 9 {
			return nil
		}
		s, err := svc.userService.IsExists(ctx, &gosionsvc_internal_user.ExistsRequest{
			UserId: in.UserId,
		})
		if err != nil {
			return errstate.ErrRequest
		}
		if !s.State.Ok {
			return s.State
		}

		repo := svc.GetRepo()
		defer repo.Close()

		exists, err := repo.IsExists(in.UserId)
		if err != nil {
			return errstate.ErrRequest
		}

		if exists {
			return errstate.ErrAlreadyLocking
		}

		t := time.Now().UnixNano()
		et := t + in.Time*1e6

		l := &lockingUser{
			UserId:       in.UserId,
			CreateUserId: auth.User,
			ExpiredTime:  et,
			CreateAt:     t,
		}

		err = repo.Add(l)
		if err != nil {
			return errstate.ErrRequest
		}

		return nil
	})
}

func (svc *lockingService) Unlock(ctx context.Context, in *external.UnlockRequest, out *gs_commons_dto.Status) error {
	return gs_commons_wrapper.ContextToAuthorize(ctx, out, func(auth *gs_commons_wrapper.WrapperUser) *gs_commons_dto.State {
		if len(in.UserId) < 9 {
			return nil
		}

		repo := svc.GetRepo()
		defer repo.Close()

		exists, err := repo.IsExists(in.UserId)
		if err != nil {
			return errstate.ErrRequest
		}

		if exists {
			err = repo.Remove(in.UserId)
			if err != nil {
				return errstate.ErrRequest
			}
			return errstate.Success
		}

		return errstate.ErrNotFound
	})
}

func NewLockingService() external.LockingHandler {
	return &lockingService{}
}
