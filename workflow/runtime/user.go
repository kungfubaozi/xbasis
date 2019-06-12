package runtime

import (
	"context"
	"konekko.me/gosion/analysis/client"
	"konekko.me/gosion/commons/indexutils"
	"konekko.me/gosion/commons/wrapper"
	"konekko.me/gosion/permission/pb/inner"
	"konekko.me/gosion/user/pb/inner"
	"konekko.me/gosion/workflow/flowerr"
	"konekko.me/gosion/workflow/models"
)

func getWrapperUser(ctx context.Context) *gs_commons_wrapper.WrapperUser {
	return ctx.Value("auth").(*gs_commons_wrapper.WrapperUser)
}

type user struct {
	client            *indexutils.Client
	log               analysisclient.LogClient
	userService       gosionsvc_internal_user.UserService
	permissionService gosionsvc_internal_permission.AccessibleService
}

func (u *user) Notify(ctx context.Context, userTask *models.UserTask) {
	panic("implement me")
}

func (u *user) IsUserMatch(ctx context.Context, userTask *models.UserTask) *flowerr.Error {
	panic("implement me")
}
