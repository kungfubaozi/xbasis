package runtime

import (
	"context"
	"konekko.me/xbasis/analysis/client"
	"konekko.me/xbasis/commons/indexutils"
	wrapper "konekko.me/xbasis/commons/wrapper"
	"konekko.me/xbasis/permission/pb/inner"
	"konekko.me/xbasis/user/pb/inner"
	"konekko.me/xbasis/workflow/flowerr"
	"konekko.me/xbasis/workflow/models"
)

func getWrapperUser(ctx context.Context) *wrapper.WrapperUser {
	return ctx.Value("auth").(*wrapper.WrapperUser)
}

type user struct {
	client            *indexutils.Client
	log               analysisclient.LogClient
	userService       xbasissvc_internal_user.UserService
	permissionService xbasissvc_internal_permission.AccessibleService
}

func (u *user) Notify(ctx context.Context, userTask *models.UserTask) {
	panic("implement me")
}

func (u *user) IsUserMatch(ctx context.Context, userTask *models.UserTask) *flowerr.Error {
	panic("implement me")
}
