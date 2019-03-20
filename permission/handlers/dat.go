package permission_handers

import (
	"context"
	"github.com/garyburd/redigo/redis"
	"konekko.me/gosion/commons/dto"
	"konekko.me/gosion/commons/wrapper"
	"konekko.me/gosion/permission/pb"
	"konekko.me/gosion/permission/repositories"
)

type durationAccessService struct {
	pool *redis.Pool
}

func (svc *durationAccessService) GetRepo() permission_repositories.DurationAccessRepo {
	return permission_repositories.DurationAccessRepo{Conn: svc.pool.Get()}
}

func (svc *durationAccessService) Try(ctx context.Context, in *gs_service_permission.DurationAccessRequest, out *gs_commons_dto.Status) error {
	return gs_commons_wrapper.ContextToAuthorize(ctx, out, func(auth *gs_commons_dto.Authorize) *gs_commons_dto.State {
		return nil
	})
}

func NewDurationAccessService(pool *redis.Pool) gs_service_permission.DurationAccessHandler {
	return &durationAccessService{pool: pool}
}
