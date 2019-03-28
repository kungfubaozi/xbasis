package authentication_handlers

import (
	"context"
	"github.com/garyburd/redigo/redis"
	"github.com/vmihailenco/msgpack"
	"konekko.me/gosion/authentication/pb/nops"
	"konekko.me/gosion/authentication/repositories"
	"konekko.me/gosion/commons/config"
	"konekko.me/gosion/commons/constants"
	"konekko.me/gosion/commons/dto"
	"konekko.me/gosion/commons/errstate"
	"konekko.me/gosion/commons/wrapper"
	"konekko.me/gosion/connection/cmd/connectioncli"
	"konekko.me/gosion/safety/pb/nops"
	"konekko.me/gosion/user/pb/nops"
)

type authService struct {
	pool               *redis.Pool
	configuration      *gs_commons_config.GosionConfiguration
	nopSecurityService gs_nops_service_safety.SecurityService
	nopUserService     gs_nops_service_user.UserService
	connectioncli      connectioncli.ConnectionClient
}

func (svc *authService) GetRepo() *authentication_repositories.TokenRepo {
	return &authentication_repositories.TokenRepo{Conn: svc.pool.Get()}
}

func (svc *authService) Verify(ctx context.Context, in *gs_nops_service_authentication.VerifyRequest, out *gs_commons_dto.Status) error {
	return gs_commons_wrapper.ContextToAuthorize(ctx, out, func(auth *gs_commons_wrapper.WrapperUser) *gs_commons_dto.State {
		//1.verify token

		if len(in.ClientId) == 0 || len(in.Token) == 0 {
			return errstate.ErrRequest
		}

		claims, err := decodeToken(in.Token, svc.configuration.TokenSecretKey)
		if err != nil {
			return errstate.ErrAccessToken
		}

		if claims.Valid() != nil {
			return errstate.ErrAccessTokenExpired
		}

		if claims.Token.Type != gs_commons_constants.AccessToken {
			return errstate.ErrAccessToken
		}

		repo := svc.GetRepo()
		defer repo.Close()

		b, err := repo.Get(claims.Token.UserId, auth.ClientId+"."+claims.Token.Relation)
		if err != nil {
			return errstate.ErrAccessTokenOrClient
		}

		var ui authentication_repositories.UserAuthorizeInfo
		err = msgpack.Unmarshal(b, &ui)
		if err != nil {
			return errstate.ErrSystem
		}

		return nil
	})
}

func NewAuthService() gs_nops_service_authentication.AuthHandler {
	return &authService{}
}
