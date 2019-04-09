package safetyhanders

import (
	"context"
	"konekko.me/gosion/commons/dto"
	"konekko.me/gosion/commons/wrapper"
	"konekko.me/gosion/safety/pb"
)

type lockingService struct {
}

func (svc *lockingService) Lock(ctx context.Context, in *gs_service_safety.LockRequest, out *gs_commons_dto.Status) error {
	return gs_commons_wrapper.ContextToAuthorize(ctx, out, func(auth *gs_commons_wrapper.WrapperUser) *gs_commons_dto.State {

		return nil
	})
}

func (svc *lockingService) Unlock(ctx context.Context, in *gs_service_safety.UnlockRequest, out *gs_commons_dto.Status) error {
	return gs_commons_wrapper.ContextToAuthorize(ctx, out, func(auth *gs_commons_wrapper.WrapperUser) *gs_commons_dto.State {

		return nil
	})
}

func NewLockingService() gs_service_safety.LockingHandler {
	return &lockingService{}
}
