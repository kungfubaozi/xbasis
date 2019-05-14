package runtime

import (
	"konekko.me/gosion/workflow/modules"
	"konekko.me/gosion/workflow/types"
)

type trigger struct {
}

func (t *trigger) Register(instanceId, nodeId string, tt types.TriggerType) {
	panic("implement me")
}

func (t *trigger) time() {
	panic("implement me")
}

func (t *trigger) event() {
	panic("implement me")
}

func (t *trigger) Do(tt types.TriggerType, value interface{}) {
	panic("implement me")
}

func newTrigger() modules.Trigger {
	return &trigger{}
}
