package userhandlers

import (
	"context"
	"gopkg.in/mgo.v2"
	"konekko.me/gosion/commons/indexutils"
	"konekko.me/gosion/user/pb"
)

type userService struct {
	session *mgo.Session
	client  *indexutils.Client
}

func (svc *userService) Search(context.Context, *gosionsvc_external_user.SearchRequest, *gosionsvc_external_user.SearchResponse) error {
	panic("implement me")
}

func NewUserService(session *mgo.Session,
	client *indexutils.Client) gosionsvc_external_user.UserHandler {
	return &userService{session: session, client: client}
}
