package userhandlers

import (
	"context"
	"gopkg.in/mgo.v2"
	"konekko.me/xbasis/analysis/client"
	commons "konekko.me/xbasis/commons/dto"
	"konekko.me/xbasis/commons/errstate"
	"konekko.me/xbasis/commons/indexutils"
	"konekko.me/xbasis/commons/wrapper"
	inner "konekko.me/xbasis/user/pb/inner"
)

type innerUserService struct {
	session *mgo.Session
	log     analysisclient.LogClient
	client  *indexutils.Client
}

func (svc *innerUserService) GetUserInfoById(ctx context.Context, in *inner.GetUserInfoByIdRequest, out *inner.SimpleUserInfo) error {
	return xbasiswrapper.ContextToAuthorize(ctx, out, func(auth *xbasiswrapper.WrapperUser) *commons.State {
		if len(in.UserId) < 10 {
			return nil
		}

		repo := svc.GetRepo()
		defer repo.Close()

		s, err := repo.FindUserInfo(in.UserId)
		if err != nil {
			return errstate.ErrRequest
		}

		out.UserId = s.UserId
		out.RealName = s.RealName
		out.Username = s.Username
		out.Icon = s.Icon
		out.UserState = int64(s.State)
		out.Invite = s.Invite
		out.FromInvite = s.FromInvite

		return errstate.Success
	})
}

func (svc *innerUserService) GetRepo() *userRepo {
	return &userRepo{session: svc.session.Clone(), Client: svc.client}
}

func (svc *innerUserService) IsExists(ctx context.Context, in *inner.ExistsRequest, out *commons.Status) error {
	return xbasiswrapper.ContextToAuthorize(ctx, out, func(auth *xbasiswrapper.WrapperUser) *commons.State {
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

func NewInnerUserService(session *mgo.Session, log analysisclient.LogClient, client *indexutils.Client) inner.UserHandler {
	return &innerUserService{session: session, log: log, client: client}
}
