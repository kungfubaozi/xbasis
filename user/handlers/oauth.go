package userhandlers

import (
	"context"
	"gopkg.in/mgo.v2"
	"konekko.me/xbasis/analysis/client"
	commons "konekko.me/xbasis/commons/dto"
	"konekko.me/xbasis/commons/indexutils"
	wrapper "konekko.me/xbasis/commons/wrapper"
	userpb "konekko.me/xbasis/user/pb"
)

type oauthService struct {
	session *mgo.Session
	client  *indexutils.Client
	log     analysisclient.LogClient
}

func (svc *oauthService) UnbindOAuth(ctx context.Context, in *userpb.UnbindOAuthRequest, out *commons.Status) error {
	return wrapper.ContextToAuthorize(ctx, out, func(auth *wrapper.WrapperUser) *commons.State {
		return nil
	})
}

func (svc *oauthService) Login(ctx context.Context, in *userpb.OAuthLoginRequest, out *userpb.OAuthLoginResponse) error {
	return wrapper.ContextToAuthorize(ctx, out, func(auth *wrapper.WrapperUser) *commons.State {
		return nil
	})
}

func (svc *oauthService) BindOAuth(ctx context.Context, in *userpb.BindOAuthRequest, out *commons.Status) error {
	return wrapper.ContextToAuthorize(ctx, out, func(auth *wrapper.WrapperUser) *commons.State {
		return nil
	})
}

func NewOAuthService(session *mgo.Session,
	client *indexutils.Client, log analysisclient.LogClient) userpb.OAuthHandler {
	return &oauthService{session: session, client: client, log: log}
}
