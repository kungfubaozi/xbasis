package safety_handers

import (
	"context"
	"gopkg.in/mgo.v2"
	"konekko.me/gosion/commons/dto"
	"konekko.me/gosion/commons/wrapper"
	"konekko.me/gosion/safety/pb"
)

type statusService struct {
	session *mgo.Session
}

func (svc *statusService) Test(ctx context.Context, in *gs_service_safety.TestRequest, out *gs_commons_dto.Status) error {
	return gs_commons_wrapper.ContextToAuthorize(ctx, out, func(auth *gs_commons_wrapper.WrapperUser) *gs_commons_dto.State {

		return nil
	})
}

func NewStatusService(session *mgo.Session) gs_service_safety.StatusHandler {
	return &statusService{session: session}
}
