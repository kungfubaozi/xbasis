package safety_handers

import (
	"context"
	"konekko.me/gosion/commons/dto"
	"konekko.me/gosion/safety/pb"
)

type frozenService struct {
}

func (svc *frozenService) Request(context.Context, *gs_service_safety.FrozenRequest, *gs_commons_dto.Status) error {
	panic("implement me")
}

func (svc *frozenService) Unblock(context.Context, *gs_service_safety.UnblockRequest, *gs_commons_dto.Status) error {
	panic("implement me")
}

func NewFrozenService() gs_service_safety.FrozenHandler {
	return &frozenService{}
}
