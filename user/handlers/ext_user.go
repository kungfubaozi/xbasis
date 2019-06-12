package userhandlers

import (
	"context"
	"gopkg.in/mgo.v2"
	"konekko.me/gosion/analysis/client"
	"konekko.me/gosion/commons/dto"
	"konekko.me/gosion/commons/errstate"
	"konekko.me/gosion/commons/wrapper"
	inner "konekko.me/gosion/user/pb/inner"
)

type extUserService struct {
	session *mgo.Session
	log     analysisclient.LogClient
}

func (svc *extUserService) GetRepo() *userRepo {
	return &userRepo{session: svc.session.Clone()}
}

func (svc *extUserService) IsExists(ctx context.Context, in *inner.ExistsRequest, out *gs_commons_dto.Status) error {
	return gs_commons_wrapper.ContextToAuthorize(ctx, out, func(auth *gs_commons_wrapper.WrapperUser) *gs_commons_dto.State {
		if len(in.UserId) < 9 {
			return nil
		}
		repo := svc.GetRepo()
		defer repo.Close()

		u, err := repo.FindById(in.UserId)
		if err != nil {
			return nil
		}

		if len(u.Id) > 10 {
			return errstate.Success
		}

		return nil
	})
}

func NewExtUserService(session *mgo.Session, log analysisclient.LogClient) inner.UserHandler {
	return &extUserService{session: session, log: log}
}
