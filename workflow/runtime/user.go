package runtime

import (
	"context"
	"konekko.me/gosion/commons/gslogrus"
	"konekko.me/gosion/commons/indexutils"
	"konekko.me/gosion/commons/wrapper"
	"konekko.me/gosion/permission/pb/ext"
	"konekko.me/gosion/user/pb/ext"
	"konekko.me/gosion/workflow/flowerr"
	"konekko.me/gosion/workflow/models"
)

func getWrapperUser(ctx context.Context) *gs_commons_wrapper.WrapperUser {
	return ctx.Value("auth").(*gs_commons_wrapper.WrapperUser)
}

type user struct {
	client            *indexutils.Client
	log               *gslogrus.Logger
	userService       gs_ext_service_user.UserService
	permissionService gs_ext_service_permission.AccessibleService
}

func (u *user) Notify(ctx context.Context, userTask *models.UserTask) {
	panic("implement me")
}

func (u *user) IsUserMatch(ctx context.Context, userTask *models.UserTask) *flowerr.Error {
	panic("implement me")
}
