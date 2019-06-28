package main

import (
	_ "github.com/micro/go-micro"
	_ "github.com/micro/go-micro/registry"
	_ "github.com/micro/go-micro/registry/consul"
	"konekko.me/xbasis/permission/service"
)

func main() {
	permissionsvc.StartService()
}
