package modules

import (
	"context"
	"konekko.me/gosion/workflow/flowerr"
	"konekko.me/gosion/workflow/models"
)

type IUser interface {
	IsUserMatch(ctx context.Context, userTask *models.UserTask) *flowerr.Error

	Notify(ctx context.Context, userTask *models.UserTask)
}
