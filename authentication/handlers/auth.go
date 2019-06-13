package authenticationhandlers

import (
	"context"
	"fmt"
	"github.com/garyburd/redigo/redis"
	"github.com/vmihailenco/msgpack"
	"konekko.me/gosion/analysis/client"
	"konekko.me/gosion/application/pb/inner"
	inner "konekko.me/gosion/authentication/pb/inner"
	"konekko.me/gosion/commons/config/call"
	"konekko.me/gosion/commons/constants"
	"konekko.me/gosion/commons/dto"
	"konekko.me/gosion/commons/errstate"
	"konekko.me/gosion/commons/indexutils"
	"konekko.me/gosion/commons/wrapper"
	"konekko.me/gosion/connection/cmd/connectioncli"
	"konekko.me/gosion/permission/pb/inner"
	"konekko.me/gosion/safety/pb"
	"konekko.me/gosion/safety/pb/inner"
	"sync"
)

type authService struct {
	pool                          *redis.Pool
	blacklistService              gosionsvc_external_safety.BlacklistService
	innerSecurityService          gosionsvc_internal_safety.SecurityService
	innerApplicationStatusService gosionsvc_internal_application.ApplicationStatusService
	innerAccessibleService        gosionsvc_internal_permission.AccessibleService
	connectioncli                 connectioncli.ConnectionClient
	*indexutils.Client
	log analysisclient.LogClient
}

func (svc *authService) GetRepo() *tokenRepo {
	return &tokenRepo{conn: svc.pool.Get()}
}

func (svc *authService) Verify(ctx context.Context, in *inner.VerifyRequest, out *inner.VerifyResponse) error {
	return gs_commons_wrapper.ContextToAuthorize(ctx, out, func(auth *gs_commons_wrapper.WrapperUser) *gs_commons_dto.State {

		configuration := serviceconfiguration.Get()

		if len(in.ClientId) == 0 || len(in.Token) == 0 {
			return errstate.ErrRequest
		}

		headers := &analysisclient.LogHeaders{
			TraceId:     auth.TraceId,
			ServiceName: gs_commons_constants.InternalAuthenticationService,
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
		var tokenApp *gosionsvc_internal_application.GetAppClientStatusResponse

		//应用检查(当为分享功能时，需要检查token内使用的clientId是否为当前应用的id)
		if in.Share && claims.Token.ClientId != auth.FromClientId {
			wg.Add(1)
			go func() {
				defer wg.Done()

				//check token side application status
				a, err := svc.innerApplicationStatusService.GetAppClientStatus(ctx, &gosionsvc_internal_application.GetAppClientStatusRequest{
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

			b, err := repo.Get(claims.Token.UserId, auth.FromClientId+"."+claims.Token.Relation)
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
			if !in.Share && (claims.Token.ClientId != uai.ClientId || auth.FromClientId != uai.ClientId || auth.AppId != uai.AppId) {
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

			s, err := svc.innerSecurityService.Get(ctx, &gosionsvc_internal_safety.GetRequest{UserId: claims.Token.UserId})
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

				//如果功能没有操作角色则忽略
				if len(in.FunctionRoles) == 0 {
					resp(errstate.Success)
					return
				}

				s, err := svc.innerAccessibleService.Check(ctx, &gosionsvc_internal_permission.CheckRequest{
					UserId:      claims.Token.UserId,
					StructureId: in.Funcs,
					//FunctionRoles: in.FunctionRoles,
					FunctionId: in.FunctionId,
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

func NewAuthService(pool *redis.Pool, innerSecurityService gosionsvc_internal_safety.SecurityService,
	connectioncli connectioncli.ConnectionClient, client *indexutils.Client, as gosionsvc_internal_application.ApplicationStatusService,
	blacklistService gosionsvc_external_safety.BlacklistService, innerAccessibleService gosionsvc_internal_permission.AccessibleService, log analysisclient.LogClient) inner.AuthHandler {
	return &authService{pool: pool, innerSecurityService: innerSecurityService, connectioncli: connectioncli, blacklistService: blacklistService,
		Client: client, innerApplicationStatusService: as, innerAccessibleService: innerAccessibleService, log: log}
}
