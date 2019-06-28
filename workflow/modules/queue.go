package modules

import "konekko.me/xbasis/workflow/flowerr"

type EventQueue interface {
	Register(event, instanceId, nodeId string) *flowerr.Error
}
