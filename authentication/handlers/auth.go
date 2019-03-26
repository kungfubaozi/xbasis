package authentication_handlers

import (
	"context"
	"github.com/garyburd/redigo/redis"
	"konekko.me/gosion/authentication/pb/nops"
	"konekko.me/gosion/commons/dto"
	"konekko.me/gosion/commons/wrapper"
	"konekko.me/gosion/safety/pb/nops"
	"konekko.me/gosion/user/pb/nops"
)

type authService struct {
	pool            *redis.Pool
	securityService gs_nops_service_safety.SecurityService
	nopUserService  gs_nops_service_user.UserService
}

func (svc *authService) Verify(ctx context.Context, in *gs_nops_service_authentication.VerifyRequest, out *gs_commons_dto.Status) error {
	return gs_commons_wrapper.ContextToAuthorize(ctx, out, func(auth *gs_commons_wrapper.WrapperUser) *gs_commons_dto.State {
		return nil
	})
}

func NewAuthService() gs_nops_service_authentication.AuthHandler {
	return &authService{}
}
