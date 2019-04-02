package permissionhandlers

import (
	"context"
	"github.com/garyburd/redigo/redis"
	"gopkg.in/mgo.v2"
	"konekko.me/gosion/commons/dto"
	"konekko.me/gosion/commons/wrapper"
	"konekko.me/gosion/permission/pb"
)

type structureService struct {
	session *mgo.Session
	pool    *redis.Pool
}

func (svc *structureService) GetRepo() *structureRepo {
	return &structureRepo{conn: svc.pool.Get(), session: svc.session.Clone()}
}

func (svc structureService) Create(ctx context.Context, in *gs_service_permission.CreateRequest, out *gs_commons_dto.Status) error {
	return gs_commons_wrapper.ContextToAuthorize(ctx, out, func(auth *gs_commons_wrapper.WrapperUser) *gs_commons_dto.State {
		return nil
	})
}

func (svc structureService) Enabled(ctx context.Context, in *gs_service_permission.EnabledRequest, out *gs_commons_dto.Status) error {
	return gs_commons_wrapper.ContextToAuthorize(ctx, out, func(auth *gs_commons_wrapper.WrapperUser) *gs_commons_dto.State {
		return nil
	})
}

func (svc structureService) GetList(ctx context.Context, in *gs_service_permission.GetListRequest, out *gs_service_permission.GetListResponse) error {
	return gs_commons_wrapper.ContextToAuthorize(ctx, out, func(auth *gs_commons_wrapper.WrapperUser) *gs_commons_dto.State {
		return nil
	})
}

func NewStructureService(session *mgo.Session, pool *redis.Pool) gs_service_permission.StructureHandler {
	return &structureService{session: session, pool: pool}
}
