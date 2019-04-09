package authenticationhandlers

import (
	"context"
	"github.com/garyburd/redigo/redis"
	"github.com/vmihailenco/msgpack"
	"konekko.me/gosion/authentication/pb/ext"
	"konekko.me/gosion/commons/config"
	"konekko.me/gosion/commons/constants"
	"konekko.me/gosion/commons/dto"
	"konekko.me/gosion/commons/errstate"
	"konekko.me/gosion/commons/generator"
	"konekko.me/gosion/commons/wrapper"
	"konekko.me/gosion/connection/cmd/connectioncli"
	"time"
)

type tokenService struct {
	pool          *redis.Pool
	configuration *gs_commons_config.GosionConfiguration
	connectioncli connectioncli.ConnectionClient
}

func (svc *tokenService) GetRepo() *tokenRepo {
	return &tokenRepo{conn: svc.pool.Get()}
}

func (svc *tokenService) Generate(ctx context.Context, in *gs_ext_service_authentication.GenerateRequest, out *gs_ext_service_authentication.GenerateResponse) error {
	return gs_commons_wrapper.ContextToAuthorize(ctx, out, func(auth *gs_commons_wrapper.WrapperUser) *gs_commons_dto.State {

		repo := svc.GetRepo()
		defer repo.Close()

		//line check
		if s := sizeCheck(svc.connectioncli, repo, in.Auth.UserId, in.Auth.ClientId); !s.Ok {
			return s
		}

		id := gs_commons_generator.NewIDG()

		relationId := id.Get()

		refresh := &simpleUserToken{
			UserId:   in.Auth.UserId,
			AppId:    in.Auth.AppId,
			ClientId: in.Auth.ClientId,
			Relation: relationId,
			Type:     gs_commons_constants.RefreshToken,
		}

		access := &simpleUserToken{
			UserId:   in.Auth.UserId,
			AppId:    in.Auth.AppId,
			ClientId: in.Auth.ClientId,
			Relation: relationId,
			Type:     gs_commons_constants.AccessToken,
		}

		ui := &userAuthorizeInfo{
			AppId:     in.Auth.AppId,
			Platform:  in.Auth.Platform,
			Relation:  relationId,
			UserId:    in.Auth.UserId,
			Device:    in.Auth.Device,
			UserAgent: in.Auth.UserAgent,
		}

		b, err := msgpack.Marshal(ui)
		if err != nil {
			return errstate.ErrSystem
		}

		err = repo.Add(in.Auth.UserId, in.Auth.ClientId, relationId, b)
		if err != nil {
			return errstate.ErrSystem
		}

		refreshToken, err := encodeToken(svc.configuration.TokenSecretKey, time.Hour*24*7, refresh)
		if err != nil {
			return errstate.ErrSystem
		}

		accessToken, err := encodeToken(svc.configuration.TokenSecretKey, time.Minute*10, access)
		if err != nil {
			return errstate.ErrSystem
		}

		if len(refreshToken) > 0 && len(accessToken) > 0 {
			out.AccessToken = accessToken
			out.RefreshToken = refreshToken
			out.State = errstate.Success
			return nil
		}

		return nil
	})
}

func NewTokenService(pool *redis.Pool) gs_ext_service_authentication.TokenHandler {
	return &tokenService{pool: pool}
}
