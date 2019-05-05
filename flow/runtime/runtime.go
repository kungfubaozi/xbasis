package runtime

import (
	"context"
	"konekko.me/gosion/commons/dto"
	"konekko.me/gosion/commons/errstate"
	"konekko.me/gosion/flow/flowstate"
	"konekko.me/gosion/flow/history"
	"konekko.me/gosion/flow/instance"
	"konekko.me/gosion/flow/process"
	"konekko.me/gosion/flow/script"
	"konekko.me/gosion/flow/types"
)

type Runtime interface {
	//ctx: contains user authentication information
	Submit(ctx context.Context, instanceId, nodeId string, value map[string]interface{}) (*gs_commons_dto.State, error)
}

type runtime struct {
	shutdown  chan error
	script    *script.LuaScript
	pipelines map[string]process.Pipeline
	hi        history.Interface
	pi        process.Interface
	ii        instance.Interface
}

//submit操作会对当前进行的节点进行分析，并判断传入的nodeId是否在当前节点中，在就继续
func (r *runtime) Submit(ctx context.Context, instanceId, nodeId string, value map[string]interface{}) (*gs_commons_dto.State, error) {
	i, err := r.ii.FindById(instanceId)
	if err != nil {
		if err == types.ErrNil {
			return flowstate.ErrInvalidInstance, nil
		}
		return errstate.ErrRequest, nil
	}

	//get pipeline
	pipe := r.pipelines[i.ProcessId]

	size := 0
	p := newProcessingTask(r)
	for _, v := range i.CurrentNodes {

		if v == nodeId {
			node, err := pipe.GetNode(nodeId)
			if err != nil {
				return errstate.ErrRequest, nil
			}
			//check submit node data and finished that node
			p.Do(ctx, instanceId, nodeId, node.CT, value)
		}

		//check finished node
		ok, err := r.ii.IsFinished(instanceId, v)
		if err != nil {
			return errstate.ErrRequest, nil
		}

		if ok {
			size++
		}

	}

	if size == len(i.CurrentNodes) {
		//next node
	}

	for {
		select {
		case v := <-ctx.Done():
			err := ctx.Err()
			if err != nil {
				return errstate.ErrRequest, nil
			}

			break
		}
	}
}
