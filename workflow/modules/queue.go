package modules

import "konekko.me/gosion/workflow/flowerr"

type EventQueue interface {
	Register(event, instanceId, nodeId string) *flowerr.Error
}
