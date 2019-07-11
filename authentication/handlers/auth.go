package authenticationhandlers

import (
	"context"
	"github.com/garyburd/redigo/redis"
	"github.com/vmihailenco/msgpack"
	"konekko.me/xbasis/analysis/client"
	"konekko.me/xbasis/application/pb/inner"
	inner "konekko.me/xbasis/authentication/pb/inner"
	"konekko.me/xbasis/commons/config/call"
	constants "konekko.me/xbasis/commons/constants"
	commons "konekko.me/xbasis/commons/dto"
	"konekko.me/xbasis/commons/errstate"
	"konekko.me/xbasis/commons/indexutils"
	wrapper "konekko.me/xbasis/commons/wrapper"
	"konekko.me/xbasis/connection/cmd/connectioncli"
	"konekko.me/xbasis/permission/pb/inner"
	"konekko.me/xbasis/safety/pb"
	"konekko.me/xbasis/safety/pb/inner"
	"sync"
)

type authService struct {
	pool                          *redis.Pool
	blacklistService              xbasissvc_external_safety.BlacklistService
	innerSecurityService          xbasissvc_internal_safety.SecurityService
	innerApplicationStatusService xbasissvc_internal_application.ApplicationStatusService
	innerAccessibleService        xbasissvc_internal_permission.AccessibleService
	connectioncli                 connectioncli.ConnectionClient
	*indexutils.Client
	log analysisclient.LogClient
}

func (svc *authService) GetRepo() *tokenRepo {
	return &tokenRepo{conn: svc.pool.Get()}
}

func (svc *authService) Verify(ctx context.Context, in *inner.VerifyRequest, out *inner.VerifyResponse) error {
	return wrapper.ContextToAuthorize(ctx, out, func(auth *wrapper.WrapperUser) *commons.State {

		configuration := serviceconfiguration.Get()

		if len(in.ClientId) == 0 || len(in.Token) == 0 {
			return errstate.ErrRequest
		}

		headers := &analysisclient.LogHeaders{
			TraceId:     auth.TraceId,
			ServiceName: constants.InternalAuthenticationService,
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
				Message: "AccessToken has been expired",
			})
			return errstate.ErrAccessTokenExpired
		}

		if claims.Token.Type != constants.AccessToken {
			svc.log.Info(&analysisclient.LogContent{
				Headers: headers,
				Action:  "ErrTokenType",
				Message: "Is not accessToken type",
			})
			return errstate.ErrAccessToken
		}

		state := errstate.Success

		resp := func(s *commons.State) {
			if state.Ok && s != nil {
				state = s
			}
		}

		var uai *userAuthorizeInfo
		var tokenApp *xbasissvc_internal_application.GetAppClientStatusResponse

		//应用检查(当为分享功能时，需要检查token内使用的clientId是否为当前应用的id)
		if in.Share && claims.Token.ClientId != auth.FromClientId {
			wg.Add(1)
			go func() {
				defer wg.Done()

				//check token side application status
				a, err := svc.innerApplicationStatusService.GetAppClientStatus(ctx, &xbasissvc_internal_application.GetAppClientStatusRequest{
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
				if a.ClientEnabled != constants.Enabled {
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
			if claims.Token.UserId != uai.UserId || claims.Token.Relation != uai.Relation || uai.Device != auth.UserDevice || claims.Token.ClientId != auth.FromClientId {
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

			s, err := svc.innerSecurityService.Get(ctx, &xbasissvc_internal_safety.GetRequest{UserId: claims.Token.UserId})
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

				s, err := svc.innerAccessibleService.Check(ctx, &xbasissvc_internal_permission.CheckRequest{
					UserId: claims.Token.UserId,
					//FunctionRoles: in.FunctionRoles,
					FunctionId: in.FunctionId,
					AppId:      in.AppId,
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

func NewAuthService(pool *redis.Pool, innerSecurityService xbasissvc_internal_safety.SecurityService,
	connectioncli connectioncli.ConnectionClient, client *indexutils.Client, as xbasissvc_internal_application.ApplicationStatusService,
	blacklistService xbasissvc_external_safety.BlacklistService, innerAccessibleService xbasissvc_internal_permission.AccessibleService, log analysisclient.LogClient) inner.AuthHandler {
	return &authService{pool: pool, innerSecurityService: innerSecurityService, connectioncli: connectioncli, blacklistService: blacklistService,
		Client: client, innerApplicationStatusService: as, innerAccessibleService: innerAccessibleService, log: log}
}
