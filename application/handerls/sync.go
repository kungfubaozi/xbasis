package applicationhanderls

import (
	"context"
	"gopkg.in/mgo.v2"
	inner "konekko.me/gosion/application/pb/inner"
	"konekko.me/gosion/commons/dto"
	"konekko.me/gosion/commons/errstate"
	"konekko.me/gosion/commons/indexutils"
	"konekko.me/gosion/commons/wrapper"
)

type syncService struct {
	*indexutils.Client
	session *mgo.Session
}

func (svc *syncService) GetRepo() *syncRepo {
	return &syncRepo{Client: svc.Client, session: svc.session.Clone()}
}

func (svc *syncService) Transport(ctx context.Context, in *inner.UserInfo, out *gs_commons_dto.Status) error {
	return gs_commons_wrapper.ContextToAuthorize(ctx, out, func(auth *gs_commons_wrapper.WrapperUser) *gs_commons_dto.State {
		return nil
	})
}

func (svc *syncService) Check(ctx context.Context, in *inner.CheckRequest, out *gs_commons_dto.Status) error {
	return gs_commons_wrapper.ContextToAuthorize(ctx, out, func(auth *gs_commons_wrapper.WrapperUser) *gs_commons_dto.State {

		if len(in.UserId) > 0 || len(in.AppId) > 0 {

			repo := svc.GetRepo()
			defer repo.Close()

			//Check if the user is synchronized to the corresponding application
			c, err := repo.IsSynced(in.UserId, in.AppId)
			if err == nil && c == 1 {
				return errstate.Success
			}

		}

		return nil
	})
}

func (svc *syncService) Update(ctx context.Context, in *inner.UserInfo, out *gs_commons_dto.Status) error {
	return gs_commons_wrapper.ContextToAuthorize(ctx, out, func(auth *gs_commons_wrapper.WrapperUser) *gs_commons_dto.State {

		return nil
	})
}

func NewSyncService(client *indexutils.Client, session *mgo.Session) inner.UsersyncHandler {
	return &syncService{Client: client, session: session}
}
