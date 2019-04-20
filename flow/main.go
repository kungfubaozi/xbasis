package main

import (
	"fmt"
	"github.com/yuin/gluamapper"
	"github.com/yuin/gopher-lua"
	"time"
)

type Role struct {
	Name string
}

type Person struct {
	Name      string
	Age       int
	WorkPlace string
	Role      []*Role
}

func main() {
	//flowinstance.InitProcess()
	a := time.Now().UnixNano()
	loadFile()
	fmt.Println("time", (time.Now().UnixNano()-a)/1e6)
}

func loadCode() {
	L := lua.NewState()
	if err := L.DoString(`
person = {
  name = "Michel",
  age  = "31", -- weakly input
  work_place = "San Jose",
  role = {
    {
      name = "Administrator"
    },
    {
      name = "Operator"
    }
  }
}
`); err != nil {
		panic(err)
	}
	var person Person
	if err := gluamapper.Map(L.GetGlobal("person").(*lua.LTable), &person); err != nil {
		panic(err)
	}
	fmt.Println(fmt.Sprintf("%s %d", person.Name, person.Age))
}

func loadFile() {
	dic := make(map[string]string)
	dic["user"] = "test"
	dic["os"] = "ios"
	dic["version"] = "1.0"
	//path := "/Users/Richard/Desktop/Development/Golang/src/konekko.me/gosion/flow/script/test.lua"
	L := lua.NewState()
	defer L.Close()
	if err := L.DoString(`
function condition(value)
    return value == "1235"
end`); err != nil {
		panic(err)
	}

	if err := L.CallByParam(lua.P{
		Fn:      L.GetGlobal("condition"),
		NRet:    1,
		Protect: true,
	}, lua.LString("123")); err != nil {
		panic(err)
	}
	ret := L.Get(-1)
	L.Pop(1)
	fmt.Println(ret.String(), ret.Type())

	//
	//
	//table := L.NewTable()
	//for k, v := range dic {
	//	L.SetTable(table, lua.LString(k), lua.LString(v))
	//}
	//
	//if err := L.CallByParam(lua.P{
	//	Fn:      L.GetGlobal("testFun"),
	//	NRet:    1,
	//	Protect: true,
	//}, table); err != nil {
	//	panic(err)
	//}
	//ret := L.Get(-1) // returned value
	//L.Pop(1)         // remove received value
	//obj := gluamapper.ToGoValue(ret, gluamapper.Option{NameFunc: printTest})
	//fmt.Println(obj)
}

func printTest(s string) string {
	return s
}
