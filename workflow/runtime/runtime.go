package runtime

import (
	"context"
	"fmt"
	"github.com/micro/go-micro/metadata"
	"konekko.me/gosion/commons/gslogrus"
	"konekko.me/gosion/commons/wrapper"
	"konekko.me/gosion/workflow/distribute"
	"konekko.me/gosion/workflow/flowerr"
	"konekko.me/gosion/workflow/models"
	"konekko.me/gosion/workflow/modules"
	"konekko.me/gosion/workflow/types"
)

type runtime struct {
	modules    modules.Modules
	shutdown   chan error
	pipelines  map[string]modules.Pipeline
	processing distribute.Handler
	next       distribute.Handler
	dataGetter distribute.Handler
	log        *gslogrus.Logger
}

func createRuntime(shutdown chan error, log *gslogrus.Logger) (modules.AddProcessToPipelineCallback, *runtime) {
	r := &runtime{
		shutdown:  shutdown,
		log:       log,
		pipelines: make(map[string]modules.Pipeline),
	}
	return func(pip modules.Pipeline) {
		r.pipelines[pip.Id()] = pip
	}, r
}

func (r *runtime) Submit(ctx1 context.Context, instanceId, nodeId string, value map[string]interface{}) *flowerr.Error {
	data, ok := metadata.FromContext(ctx1)
	if ok {

		ctx := context.Background()
		user := gs_commons_wrapper.GetData(data)
		ctx = context.WithValue(ctx, "auth", user)

		i, err := r.modules.Instance().FindById(instanceId)
		if err != nil {
			if err == flowerr.ErrNil {
				return flowerr.ErrInvalidInstance
			}
			return err
		}

		//get pipeline
		pipe := r.pipelines[i.ProcessId]

		size := 0
		for _, v := range i.CurrentNodes {

			node, err := pipe.GetNode(nodeId)

			if v == nodeId {

				if err != nil {
					return flowerr.ErrRequest
				}

				//这里检查的node类型只有event和task，gateway不在考虑范围内，因为currentNodes不会包含gateway类型的node
				//check submit node data and finished that node
				r, err := r.processing.Do(ctx, i, node, node.CT, value)

				ctx = r
				if err != nil {
					if err == flowerr.NextFlow {
						size++
					} else {
						return err
					}
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
			nodes, err := r.again(ctx, currentNodes, i, pipe, v)
			if err != nil {
				return err
			}
			if len(nodes) > 0 {
				return nil
			}
			if len(nodes) > 0 {
				currentNodes = append(currentNodes, nodes...)
			}
		}
		//update instance node
		err = r.modules.Instance().UpdateInstanceCurrentNodes(instanceId, currentNodes...)
		if err != nil {
			fmt.Println("update error", err)
			return err
		}
		//}

		r.next.Restore()
		r.dataGetter.Restore()

	}

	return nil
}

func (r *runtime) again(ctx context.Context, currentNodes []string, i *models.Instance, pipe modules.Pipeline, nodeId string) ([]string, *flowerr.Error) {
	var ns []string
	if len(currentNodes) > 0 {
		ns = append(ns, currentNodes...)
	}
	flows, err := pipe.Flows(nodeId)
	if err != nil {
		return nil, err
	}
	for _, f := range flows {
		//需要注意的是，不加入网关会忽略flow上的script
		//需要注意的是，不加入网关会忽略flow上的script
		//需要注意的是，不加入网关会忽略flow上的script

		var connects interface{}
		switch f.EndType {
		case types.CTParallelGateway: //当是ParallelGateway时，需要获取与之关联的task节点
			connects = pipe.FindParallelNodes(f.End)
			break
		case types.CTExclusiveGateway, types.CTInclusiveGateway: //获取所有flow start与当前节点连接的
			connects, err = pipe.Flows(f.Start)
			if err != nil {
				return nil, err
			}
			//获取节点提交的数据
			c, err := r.getSubmitData(i, ctx, f.Start, pipe)
			if err != nil {
				return nil, err
			}
			//新的context
			ctx = c
			break
		}
		//获取当前flow末尾连线的节点
		n, err := pipe.GetNode(f.End)
		if err != nil {
			return nil, err
		}
		ctx, err = r.next.Do(ctx, i, n, f.EndType, f, connects)
		if err != nil {
			return nil, err
		}
	}
	nodes := r.next.Data().(*models.NextStatus)
	for _, v1 := range nodes.CurrentNodes {
		ns = append(ns, v1)
	}
	if len(nodes.Again) > 0 {
		return r.again(ctx, ns, i, pipe, nodes.Again)
	}
	return ns, nil
}

//不同于其他的是，此操作是向前查找
func (r *runtime) getSubmitData(i *models.Instance, ctx context.Context, fromNodeId string, pipe modules.Pipeline) (context.Context, *flowerr.Error) {
	nodes := pipe.GetNodeBackwardRelations(fromNodeId)
	for _, v := range nodes {
		ctx1, err := r.dataGetter.Do(ctx, i, nil, v.CT, v)
		if err != nil {
			return nil, err
		}
		ctx = ctx1
	}
	return ctx, nil
}

func (r *runtime) RunningProcessSize() int {
	panic("implement me")
}
