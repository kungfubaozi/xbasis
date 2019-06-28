package script

import (
	"fmt"
	"github.com/yuin/gopher-lua"
	"konekko.me/xbasis/workflow/flowerr"
	"reflect"
)

type LuaScript struct {
	l *lua.LState
}

func NewScript() *LuaScript {
	s := &LuaScript{}
	s.l = lua.NewState()
	return s
}

func (l *LuaScript) Run(script string, value map[string]interface{}) (bool, *flowerr.Error) {
	if len(script) > 5 {
		code := fmt.Sprintf("function test(flow) return %s end", script)
		table := l.l.NewTable()
		for k, v := range value {
			kind := reflect.TypeOf(v).Kind()
			if kind == reflect.Bool {
				l.l.SetTable(table, lua.LString(k), lua.LBool(v.(bool)))
			} else if kind == reflect.String {
				l.l.SetTable(table, lua.LString(k), lua.LString(v.(string)))
			} else if kind == reflect.Int64 {
				l.l.SetTable(table, lua.LString(k), lua.LNumber(v.(int64)))
			} else if kind == reflect.Int32 {
				l.l.SetTable(table, lua.LString(k), lua.LNumber(v.(int32)))
			} else if kind == reflect.Float64 {
				l.l.SetTable(table, lua.LString(k), lua.LNumber(v.(float64)))
			} else if kind == reflect.Float32 {
				l.l.SetTable(table, lua.LString(k), lua.LNumber(v.(float32)))
			}
		}
		l.l.DoString(code)
		if err := l.l.CallByParam(lua.P{
			Fn:      l.l.GetGlobal("test"),
			NRet:    1,
			Protect: true,
		}, table); err != nil {
			return false, flowerr.FromError(err)
		}
		ret := l.l.Get(-1)
		l.l.Pop(1)
		if ret.Type() != lua.LTBool {
			return false, flowerr.ErrScriptResult
		}
		return ret.String() == "true", nil
	}
	return false, nil
}
