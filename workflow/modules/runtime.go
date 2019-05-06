package modules

import (
	"context"
	"github.com/micro/go-micro/metadata"
	"konekko.me/gosion/commons/dto"
	"konekko.me/gosion/commons/errstate"
	"konekko.me/gosion/commons/gslogrus"
	"konekko.me/gosion/workflow/flowstate"
	"konekko.me/gosion/workflow/script"
	"konekko.me/gosion/workflow/types"
)

type IRuntime interface {
	Submit(ctx context.Context, instanceId, nodeId string, value map[string]interface{}) (*gs_commons_dto.State, error)

	RunningProcessSize() int
}

type runtime struct {
	modules    Modules
	shutdown   chan error
	script     *script.LuaScript
	pipelines  map[string]Pipeline
	processing distribution
	next       distribution
	log        *gslogrus.Logger
}

func createRuntime(shutdown chan error, log *gslogrus.Logger) (AddProcessToPipelineCallback, *runtime) {
	r := &runtime{
		shutdown:  shutdown,
		log:       log,
		script:    script.NewScript(),
		pipelines: make(map[string]Pipeline),
	}
	return func(pip Pipeline) {
		r.pipelines[pip.Id()] = pip
	}, r
}

func (r *runtime) Submit(ctx context.Context, instanceId, nodeId string, value map[string]interface{}) (*gs_commons_dto.State, error) {
	ctx = metadata.NewContext(ctx, map[string]string{})

	i, err := r.modules.Instance().FindById(instanceId)
	if err != nil {
		if err == types.ErrNil {
			return flowstate.ErrInvalidInstance, nil
		}
		return errstate.ErrRequest, nil
	}

	//get pipeline
	pipe := r.pipelines[i.ProcessId]

	size := 0
	for _, v := range i.CurrentNodes {

		node, err := pipe.GetNode(nodeId)

		//执行的node类型不能是event/task之外的类型
		switch node.ct {
		case types.CTInclusiveGateway, types.CTExclusiveGateway, types.CTParallelGateway:
			return errstate.ErrRequest, nil
		}

		if v == nodeId {

			if err != nil {
				return errstate.ErrRequest, nil
			}
			//这里检查的node类型只有event和task，gateway不在考虑范围内，因为currentNodes不会包含gateway类型的node
			//check submit node data and finished that node
			r.processing.Do(ctx, instanceId, node, node.ct, value)
		}

		//check finished node
		ok, err := r.modules.Instance().IsFinished(instanceId, v)
		if err != nil {
			return errstate.ErrRequest, nil
		}

		if ok {
			size++
		}

	}

	r.processing.Restore()

	if size == len(i.CurrentNodes) {
		//next node
		var currentNodes []string
		again := func(nodeId string) (string, error) {
			flows, err := pipe.Flows(nodeId)
			if err != nil {
				return "", err
			}
			for _, f := range flows {
				r.next.Do(ctx, instanceId, nil, f.EndType, f)
			}
			nodes := r.next.Data().(*nextstatus)
			for _, v1 := range nodes.currentNodes {
				currentNodes = append(currentNodes, v1)
			}
			return nodes.again, nil
		}
		for _, v := range i.CurrentNodes {
			//again的作用是: 当前的节点的另一端不是一个有效的task/event，需要再进行一次查询，直到出现有效的task/event
			a, err := again(v)
			if err != nil {
				return errstate.ErrRequest, nil
			}
			if len(a) > 0 {
				again(a)
			}
		}
		//update instance node
		r.modules.Instance().UpdateInstanceCurrentNodes(instanceId, currentNodes...)
	}

	r.next.Restore()

	return errstate.ErrRequest, nil
}

func (r *runtime) RunningProcessSize() int {
	panic("implement me")
}
