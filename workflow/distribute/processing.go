package distribute

import (
	"context"
	"konekko.me/gosion/commons/gslogrus"
	"konekko.me/gosion/workflow/flowerr"
	"konekko.me/gosion/workflow/models"
	"konekko.me/gosion/workflow/modules"
	"konekko.me/gosion/workflow/types"
	"reflect"
	"regexp"
	"sync"
)

type processing struct {
	modules  modules.Modules
	instance *models.Instance
	appId    string
	ctx      context.Context
	log      *gslogrus.Logger
	node     *models.Node
	sd       map[string]interface{}
	value    interface{}
}

func (f *processing) eventGateway() *flowerr.Error {
	panic("implement me")
}

func (f *processing) timerStartEvent() *flowerr.Error {
	panic("implement me")
}

func (f *processing) messageStartEvent() *flowerr.Error {
	panic("implement me")
}

func (f *processing) cancelEndEvent() *flowerr.Error {
	panic("implement me")
}

func (f *processing) terminateEndEvent() *flowerr.Error {
	panic("implement me")
}

func (f *processing) Data() interface{} {
	panic("implement me")
}

func (f *processing) Do(ctx context.Context, instance *models.Instance, node *models.Node, ct types.ConnectType, value ...interface{}) (context.Context, *flowerr.Error) {
	f.instance = instance
	f.value = value[0]
	f.node = node
	f.ctx = ctx
	return handler(ctx, ct, f)
}

//执行的node类型不能是除Event/Task之外的类型
func (f *processing) exclusiveGateway() *flowerr.Error {
	panic("implement me")
}

//执行的node类型不能是除Event/Task之外的类型
func (f *processing) parallelGateway() *flowerr.Error {
	panic("implement me")
}

//执行的node类型不能是除Event/Task之外的类型
func (f *processing) inclusiveGateway() *flowerr.Error {
	panic("implement me")
}

func (f *processing) startEvent() *flowerr.Error {
	start, err := f.modules.Instance().IsStarted(f.instance.Id)
	if err != nil {
		return err
	}
	//没开始
	if !start {
		e := f.node.Data.(*models.StartEvent)
		//检查form
		return f.formCheck(e.FormRef, func() *flowerr.Error {
			//开始流程
			return flowerr.NextFlow
		})
	}
	//开始肯定错了
	return flowerr.ErrNode
}

func (f *processing) endEvent() *flowerr.Error {
	panic("implement me")
}

func (f *processing) apiStartEvent() *flowerr.Error {
	panic("implement me")
}

func (f *processing) userTask() *flowerr.Error {
	e := f.node.Data.(*models.UserTask)
	//检查当前操作人是否对应task设置的类型
	err := f.modules.User().IsUserMatch(f.ctx, e)
	if err != nil {
		return err
	}
	//检查form
	return f.formCheck(e.FormRef, func() *flowerr.Error {
		return flowerr.NextFlow
	})
}

func (f *processing) notifyTask() *flowerr.Error {
	panic("implement me")
}

func (f *processing) triggerStartEvent() *flowerr.Error {
	panic("implement me")
}

func (f *processing) context(ctx context.Context) context.Context {
	if ctx != nil {
		f.ctx = ctx
	}
	return f.ctx
}

func (f *processing) metadata(key, data interface{}) {
	f.ctx = context.WithValue(f.ctx, key, data)
}

func (f *processing) Restore() {
	panic("implement me")
}

func (f *processing) formCheck(formId string, callback types.ErrCallback) *flowerr.Error {
	if len(formId) > 0 {
		form, err := f.modules.Form().FindById(formId)
		if err != nil {
			return err
		}
		f.sd = make(map[string]interface{})
		//检查提交的value与form
		usv := f.value.(map[string]interface{})
		for _, v := range form.Fields {
			value := usv[v.Name]
			if value == nil && v.Require {
				return flowerr.ErrSubmitFormFieldNil
			}
			//如果填写了就必须满足要求
			prop := v.Properties
			switch v.Type {
			case types.FTEditView:
				if !f.typeCheck(v, reflect.String) {
					return flowerr.ErrSubmitFormFieldType
				}
				vtp := prop.(*models.EditViewProperties)
				tv := value.(string)
				if len(tv) >= vtp.MinLength || len(tv) <= vtp.MaxLength {
					ok, err := regexp.MatchString(vtp.Regx, tv)
					if err != nil {
						return flowerr.ErrSubmitFormFieldRegex
					}
					if ok {
						f.sd[v.Name] = tv
						break
					}
				}
				return flowerr.ErrSubmitFormFieldType
			case types.FTNumberView:
				vtp := prop.(*models.NumberViewProperties)
				if vtp.Decimal {
					if !f.typeCheck(v, reflect.Float64) {
						return flowerr.ErrSubmitFormFieldType
					}
					max := vtp.Max.(float64)
					min := vtp.Min.(float64)
					tv := value.(float64)
					if tv >= min || tv <= max {
						f.sd[v.Name] = tv
						break
					}
				} else {
					if !f.typeCheck(v, reflect.Int64) {
						return flowerr.ErrSubmitFormFieldType
					}
					max := vtp.Max.(int64)
					min := vtp.Min.(int64)
					tv := value.(int64)
					if tv <= max || tv >= min {
						f.sd[v.Name] = tv
						break
					}
				}
				return flowerr.ErrSubmitFormFieldValue
			case types.FTRadioGroupView:
				break
			}
		}
		if len(f.sd) > 0 {
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
				resp(f.modules.Form().Submit(f.ctx, f.instance.Id, f.node.Id, formId, form.Encryption, f.sd))
			}()

			go func() {
				defer wg.Done()
				resp(f.modules.History().Record(f.ctx, &models.History{
					InstanceId: f.instance.Id,
					NodeId:     f.node.Id,
					Operate:    types.OPSubmitForm,
				}))
			}()
			return s
		}
	}
	return callback()
}

func (f *processing) typeCheck(v interface{}, kind reflect.Kind) bool {
	return reflect.TypeOf(v).Kind() == kind
}

func NewProcessing(modules modules.Modules, log *gslogrus.Logger) Handler {
	return &processing{modules: modules, log: log}
}
