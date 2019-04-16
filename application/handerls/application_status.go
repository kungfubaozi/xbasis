package applicationhanderls

import (
	"context"
	"github.com/garyburd/redigo/redis"
	"konekko.me/gosion/application/pb/ext"
	"konekko.me/gosion/commons/constants"
	"konekko.me/gosion/commons/dto"
	"konekko.me/gosion/commons/errstate"
	"konekko.me/gosion/commons/indexutils"
	"konekko.me/gosion/commons/wrapper"
)

type applicationStatusService struct {
	//session *mgo.Session
	pool *redis.Pool
	*indexutils.Client
}

func (svc *applicationStatusService) GetAppClientStatus(ctx context.Context, in *gs_ext_service_application.GetAppClientStatusRequest,
	out *gs_ext_service_application.GetAppClientStatusResponse) error {
	return gs_commons_wrapper.ContextToAuthorize(ctx, out, func(auth *gs_commons_wrapper.WrapperUser) *gs_commons_dto.State {
		if len(in.ClientId) == 0 {
			return errstate.ErrRequest
		}
		repo := svc.GetRepo()
		defer repo.Close()

		a, err := repo.FindByClientId(in.ClientId)
		if err != nil {
			return errstate.ErrInvalidClientId
		}

		for _, v := range a.Clients {
			if v.Id == in.ClientId {

				if v.Platform == gs_commons_constants.PlatformOfWeb {
					if in.Redirect == a.Settings.RedirectURL {
						out.CanRedirect = true
					}
				}

				out.State = errstate.Success
				out.ClientPlatform = v.Platform
				out.ClientEnabled = v.Enabled
				out.AppId = a.Id
				out.AppQuarantine = a.Settings.Quarantine
				out.UserStructure = a.UserS.Id
				out.FunctionStructure = a.FunctionS.Id
				out.Mustsync = a.Settings.MustSync
				out.Main = a.Main == 101

				return nil
			}
		}

		return nil

	})
}

func (svc *applicationStatusService) GetRepo() *applicationRepo {
	return &applicationRepo{Client: svc.Client}
}

func NewApplicationStatusService(client *indexutils.Client) gs_ext_service_application.ApplicationStatusHandler {
	return &applicationStatusService{Client: client}
}
