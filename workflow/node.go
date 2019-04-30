package workflow

type flowNodes interface {
	UserTask(flow *sequenceFlow, node *node)

	HttpTask(flow *sequenceFlow, node *node)

	DecisionTask(flow *sequenceFlow, node *node)

	SendTask(flow *sequenceFlow, node *node)

	GRPCTask(flow *sequenceFlow, node *node)

	MailTask(flow *sequenceFlow, node *node)

	StartEvent(flow *sequenceFlow, node *node)

	TimerStartEvent(flow *sequenceFlow, node *node)

	MessageStartEvent(flow *sequenceFlow, node *node)

	EndEvent(flow *sequenceFlow, node *node)

	EndErrorEvent(flow *sequenceFlow, node *node)

	EndCancelEvent(flow *sequenceFlow, node *node)

	TerminateEvent(flow *sequenceFlow, node *node)

	ExclusiveGateway(flow *sequenceFlow, node *node)

	ParallelGateway(flow *sequenceFlow, node *node)

	InclusiveGateway(flow *sequenceFlow, node *node)

	EventGateway(flow *sequenceFlow, node *node)

	Process(flows []*sequenceFlow, node *node)
}

type nextFlowControl struct {
	ins     *instance
	pip     *pipeline
	records map[string]bool
	end     bool
}

func (n *nextFlowControl) Process(flows []*sequenceFlow, node *node) {
	for _, v := range flows {
		//find the node that connects to the end of the flow
		cn := n.pip.nodes[v.End]
		switch v.EndType {
		case CTDecisionTask:
			n.DecisionTask(v, cn)
			break
		case CTGRPCTask:
			n.GRPCTask(v, cn)
			break
		case CTHttpTask:
			n.HttpTask(v, cn)
			break
		case CTMailTask:
			n.MailTask(v, cn)
			break
		case CTSendTask:
			n.SendTask(v, cn)
			break
		case CTUserTask:
			n.UserTask(v, cn)
			break
		case CTEndCancelEvent:
			n.EndCancelEvent(v, cn)
			break
		case CTEndErrorEvent:
			n.EndErrorEvent(v, cn)
			break
		case CTEndEvent:
			n.EndEvent(v, cn)
			break
		case CTTerminateEvent:
			n.TerminateEvent(v, cn)
			break
		case CTEventGateway:
			n.EventGateway(v, cn)
			break
		case CTExclusiveGateway:
			n.ExclusiveGateway(v, cn)
			break
		case CTInclusiveGateway:
			n.InclusiveGateway(v, cn)
			break
		case CTParallelGateway:
			n.ParallelGateway(v, cn)
			break
		case CTStartEvent:
			n.StartEvent(v, cn)
			break
		case CTMessageStartEvent:
			n.MessageStartEvent(v, cn)
			break
		case CTTimerStartEvent:
			n.TimerStartEvent(v, cn)
			break
		}
	}
}

func (n *nextFlowControl) processed(id string) bool {
	return n.records[id]
}

func (n *nextFlowControl) UserTask(flow *sequenceFlow, node *node) {
	if !n.processed(node.id) {

	}
}

func (n *nextFlowControl) HttpTask(flow *sequenceFlow, node *node) {
	if !n.processed(node.id) {

	}
}

func (n *nextFlowControl) DecisionTask(flow *sequenceFlow, node *node) {
	if !n.processed(node.id) {

	}
}

func (n *nextFlowControl) SendTask(flow *sequenceFlow, node *node) {
	if !n.processed(node.id) {

	}
}

func (n *nextFlowControl) GRPCTask(flow *sequenceFlow, node *node) {
	if !n.processed(node.id) {

	}
}

func (n *nextFlowControl) MailTask(flow *sequenceFlow, node *node) {
	if !n.processed(node.id) {

	}
}

func (n *nextFlowControl) StartEvent(flow *sequenceFlow, node *node) {
	panic("implement me")
}

func (n *nextFlowControl) TimerStartEvent(flow *sequenceFlow, node *node) {
	panic("implement me")
}

func (n *nextFlowControl) MessageStartEvent(flow *sequenceFlow, node *node) {
	panic("implement me")
}

func (n *nextFlowControl) EndEvent(flow *sequenceFlow, node *node) {
	panic("implement me")
}

func (n *nextFlowControl) EndErrorEvent(flow *sequenceFlow, node *node) {
	panic("implement me")
}

func (n *nextFlowControl) EndCancelEvent(flow *sequenceFlow, node *node) {
	panic("implement me")
}

func (n *nextFlowControl) TerminateEvent(flow *sequenceFlow, node *node) {
	panic("implement me")
}

func (n *nextFlowControl) ExclusiveGateway(flow *sequenceFlow, node *node) {
	panic("implement me")
}

func (n *nextFlowControl) ParallelGateway(flow *sequenceFlow, node *node) {
	panic("implement me")
}

func (n *nextFlowControl) InclusiveGateway(flow *sequenceFlow, node *node) {
	panic("implement me")
}

func (n *nextFlowControl) EventGateway(flow *sequenceFlow, node *node) {
	panic("implement me")
}
