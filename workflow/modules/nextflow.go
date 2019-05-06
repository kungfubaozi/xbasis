package modules

import (
	"context"
	"konekko.me/gosion/workflow/models"
	"konekko.me/gosion/workflow/types"
)

type nextflow struct {
	modules  Modules
	finished map[string]bool
	value    interface{}
	status   *nextstatus
}

type nextstatus struct {
	again        string
	currentNodes []string
}

func (n *nextflow) TriggerStartEvent() {
	panic("implement me")
}

func (n *nextflow) Do(ctx context.Context, instanceId string, node *node, ct types.ConnectType, value interface{}) {
	n.finished = make(map[string]bool)
	n.value = value
	n.status = &nextstatus{}
	distribute(ct, n)
}

func (n *nextflow) Data() interface{} {
	panic("implement me")
}

func (n *nextflow) ExclusiveGateway() {
	panic("implement me")
}

func (n *nextflow) ParallelGateway() {
	panic("implement me")
}

func (n *nextflow) InclusiveGateway() {
	panic("implement me")
}

func (n *nextflow) StartEvent() {
	panic("implement me")
}

func (n *nextflow) EndEvent() {
	panic("implement me")
}

func (n *nextflow) UserTask() {
	panic("implement me")
}

func (n *nextflow) Restore() {
	panic("implement me")
}

func (n *nextflow) flow() *models.SequenceFlow {
	return n.value.(*models.SequenceFlow)
}

func newNextflow(modules Modules) distribution {
	return &nextflow{modules: modules}
}
