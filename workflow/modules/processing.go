package modules

import (
	"context"
	"konekko.me/gosion/workflow/types"
)

type processing struct {
	modules    Modules
	instanceId string
	node       *node
	value      interface{}
}

func (n *processing) TriggerStartEvent() {
	panic("implement me")
}

func (n *processing) Do(ctx context.Context, instanceId string, node *node, ct types.ConnectType, value interface{}) {
	n.instanceId = instanceId
	n.value = value
	n.node = node
	distribute(ct, n)
}

func (n *processing) Data() interface{} {
	panic("implement me")
}

func (n *processing) ExclusiveGateway() {
	panic("implement me")
}

func (n *processing) ParallelGateway() {
	panic("implement me")
}

func (n *processing) InclusiveGateway() {
	panic("implement me")
}

func (n *processing) StartEvent() {
	panic("implement me")
}

func (n *processing) EndEvent() {
	panic("implement me")
}

func (n *processing) UserTask() {
	panic("implement me")
}

func (n *processing) Restore() {
	panic("implement me")
}

func newProcessing(modules Modules) distribution {
	return &processing{modules: modules}
}
