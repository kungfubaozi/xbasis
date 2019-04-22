package workflow

import (
	"errors"
	"fmt"
	"github.com/yuin/gopher-lua"
	"reflect"
)

type LuaScript struct {
	l *lua.LState
}

func InitScript() *LuaScript {
	s := &LuaScript{}
	s.l = lua.NewState()
	return s
}

func (l *LuaScript) Run(script string, value map[string]interface{}) (bool, error) {
	if len(script) > 5 {
		code := fmt.Sprintf("function test(flow) %s end", script)
		table := l.l.NewTable()
		for k, v := range value {
			kind := reflect.TypeOf(v).Kind()
			if kind == reflect.Bool {
				l.l.SetTable(table, lua.LString(k), lua.LBool(v.(bool)))
			} else if kind == reflect.String {
				l.l.SetTable(table, lua.LString(k), lua.LString(v.(string)))
			} else if kind == reflect.Int32 {
				l.l.SetTable(table, lua.LString(k), lua.LNumber(v.(int32)))
			} else if kind == reflect.Int64 {
				l.l.SetTable(table, lua.LString(k), lua.LNumber(v.(int64)))
			}
		}
		l.l.DoString(code)
		if err := l.l.CallByParam(lua.P{
			Fn:      l.l.GetGlobal("test"),
			NRet:    1,
			Protect: true,
		}, table); err != nil {
			return false, err
		}
		ret := l.l.Get(-1)
		l.l.Pop(1)
		if ret.Type() != lua.LTBool {
			return false, errors.New("result value error.")
		}
		return ret.String() == "true", nil
	}
	return false, nil
}
