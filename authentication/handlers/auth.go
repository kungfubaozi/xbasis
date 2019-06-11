package authenticationhandlers

import (
	"context"
	"fmt"
	"github.com/garyburd/redigo/redis"
	"github.com/vmihailenco/msgpack"
	"konekko.me/gosion/analysis/client"
	"konekko.me/gosion/application/pb/ext"
	"konekko.me/gosion/authentication/pb/ext"
	"konekko.me/gosion/commons/config/call"
	"konekko.me/gosion/commons/constants"
	"konekko.me/gosion/commons/dto"
	"konekko.me/gosion/commons/errstate"
	"konekko.me/gosion/commons/indexutils"
	"konekko.me/gosion/commons/wrapper"
	"konekko.me/gosion/connection/cmd/connectioncli"
	"konekko.me/gosion/permission/pb/ext"
	"konekko.me/gosion/safety/pb"
	"konekko.me/gosion/safety/pb/ext"
	"sync"
)

type authService struct {
	pool                        *redis.Pool
	blacklistService            gs_service_safety.BlacklistService
	extSecurityService          gs_ext_service_safety.SecurityService
	extApplicationStatusService gs_ext_service_application.ApplicationStatusService
	extAccessibleService        gs_ext_service_permission.AccessibleService
	connectioncli               connectioncli.ConnectionClient
	*indexutils.Client
	log analysisclient.LogClient
}

func (svc *authService) GetRepo() *tokenRepo {
	return &tokenRepo{conn: svc.pool.Get()}
}

