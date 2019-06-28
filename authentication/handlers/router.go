package authenticationhandlers

import (
	"context"
	"fmt"
	"github.com/garyburd/redigo/redis"
	"github.com/vmihailenco/msgpack"
	"konekko.me/xbasis/application/pb/inner"
	external "konekko.me/xbasis/authentication/pb"
	"konekko.me/xbasis/authentication/pb/inner"
	"konekko.me/xbasis/commons/config/call"
	constants "konekko.me/xbasis/commons/constants"
	commons "konekko.me/xbasis/commons/dto"
	"konekko.me/xbasis/commons/errstate"
	"konekko.me/xbasis/commons/generator"
	"konekko.me/xbasis/commons/indexutils"
	wrapper "konekko.me/xbasis/commons/wrapper"
	"konekko.me/xbasis/connection/cmd/connectioncli"
	"konekko.me/xbasis/permission/pb/inner"
	"konekko.me/xbasis/user/pb/inner"
	"time"
)

type routeService struct {
	innerApplicationStatusService xbasissvc_internal_application.ApplicationStatusService
	innerUserSyncService          xbasissvc_internal_application.UserSyncService
	innerTokenService             xbasissvc_internal_authentication.TokenService
	innerUserService              xbasissvc_internal_user.UserService
	innerAccessible               xbasissvc_internal_permission.AccessibleService
	connectioncli                 connectioncli.ConnectionClient
	*indexutils.Client
	pool *redis.Pool
	id   xbasisgenerator.IDGenerator
}

func (svc *routeService) Authorize(ctx context.Context, in *external.AuthorizeRequest, out *commons.Status) error {
	return wrapper.ContextToAuthorize(ctx, out, func(auth *wrapper.WrapperUser) *commons.State {
		if len(auth.Token.Relation) < 16 || len(in.ClientId) < 8 {
			return nil
		}
		s, err := svc.innerApplicationStatusService.GetAppClientStatus(ctx, &xbasissvc_internal_application.GetAppClientStatusRequest{
			ClientId: in.ClientId,
		})
		if err != nil {
			return nil
		}
		if !s.State.Ok {
			return s.State
		}
		//如果用户有权限访问可进入
		if s.AppQuarantine {
			s, err := svc.innerAccessible.HasGrant(ctx, &xbasissvc_internal_permission.HasGrantRequest{
				UserId: auth.Token.UserId,
				AppId:  s.AppId,
			})
			if err != nil {
				return nil
			}
			if !s.State.Ok {
				return nil
			}
		}
		info, err := svc.innerUserService.GetUserInfoById(ctx, &xbasissvc_internal_user.GetUserInfoByIdRequest{
			UserId: auth.Token.UserId,
		})
		if err != nil {
			return nil
		}
		if !info.State.Ok {
			return info.State
		}

		s1, err := svc.innerUserSyncService.Update(ctx, &xbasissvc_internal_application.UserInfo{
			Username: info.Username,
			GId:      info.UserId,
			RealName: info.RealName,
			Icon:     info.Icon,
			AppId:    s.AppId,
		})
		if err != nil {
			return nil
		}
		return s1.State
	})
}

func (svc *routeService) Logout(ctx context.Context, in *external.LogoutRequest, out *commons.Status) error {
	return wrapper.ContextToAuthorize(ctx, out, func(auth *wrapper.WrapperUser) *commons.State {
		repo := svc.GetRepo()
		defer repo.Close()

		configuration := serviceconfiguration.Get()

		claims, err := decodeToken(in.RefreshToken, configuration.TokenSecretKey)
		if err != nil {
			return nil
		}

		if claims.Token.Type != constants.RefreshToken {
			return errstate.ErrRefreshToken
		}

		v, err := repo.SizeOf(claims.Token.UserId)

		//离线与之相关的所有登录信息
		return offlineRelation(svc.connectioncli, v, repo, claims.Token.UserId, claims.Token.Relation)
	})
}

func (svc *routeService) GetRepo() *tokenRepo {
	return &tokenRepo{conn: svc.pool.Get()}
}

