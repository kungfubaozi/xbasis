package safetyhanders

import (
	"context"
	"fmt"
	"github.com/garyburd/redigo/redis"
	"gopkg.in/mgo.v2"
	"konekko.me/gosion/commons/constants"
	"konekko.me/gosion/commons/dto"
	"konekko.me/gosion/commons/errstate"
	"konekko.me/gosion/commons/wrapper"
	"konekko.me/gosion/safety/pb/ext"
)

type securityService struct {
	session *mgo.Session
	pool    *redis.Pool
}

func (svc *securityService) GetRepo() *securityRepo {
	return &securityRepo{session: svc.session.Clone(), conn: svc.pool.Get()}
}

//locking
//different places
func (svc *securityService) Get(ctx context.Context, in *gs_ext_service_safety.GetRequest, out *gs_ext_service_safety.GetResponse) error {
	return gs_commons_wrapper.ContextToAuthorize(ctx, out, func(auth *gs_commons_wrapper.WrapperUser) *gs_commons_dto.State {

		fmt.Println("check user security")

		if len(in.UserId) == 0 {
			return nil
		}
		out.State = errstate.Success
		//resp := func(s *gs_commons_dto.State) {
		//		//	if state.Ok {
		//		//		state = s
		//		//	}
		//		//}
		//var wg sync.WaitGroup
		//wg.Add(3)
		//
		////check user blacklist
		//go func() {
		//	defer wg.Done()
		//}()
		//
		////check user frozen
		//go func() {
		//	defer wg.Done()
		//}()
		//
		////check user locking
		//go func() {
		//	defer wg.Done()
		//}()

		out.Current = gs_commons_constants.UserStateOfClear
		return nil
	})
}

func NewSecurityService(session *mgo.Session) gs_ext_service_safety.SecurityHandler {
	return &securityService{session: session}
}
