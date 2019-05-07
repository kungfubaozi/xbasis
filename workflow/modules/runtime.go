package modules

import (
	"context"
	"fmt"
	"github.com/micro/go-micro/metadata"
	"konekko.me/gosion/commons/dto"
	"konekko.me/gosion/commons/errstate"
	"konekko.me/gosion/commons/gslogrus"
	"konekko.me/gosion/commons/wrapper"
	"konekko.me/gosion/workflow/flowstate"
	"konekko.me/gosion/workflow/models"
	"konekko.me/gosion/workflow/types"
)

type IRuntime interface {
	Submit(ctx context.Context, instanceId, nodeId string, value map[string]interface{}) (*gs_commons_dto.State, error)

	RunningProcessSize() int
}

type runtime struct {
	modules    Modules
	shutdown   chan error
	pipelines  map[string]Pipeline
	processing distribution
	next       distribution
	log        *gslogrus.Logger
}

func createRuntime(shutdown chan error, log *gslogrus.Logger) (AddProcessToPipelineCallback, *runtime) {
	r := &runtime{
		shutdown:  shutdown,
		log:       log,
		pipelines: make(map[string]Pipeline),
	}
	return func(pip Pipeline) {
		r.pipelines[pip.Id()] = pip
	}, r
}

func (r *runtime) Submit(ctx1 context.Context, instanceId, nodeId string, value map[string]interface{}) (*gs_commons_dto.State, error) {
	data, ok := metadata.FromContext(ctx1)
	if ok {
		ctx := context.Background()
		user := gs_commons_wrapper.GetData(data)
		ctx = context.WithValue(ctx, "auth", user)

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

			//执行的node类型不能是除Event/Task之外的类型
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
				r, err := r.processing.Do(ctx, i, node, node.ct, value)
				if err != nil {
					return r, err
				}
				if r.Code == types.NFNextFlow { //next flow
					size++
				} else {
					return r, nil
				}
			}
			//
			////check finished node
			//ok, err := r.modules.Instance().IsFinished(instanceId, v)
			//if err != nil {
			//	return errstate.ErrRequest, nil
			//}
			//
			//if ok {
			//	nextNodes = append(nextNodes, v)
			//	size++
			//}

		}

		r.processing.Restore()

		r.next.Restore()

		//if size == len(i.CurrentNodes) {
		//next node
		var currentNodes []string
		for _, v := range i.CurrentNodes {
			//again的作用是: 当前的节点的另一端不是一个有效的task/event，需要再进行一次查询，直到出现有效的task/event
			nodes, ok, err := r.again(ctx, currentNodes, i, pipe, v)
			if err != nil {
				return errstate.ErrRequest, nil
			}
			if !ok.Ok {
				return ok, nil
			}
			if len(nodes) > 0 {
				currentNodes = append(currentNodes, nodes...)
			}
		}
		//update instance node
		s, err := r.modules.Instance().UpdateInstanceCurrentNodes(instanceId, currentNodes...)
		if err != nil {
			fmt.Println("update error", err)
			return errstate.ErrRequest, nil
		}
		if s.Ok {
			return s, nil
		}
		//}

		r.next.Restore()

		return errstate.ErrRequest, nil
	}
}

func (r *runtime) again(ctx context.Context, currentNodes []string, i *models.Instance, pipe Pipeline, nodeId string) ([]string, *gs_commons_dto.State, error) {
	var ns []string
	if len(currentNodes) > 0 {
		ns = append(ns, currentNodes...)
	}
	flows, err := pipe.Flows(nodeId)
	if err != nil {
		return nil, nil, err
	}
	for _, f := range flows {
		//当是ParallelGateway时，需要获取与之关联的task节点
		var connects []string
		switch f.EndType {
		case types.CTParallelGateway:
			connects = pipe.FindEndConnectNodes(f.End)
			break
		}
		n, err := pipe.GetNode(f.End)
		if err != nil {
			return nil, nil, err
		}
		s, err := r.next.Do(ctx, i, n, f.EndType, f, connects)
		if err != nil {
			return nil, s, err
		}
		if s.Ok {

		}
	}
	nodes := r.next.Data().(*nextstatus)
	for _, v1 := range nodes.currentNodes {
		ns = append(ns, v1)
	}
	if len(nodes.again) > 0 {
		return r.again(ctx, ns, i, pipe, nodes.again)
	}
	return ns, errstate.Success, nil
}

func (r *runtime) RunningProcessSize() int {
	panic("implement me")
}
