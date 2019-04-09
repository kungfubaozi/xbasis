package safetyhanders

import (
	"context"
	"konekko.me/gosion/commons/dto"
	"konekko.me/gosion/commons/wrapper"
	"konekko.me/gosion/safety/pb"
)

type frozenService struct {
}

func (svc *frozenService) Request(ctx context.Context, in *gs_service_safety.FrozenRequest, out *gs_commons_dto.Status) error {
	return gs_commons_wrapper.ContextToAuthorize(ctx, out, func(auth *gs_commons_wrapper.WrapperUser) *gs_commons_dto.State {

		return nil
	})
}

func (svc *frozenService) Unblock(ctx context.Context, in *gs_service_safety.UnblockRequest, out *gs_commons_dto.Status) error {
	return gs_commons_wrapper.ContextToAuthorize(ctx, out, func(auth *gs_commons_wrapper.WrapperUser) *gs_commons_dto.State {

		return nil
	})
}

func NewFrozenService() gs_service_safety.FrozenHandler {
	return &frozenService{}
}
