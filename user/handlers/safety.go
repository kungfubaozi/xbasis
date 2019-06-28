package userhandlers

import (
	"context"
	"gopkg.in/mgo.v2"
	commons "konekko.me/xbasis/commons/dto"
	wrapper "konekko.me/xbasis/commons/wrapper"
	external "konekko.me/xbasis/user/pb"
)

type safetyService struct {
	session *mgo.Session
}

func (svc *safetyService) ForgetPassword(ctx context.Context, in *external.ForgetPasswordRequest, out *commons.Status) error {
	return wrapper.ContextToAuthorize(ctx, out, func(auth *wrapper.WrapperUser) *commons.State {
		return nil
	})
}

func NewSafetyService(session *mgo.Session) external.SafetyHandler {
	return &safetyService{session: session}
}
