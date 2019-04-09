package userhandlers

import (
	"context"
	"gopkg.in/mgo.v2"
	"konekko.me/gosion/commons/dto"
	"konekko.me/gosion/commons/wrapper"
	"konekko.me/gosion/user/pb"
)

type safetyService struct {
	session *mgo.Session
}

func (svc *safetyService) ForgetPassword(ctx context.Context, in *gs_service_user.ForgetPasswordRequest, out *gs_commons_dto.Status) error {
	return gs_commons_wrapper.ContextToAuthorize(ctx, out, func(auth *gs_commons_wrapper.WrapperUser) *gs_commons_dto.State {
		return nil
	})
}

func NewSafetyService(session *mgo.Session) gs_service_user.SafetyHandler {
	return &safetyService{session: session}
}
