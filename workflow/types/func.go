package types

import (
	"konekko.me/gosion/workflow/flowerr"
)

type GetterCommand int64

const (
	GCBackwardRelations GetterCommand = iota

	GCNodeFlows

	GCNodeSubmitData

	GCNode
)

type ErrCallback func() *flowerr.Error

type CommandDataGetter func(command GetterCommand, values ...interface{}) (interface{}, *flowerr.Error)
