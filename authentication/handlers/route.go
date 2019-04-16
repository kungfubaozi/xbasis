package authenticationhandlers

import (
	"context"
	"github.com/garyburd/redigo/redis"
	"konekko.me/gosion/application/pb/ext"
	"konekko.me/gosion/authentication/pb"
	"konekko.me/gosion/commons/constants"
	"konekko.me/gosion/commons/dto"
	"konekko.me/gosion/commons/errstate"
	"konekko.me/gosion/commons/indexutils"
	"konekko.me/gosion/commons/wrapper"
)

type routeService struct {
	extApplicationStatusService gs_ext_service_application.ApplicationStatusService
	extUsersyncService          gs_ext_service_application.UsersyncService
	*indexutils.Client
	pool *redis.Pool
}

func (svc *routeService) Logout(ctx context.Context, in *gs_service_authentication.LogoutRequest, out *gs_commons_dto.Status) error {
	panic("implement me")
}

func (svc *routeService) Refresh(ctx context.Context, in *gs_service_authentication.RefreshRequest, out *gs_service_authentication.RefreshResponse) error {
	panic("implement me")
}

//just support root application web client
//It passes on to the caller new accessToken and refreshToken!
func (svc *routeService) Push(ctx context.Context, in *gs_service_authentication.PushRequest, out *gs_service_authentication.PushResponse) error {
	return gs_commons_wrapper.ContextToAuthorize(ctx, out, func(auth *gs_commons_wrapper.WrapperUser) *gs_commons_dto.State {
		if len(in.Redirect) == 0 || len(in.RouteTo) == 0 {
			return nil
		}

		app, err := svc.extApplicationStatusService.GetAppClientStatus(ctx, &gs_ext_service_application.GetAppClientStatusRequest{
			ClientId: in.RouteTo,
		})

		if err != nil {
			return nil
		}

		if app.State.Ok {

			if app.ClientEnabled != gs_commons_constants.Enabled {
				return errstate.ErrClientClosed
			}

			if app.Mustsync {

				//check sync log

				s, err := svc.extUsersyncService.Check(ctx, &gs_ext_service_application.CheckRequest{UserId: auth.User, AppId: app.AppId})
				if err != nil {
					return nil
				}

				if !s.State.Ok {
					return s.State
				}

			}

			//push op
			//The application to be transferred must have the following two structures
			if len(app.FunctionStructure) > 0 && len(app.UserStructure) > 0 {

				// check permission
				c, err := svc.Client.Count("gs_user_ort", map[string]interface{}{"link_structure_roles.structure_id": app.FunctionStructure, "user_id": auth.User})
				if err != nil {
					return nil
				}

				if c == 0 {
					return errstate.ErrUserAppPermission
				} else if c > 1 {
					return errstate.ErrSystem
				} else if c == 1 {
					//process

					//must same platform
					if auth.Token.ClientPlatform != app.ClientPlatform && app.ClientPlatform != auth.Platform {
						return errstate.ErrRoutePlatform
					}

					//Not the same application
					if app.AppId == auth.Token.AppId {
						return errstate.ErrRouteSameApplication
					}

				}
			}

		}

		return nil
	})
}

func NewRouteService(client *indexutils.Client, pool *redis.Pool, extApplicationStatusService gs_ext_service_application.ApplicationStatusService,
	extUsersyncService gs_ext_service_application.UsersyncService) gs_service_authentication.RouterHandler {
	return &routeService{Client: client, pool: pool, extUsersyncService: extUsersyncService, extApplicationStatusService: extApplicationStatusService}
}
