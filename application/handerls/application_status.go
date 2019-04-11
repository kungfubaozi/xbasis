package applicationhanderls

import (
	"context"
	"fmt"
	"gopkg.in/mgo.v2"
	"konekko.me/gosion/application/pb/ext"
	"konekko.me/gosion/commons/dto"
	"konekko.me/gosion/commons/errstate"
	"konekko.me/gosion/commons/indexutils"
	"konekko.me/gosion/commons/wrapper"
	"time"
)

type applicationStatusService struct {
	//session *mgo.Session
	//pool    *redis.Pool
	*indexutils.Client
}

func (svc *applicationStatusService) GetAppClientStatus(ctx context.Context, in *gs_ext_service_application.GetAppClientStatusRequest,
	out *gs_ext_service_application.GetAppClientStatusResponse) error {
	return gs_commons_wrapper.ContextToAuthorize(ctx, out, func(auth *gs_commons_wrapper.WrapperUser) *gs_commons_dto.State {
		if len(in.ClientId) == 0 {
			return errstate.ErrRequest
		}
		s := time.Now().UnixNano()

		repo := svc.GetRepo()
		defer repo.Close()

		a, err := repo.FindByClientId(in.ClientId)
		if err != nil && err == mgo.ErrNotFound {
			return errstate.ErrInvalidClientId
		}
		if err != nil {
			fmt.Println("err", err)
			return nil
		}

		for _, v := range a.Clients {
			if v.Id == in.ClientId {
				fmt.Println("ok, find")
				out.State = errstate.Success
				out.ClientPlatform = v.Platform
				out.ClientEnabled = v.Enabled
				out.AppId = a.Id
				out.AppOpenMode = a.Settings.OpenMode
				out.AppQuarantine = a.Settings.Quarantine
				fmt.Println("time.now-wrapper", (time.Now().UnixNano()-s)/1e6)
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
