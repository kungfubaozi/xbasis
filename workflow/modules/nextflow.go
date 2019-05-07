package modules

import (
	"context"
	"konekko.me/gosion/commons/dto"
	"konekko.me/gosion/commons/errstate"
	"konekko.me/gosion/commons/gslogrus"
	"konekko.me/gosion/workflow/flowstate"
	"konekko.me/gosion/workflow/models"
	"konekko.me/gosion/workflow/script"
	"konekko.me/gosion/workflow/types"
	"sync"
)

//check next flow conditions, set currentNodes to process status.

type nextflow struct {
	modules  Modules
	finished map[string]bool
	log      *gslogrus.Logger
	values   []interface{}
	status   *nextstatus
	node     *node
	ctx      context.Context
	instance *models.Instance
	script   distribution
}

type nextstatus struct {
	again        string
	currentNodes []string
}

func (n *nextflow) Do(ctx context.Context, instance *models.Instance, node *node, ct types.ConnectType, value ...interface{}) (*gs_commons_dto.State, error) {
	n.values = value
	n.node = node
	n.ctx = ctx
	n.instance = instance
	n.status = &nextstatus{}
	return distribute(ct, n)
}

func (n *nextflow) Data() interface{} {
	return n.status
}

func (n *nextflow) ExclusiveGateway() (*gs_commons_dto.State, error) {
	flow := n.flow()
	state, err := n.script.Do(n.ctx, n.instance, n.node, flow.StartType, n.values[0])
	if err != nil {
		return nil, err
	}
	if state.Code == types.NFNextFlow {
		n.again(flow.End)
	}
	return state, nil
}

//等待关联节点完成
func (n *nextflow) ParallelGateway() (*gs_commons_dto.State, error) {
	tasks := n.parallel()
	if tasks == nil {
		return flowstate.ErrNode, nil
	}
	size := 0
	var err error
	e := func(er error) {
		if err == nil {
			err = er
		}
	}
	var wg sync.WaitGroup
	wg.Add(len(tasks))
	for _, v := range tasks {
		go func() {
			defer wg.Done()
			ok, err := n.modules.Instance().IsFinished(n.instance.Id, v)
			if err != nil {
				e(err)
				return
			}
			if ok {
				size++
			}
		}()
	}
	wg.Wait()
	if err != nil {
		return nil, err
	}
	if size == len(tasks) {
		n.again(n.flow().End) //如果关联的task都完成了，那么继续下一步流程
	} else {
		n.next() //没有完成当前任务，设置当前gateway为currentNode等待完成
	}
	return errstate.Success, nil
}

func (n *nextflow) InclusiveGateway() (*gs_commons_dto.State, error) {
	panic("implement me")
}

//Unsupported, is wrong!
func (n *nextflow) TriggerStartEvent() (*gs_commons_dto.State, error) {
	return flowstate.ErrNode, nil
}

//Unsupported, is wrong!
func (n *nextflow) StartEvent() (*gs_commons_dto.State, error) {
	return flowstate.ErrNode, nil
}

func (n *nextflow) EndEvent() (*gs_commons_dto.State, error) {
	panic("implement me")
}

func (n *nextflow) NotifyTask() (*gs_commons_dto.State, error) {
	panic("implement me")
}

//just notify users
func (n *nextflow) UserTask() (*gs_commons_dto.State, error) {
	n.next()
	e := n.node.data.(*models.UserTask)
	//notify
	n.modules.User().Notify(n.ctx, e)
	panic("")
}

func (n *nextflow) Restore() {
	n.status = nil
	n.finished = make(map[string]bool)
}

func (n *nextflow) again(id string) {
	n.status.again = id
}

func (n *nextflow) next() {
	n.status.currentNodes = append(n.status.currentNodes, n.node.id)
}

func (n *nextflow) parallel() []string {
	v := n.values[1]
	if v != nil {
		return v.([]string)
	}
	return nil
}

func (n *nextflow) flow() *models.SequenceFlow {
	return n.values[0].(*models.SequenceFlow)
}

func newNextflow(modules Modules, log *gslogrus.Logger, script *script.LuaScript) distribution {
	return &nextflow{modules: modules, log: log, script: newScript(modules, log, script)}
}
