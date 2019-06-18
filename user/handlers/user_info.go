package userhandlers

import (
	"context"
	"gopkg.in/mgo.v2"
	"konekko.me/gosion/analysis/client"
	"konekko.me/gosion/commons/dto"
	"konekko.me/gosion/commons/errstate"
	"konekko.me/gosion/commons/wrapper"
	"konekko.me/gosion/user/pb"
)

type userInfoService struct {
	session *mgo.Session
	log     analysisclient.LogClient
}

func (svc *userInfoService) GetRepo() *userRepo {
	return &userRepo{session: svc.session.Clone()}
}

func (svc *userInfoService) GetLocalInfo(ctx context.Context, in *gosionsvc_external_user.GetInfoByIdRequest, out *gosionsvc_external_user.GetInfoResponse) error {
	return gs_commons_wrapper.ContextToAuthorize(ctx, out, func(auth *gs_commons_wrapper.WrapperUser) *gs_commons_dto.State {

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

func NewUserInfoService(session *mgo.Session, log analysisclient.LogClient) gosionsvc_external_user.UserInfoHandler {
	return &userInfoService{session: session, log: log}
}
