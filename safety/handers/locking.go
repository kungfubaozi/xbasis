package safety_handers

import (
	"context"
	"konekko.me/gosion/commons/dto"
	"konekko.me/gosion/safety/pb"
)

type lockingService struct {
}

func (svc *lockingService) Lock(context.Context, *gs_service_safety.LockRequest, *gs_commons_dto.Status) error {
	panic("implement me")
}

func (svc *lockingService) Unlock(context.Context, *gs_service_safety.UnlockRequest, *gs_commons_dto.Status) error {
	panic("implement me")
}

func NewLockingService() gs_service_safety.LockingHandler {
	return &lockingService{}
}
