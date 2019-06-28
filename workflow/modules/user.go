package modules

import (
	"context"
	"konekko.me/xbasis/workflow/flowerr"
	"konekko.me/xbasis/workflow/models"
)

type IUser interface {
	IsUserMatch(ctx context.Context, userTask *models.UserTask) *flowerr.Error

	Notify(ctx context.Context, userTask *models.UserTask)
}
