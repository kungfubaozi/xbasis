package safetyhanders

import (
	"context"
	"gopkg.in/mgo.v2"
	"konekko.me/gosion/commons/constants"
	"konekko.me/gosion/commons/dto"
	"konekko.me/gosion/commons/errstate"
	"konekko.me/gosion/commons/indexutils"
	"konekko.me/gosion/commons/wrapper"
	external "konekko.me/gosion/safety/pb"
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

func (svc *blacklistService) Check(ctx context.Context, in *external.CheckRequest, out *gs_commons_dto.Status) error {
	return gs_commons_wrapper.ContextToAuthorize(ctx, out, func(auth *gs_commons_wrapper.WrapperUser) *gs_commons_dto.State {

		if len(in.Content) > 0 {
			return errstate.Success
		}

		if in.Type == gs_commons_constants.BlacklistOfIP || in.Type == gs_commons_constants.BlacklistOfUserDevice {

			repo := svc.GetRepo()
			defer repo.Close()

			if !repo.Exists(in.Type, in.Content) {
				return errstate.Success
			}

		}
		return nil
	})
}

func (svc *blacklistService) Add(ctx context.Context, in *external.AddRequest, out *gs_commons_dto.Status) error {
	return gs_commons_wrapper.ContextToAuthorize(ctx, out, func(auth *gs_commons_wrapper.WrapperUser) *gs_commons_dto.State {

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

func (svc *blacklistService) Remove(ctx context.Context, in *external.RemoveRequest, out *gs_commons_dto.Status) error {
	return gs_commons_wrapper.ContextToAuthorize(ctx, out, func(auth *gs_commons_wrapper.WrapperUser) *gs_commons_dto.State {

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
