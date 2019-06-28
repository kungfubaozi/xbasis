package safetyhanders

import (
	"context"
	"github.com/garyburd/redigo/redis"
	"gopkg.in/mgo.v2"
	constants "konekko.me/xbasis/commons/constants"
	commons "konekko.me/xbasis/commons/dto"
	"konekko.me/xbasis/commons/errstate"
	"konekko.me/xbasis/commons/wrapper"
	inner "konekko.me/xbasis/safety/pb/inner"
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
	return xbasiswrapper.ContextToAuthorize(ctx, out, func(auth *xbasiswrapper.WrapperUser) *commons.State {

		if len(in.UserId) == 0 {
			return nil
		}
		state := errstate.Success
		sc := constants.UserStateOfClear
		resp := func(s *commons.State) {
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
				sc = constants.UserStateOfLocking
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
