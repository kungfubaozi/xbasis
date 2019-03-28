package authentication_handlers

import (
	"context"
	"konekko.me/gosion/authentication/pb"
	"konekko.me/gosion/commons/dto"
	"konekko.me/gosion/commons/wrapper"
)

type routeService struct {
}

func (svc *routeService) Logout(ctx context.Context, in *gs_service_authentication.LogoutRequest, out *gs_commons_dto.Status) error {
	panic("implement me")
}

func (svc *routeService) Refresh(ctx context.Context, in *gs_service_authentication.RefreshRequest, out *gs_service_authentication.RefreshResponse) error {
	panic("implement me")
}

//just support web platform
//just support web platform
//just support web platform
//It passes on to the caller new accessToken and refreshToken!
func (svc *routeService) Push(ctx context.Context, in *gs_service_authentication.PushRequest, out *gs_service_authentication.PushResponse) error {
	return gs_commons_wrapper.ContextToAuthorize(ctx, out, func(auth *gs_commons_wrapper.WrapperUser) *gs_commons_dto.State {
		return nil
	})
}

func NewRouteService() gs_service_authentication.RouterHandler {
	return &routeService{}
}
