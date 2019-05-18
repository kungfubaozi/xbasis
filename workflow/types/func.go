package types

import (
	"konekko.me/gosion/workflow/flowerr"
)

type GetterCommand int64

const (
	GCBackwardRelations GetterCommand = iota //获取此节点数据流执行后面的有效节点

	GCNodeFlows //获取节点关联的flow

	GCNodeSubmitData //获取节点提交的数据

	GCNode //获取对应节点

	GCForwardRelationNodes //获取此节点数据流执行前面的有效节点
)

type ErrCallback func() *flowerr.Error

type CommandDataGetter func(command GetterCommand, values ...interface{}) (interface{}, *flowerr.Error)
