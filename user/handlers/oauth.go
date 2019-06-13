package userhandlers

import (
	"context"
	"gopkg.in/mgo.v2"
	"konekko.me/gosion/commons/dto"
	"konekko.me/gosion/commons/wrapper"
	"konekko.me/gosion/user/pb"
)

type oauthService struct {
	session *mgo.Session
}

func (svc *oauthService) UnbindOAuth(ctx context.Context, in *gosionsvc_external_user.UnbindOAuthRequest, out *gs_commons_dto.Status) error {
	return gs_commons_wrapper.ContextToAuthorize(ctx, out, func(auth *gs_commons_wrapper.WrapperUser) *gs_commons_dto.State {
		return nil
	})
}

func (svc *oauthService) Login(ctx context.Context, in *gosionsvc_external_user.OAuthLoginRequest, out *gosionsvc_external_user.OAuthLoginResponse) error {
	return gs_commons_wrapper.ContextToAuthorize(ctx, out, func(auth *gs_commons_wrapper.WrapperUser) *gs_commons_dto.State {
		return nil
	})
}

func (svc *oauthService) BindOAuth(ctx context.Context, in *gosionsvc_external_user.BindOAuthRequest, out *gs_commons_dto.Status) error {
	return gs_commons_wrapper.ContextToAuthorize(ctx, out, func(auth *gs_commons_wrapper.WrapperUser) *gs_commons_dto.State {
		return nil
	})
}

func NewOAuthService() gosionsvc_external_user.OAuthHandler {
	return &oauthService{}
}
