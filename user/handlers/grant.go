package userhandlers

import (
	"context"
	"github.com/garyburd/redigo/redis"
	"gopkg.in/mgo.v2"
	"konekko.me/gosion/commons/dto"
	"konekko.me/gosion/commons/wrapper"
	"konekko.me/gosion/user/pb"
)

type grantService struct {
	session *mgo.Session
	pool    *redis.Pool
}

func (svc *grantService) Grant(ctx context.Context, in *gs_service_user.GrantRequest, out *gs_commons_dto.Status) error {
	return gs_commons_wrapper.ContextToAuthorize(ctx, out, func(auth *gs_commons_wrapper.WrapperUser) *gs_commons_dto.State {

		return nil
	})
}

func (svc *grantService) Status(ctx context.Context, in *gs_service_user.StatusRequest, out *gs_commons_dto.Status) error {
	return gs_commons_wrapper.ContextToAuthorize(ctx, out, func(auth *gs_commons_wrapper.WrapperUser) *gs_commons_dto.State {

		return nil
	})
}

func NewGrantService(session *mgo.Session, pool *redis.Pool) gs_service_user.ApplicationGrantHandler {
	return &grantService{session: session, pool: pool}
}
