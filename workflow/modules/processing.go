package modules

import (
	"context"
	"konekko.me/gosion/commons/dto"
	"konekko.me/gosion/commons/errstate"
	"konekko.me/gosion/commons/gslogrus"
	"konekko.me/gosion/workflow/flowstate"
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

/*
processing做的事情是控制submit操作是否可以继续向下运行，返回是flowstate.NextFlow
返回的是errstate.Success则表示当前节点提交成功
*/
func (n *processing) Do(ctx context.Context, instance *models.Instance, node *node, ct types.ConnectType, value ...interface{}) (*gs_commons_dto.State, error) {
	n.instance = instance
	n.value = value[0]
	n.node = node
	n.ctx = ctx
	return distribute(ct, n)
}

func (n *processing) Data() interface{} {
	panic("implement me")
}

func (n *processing) ExclusiveGateway() (*gs_commons_dto.State, error) {
	return flowstate.ErrNode, nil
}

func (n *processing) ParallelGateway() (*gs_commons_dto.State, error) {
	return flowstate.ErrNode, nil
}

func (n *processing) InclusiveGateway() (*gs_commons_dto.State, error) {
	return flowstate.ErrNode, nil
}

func (n *processing) TriggerStartEvent() (*gs_commons_dto.State, error) {
	panic("implement me")
}

//实例启动时的currentNodes就是startEvent，提交startEvent就是开始流程，并在此处理
func (n *processing) StartEvent() (*gs_commons_dto.State, error) {
	start, err := n.modules.Instance().IsStarted(n.instance.Id)
	if err != nil {
		return nil, err
	}
	//没开始
	if !start {
		e := n.node.data.(*models.StartEvent)
		//检查form
		return n.formCheck(e.FormRef, func() (*gs_commons_dto.State, error) {
			//开始流程
			return flowstate.NextFlow, nil
		})
	}
	//开始肯定错了
	return flowstate.ErrNode, nil
}

//stop that instance
func (n *processing) EndEvent() (*gs_commons_dto.State, error) {
	panic("implement me")
}

//notify users
func (n *processing) NotifyTask() (*gs_commons_dto.State, error) {
	return flowstate.ErrNode, nil
}

func (n *processing) UserTask() (*gs_commons_dto.State, error) {
	e := n.node.data.(*models.UserTask)
	//检查当前操作人是否对应task设置的类型
	s, err := n.modules.User().IsUserMatch(n.ctx, e)
	if err != nil {
		return nil, err
	}
	if !s.Ok {
		return s, nil
	}
	//检查form
	return n.formCheck(e.FormRef, func() (*gs_commons_dto.State, error) {
		return flowstate.NextFlow, nil
	})
}

func (n *processing) Restore() {

}

func (n *processing) formCheck(formId string, callback types.StateCallback) (*gs_commons_dto.State, error) {
	if len(formId) > 0 {
		f, err := n.modules.Form().FindById(formId)
		if err != nil {
			return nil, err
		}
		n.sd = make(map[string]interface{})
		//检查提交的value与form
		usv := n.value.(map[string]interface{})
		for _, v := range f.Fields {
			value := usv[v.Name]
			if value == nil && v.Require {
				return flowstate.ErrSubmitFormFieldNil, nil
			}
			//如果填写了就必须满足要求
			prop := v.Properties
			switch v.Type {
			case types.FTEditView:
				if !n.typeCheck(v, reflect.String) {
					return flowstate.ErrSubmitFormFieldType, nil
				}
				vtp := prop.(*models.EditViewProperties)
				tv := value.(string)
				if len(tv) >= vtp.MinLength || len(tv) <= vtp.MaxLength {
					ok, err := regexp.MatchString(vtp.Regx, tv)
					if err != nil {
						return flowstate.ErrSubmitFormFieldRegex, nil
					}
					if ok {
						n.sd[v.Name] = tv
						break
					}
				}
				return flowstate.ErrSubmitFormFieldValue, nil
			case types.FTNumberView:
				vtp := prop.(*models.NumberViewProperties)
				if vtp.Decimal {
					if !n.typeCheck(v, reflect.Float64) {
						return flowstate.ErrSubmitFormFieldType, nil
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
						return flowstate.ErrSubmitFormFieldType, nil
					}
					max := vtp.Max.(int64)
					min := vtp.Min.(int64)
					tv := value.(int64)
					if tv <= max || tv >= min {
						n.sd[v.Name] = tv
						break
					}
				}
				return flowstate.ErrSubmitFormFieldValue, nil
			case types.FTRadioGroupView:
				break
			}
		}
		if len(n.sd) > 0 {
			//提交到form记录
			var wg sync.WaitGroup
			wg.Add(2)
			s := errstate.Success
			resp := func(s1 *gs_commons_dto.State, err error) {
				if err != nil {
					s = errstate.ErrRequest
					return
				}
				if !s1.Ok {
					s = s1
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
			return s, nil
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
