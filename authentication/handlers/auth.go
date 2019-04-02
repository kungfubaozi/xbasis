package authenticationhandlers

import (
	"context"
	"github.com/garyburd/redigo/redis"
	"github.com/vmihailenco/msgpack"
	"konekko.me/gosion/authentication/pb/ext"
	"konekko.me/gosion/authentication/repositories"
	"konekko.me/gosion/commons/config"
	"konekko.me/gosion/commons/constants"
	"konekko.me/gosion/commons/dto"
	"konekko.me/gosion/commons/errstate"
	"konekko.me/gosion/commons/wrapper"
	"konekko.me/gosion/connection/cmd/connectioncli"
	"konekko.me/gosion/safety/pb/ext"
)

type authService struct {
	pool               *redis.Pool
	configuration      *gs_commons_config.GosionConfiguration
	extSecurityService gs_ext_service_safety.SecurityService
	connectioncli      connectioncli.ConnectionClient
}

func (svc *authService) GetRepo() *authentication_repositories.TokenRepo {
	return &authentication_repositories.TokenRepo{Conn: svc.pool.Get()}
}

func (svc *authService) Verify(ctx context.Context, in *gs_ext_service_authentication.VerifyRequest, out *gs_commons_dto.Status) error {
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

		var uai authentication_repositories.UserAuthorizeInfo
		err = msgpack.Unmarshal(b, &uai)
		if err != nil {
			return errstate.ErrSystem
		}

		//check
		if claims.Token.UserId != uai.UserId || claims.Token.ClientId != uai.ClientId ||
			claims.Token.Relation != uai.Relation || claims.Token.AppId != uai.AppId ||
			uai.Device != auth.UserDevice || uai.UserAgent != auth.UserAgent ||
			auth.ClientId != uai.ClientId || auth.User != uai.UserId || auth.AppId != uai.AppId {
			return errstate.ErrAccessToken
		}

		s, err := svc.extSecurityService.Get(ctx, &gs_ext_service_safety.GetRequest{UserId: uai.UserId})
		if err != nil {
			return errstate.ErrSystem
		}

		return s.State
	})
}

func NewAuthService() gs_ext_service_authentication.AuthHandler {
	return &authService{}
}
