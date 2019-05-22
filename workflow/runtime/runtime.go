package runtime

import (
	"context"
	"fmt"
	"github.com/micro/go-micro/metadata"
	"github.com/samuel/go-zookeeper/zk"
	"konekko.me/gosion/commons/gslogrus"
	"konekko.me/gosion/commons/wrapper"
	"konekko.me/gosion/workflow/distribute"
	"konekko.me/gosion/workflow/flowerr"
	"konekko.me/gosion/workflow/models"
	"konekko.me/gosion/workflow/modules"
	"konekko.me/gosion/workflow/types"
	"sync"
)

type runtime struct {
	modules    modules.Modules
	shutdown   chan error
	pipelines  modules.Pipelines
	processing distribute.Handler
	next       distribute.Handler
	dataGetter distribute.Handler
	log        *gslogrus.Logger
	conn       *zk.Conn
}

func (r *runtime) Submit(ctx1 context.Context, instanceId, nodeId string, value map[string]interface{}) *flowerr.Error {
	data, ok := metadata.FromContext(ctx1)
	if ok {

		//lock instanceId
		path := fmt.Sprintf("lock-%s", instanceId)
		_, e := r.conn.Create(path, nil, zk.FlagEphemeral, zk.WorldACL(zk.PermCreate))
		if e != nil {
			return flowerr.FromError(e)
		}
		defer r.conn.Delete(path, 0)

		//progress...
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
		pipe, e := r.pipelines.Get(i.ProcessId)
		if err != nil {
			return flowerr.FromError(e)
		}

		commandFunc := func(command types.GetterCommand, values ...interface{}) (interface{}, *flowerr.Error) {
			switch command {
			case types.GCBackwardRelations:
				return pipe.GetNodeBackwardRelations(values[0].(string)), nil
			case types.GCNodeFlows:
				connects, err := pipe.Flows(values[0].(string))
				if err != nil {
					return nil, err
				}
				return connects, err
			case types.GCNodeSubmitData:
				data, err := r.getSubmitData(i, context.Background(), values[0].(string), pipe)
				if err != nil {
					return nil, err
				}
				return data, nil
			case types.GCNode:
				node, err := pipe.GetNode(values[0].(string))
				if err != nil {
					return nil, err
				}
				return node, nil
			case types.GCForwardRelationNodes:
				return pipe.GetNodeForwardRelations(values[0].(string)), nil
			}
			return nil, flowerr.ErrNil
		}

		r.processing.SetCommandFunc(commandFunc)

		size := 0
		for _, v := range i.CurrentNodes {

			node, err := pipe.GetNode(nodeId)

			if v == nodeId {

				if err != nil {
					return flowerr.ErrRequest
				}

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

		}

		r.processing.Restore()

		r.next.Restore()

		r.next.SetCommandFunc(commandFunc)

		//不需要遍历CurrentNodes，直接走提交的节点即可
		err = r.again(ctx, i, pipe, nodeId)
		if err != nil {
			return err
		}

		//执行
		data, err := r.next.RunActions(i.CurrentNodes, nodeId)
		if err != nil {
			return err
		}

		d, ok := data.([]string)

		r.next.Restore()
		r.dataGetter.Restore()

		if ok {
			err = r.modules.Instance().UpdateInstanceCurrentNodes(instanceId, d...)
			if err != nil {
				fmt.Println("update error", err)
				return err
			}
			return nil
		}

		return flowerr.ErrUnknow
	}

	return nil
}

func (r *runtime) again(ctx context.Context, i *models.Instance, pipe modules.Pipeline, nodeId string) *flowerr.Error {

	//nodeId是当前执行node
	flows, err := pipe.Flows(nodeId)
	if err != nil {
		return err
	}

	var wg sync.WaitGroup
	wg.Add(len(flows))
	resp := func(e *flowerr.Error) {
		if err == nil {
			err = e
		}
	}

	//处理的是当前节后后面的所有节点, 不是当前node
	//这里还有问题没有解决：冗余和多余请求的问题，当前节点连接的可能是同一节点
	//以及包括还有提交的数据都是一样的
	for _, f := range flows {
		go func() {
			defer wg.Done()
			//获取当前flow末尾连线的节点
			n, err := pipe.GetNode(f.End)
			if err != nil {
				resp(err)
				return
			}

			ctx, err = r.next.Do(ctx, i, n, n.CT, f)
			if err != nil {
				resp(err)
				return
			}
		}()
	}

	wg.Wait()

	if err != nil {
		return err
	}

	nodes := r.next.Data().(*models.NextStatus)

	v, ok := ctx.Value(types.RollbackKey).(bool)
	if !ok {
		v = false
	}
	if v {
		//如果没有继续往下的则设置rb为false
		if len(nodes.Again) == 0 {
			ctx = context.WithValue(ctx, types.RollbackKey, false)
		}
	}

	if len(nodes.Again) > 0 {

		for _, v := range nodes.Again {
			e := r.again(ctx, i, pipe, v)
			if e != nil {
				return e
			}
		}

		return err
	}
	return nil
}

//不同于其他的是，此操作是向前查找
func (r *runtime) getSubmitData(i *models.Instance, ctx context.Context, fromNodeId string, pipe modules.Pipeline) (interface{}, *flowerr.Error) {
	nodes := pipe.GetNodeBackwardRelations(fromNodeId)
	var err *flowerr.Error
	resp := func(e *flowerr.Error) {
		if err == nil {
			err = e
		}
	}
	var wg sync.WaitGroup
	wg.Add(len(nodes))
	for _, v := range nodes {
		go func() {
			defer wg.Done()
			_, err := r.dataGetter.Do(ctx, i, nil, v.CT, v)
			resp(err)
		}()
	}
	wg.Wait()
	if err != nil {
		return nil, err
	}
	return r.dataGetter.Data(), nil
}

func (r *runtime) RunningProcessSize() int {
	panic("implement me")
}
