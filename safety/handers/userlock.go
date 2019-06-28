package safetyhanders

import (
	"context"
	"konekko.me/xbasis/analysis/client"
	commons "konekko.me/xbasis/commons/dto"
	external "konekko.me/xbasis/safety/pb"
)

type userlockService struct {
	log analysisclient.LogClient
}

//暂时没有时间做
func (svc *userlockService) Lock(context.Context, *external.UserLockRequest, *commons.Status) error {
	panic("implement me")
}

func (svc *userlockService) Unlock(context.Context, *external.UserUnlockRequest, *commons.Status) error {
	panic("implement me")
}

func NewUserlockService() external.UserlockHandler {
	return &userlockService{}
}
