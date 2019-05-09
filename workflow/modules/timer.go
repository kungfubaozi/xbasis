package modules

import "konekko.me/gosion/workflow/flowerr"

type Timer interface {
	Register(corn string, instanceId, nodeId string) *flowerr.Error
}
