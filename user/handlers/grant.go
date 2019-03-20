package user_handlers

import (
	"context"
	"github.com/garyburd/redigo/redis"
	"gopkg.in/mgo.v2"
	"konekko.me/gosion/commons/dto"
	"konekko.me/gosion/user/pb"
)

type grantService struct {
	session *mgo.Session
	pool    *redis.Pool
}

func (svc *grantService) Grant(ctx context.Context, in *gs_service_user.GrantRequest, out *gs_commons_dto.Status) error {
	panic("implement me")
}

func (svc *grantService) Status(ctx context.Context, in *gs_service_user.StatusRequest, out *gs_commons_dto.Status) error {
	panic("implement me")
}

func NewGrantService(session *mgo.Session, pool *redis.Pool) gs_service_user.ApplicationGrantHandler {
	return &grantService{session: session, pool: pool}
}
