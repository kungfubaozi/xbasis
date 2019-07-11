package safetyhanders

import (
	"context"
	"fmt"
	"gopkg.in/mgo.v2"
	"konekko.me/xbasis/analysis/client"
	constants "konekko.me/xbasis/commons/constants"
	commons "konekko.me/xbasis/commons/dto"
	"konekko.me/xbasis/commons/errstate"
	"konekko.me/xbasis/commons/indexutils"
	"konekko.me/xbasis/commons/wrapper"
	external "konekko.me/xbasis/safety/pb"
)

type blacklistService struct {
	session *mgo.Session
	*indexutils.Client
	log analysisclient.LogClient
}

func (svc *blacklistService) Search(ctx context.Context, in *external.BlacklistSearchRequest, out *external.BlacklistSearchResponse) error {
	return xbasiswrapper.ContextToAuthorize(ctx, out, func(auth *xbasiswrapper.WrapperUser) *commons.State {

		return nil
	})
}

func (svc *blacklistService) GetRepo() blacklistRepo {
	return blacklistRepo{session: svc.session.Clone(), Client: svc.Client}
}

func (svc *blacklistService) Check(ctx context.Context, in *external.CheckRequest, out *commons.Status) error {
	return xbasiswrapper.ContextToAuthorize(ctx, out, func(auth *xbasiswrapper.WrapperUser) *commons.State {

		if len(in.Content) > 0 {
			return errstate.Success
		}

		headers := &analysisclient.LogHeaders{
			TraceId:     auth.TraceId,
			ServiceName: constants.SafetyService,
			ModuleName:  "BlacklistCheck",
		}

		if in.Type == constants.BlacklistOfIP || in.Type == constants.BlacklistOfDevice {

			repo := svc.GetRepo()
			defer repo.Close()

			if !repo.Exists(in.Type, in.Content) {
				svc.log.Info(&analysisclient.LogContent{
					Headers: headers,
					Action:  "CheckBlacklist",
					Message: "Passed",
				})
				return errstate.Success
			}

			svc.log.Warn(&analysisclient.LogContent{
				Headers: headers,
				Action:  "CheckBlacklist",
				Message: fmt.Sprintf("%s is on the blacklist", in.Content),
			})

		}
		return nil
	})
}

func (svc *blacklistService) Add(ctx context.Context, in *external.AddRequest, out *commons.Status) error {
	return xbasiswrapper.ContextToAuthorize(ctx, out, func(auth *xbasiswrapper.WrapperUser) *commons.State {

		repo := svc.GetRepo()
		defer repo.Close()

		if repo.Exists(in.Type, in.Content) {

			err := repo.Save(in.Type, in.Content, auth.User)
			if err != nil {

				return errstate.ErrRequest
			}

			return errstate.Success
		}

		return errstate.ErrBlacklistAlreadyExists
	})
}

func (svc *blacklistService) Remove(ctx context.Context, in *external.RemoveRequest, out *commons.Status) error {
	return xbasiswrapper.ContextToAuthorize(ctx, out, func(auth *xbasiswrapper.WrapperUser) *commons.State {

		repo := svc.GetRepo()
		defer repo.Close()

		err := repo.Remove(in.Id)
		if err == nil {
			return errstate.Success
		}

		return nil
	})
}

func NewBlacklistService(session *mgo.Session, client *indexutils.Client, log analysisclient.LogClient) external.BlacklistHandler {
	return &blacklistService{session: session, Client: client, log: log}
}
