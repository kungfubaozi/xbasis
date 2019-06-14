package safetyhanders

import (
	"context"
	"github.com/garyburd/redigo/redis"
	"gopkg.in/mgo.v2"
	"konekko.me/gosion/commons/constants"
	"konekko.me/gosion/commons/dto"
	"konekko.me/gosion/commons/errstate"
	"konekko.me/gosion/commons/wrapper"
	inner "konekko.me/gosion/safety/pb/inner"
	"sync"
)

type securityService struct {
	session *mgo.Session
	pool    *redis.Pool
}

func (svc *securityService) GetRepo() *lockingRepo {
	return &lockingRepo{session: svc.session.Clone(), conn: svc.pool.Get()}
}

//locking
//different places
func (svc *securityService) Get(ctx context.Context, in *inner.GetRequest, out *inner.GetResponse) error {
	return gs_commons_wrapper.ContextToAuthorize(ctx, out, func(auth *gs_commons_wrapper.WrapperUser) *gs_commons_dto.State {

		if len(in.UserId) == 0 {
			return nil
		}
		state := errstate.Success
		sc := gs_commons_constants.UserStateOfClear
		resp := func(s *gs_commons_dto.State) {
			if state.Ok {
				state = s
			}
		}
		var wg sync.WaitGroup
		wg.Add(1)

		repo := svc.GetRepo()
		defer repo.Close()

		//check user locking
		go func() {
			defer wg.Done()
			e, err := repo.IsExists(in.UserId)
			if err != nil {
				resp(errstate.ErrRequest)
				return
			}
			if e {
				sc = gs_commons_constants.UserStateOfLocking
				return
			}
		}()

		wg.Wait()

		out.Current = int64(sc)
		out.State = state
		return nil
	})
}

func NewSecurityService(session *mgo.Session, pool *redis.Pool) inner.SecurityHandler {
	return &securityService{session: session, pool: pool}
}
