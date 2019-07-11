package safetyhanders

import (
	"context"
	"github.com/garyburd/redigo/redis"
	"gopkg.in/mgo.v2"
	"konekko.me/xbasis/analysis/client"
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
	log     analysisclient.LogClient
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

		headers := &analysisclient.LogHeaders{
			TraceId:     auth.TraceId,
			ServiceName: constants.InternalSafetyService,
			ModuleName:  "UserSecurity",
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
				svc.log.Info(&analysisclient.LogContent{
					Headers: headers,
					Action:  "UserInLockList",
				})
				return
			}
			if e {
				sc = constants.UserStateOfLocking
				return
			}
		}()

		wg.Wait()

		svc.log.Info(&analysisclient.LogContent{
			Headers: headers,
			Action:  "UserSecurityCheck",
			Message: "Passed",
		})

		out.Current = int64(sc)
		out.State = state
		return nil
	})
}

func NewSecurityService(session *mgo.Session, pool *redis.Pool, log analysisclient.LogClient) inner.SecurityHandler {
	return &securityService{session: session, pool: pool, log: log}
}
