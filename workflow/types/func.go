package types

import (
	"context"
	"konekko.me/gosion/workflow/flowerr"
)

type ErrCallback func() (context.Context, *flowerr.Error)
