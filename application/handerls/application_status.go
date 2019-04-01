package application_handerls

import (
	"context"
	"github.com/garyburd/redigo/redis"
	"gopkg.in/mgo.v2"
	"konekko.me/gosion/application/pb/ext"
	"konekko.me/gosion/application/repositories"
	"konekko.me/gosion/commons/dto"
	"konekko.me/gosion/commons/errstate"
	"konekko.me/gosion/commons/wrapper"
)

type applicationStatusService struct {
	session *mgo.Session
	pool    *redis.Pool
}

func (svc *applicationStatusService) GetAppClientStatus(ctx context.Context, in *gs_ext_service_application.GetAppClientStatusRequest,
	out *gs_ext_service_application.GetAppClientStatusResponse) error {
	return gs_commons_wrapper.ContextToAuthorize(ctx, out, func(auth *gs_commons_wrapper.WrapperUser) *gs_commons_dto.State {
		repo := svc.GetRepo()
		defer repo.Close()

		if len(in.ClientId) == 0 {
			return errstate.ErrRequest
		}

		a, err := repo.FindByClientId(in.ClientId)
		if err != nil {
			return nil
		}

		for _, v := range a.Clients {
			if v.Id == in.ClientId {
				out.State = errstate.Success
				out.ClientPlatform = v.Platform
				out.ClientEnabled = v.Enabled
				out.AppId = a.Id
				out.AppOpenMode = a.Settings.OpenMode
				out.AppQuarantine = a.Settings.Quarantine
				return nil
			}
		}

		return nil

	})
}

func (svc *applicationStatusService) GetRepo() application_repositories.ApplicationRepo {
	return application_repositories.ApplicationRepo{Session: svc.session.Clone(), Conn: svc.pool.Get()}
}

func NewApplicationStatusServie() gs_nops_service_application.ApplicationStatusHandler {
	return &applicationStatusService{}
}
