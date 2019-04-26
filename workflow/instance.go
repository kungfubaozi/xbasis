package workflow

import (
	"github.com/pkg/errors"
)

//single
type instance struct {
	script   *luaScript
	id       string
	scripts  map[string]string
	forms    map[string]*typeForm
	running  bool
	start    chan bool
	nextTask chan bool
	reload   chan bool
	index    int //当前位置

	tasks map[string]interface{}

	gateways map[string]interface{}

	events map[string]interface{}

	flows [][]*sequenceFlow

	whoStarted string

	resp chan interface{}
}

type flownode struct {
	start  bool
	data   interface{}
	ft     ConnectionType
	script string
	i      *instance
}

func (i *instance) lookupForwardTaskData(end string) {
	var startTasks []string
	for _, v := range i.flows {
		for _, sv := range v {
			if sv.End == end {
				startTasks = append(startTasks, sv.Start)
			}
		}
	}
}

func (i *instance) Start() {
	for {
		select {
		case s := <-i.start:
			if !s {
				i.running = false
				return
			}
			i.LockInstance()
			break
		case s := <-i.nextTask:
			if s && i.running {
				i.index = i.index + 1

				var waitFlows []*sequenceFlow
				for _, v := range i.flows[i.index] {

					sf := v
					startConn := i.nodes[sf.Start]
					endConn := i.nodes[sf.End]

					if v.StartType == FTExclusiveGateway { //如果都为true，则选择优先级高的执行

					} else if v.StartType == FTParallelGateway {

					} else if v.StartType == FTInclusiveGateway {

					} else if v.StartType == FTEventGateway {

					}

					startFN := &flownode{
						start: true,
						i:     i,
						data:  startConn,
						ft:    sf.StartType,
					}

					sfns := startFN.Swi()

					if sfns.Ok {

						endFN := &flownode{
							start: false,
							i:     i,
							data:  endConn,
							ft:    sf.EndType,
						}

						endFN.Swi()

						break
					}

					i.resp <- sfns

				}
			}

			break

		case s := <-i.reload:
			if s {

			}
			break
		}

	}
}

func (i *instance) LockInstance() {

}

func (i *instance) SaveFlowStatus() {

}

func (fn *flownode) Swi() *state {
	if fn.ft == FTStartEvent {
		return fn.startEvent()
	} else if fn.ft == FTMessageStartEvent {
		return fn.messageStartEvent()
	} else if fn.ft == FTTimerStartEvent {
		return fn.timerStartEvent()
	} else if fn.ft == FTUserTask {
		return fn.userTask()
	} else if fn.ft == FTDecisionTask {
		return fn.decisionTask()
	} else if fn.ft == FTHttpTask {
		return fn.httpTask()
	} else if fn.ft == FTHeader {
		return fn.process()
	} else if fn.ft == FTEndEvent {
		return fn.endEvent()
	} else if fn.ft == FTEndCancelEvent {
		return fn.endCancelEvent()
	} else if fn.ft == FTEndErrorEvent {
		return fn.endErrorEvent()
	} else if fn.ft == FTTerminateEvent {
		return fn.terminateEvent()
	} else {
		panic(errors.New("err type."))
	}
}

func (fn *flownode) process() *state {

}

func (fn *flownode) startEvent() *state {
	//检查是否有form存在
	//获取谁发起的流程
	se := fn.data.(*startEvent)
	if len(se.FormRef) > 0 {

	}
	return Success
}

func (fn *flownode) messageStartEvent() *state {

}

func (fn *flownode) timerStartEvent() *state {

}

func (fn *flownode) userTask() *state {

}

func (fn *flownode) httpTask() *state {

}

func (fn *flownode) decisionTask() *state {

}

func (fn *flownode) endEvent() *state {

}

func (fn *flownode) endCancelEvent() *state {

}

func (fn *flownode) endErrorEvent() *state {

}

func (fn *flownode) terminateEvent() *state {

}

func (fn *flownode) testScript() *state {
	if len(fn.script) > 0 {

	}
}
