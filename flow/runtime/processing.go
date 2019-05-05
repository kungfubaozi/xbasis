package runtime

import (
	"context"
	"konekko.me/gosion/flow/types"
)

type processing struct {
	r          *runtime
	instanceId string
	nodeId     string
	ct         types.ConnectType
	value      map[string]interface{}
}

func (t *processing) ExclusiveGateway() {
	panic("implement me")
}

func (t *processing) ParallelGateway() {
	panic("implement me")
}

func (t *processing) InclusiveGateway() {
	panic("implement me")
}

func (t *processing) Do(ctx context.Context, instanceId, nodeId string, ct types.ConnectType, value map[string]interface{}) {
	t.instanceId = instanceId
	t.nodeId = nodeId
	t.ct = ct
	t.value = value
	switch ct {
	case types.CTStartEvent:
		t.StartEvent()
		break
	case types.CTEndEvent:
		t.EndEvent()
		break
	case types.CTUserTask:
		t.UserTask()
		break
	case types.CTExclusiveGateway:
		break
	case types.CTParallelGateway:
		break
	case types.CTInclusiveGateway:
		break
	}
}

func (t *processing) StartEvent() {
	panic("implement me")
}

func (t *processing) EndEvent() {
	panic("implement me")
}

func (t *processing) UserTask() {
	panic("implement me")
}

func newProcessingTask(r *runtime) types.TypeTasks {
	return &processing{r: r}
}
