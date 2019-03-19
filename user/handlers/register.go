package user_handlers

import (
	"context"
	"gopkg.in/mgo.v2"
	"konekko.me/gosion/commons/dto"
	"konekko.me/gosion/commons/wrapper"
	"konekko.me/gosion/user/pb"
)

type registerService struct {
	session *mgo.Session
}

func (svc *registerService) New(ctx context.Context, in *gs_service_user.NewRequest, out *gs_commons_dto.Status) error {
	return gs_commons_wrapper.ContextToAuthorize(ctx, out, func(auth *gs_commons_dto.Authorize) *gs_commons_dto.State {
		return nil
	})
}

func NewRegisterService() gs_service_user.RegisterHandler {
	return &registerService{}
}
