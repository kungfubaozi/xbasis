package authenticationhandlers

import (
	"context"
	"github.com/garyburd/redigo/redis"
	"github.com/vmihailenco/msgpack"
	"konekko.me/gosion/application/pb/ext"
	"konekko.me/gosion/authentication/pb/ext"
	"konekko.me/gosion/commons/config/call"
	"konekko.me/gosion/commons/constants"
	"konekko.me/gosion/commons/dto"
	"konekko.me/gosion/commons/errstate"
	"konekko.me/gosion/commons/gslogrus"
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
	*gslogrus.Logger
	*indexutils.Client
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

		var wg sync.WaitGroup
		wg.Add(3)

		claims, err := decodeToken(in.Token, configuration.TokenSecretKey)
		if err != nil {
			return errstate.ErrAccessToken
		}

		if claims.Valid() != nil {
			return errstate.ErrAccessTokenExpired
		}

		if claims.Token.Type != gs_commons_constants.AccessToken {
			return errstate.ErrAccessToken
		}

		state := errstate.Success

		resp := func(s *gs_commons_dto.State) {
			if state.Ok {
				state = s
			}
		}

		var uai userAuthorizeInfo
		var tokenApp *gs_ext_service_application.GetAppClientStatusResponse

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
					resp(errstate.ErrApplicationClosed)
					return
				}
				tokenApp = a
			}()
		}

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

			//check
			if claims.Token.UserId != uai.UserId || claims.Token.Relation != uai.Relation || uai.Device != auth.UserDevice ||
				uai.UserAgent != auth.UserAgent {
				resp(errstate.ErrAccessToken)
				return
			}

			//share function
			if !in.Share && (claims.Token.ClientId != uai.ClientId || auth.ClientId != uai.ClientId || auth.AppId != uai.AppId) {
				resp(errstate.ErrAccessToken)
			}

		}()

		go func() {
			defer wg.Done()
			s, err := svc.extSecurityService.Get(ctx, &gs_ext_service_safety.GetRequest{UserId: claims.Token.UserId})
			if err != nil {
				resp(errstate.ErrSystem)
				return
			}
			resp(s.State)
		}()

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
			if len(uai.UserId) == 0 {
				return errstate.ErrSystem
			}
			out.State = state
			out.UserId = claims.Token.UserId
			out.ClientPlatform = uai.Platform
			out.ClientId = uai.ClientId
			out.AppId = claims.Token.AppId
			out.Relation = claims.Token.Relation
			out.AppType = tokenApp.Type
			return errstate.Success
		}

		return state
	})
}

func NewAuthService(pool *redis.Pool, extSecurityService gs_ext_service_safety.SecurityService,
	connectioncli connectioncli.ConnectionClient, client *indexutils.Client, as gs_ext_service_application.ApplicationStatusService,
	blacklistService gs_service_safety.BlacklistService, extAccessibleService gs_ext_service_permission.AccessibleService, logger *gslogrus.Logger) gs_ext_service_authentication.AuthHandler {
	return &authService{pool: pool, extSecurityService: extSecurityService, connectioncli: connectioncli, blacklistService: blacklistService,
		Client: client, extApplicationStatusService: as, extAccessibleService: extAccessibleService, Logger: logger}
}
