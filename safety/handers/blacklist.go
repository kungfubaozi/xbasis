package safetyhanders

import (
	"context"
	"gopkg.in/mgo.v2"
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
}

func (svc *blacklistService) Search(context.Context, *external.BlacklistSearchRequest, *external.BlacklistSearchResponse) error {
	panic("implement me")
}

func (svc *blacklistService) GetRepo() blacklistRepo {
	return blacklistRepo{session: svc.session.Clone(), Client: svc.Client}
}

func (svc *blacklistService) Check(ctx context.Context, in *external.CheckRequest, out *commons.Status) error {
	return xbasiswrapper.ContextToAuthorize(ctx, out, func(auth *xbasiswrapper.WrapperUser) *commons.State {

		if len(in.Content) > 0 {
			return errstate.Success
		}

		if in.Type == constants.BlacklistOfIP || in.Type == constants.BlacklistOfUserDevice {

			repo := svc.GetRepo()
			defer repo.Close()

			if !repo.Exists(in.Type, in.Content) {
				return errstate.Success
			}

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

func NewBlacklistService(session *mgo.Session, client *indexutils.Client) external.BlacklistHandler {
	return &blacklistService{session: session, Client: client}
}
