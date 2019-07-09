package userhandlers

import (
	"context"
	"gopkg.in/mgo.v2"
	"konekko.me/xbasis/analysis/client"
	commons "konekko.me/xbasis/commons/dto"
	"konekko.me/xbasis/commons/errstate"
	"konekko.me/xbasis/commons/indexutils"
	wrapper "konekko.me/xbasis/commons/wrapper"
	userpb "konekko.me/xbasis/user/pb"
)

type userInfoService struct {
	session *mgo.Session
	log     analysisclient.LogClient
	client  *indexutils.Client
}

func (svc *userInfoService) GetRepo() *userRepo {
	return &userRepo{session: svc.session.Clone(), Client: svc.client}
}

func (svc *userInfoService) GetLocalInfo(ctx context.Context, in *userpb.GetInfoByIdRequest, out *userpb.GetInfoResponse) error {
	return wrapper.ContextToAuthorize(ctx, out, func(auth *wrapper.WrapperUser) *commons.State {

		repo := svc.GetRepo()
		defer repo.Close()

		info, err := repo.FindUserInfo(auth.Token.UserId)
		if err != nil {
			return errstate.ErrRequest
		}

		out.Icon = info.Icon
		out.Username = info.Username

		return errstate.Success
	})
}

func NewUserInfoService(session *mgo.Session, log analysisclient.LogClient, client *indexutils.Client) userpb.UserInfoHandler {
	return &userInfoService{session: session, log: log, client: client}
}
