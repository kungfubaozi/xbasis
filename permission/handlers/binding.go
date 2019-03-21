package permission_handlers

import (
	"context"
	"github.com/garyburd/redigo/redis"
	"gopkg.in/mgo.v2"
	"konekko.me/gosion/commons/dto"
	"konekko.me/gosion/commons/wrapper"
	"konekko.me/gosion/permission/pb"
)

type bindingService struct {
	pool    *redis.Pool
	session *mgo.Session
}

func (svc *bindingService) UserRole(ctx context.Context, in *gs_service_permission.BindingRoleRequest, out *gs_commons_dto.Status) error {
	return gs_commons_wrapper.ContextToAuthorize(ctx, out, func(auth *gs_commons_wrapper.WrapperUser) *gs_commons_dto.State {
		return nil
	})
}

func (svc *bindingService) FunctionRole(ctx context.Context, in *gs_service_permission.BindingRoleRequest, out *gs_commons_dto.Status) error {
	return gs_commons_wrapper.ContextToAuthorize(ctx, out, func(auth *gs_commons_wrapper.WrapperUser) *gs_commons_dto.State {
		return nil
	})
}

func (svc *bindingService) UnbindUserRole(ctx context.Context, in *gs_service_permission.BindingRoleRequest, out *gs_commons_dto.Status) error {
	return gs_commons_wrapper.ContextToAuthorize(ctx, out, func(auth *gs_commons_wrapper.WrapperUser) *gs_commons_dto.State {
		return nil
	})
}

func (svc *bindingService) UnbindFunctionRole(ctx context.Context, in *gs_service_permission.BindingRoleRequest, out *gs_commons_dto.Status) error {
	return gs_commons_wrapper.ContextToAuthorize(ctx, out, func(auth *gs_commons_wrapper.WrapperUser) *gs_commons_dto.State {
		return nil
	})
}

func NewBindingService(pool *redis.Pool, session *mgo.Session) gs_service_permission.BindingHandler {
	return &bindingService{pool: pool, session: session}
}
