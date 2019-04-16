package userhandlers

import (
	"context"
	"konekko.me/gosion/commons/dto"
	"konekko.me/gosion/commons/wrapper"
	"konekko.me/gosion/user/pb"
)

type authorizationService struct {
}

func (svc *authorizationService) Sync(ctx context.Context, in *gs_service_user.SyncRequest, out *gs_commons_dto.Status) error {
	return gs_commons_wrapper.ContextToAuthorize(ctx, out, func(auth *gs_commons_wrapper.WrapperUser) *gs_commons_dto.State {
		return nil
	})
}

func NewAuthorizationService() gs_service_user.AuthorizationHandler {
	return &authorizationService{}
}
