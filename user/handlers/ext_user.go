package userhandlers

import (
	"context"
	"konekko.me/gosion/commons/dto"
	"konekko.me/gosion/user/pb/ext"
)

type extUserService struct {
}

func (svc *extUserService) IsExists(context.Context, *gs_ext_service_user.ExistsRequest, *gs_commons_dto.Status) error {
	panic("implement me")
}

func NewExtUserService() gs_ext_service_user.UserHandler {
	return &extUserService{}
}
