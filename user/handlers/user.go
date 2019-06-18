package userhandlers

import (
	"context"
	"gopkg.in/mgo.v2"
	"konekko.me/gosion/commons/dto"
	"konekko.me/gosion/commons/indexutils"
	"konekko.me/gosion/commons/wrapper"
	"konekko.me/gosion/user/pb"
)

type userService struct {
	session *mgo.Session
	client  *indexutils.Client
}

func (svc *userService) Search(ctx context.Context, in *gosionsvc_external_user.SearchRequest, out *gosionsvc_external_user.SearchResponse) error {
	return gs_commons_wrapper.ContextToAuthorize(ctx, out, func(auth *gs_commons_wrapper.WrapperUser) *gs_commons_dto.State {
		return nil
	})
}

func NewUserService(session *mgo.Session,
	client *indexutils.Client) gosionsvc_external_user.UserHandler {
	return &userService{session: session, client: client}
}
