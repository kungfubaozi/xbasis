package safetyhanders

import (
	"context"
	"konekko.me/gosion/analysis/client"
	"konekko.me/gosion/commons/dto"
	external "konekko.me/gosion/safety/pb"
)

type userlockService struct {
	log analysisclient.LogClient
}

//暂时没有时间做
func (svc *userlockService) Lock(context.Context, *external.UserLockRequest, *gs_commons_dto.Status) error {
	panic("implement me")
}

func (svc *userlockService) Unlock(context.Context, *external.UserUnlockRequest, *gs_commons_dto.Status) error {
	panic("implement me")
}

func NewUserlockService() external.UserlockHandler {
	return &userlockService{}
}
