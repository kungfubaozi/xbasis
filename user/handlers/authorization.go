package userhandlers

import (
	"context"
	"konekko.me/gosion/application/pb/ext"
	"konekko.me/gosion/commons/dto"
	"konekko.me/gosion/commons/wrapper"
	"konekko.me/gosion/user/pb"
)

type authorizationService struct {
	usersyncService gs_ext_service_application.UsersyncService
}

func (svc *authorizationService) Sync(ctx context.Context, in *gs_service_user.SyncRequest, out *gs_commons_dto.Status) error {
	return gs_commons_wrapper.ContextToAuthorize(ctx, out, func(auth *gs_commons_wrapper.WrapperUser) *gs_commons_dto.State {

		if len(in.ClientId) > 8 {
			c, err := svc.usersyncService.Check(ctx, &gs_ext_service_application.CheckRequest{
				UserId: auth.Token.UserId,
				AppId:  in.ClientId,
			})
			if err == nil {
				if !c.State.Ok && c.State.Ok {

				}
			}
		}

		return nil
	})
}

func NewAuthorizationService() gs_service_user.AuthorizationHandler {
	return &authorizationService{}
}
