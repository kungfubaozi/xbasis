package userhandlers

import (
	"context"
	"konekko.me/gosion/application/pb/inner"
	"konekko.me/gosion/commons/dto"
	"konekko.me/gosion/commons/wrapper"
	external "konekko.me/gosion/user/pb"
)

type authorizationService struct {
	usersyncService gosionsvc_internal_application.UsersyncService
}

func (svc *authorizationService) Sync(ctx context.Context, in *external.SyncRequest, out *gs_commons_dto.Status) error {
	return gs_commons_wrapper.ContextToAuthorize(ctx, out, func(auth *gs_commons_wrapper.WrapperUser) *gs_commons_dto.State {

		if len(in.ClientId) > 8 {
			c, err := svc.usersyncService.Check(ctx, &gosionsvc_internal_application.CheckRequest{
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

func NewAuthorizationService() external.AuthorizationHandler {
	return &authorizationService{}
}
