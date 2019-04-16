package authenticationhandlers

import (
	"context"
	"github.com/garyburd/redigo/redis"
	"github.com/vmihailenco/msgpack"
	"konekko.me/gosion/authentication/pb/ext"
	"konekko.me/gosion/commons/config/call"
	"konekko.me/gosion/commons/constants"
	"konekko.me/gosion/commons/dto"
	"konekko.me/gosion/commons/errstate"
	"konekko.me/gosion/commons/indexutils"
	"konekko.me/gosion/commons/wrapper"
	"konekko.me/gosion/connection/cmd/connectioncli"
	"konekko.me/gosion/safety/pb/ext"
	"sync"
)

type authService struct {
	pool               *redis.Pool
	extSecurityService gs_ext_service_safety.SecurityService
	connectioncli      connectioncli.ConnectionClient
	*indexutils.Client
}

func (svc *authService) GetRepo() *tokenRepo {
	return &tokenRepo{conn: svc.pool.Get()}
}

func (svc *authService) Verify(ctx context.Context, in *gs_ext_service_authentication.VerifyRequest, out *gs_ext_service_authentication.VerifyResponse) error {
	return gs_commons_wrapper.ContextToAuthorize(ctx, out, func(auth *gs_commons_wrapper.WrapperUser) *gs_commons_dto.State {
		//1.verify token

		var wg sync.WaitGroup
		wg.Add(3)

		configuration := serviceconfiguration.Get()

		if len(in.ClientId) == 0 || len(in.Token) == 0 {

			return errstate.ErrRequest
		}

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

			//check
			if claims.Token.UserId != uai.UserId || claims.Token.ClientId != uai.ClientId ||
				claims.Token.Relation != uai.Relation || claims.Token.AppId != uai.AppId ||
				uai.Device != auth.UserDevice || uai.UserAgent != auth.UserAgent ||
				auth.ClientId != uai.ClientId || auth.AppId != uai.AppId {
				resp(errstate.ErrAccessToken)
				return
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

			var userRoles []interface{}

			//get user require roles
			var userroles map[string]interface{}

			ok, err := svc.Client.QueryFirst("gs_user_ort", map[string]interface{}{"link_structure_roles.structure_id": in.Funcs, "user_id": claims.Token.UserId}, &userroles, "link_structure_roles.roles")
			if err != nil || !ok {

				resp(errstate.ErrRequest)
				return
			}

			userRoles = userroles["link_structure_roles"].([]interface{})[0].(map[string]interface{})["roles"].([]interface{})

			if userRoles != nil && len(userRoles) > 0 && len(in.FunctionRoles) > 0 {

				roles := make(map[string]string)
				ok := false
				for _, v := range userRoles {
					b := v.(string)
					roles[b] = "ok"
				}
				for _, v := range in.FunctionRoles {
					if roles[v] == "ok" {
						ok = true
						break
					}
				}

				if ok {
					resp(errstate.Success)
					return
				} else {
					resp(errstate.ErrUserPermission)
					return
				}
			}

			resp(errstate.ErrUserPermission)
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
			return nil
		}

		return state
	})
}

func NewAuthService(pool *redis.Pool, extSecurityService gs_ext_service_safety.SecurityService,
	connectioncli connectioncli.ConnectionClient, client *indexutils.Client) gs_ext_service_authentication.AuthHandler {
	return &authService{pool: pool, extSecurityService: extSecurityService, connectioncli: connectioncli, Client: client}
}
