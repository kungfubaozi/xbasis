package history

type Interface interface {
	GetInstanceNodeHistory(instanceId, nodeId string)

	GetInstanceOperateHistory(instanceId string)

	GetInstanceStatus()
}

type History struct {
}

func (h *History) GetInstanceNodeHistory(instanceId, nodeId string) {
	panic("implement me")
}

func (h *History) GetInstanceOperateHistory(instanceId string) {
	panic("implement me")
}

func (h *History) GetInstanceStatus() {
	panic("implement me")
}