func (svc *authService) Verify(ctx context.Context, in *gs_ext_service_authentication.VerifyRequest, out *gs_ext_service_authentication.VerifyResponse) error {
	return gs_commons_wrapper.ContextToAuthorize(ctx, out, func(auth *gs_commons_wrapper.WrapperUser) *gs_commons_dto.State {

		configuration := serviceconfiguration.Get()

		if len(in.ClientId) == 0 || len(in.Token) == 0 {
			return errstate.ErrRequest
		}

		headers := &analysisclient.LogHeaders{
			TraceId:     auth.TraceId,
			ServiceName: gs_commons_constants.ExtAuthenticationService,
			ModuleName:  "Auth",
		}

		var wg sync.WaitGroup
		wg.Add(3)

		claims, err := decodeToken(in.Token, configuration.TokenSecretKey)
		if err != nil {
			svc.log.Info(&analysisclient.LogContent{
				Headers: headers,
				Action:  "DecodeTokenError",
				Message: err.Error(),
			})
			return errstate.ErrAccessToken
		}

		if claims.Valid() != nil {
			svc.log.Info(&analysisclient.LogContent{
				Headers: headers,
				Action:  "AccessTokenExpired",
				Message: fmt.Sprintf("%d", claims.ExpiresAt),
			})
			return errstate.ErrAccessTokenExpired
		}

		if claims.Token.Type != gs_commons_constants.AccessToken {
			svc.log.Info(&analysisclient.LogContent{
				Headers: headers,
				Action:  "ErrTokenType",
			})
			return errstate.ErrAccessToken
		}

		state := errstate.Success

		resp := func(s *gs_commons_dto.State) {
			if state.Ok && s != nil {
				state = s
			}
		}

		var uai *userAuthorizeInfo
		var tokenApp *gs_ext_service_application.GetAppClientStatusResponse

		//应用检查(当为分享功能时，需要检查token内使用的clientId是否为当前应用的id)
		if in.Share && claims.Token.ClientId != auth.ClientId {
			wg.Add(1)
			go func() {
				defer wg.Done()

				//check token side application status
				a, err := svc.extApplicationStatusService.GetAppClientStatus(ctx, &gs_ext_service_application.GetAppClientStatusRequest{
					ClientId: claims.Token.ClientId,
				})

				if err != nil {
					resp(errstate.ErrSystem)
					return
				}
				if !a.State.Ok {
					resp(a.State)
					return
				}
				if a.ClientEnabled != gs_commons_constants.Enabled {
					svc.log.Info(&analysisclient.LogContent{
						Headers:   headers,
						Action:    "ApplicationClosed",
						StateCode: errstate.ErrApplicationClosed.Code,
					})
					resp(errstate.ErrApplicationClosed)
					return
				}
				tokenApp = a
			}()
		}

		//token检查
		go func() {
			defer wg.Done()

			repo := svc.GetRepo()
			defer repo.Close()

			b, err := repo.Get(claims.Token.UserId, auth.ClientId+"."+claims.Token.Relation)
			if err != nil || b == nil {
				resp(errstate.ErrAccessTokenOrClient)
				return
			}

			err = msgpack.Unmarshal(b, &uai)
			if err != nil {
				resp(errstate.ErrSystem)
				return
			}

			if claims.Token.Id != uai.AccessId {
				resp(errstate.ErrAccessToken)
				return
			}

			//check //uai.UserAgent != auth.UserAgent
			if claims.Token.UserId != uai.UserId || claims.Token.Relation != uai.Relation || uai.Device != auth.UserDevice {
				resp(errstate.ErrAccessToken)
				svc.log.Info(&analysisclient.LogContent{
					Headers: headers,
					Action:  "ClaimsInfoCheckError",
				})
				return
			}

			//share function
			if !in.Share && (claims.Token.ClientId != uai.ClientId || auth.ClientId != uai.ClientId || auth.AppId != uai.AppId) {
				resp(errstate.ErrAccessToken)
				svc.log.Info(&analysisclient.LogContent{
					Headers: headers,
					Action:  "ShareFunctionClaimsInfoCheckError",
				})
			}

		}()

		//安全检查
		go func() {
			defer wg.Done()

			s, err := svc.extSecurityService.Get(ctx, &gs_ext_service_safety.GetRequest{UserId: claims.Token.UserId})
			if err != nil {
				resp(errstate.ErrSystem)
				return
			}

			resp(s.State)
		}()

		//功能权限检查
		go func() {
			defer wg.Done()

			if !in.Share {

				if len(in.FunctionRoles) == 0 {
					resp(errstate.Success)
					return
				}

				s, err := svc.extAccessibleService.Check(ctx, &gs_ext_service_permission.CheckRequest{
					UserId:        claims.Token.UserId,
					StructureId:   in.Funcs,
					FunctionRoles: in.FunctionRoles,
				})

				if err != nil {
					resp(errstate.ErrSystem)
					return
				}

				resp(s.State)
				return
			}

			resp(errstate.Success)
			return

		}()

		wg.Wait()

		if state.Ok {

			if uai == nil || len(uai.UserId) == 0 {
				svc.log.Info(&analysisclient.LogContent{
					Headers: headers,
					Action:  "ErrUserId",
				})
				return errstate.ErrSystem
			}

			out.State = state
			out.UserId = claims.Token.UserId
			out.ClientPlatform = uai.Platform
			out.ClientId = uai.ClientId
			out.AppId = claims.Token.AppId
			out.Relation = claims.Token.Relation
			if tokenApp != nil {
				out.AppType = tokenApp.Type
			}
			return nil
		}

		return state
	})
}

func NewAuthService(pool *redis.Pool, extSecurityService gs_ext_service_safety.SecurityService,
	connectioncli connectioncli.ConnectionClient, client *indexutils.Client, as gs_ext_service_application.ApplicationStatusService,
	blacklistService gs_service_safety.BlacklistService, extAccessibleService gs_ext_service_permission.AccessibleService, log analysisclient.LogClient) gs_ext_service_authentication.AuthHandler {
	return &authService{pool: pool, extSecurityService: extSecurityService, connectioncli: connectioncli, blacklistService: blacklistService,
		Client: client, extApplicationStatusService: as, extAccessibleService: extAccessibleService, log: log}
}
