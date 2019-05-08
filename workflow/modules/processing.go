package modules

import (
	"context"
	"konekko.me/gosion/commons/gslogrus"
	"konekko.me/gosion/workflow/flowerr"
	"konekko.me/gosion/workflow/models"
	"konekko.me/gosion/workflow/types"
	"reflect"
	"regexp"
	"sync"
)

//check submit data and finished that node

type processing struct {
	modules  Modules
	instance *models.Instance
	appId    string
	ctx      context.Context
	log      *gslogrus.Logger
	node     *node
	sd       map[string]interface{}
	value    interface{}
}

func (n *processing) ApiStartEvent() (context.Context, *flowerr.Error) {
	panic("implement me")
}

/*
processing做的事情是控制submit操作是否可以继续向下运行，返回是flowerr.NextFlow
返回的是errstate.Success则表示当前节点提交成功
*/
func (n *processing) Do(ctx context.Context, instance *models.Instance, node *node, ct types.ConnectType, value ...interface{}) (context.Context, *flowerr.Error) {
	n.instance = instance
	n.value = value[0]
	n.node = node
	n.ctx = ctx
	return distribute(ctx, ct, n)
}

func (n *processing) Data() interface{} {
	panic("implement me")
}

//执行的node类型不能是除Event/Task之外的类型
func (n *processing) ExclusiveGateway() (context.Context, *flowerr.Error) {
	return n.ctx, flowerr.ErrNode
}

//执行的node类型不能是除Event/Task之外的类型
func (n *processing) ParallelGateway() (context.Context, *flowerr.Error) {
	return n.ctx, flowerr.ErrNode
}

//执行的node类型不能是除Event/Task之外的类型
func (n *processing) InclusiveGateway() (context.Context, *flowerr.Error) {
	return n.ctx, flowerr.ErrNode
}

func (n *processing) TriggerStartEvent() (context.Context, *flowerr.Error) {
	panic("implement me")
}

//实例启动时的currentNodes就是startEvent，提交startEvent就是开始流程，并在此处理
func (n *processing) StartEvent() (context.Context, *flowerr.Error) {
	start, err := n.modules.Instance().IsStarted(n.instance.Id)
	if err != nil {
		return nil, err
	}
	//没开始
	if !start {
		e := n.node.data.(*models.StartEvent)
		//检查form
		return n.formCheck(e.FormRef, func() (context.Context, *flowerr.Error) {
			//开始流程
			return n.ctx, flowerr.NextFlow
		})
	}
	//开始肯定错了
	return n.ctx, flowerr.ErrNode
}

//stop that instance
func (n *processing) EndEvent() (context.Context, *flowerr.Error) {
	panic("implement me")
}

//notify users
func (n *processing) NotifyTask() (context.Context, *flowerr.Error) {
	return n.ctx, flowerr.ErrNode
}

func (n *processing) UserTask() (context.Context, *flowerr.Error) {
	e := n.node.data.(*models.UserTask)
	//检查当前操作人是否对应task设置的类型
	err := n.modules.User().IsUserMatch(n.ctx, e)
	if err != nil {
		return n.ctx, err
	}
	//检查form
	return n.formCheck(e.FormRef, func() (context.Context, *flowerr.Error) {
		return n.ctx, flowerr.NextFlow
	})
}

func (n *processing) Restore() {

}

func (n *processing) formCheck(formId string, callback types.ErrCallback) (context.Context, *flowerr.Error) {
	if len(formId) > 0 {
		f, err := n.modules.Form().FindById(formId)
		if err != nil {
			return n.ctx, err
		}
		n.sd = make(map[string]interface{})
		//检查提交的value与form
		usv := n.value.(map[string]interface{})
		for _, v := range f.Fields {
			value := usv[v.Name]
			if value == nil && v.Require {
				return n.ctx, flowerr.ErrSubmitFormFieldNil
			}
			//如果填写了就必须满足要求
			prop := v.Properties
			switch v.Type {
			case types.FTEditView:
				if !n.typeCheck(v, reflect.String) {
					return n.ctx, flowerr.ErrSubmitFormFieldType
				}
				vtp := prop.(*models.EditViewProperties)
				tv := value.(string)
				if len(tv) >= vtp.MinLength || len(tv) <= vtp.MaxLength {
					ok, err := regexp.MatchString(vtp.Regx, tv)
					if err != nil {
						return n.ctx, flowerr.ErrSubmitFormFieldRegex
					}
					if ok {
						n.sd[v.Name] = tv
						break
					}
				}
				return n.ctx, flowerr.ErrSubmitFormFieldType
			case types.FTNumberView:
				vtp := prop.(*models.NumberViewProperties)
				if vtp.Decimal {
					if !n.typeCheck(v, reflect.Float64) {
						return n.ctx, flowerr.ErrSubmitFormFieldType
					}
					max := vtp.Max.(float64)
					min := vtp.Min.(float64)
					tv := value.(float64)
					if tv >= min || tv <= max {
						n.sd[v.Name] = tv
						break
					}
				} else {
					if !n.typeCheck(v, reflect.Int64) {
						return n.ctx, flowerr.ErrSubmitFormFieldType
					}
					max := vtp.Max.(int64)
					min := vtp.Min.(int64)
					tv := value.(int64)
					if tv <= max || tv >= min {
						n.sd[v.Name] = tv
						break
					}
				}
				return n.ctx, flowerr.ErrSubmitFormFieldValue
			case types.FTRadioGroupView:
				break
			}
		}
		if len(n.sd) > 0 {
			//提交到form记录
			var wg sync.WaitGroup
			wg.Add(2)
			var s *flowerr.Error
			resp := func(s1 *flowerr.Error) {
				if s == nil {
					s = s1
					return
				}
			}
			go func() {
				defer wg.Done()
				resp(n.modules.Form().Submit(n.ctx, n.instance.Id, n.node.id, formId, n.sd))
			}()

			go func() {
				defer wg.Done()
				resp(n.modules.History().Record(n.ctx, &models.History{
					InstanceId: n.instance.Id,
					NodeId:     n.node.id,
					Operate:    types.OPSubmitForm,
				}))
			}()
			return n.ctx, s
		}
	}
	return callback()
}

func (n *processing) typeCheck(v interface{}, kind reflect.Kind) bool {
	return reflect.TypeOf(v).Kind() == kind
}

func newProcessing(modules Modules, log *gslogrus.Logger) distribution {
	return &processing{modules: modules, log: log}
}
