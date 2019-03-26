package authentication_handlers

import (
	"context"
	"github.com/garyburd/redigo/redis"
	"konekko.me/gosion/authentication/pb/nops"
	"konekko.me/gosion/authentication/repositories"
	"konekko.me/gosion/commons/config"
	"konekko.me/gosion/commons/dto"
	"konekko.me/gosion/commons/generator"
	"konekko.me/gosion/commons/wrapper"
	"konekko.me/gosion/connection/cmd/connectioncli"
)

type tokenService struct {
	pool          *redis.Pool
	configuration *gs_commons_config.GosionConfiguration
	connection    connectioncli.ConnectionClient
}

func (svc *tokenService) GetRepo() *authentication_repositories.TokenRepo {
	return &authentication_repositories.TokenRepo{Conn: svc.pool.Get()}
}

func (svc *tokenService) Generate(ctx context.Context, in *gs_commons_dto.Authorize, out *gs_nops_service_authentication.GenerateResponse) error {
	return gs_commons_wrapper.ContextToAuthorize(ctx, out, func(auth *gs_commons_wrapper.WrapperUser) *gs_commons_dto.State {

		id := gs_commons_generator.NewIDG()

		relationId := id.Get()

		sut := &authentication_repositories.SimpleUserToken{
			UserId:   in.UserId,
			AppId:    in.AppId,
			ClientId: in.ClientId,
			Relation: id.Get(),
		}

		return nil
	})
}

func NewTokenService(pool *redis.Pool) gs_nops_service_authentication.TokenHandler {
	return &tokenService{pool: pool}
}
