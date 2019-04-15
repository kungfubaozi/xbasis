package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"konekko.me/gosion/permission/handlers"
)

func main() {
	buffer, err := ioutil.ReadFile("./init.json")
	if err != nil {
		panic(err)
	}
	var fc *permissionhandlers.FunctionsConfig

	err = json.Unmarshal(buffer, &fc)
	if err != nil {
		panic(err)
	}
	fmt.Println("datas", len(fc.Data))
}
