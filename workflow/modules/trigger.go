package modules

import "konekko.me/gosion/workflow/types"

type Trigger interface {
	Register(instanceId, nodeId string, tt types.TriggerType)

	time()

	event()

	Do(tt types.TriggerType, value interface{})
}
