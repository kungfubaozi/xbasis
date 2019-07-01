package authenticationhandlers

import (
	"context"
	"github.com/garyburd/redigo/redis"
	"github.com/vmihailenco/msgpack"
	"konekko.me/xbasis/analysis/client"
	inner "konekko.me/xbasis/authentication/pb/inner"
	"konekko.me/xbasis/commons/actions"
	"konekko.me/xbasis/commons/config/call"
	constants "konekko.me/xbasis/commons/constants"
	commons "konekko.me/xbasis/commons/dto"
	"konekko.me/xbasis/commons/errstate"
	geneartor "konekko.me/xbasis/commons/generator"
	wrapper "konekko.me/xbasis/commons/wrapper"
	"konekko.me/xbasis/connection/cmd/connectioncli"
	"time"
)

type tokenService struct {
	pool          *redis.Pool
	connectioncli connectioncli.ConnectionClient
	log           analysisclient.LogClient
	id            geneartor.IDGenerator
}

func (svc *tokenService) GetRepo() *tokenRepo {
	return &tokenRepo{conn: svc.pool.Get()}
}

func (svc *tokenService) Generate(ctx context.Context, in *inner.GenerateRequest, out *inner.GenerateResponse) error {
	return wrapper.ContextToAuthorize(ctx, out, func(auth *wrapper.WrapperUser) *commons.State {

		if len(auth.FromClientId) <= 10 {
			return nil
		}

		headers := &analysisclient.LogHeaders{
			TraceId:     auth.TraceId,
			ModuleName:  "Token",
			ServiceName: constants.InternalAuthenticationService,
		}

		repo := svc.GetRepo()
		defer repo.Close()

		configuration := serviceconfiguration.Get()

		//line check
		if s := offlineUser(svc.connectioncli, repo, in.Auth.UserId, in.Auth.ClientId); !s.Ok {
			return s
		}

		id := svc.id

		if len(in.RelationId) <= 20 {

			in.RelationId = id.Get()
		}

		refresh := &simpleUserToken{
			Id:       id.Get(),
			UserId:   in.Auth.UserId,
			AppId:    in.Auth.AppId,
			ClientId: in.Auth.ClientId,
			Relation: in.RelationId,
			Type:     constants.RefreshToken,
		}

		access := &simpleUserToken{
			Id:       id.Get(),
			UserId:   in.Auth.UserId,
			AppId:    in.Auth.AppId,
			ClientId: in.Auth.ClientId,
			Relation: in.RelationId,
			Type:     constants.AccessToken,
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

			svc.log.Info(&analysisclient.LogContent{
				Headers: headers,
				Action:  loggeractions.VisitApplication,
				Message: "generate token",
				Fields: &analysisclient.LogFields{
					"userId":   access.UserId,
					"clientId": access.ClientId,
					"appId":    access.AppId,
					"platform": ui.Platform,
				},
			})

			out.AccessToken = accessToken
			out.RefreshToken = refreshToken
			return errstate.Success
		}

		return nil
	})
}

func NewTokenService(pool *redis.Pool, connectioncli connectioncli.ConnectionClient, log analysisclient.LogClient) inner.TokenHandler {
	return &tokenService{pool: pool, connectioncli: connectioncli, log: log, id: geneartor.NewIDG()}
}
