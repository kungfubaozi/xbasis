package userhandlers

import (
	"context"
	"gopkg.in/mgo.v2"
	commons "konekko.me/xbasis/commons/dto"
	"konekko.me/xbasis/commons/indexutils"
	wrapper "konekko.me/xbasis/commons/wrapper"
	userpb "konekko.me/xbasis/user/pb"
)

type userService struct {
	session *mgo.Session
	client  *indexutils.Client
}

func (svc *userService) Search(ctx context.Context, in *userpb.SearchRequest, out *userpb.SearchResponse) error {
	return wrapper.ContextToAuthorize(ctx, out, func(auth *wrapper.WrapperUser) *commons.State {
		return nil
	})
}

func NewUserService(session *mgo.Session,
	client *indexutils.Client) userpb.UserHandler {
	return &userService{session: session, client: client}
}
