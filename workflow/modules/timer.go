package modules

import "konekko.me/xbasis/workflow/flowerr"

type Timer interface {
	Register(corn string, instanceId, nodeId string) *flowerr.Error
}
