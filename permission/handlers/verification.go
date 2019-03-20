package permission_handers

import (
	"context"
	"github.com/garyburd/redigo/redis"
	"gopkg.in/mgo.v2"
	"konekko.me/gosion/application/pb"
	"konekko.me/gosion/commons/constants"
	"konekko.me/gosion/commons/dto"
	"konekko.me/gosion/commons/errstate"
	"konekko.me/gosion/commons/transport"
	"konekko.me/gosion/commons/wrapper"
	"konekko.me/gosion/permission/pb"
	"konekko.me/gosion/safety/pb"
	"sync"
)

type verificationService struct {
	pool               *redis.Pool
	session            *mgo.Session
	applicationService gs_service_application.ApplicationService
	blacklistService   gs_service_safety.BlacklistService
	functionService    gs_service_permission.FunctionService
}

//application verify
//ip, userDevice blacklist verify
//api exists and authType verify
func (svc *verificationService) Test(ctx context.Context, in *gs_service_permission.HasPermissionRequest, out *gs_commons_dto.Status) error {
	return gs_commons_wrapper.ContextToAuthorize(ctx, out, func(auth *gs_commons_dto.Authorize) *gs_commons_dto.State {
		var wg sync.WaitGroup
		wg.Add(4)

		state := errstate.Success

		resp := func(s *gs_commons_dto.State) {
			if state.Ok {
				state = s
			}
			wg.Done()
		}

		//blacklist(ip)
		go func() {
			s, err := svc.blacklistService.Check(gs_commons_transport.InsideContext(gs_commons_constants.PermissionService),
				&gs_service_safety.CheckRequest{
					Type: gs_commons_constants.BlacklistOfIP,
				})
			if err != nil {
				resp(errstate.ErrRequest)
				return
			}
			resp(s.State)
		}()

		//blacklist(userDevice)
		go func() {
			s, err := svc.blacklistService.Check(gs_commons_transport.InsideContext(gs_commons_constants.PermissionService),
				&gs_service_safety.CheckRequest{
					Type: gs_commons_constants.BlacklistOfUserDevice,
				})
			if err != nil {
				resp(errstate.ErrRequest)
				return
			}
			resp(s.State)
		}()

		//api
		go func() {

		}()

		//application
		go func() {
			s, err := svc.applicationService.Status(gs_commons_transport.InsideContext(gs_commons_constants.PermissionService),
				&gs_service_application.FindRequest{
					Content: in.ClientId,
				})
			if err != nil {
				resp(errstate.ErrRequest)
				return
			}
			resp(s.State)
		}()

		wg.Wait()

		if !state.Ok {
			out.State = state
			return nil
		}

		return nil
	})
}

func NewVerificationService(pool *redis.Pool, session *mgo.Session) gs_service_permission.VerificationHandler {
	return &verificationService{pool: pool, session: session}
}
