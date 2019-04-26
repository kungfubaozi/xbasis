package authenticationhandlers

import (
	"context"
	"fmt"
	"github.com/garyburd/redigo/redis"
	"github.com/vmihailenco/msgpack"
	"konekko.me/gosion/application/pb/ext"
	"konekko.me/gosion/authentication/pb"
	"konekko.me/gosion/authentication/pb/ext"
	"konekko.me/gosion/commons/config/call"
	"konekko.me/gosion/commons/constants"
	"konekko.me/gosion/commons/dto"
	"konekko.me/gosion/commons/errstate"
	"konekko.me/gosion/commons/generator"
	"konekko.me/gosion/commons/indexutils"
	"konekko.me/gosion/commons/wrapper"
	"konekko.me/gosion/connection/cmd/connectioncli"
	"time"
)

type routeService struct {
	extApplicationStatusService gs_ext_service_application.ApplicationStatusService
	extUsersyncService          gs_ext_service_application.UsersyncService
	extTokenService             gs_ext_service_authentication.TokenService
	connectioncli               connectioncli.ConnectionClient
	*indexutils.Client
	pool *redis.Pool
}

func (svc *routeService) Logout(ctx context.Context, in *gs_service_authentication.LogoutRequest, out *gs_commons_dto.Status) error {
	return gs_commons_wrapper.ContextToAuthorize(ctx, out, func(auth *gs_commons_wrapper.WrapperUser) *gs_commons_dto.State {
		repo := svc.GetRepo()
		defer repo.Close()

		return offlineUser(svc.connectioncli, repo, auth.Token.UserId, auth.Token.ClientId)
	})
}

func (svc *routeService) GetRepo() *tokenRepo {
	return &tokenRepo{conn: svc.pool.Get()}
}

func (svc *routeService) Refresh(ctx context.Context, in *gs_service_authentication.RefreshRequest, out *gs_service_authentication.RefreshResponse) error {
	return gs_commons_wrapper.ContextToAuthorize(ctx, out, func(auth *gs_commons_wrapper.WrapperUser) *gs_commons_dto.State {
		if len(in.RefreshToken) >= 100 {

			configuration := serviceconfiguration.Get()

			claims, err := decodeToken(in.RefreshToken, configuration.TokenSecretKey)
			if err == nil {

				repo := svc.GetRepo()
				defer repo.Close()

				fmt.Println("expire", claims.VerifyExpiresAt(time.Now().Unix(), true))

				if claims.Valid() != nil {
					//offline
					s := offlineUser(svc.connectioncli, repo, claims.Token.UserId, auth.ClientId)
					if s.Ok {
						return errstate.ErrRefreshTokenExpired
					}
					return s
				}

				//limit 1 minute refresh
				if time.Now().UnixNano()-claims.IssuedAt <= 60*1e9 {
					return errstate.ErrOperateBusy
				}

				id := gs_commons_generator.NewIDG()

				access := &simpleUserToken{
					Id:       id.Get(),
					UserId:   claims.Token.UserId,
					AppId:    claims.Token.AppId,
					ClientId: claims.Token.ClientId,
					Relation: claims.Token.Relation,
					Type:     gs_commons_constants.AccessToken,
				}

				token, err := encodeToken(configuration.TokenSecretKey, time.Minute*10, access)
				if err != nil {
					return errstate.ErrSystem
				}

				//override
				b, err := repo.Get(claims.Token.UserId, claims.Token.ClientId+"."+claims.Token.Relation)
				if err != nil {
					return errstate.ErrSystem
				}

				var uai userAuthorizeInfo
				err = msgpack.Unmarshal(b, &uai)
				if err != nil {
					return errstate.ErrSystem
				}

				uai.AccessId = access.Id

				b, err = msgpack.Marshal(uai)
				if err != nil {
					return errstate.ErrSystem
				}

				err = repo.Add(access.UserId, access.ClientId, access.Relation, b)
				if err != nil {
					return errstate.ErrSystem
				}

				out.AccessToken = token
				return errstate.Success
			}
		}

		return nil
	})
}

//just support root application web client
//It passes on to the caller new accessToken and refreshToken!
func (svc *routeService) Push(ctx context.Context, in *gs_service_authentication.PushRequest, out *gs_service_authentication.PushResponse) error {
	return gs_commons_wrapper.ContextToAuthorize(ctx, out, func(auth *gs_commons_wrapper.WrapperUser) *gs_commons_dto.State {
		if len(in.RouteTo) == 0 {
			return nil
		}

		//just main client can be route to other application client
		if auth.AppType != gs_commons_constants.AppTypeRoute && auth.Token.AppType != gs_commons_constants.AppTypeRoute {
			return errstate.ErrRouteNotMainClient
		}

		app, err := svc.extApplicationStatusService.GetAppClientStatus(ctx, &gs_ext_service_application.GetAppClientStatusRequest{
			ClientId: in.RouteTo,
			Redirect: in.Redirect,
		})

		if err != nil {
			return nil
		}

		if app.State.Ok {

			fmt.Println("push")

			//can't jump to the main application
			if app.Type == gs_commons_constants.AppTypeRoute {
				return errstate.ErrRequest
			}

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

				fmt.Println("functionStructureId", app.FunctionStructure)
				fmt.Println("userId", auth.User)

				// check permission
				c, err := svc.Client.Count("gs-user-ort", map[string]interface{}{"link_structure_roles.structure_id": app.FunctionStructure, "user_id": auth.User})
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

					//jump to app
					token, err := svc.extTokenService.Generate(ctx, &gs_ext_service_authentication.GenerateRequest{
						RelationId: auth.Token.Relation,
						Route:      true,
						Auth: &gs_commons_dto.Authorize{
							ClientId: in.RouteTo,
							UserId:   auth.Token.UserId,
							Ip:       auth.IP,
							Device:   auth.UserDevice,
							Platform: app.ClientPlatform,
							AppId:    app.AppId,
						},
					})

					if err != nil {
						return errstate.ErrRequest
					}

					if !token.State.Ok {
						return token.State
					}

					out.RefreshToken = token.RefreshToken
					out.AccessToken = token.AccessToken

					return errstate.Success

				}
			}

		}

		return nil
	})
}

func NewRouteService(client *indexutils.Client, pool *redis.Pool, extApplicationStatusService gs_ext_service_application.ApplicationStatusService,
	extUsersyncService gs_ext_service_application.UsersyncService, extTokenService gs_ext_service_authentication.TokenService, connectioncli connectioncli.ConnectionClient) gs_service_authentication.RouterHandler {
	return &routeService{Client: client, pool: pool, extTokenService: extTokenService, extUsersyncService: extUsersyncService, extApplicationStatusService: extApplicationStatusService, connectioncli: connectioncli}
}
