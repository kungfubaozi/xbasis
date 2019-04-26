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
	"konekko.me/gosion/commons/generator"
	"konekko.me/gosion/commons/wrapper"
	"konekko.me/gosion/connection/cmd/connectioncli"
	"time"
)

type tokenService struct {
	pool          *redis.Pool
	connectioncli connectioncli.ConnectionClient
}

func (svc *tokenService) GetRepo() *tokenRepo {
	return &tokenRepo{conn: svc.pool.Get()}
}

func (svc *tokenService) Generate(ctx context.Context, in *gs_ext_service_authentication.GenerateRequest, out *gs_ext_service_authentication.GenerateResponse) error {
	return gs_commons_wrapper.ContextToAuthorize(ctx, out, func(auth *gs_commons_wrapper.WrapperUser) *gs_commons_dto.State {

		if len(auth.ClientId) <= 10 {
			return nil
		}

		repo := svc.GetRepo()
		defer repo.Close()

		configuration := serviceconfiguration.Get()

		//line check
		if s := offlineUser(svc.connectioncli, repo, in.Auth.UserId, in.Auth.ClientId); !s.Ok {
			return s
		}

		id := gs_commons_generator.NewIDG()

		if len(in.RelationId) <= 30 {

			in.RelationId = id.Get()
		}

		refresh := &simpleUserToken{
			Id:       id.Get(),
			UserId:   in.Auth.UserId,
			AppId:    in.Auth.AppId,
			ClientId: in.Auth.ClientId,
			Relation: in.RelationId,
			Type:     gs_commons_constants.RefreshToken,
		}

		access := &simpleUserToken{
			Id:       id.Get(),
			UserId:   in.Auth.UserId,
			AppId:    in.Auth.AppId,
			ClientId: in.Auth.ClientId,
			Relation: in.RelationId,
			Type:     gs_commons_constants.AccessToken,
		}

		ui := &userAuthorizeInfo{
			AppId:     in.Auth.AppId,
			Platform:  in.Auth.Platform,
			Relation:  in.RelationId,
			UserId:    in.Auth.UserId,
			Device:    in.Auth.Device,
			UserAgent: in.Auth.UserAgent,
			ClientId:  in.Auth.ClientId,
			Ip:        in.Auth.Ip,
			RefreshId: refresh.Id,
			AccessId:  access.Id,
		}

		b, err := msgpack.Marshal(ui)
		if err != nil {
			return errstate.ErrSystem
		}

		err = repo.Add(in.Auth.UserId, in.Auth.ClientId, in.RelationId, b)
		if err != nil {
			return errstate.ErrSystem
		}

		refreshToken, err := encodeToken(configuration.TokenSecretKey, time.Hour*24*3, refresh)
		if err != nil {
			return errstate.ErrSystem
		}

		accessToken, err := encodeToken(configuration.TokenSecretKey, time.Minute*10, access)
		if err != nil {
			return errstate.ErrSystem
		}

		if len(refreshToken) > 0 && len(accessToken) > 0 {
			out.AccessToken = accessToken
			out.RefreshToken = refreshToken
			return errstate.Success
		}

		return nil
	})
}

func NewTokenService(pool *redis.Pool, connectioncli connectioncli.ConnectionClient) gs_ext_service_authentication.TokenHandler {
	return &tokenService{pool: pool, connectioncli: connectioncli}
}
