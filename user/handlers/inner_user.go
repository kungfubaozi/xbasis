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

type innerUserService struct {
	session *mgo.Session
	log     analysisclient.LogClient
}

func (svc *innerUserService) GetUserInfoById(ctx context.Context, in *inner.GetUserInfoByIdRequest, out *inner.SimpleUserInfo) error {
	return gs_commons_wrapper.ContextToAuthorize(ctx, out, func(auth *gs_commons_wrapper.WrapperUser) *gs_commons_dto.State {
		if len(in.UserId) < 10 {
			return nil
		}

		repo := svc.GetRepo()
		defer repo.Close()

		s, err := repo.FindUserInfo(in.UserId)
		if err != nil {
			return errstate.ErrRequest
		}

		out.RealName = s.RealName
		out.Username = s.Username
		out.Icon = s.Icon

		return errstate.Success
	})
}

func (svc *innerUserService) GetRepo() *userRepo {
	return &userRepo{session: svc.session.Clone()}
}

func (svc *innerUserService) IsExists(ctx context.Context, in *inner.ExistsRequest, out *gs_commons_dto.Status) error {
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

func NewInnerUserService(session *mgo.Session, log analysisclient.LogClient) inner.UserHandler {
	return &innerUserService{session: session, log: log}
}
