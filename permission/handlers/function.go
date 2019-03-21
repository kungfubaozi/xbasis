package permission_handlers

import (
	"context"
	"github.com/garyburd/redigo/redis"
	"gopkg.in/mgo.v2"
	"konekko.me/gosion/commons/constants"
	"konekko.me/gosion/commons/dto"
	"konekko.me/gosion/commons/encrypt"
	"konekko.me/gosion/commons/errstate"
	"konekko.me/gosion/commons/generator"
	"konekko.me/gosion/commons/wrapper"
	"konekko.me/gosion/permission/pb"
	"konekko.me/gosion/permission/repositories"
	"time"
)

type functionService struct {
	pool    *redis.Pool
	session *mgo.Session
}

func (svc *functionService) GetRepo() permission_repositories.FunctionRepo {
	return permission_repositories.FunctionRepo{Session: svc.session.Clone(), Conn: svc.pool.Get()}
}

func (svc *functionService) Add(ctx context.Context, in *gs_service_permission.FunctionRequest, out *gs_commons_dto.Status) error {
	return gs_commons_wrapper.ContextToAuthorize(ctx, out, func(auth *gs_commons_wrapper.WrapperUser) *gs_commons_dto.State {

		repo := svc.GetRepo()
		defer repo.Close()

		if len(in.AppId) == 0 && len(in.BindGroupId) == 0 && len(in.Name) == 0 {
			return errstate.ErrRequest
		}

		//find group exists
		if repo.FindGroupExits(in.BindGroupId) {
			return errstate.ErrFunctionBindGroupId
		}

		//check authTypes
		if in.AuthTypes != nil && len(in.AuthTypes) > 0 {
			for _, v := range in.AuthTypes {
				switch v {
				case gs_commons_constants.AuthTypeOfFace:
				case gs_commons_constants.AuthTypeOfToken:
				case gs_commons_constants.AuthTypeOfValcode:
				case gs_commons_constants.AuthTypeOfMobileConfirm:
				default:
					return errstate.ErrFunctionAuthType
				}
			}
		}

		_, err := repo.FindApi(in.AppId, in.Api)
		if err != nil && err == mgo.ErrNotFound {

			f := &permission_repositories.Function{
				Id:           gs_commons_generator.ID().Generate().String(),
				Name:         in.Name,
				Type:         in.Type,
				CreateUserId: auth.User,
				CreateAt:     time.Now().UnixNano(),
				BindGroupId:  in.BindGroupId,
				AppId:        in.AppId,
				ApiTag:       gs_commons_encrypt.SHA1(in.Api),
				Api:          in.Api,
				AuthTypes:    in.AuthTypes,
			}

			err := repo.AddFunction(f)

			if err == nil {
				return errstate.Success
			}

		}

		return nil
	})
}

func (svc *functionService) Rename(ctx context.Context, in *gs_service_permission.FunctionRequest, out *gs_commons_dto.Status) error {
	return gs_commons_wrapper.ContextToAuthorize(ctx, out, func(auth *gs_commons_wrapper.WrapperUser) *gs_commons_dto.State {
		return nil
	})
}

func (svc *functionService) Move(ctx context.Context, in *gs_service_permission.FunctionRequest, out *gs_commons_dto.Status) error {
	return gs_commons_wrapper.ContextToAuthorize(ctx, out, func(auth *gs_commons_wrapper.WrapperUser) *gs_commons_dto.State {
		return nil
	})
}

//one application one root group, bindGroupId = appId
func (svc *functionService) AddGroup(ctx context.Context, in *gs_service_permission.FunctionGroupRequest, out *gs_commons_dto.Status) error {
	return gs_commons_wrapper.ContextToAuthorize(ctx, out, func(auth *gs_commons_wrapper.WrapperUser) *gs_commons_dto.State {

		if len(in.Name) == 0 && len(in.AppId) == 0 {
			return nil
		}

		return nil
	})
}

func (svc *functionService) MoveGroup(ctx context.Context, in *gs_service_permission.FunctionGroupRequest, out *gs_commons_dto.Status) error {
	return gs_commons_wrapper.ContextToAuthorize(ctx, out, func(auth *gs_commons_wrapper.WrapperUser) *gs_commons_dto.State {
		return nil
	})
}

func (svc *functionService) RenameGroup(ctx context.Context, in *gs_service_permission.FunctionGroupRequest, out *gs_commons_dto.Status) error {
	return gs_commons_wrapper.ContextToAuthorize(ctx, out, func(auth *gs_commons_wrapper.WrapperUser) *gs_commons_dto.State {
		return nil
	})
}

func NewFunctionService(pool *redis.Pool, session *mgo.Session) gs_service_permission.FunctionHandler {
	return &functionService{pool: pool, session: session}
}