func (svc *routeService) Refresh(ctx context.Context, in *external.RefreshRequest, out *external.RefreshResponse) error {
	return wrapper.ContextToAuthorize(ctx, out, func(auth *wrapper.WrapperUser) *commons.State {
		if len(in.RefreshToken) >= 100 {

			configuration := serviceconfiguration.Get()

			claims, err := decodeToken(in.RefreshToken, configuration.TokenSecretKey)
			if err == nil {

				repo := svc.GetRepo()
				defer repo.Close()

				if claims.Valid() != nil {
					//offline
					s := offlineUser(svc.connectioncli, repo, claims.Token.UserId, auth.FromClientId)
					if s.Ok {
						return errstate.ErrRefreshTokenExpired
					}
					return s
				}

				//刷新必须是当前的clientId对应token里的clientId
				if auth.FromClientId != claims.Token.ClientId {
					return errstate.ErrAccessTokenOrClient
				}

				//limit 1 minute refresh
				if time.Now().UnixNano()-claims.IssuedAt <= 60*1e9 {
					return errstate.ErrOperateBusy
				}

				id := svc.id

				access := &simpleUserToken{
					Id:       id.Get(),
					UserId:   claims.Token.UserId,
					AppId:    claims.Token.AppId,
					ClientId: claims.Token.ClientId,
					Relation: claims.Token.Relation,
					Type:     constants.AccessToken,
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
func (svc *routeService) Push(ctx context.Context, in *external.PushRequest, out *external.PushResponse) error {
	return wrapper.ContextToAuthorize(ctx, out, func(auth *wrapper.WrapperUser) *commons.State {
		if len(in.RouteTo) == 0 {
			return nil
		}

		//just main client can be route to other application client
		if auth.AppType != constants.AppTypeRoute && auth.Token.AppType != constants.AppTypeRoute {
			return errstate.ErrRouteNotMainClient
		}

		app, err := svc.innerApplicationStatusService.GetAppClientStatus(ctx, &xbasissvc_internal_application.GetAppClientStatusRequest{
			ClientId: in.RouteTo,
			Redirect: in.Redirect,
		})

		if err != nil {
			return nil
		}

		if app.State.Ok {

			//can't jump to the main application
			if app.Type == constants.AppTypeRoute {
				return errstate.ErrRequest
			}

			if app.ClientEnabled != constants.Enabled {
				return errstate.ErrClientClosed
			}

			//如果用户有权限访问可进入
			if app.AppQuarantine {
				s, err := svc.innerAccessible.HasGrant(ctx, &xbasissvc_internal_permission.HasGrantRequest{
					UserId: auth.Token.UserId,
					AppId:  app.AppId,
				})
				if err != nil {
					return nil
				}
				if !s.State.Ok {
					return nil
				}
			}

			s, err := svc.innerUserSyncService.Check(ctx, &xbasissvc_internal_application.CheckRequest{UserId: auth.User, AppId: app.AppId})
			if err != nil {
				return nil
			}

			if !s.State.Ok {
				return s.State
			}

			//push op
			if len(app.AppId) > 0 {

				//Todo 等待优化
				// check permission

				fmt.Println("function", auth.FunctionId, "userId", auth.Token.UserId)

				c, err := svc.Client.Count("xbs-urf-relations.*", map[string]interface{}{"functionId": auth.FunctionId, "userId": auth.Token.UserId})
				if err != nil {
					return nil
				}

				if c == 0 {
					return errstate.ErrUserAppPermission
				} else if c >= 1 {
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
					token, err := svc.innerTokenService.Generate(ctx, &xbasissvc_internal_authentication.GenerateRequest{
						RelationId: auth.Token.Relation,
						Route:      true,
						Auth: &commons.Authorize{
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
						fmt.Println("token", token.State)
						return token.State
					}

					fmt.Println("1")

					out.RefreshToken = token.RefreshToken

					//当web端登录时，不传入accessToken，需要进行refresh，保证其refreshToken是有效的
					if app.ClientPlatform != constants.PlatformOfWeb {
						out.AccessToken = token.AccessToken
					} else {
						if app.CanRedirect {

						}
					}

					return errstate.Success

				}
			}

		}

		return nil
	})
}

func NewRouteService(client *indexutils.Client, pool *redis.Pool, innerApplicationStatusService xbasissvc_internal_application.ApplicationStatusService,
	innerUserSyncService xbasissvc_internal_application.UserSyncService,
	innerTokenService xbasissvc_internal_authentication.TokenService,
	connectioncli connectioncli.ConnectionClient, innerUserService xbasissvc_internal_user.UserService, innerAccessible xbasissvc_internal_permission.AccessibleService) external.RouterHandler {
	return &routeService{Client: client, pool: pool, innerTokenService: innerTokenService,
		innerUserSyncService: innerUserSyncService, innerUserService: innerUserService, id: xbasisgenerator.NewIDG(),
		innerApplicationStatusService: innerApplicationStatusService, connectioncli: connectioncli, innerAccessible: innerAccessible}
}
