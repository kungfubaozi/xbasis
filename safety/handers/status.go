package safety_handers

import (
	"context"
	"gopkg.in/mgo.v2"
	"konekko.me/gosion/commons/dto"
	"konekko.me/gosion/safety/pb"
)

type statusService struct {
	session *mgo.Session
}

func (svc *statusService) Test(context.Context, *gs_service_safety.TestRequest, *gs_commons_dto.Status) error {
	panic("implement me")
}

func NewStatusService(session *mgo.Session) gs_service_safety.StatusHandler {
	return &statusService{session: session}
}
