package applicationhanderls

import (
	"context"
	"gopkg.in/mgo.v2"
	external "konekko.me/xbasis/application/pb"
	constants "konekko.me/xbasis/commons/constants"
	commons "konekko.me/xbasis/commons/dto"
	"konekko.me/xbasis/commons/errstate"
	"konekko.me/xbasis/commons/indexutils"
	wrapper "konekko.me/xbasis/commons/wrapper"
)

type settingsService struct {
	session *mgo.Session
	*indexutils.Client
}

func (svc *settingsService) GetRepo() *applicationRepo {
	return &applicationRepo{session: svc.session.Clone(), Client: svc.Client}
}

func (svc *settingsService) Update(ctx context.Context, in *external.UpdateRequest, out *commons.Status) error {
	return wrapper.ContextToAuthorize(ctx, out, func(auth *wrapper.WrapperUser) *commons.State {

		repo := svc.GetRepo()
		defer repo.Close()

		info, err := repo.FindByApplicationId(in.AppId)
		if err != nil {
			return nil
		}

		if in.Setting.Enabled > 0 && in.Setting.Enabled == constants.Closed ||
			in.Setting.Enabled == constants.Enabled {
			info.Settings.Enabled = in.Setting.Enabled
		} else {
			return nil
		}

		//err = repo.Upsert(info)
		//
		//if err != nil {
		//	return nil
		//}

		return errstate.Success
	})
}

func (svc *settingsService) EnabledClient(ctx context.Context, in *external.EnabledRequest, out *commons.Status) error {
	return wrapper.ContextToAuthorize(ctx, out, func(auth *wrapper.WrapperUser) *commons.State {

		repo := svc.GetRepo()
		defer repo.Close()

		if len(in.AppId) > 0 && len(in.Id) > 0 && in.Enabled > 0 {
			return nil
		}

		info, err := repo.FindByApplicationId(in.AppId)
		if err != nil {
			return nil
		}

		if len(info.Clients) > 0 {
			for _, v := range info.Clients {
				if v.Id == in.Id {
					v.Enabled = in.Enabled

					//err = repo.Upsert(info)
					//if err != nil {
					//	return nil
					//}
					return errstate.Success
				}
			}
		}

		return nil
	})
}

func NewSettingsService(session *mgo.Session, client *indexutils.Client) external.SettingsHandler {
	return &settingsService{session: session, Client: client}
}
