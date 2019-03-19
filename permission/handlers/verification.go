package permission_handers

import (
	"context"
	"github.com/garyburd/redigo/redis"
	"gopkg.in/mgo.v2"
	"konekko.me/gosion/commons/dto"
	"konekko.me/gosion/commons/wrapper"
	"konekko.me/gosion/permission/pb"
)

type verificationService struct {
	pool    *redis.Pool
	session *mgo.Session
}

func (svc *verificationService) Test(ctx context.Context, in *gs_service_permission.HasPermissionRequest, out *gs_commons_dto.Status) error {
	return gs_commons_wrapper.ContextToAuthorize(ctx, out, func(auth *gs_commons_dto.Authorize) *gs_commons_dto.State {
		return nil
	})
}

func NewVerificationService(pool *redis.Pool, session *mgo.Session) gs_service_permission.VerificationHandler {
	return &verificationService{pool: pool, session: session}
}
