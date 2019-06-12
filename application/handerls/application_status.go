package applicationhanderls

import (
	"context"
	"github.com/garyburd/redigo/redis"
	"konekko.me/gosion/analysis/client"
	inner "konekko.me/gosion/application/pb/inner"
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
	log analysisclient.LogClient
}

func (svc *applicationStatusService) GetAppClientStatus(ctx context.Context, in *inner.GetAppClientStatusRequest,
	out *inner.GetAppClientStatusResponse) error {
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
				out.SecretKey = a.SecretKey
				if a.UserS == nil || a.FunctionS == nil {

					return errstate.ErrSystem
				}

				out.UserStructure = a.UserS.Id
				out.FunctionStructure = a.FunctionS.Id
				out.Mustsync = a.Settings.MustSync
				out.Type = a.Type

				return nil
			}
		}

		return nil

	})
}

func (svc *applicationStatusService) GetRepo() *applicationRepo {
	return &applicationRepo{Client: svc.Client}
}

func NewApplicationStatusService(client *indexutils.Client, pool *redis.Pool, log analysisclient.LogClient) inner.ApplicationStatusHandler {
	return &applicationStatusService{Client: client, pool: pool, log: log}
}
