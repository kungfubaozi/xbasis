package applicationhanderls

import (
	"context"
	"fmt"
	"github.com/garyburd/redigo/redis"
	"konekko.me/xbasis/analysis/client"
	inner "konekko.me/xbasis/application/pb/inner"
	constants "konekko.me/xbasis/commons/constants"
	commons "konekko.me/xbasis/commons/dto"
	"konekko.me/xbasis/commons/errstate"
	"konekko.me/xbasis/commons/indexutils"
	wrapper "konekko.me/xbasis/commons/wrapper"
)

type applicationStatusService struct {
	//session *mgo.Session
	pool *redis.Pool
	*indexutils.Client
	log analysisclient.LogClient
}

func (svc *applicationStatusService) GetAppClientStatus(ctx context.Context, in *inner.GetAppClientStatusRequest,
	out *inner.GetAppClientStatusResponse) error {
	return wrapper.ContextToAuthorize(ctx, out, func(auth *wrapper.WrapperUser) *commons.State {
		headers := &analysisclient.LogHeaders{
			TraceId:     auth.TraceId,
			ServiceName: constants.InternalApplicationService,
			ModuleName:  "GetAppClientStatus",
		}

		if len(in.ClientId) == 0 {
			svc.log.Info(&analysisclient.LogContent{
				Headers: headers,
				Action:  "GetAppByClientId",
				Message: "Invalid input clientId",
			})
			return errstate.ErrRequest
		}

		repo := svc.GetRepo()
		defer repo.Close()

		a, err := repo.FindByClientId(in.ClientId)
		if err != nil {
			svc.log.Info(&analysisclient.LogContent{
				Headers: headers,
				Action:  "GetAppByClientId",
				Message: fmt.Sprintf("Not found application as clientId %s", in.ClientId),
			})
			return errstate.ErrInvalidClientId
		}

		for _, v := range a.Clients {
			if v.Id == in.ClientId {

				svc.log.Info(&analysisclient.LogContent{
					Headers: headers,
					Action:  "GetAppByClientId",
				})

				if v.Platform == constants.PlatformOfWeb {
					if in.Redirect == a.Settings.RedirectURL {
						out.CanRedirect = true
					}
				}

				out.ClientPlatform = v.Platform
				out.ClientEnabled = v.Enabled
				out.AppId = a.Id
				out.AppQuarantine = a.Settings.Quarantine
				out.SecretKey = a.SecretKey
				out.Type = a.Type
				out.ServiceName = a.Settings.ServiceName
				return errstate.Success
			}
		}

		return nil

	})
}

func (svc *applicationStatusService) GetRepo() *applicationRepo {
	return getApplicationRepo(nil, svc.Client, svc.pool.Get())
}

func NewApplicationStatusService(client *indexutils.Client, pool *redis.Pool, log analysisclient.LogClient) inner.ApplicationStatusHandler {
	return &applicationStatusService{Client: client, pool: pool, log: log}
}
