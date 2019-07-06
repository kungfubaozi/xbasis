package main

import (
	"github.com/davecgh/go-spew/spew"
	"github.com/micro/go-micro/registry"
	"github.com/micro/go-micro/registry/consul"
)

var apps map[string][]string

func main() {
	cr := consul.NewRegistry(registry.Addrs("192.168.80.67:8500"),
		registry.Secure(false))

	apps = make(map[string][]string)

	s, err := cr.GetService("go.micro.api")
	if err != nil {
		panic(err)
	}

	for _, v := range s {
		p(v.Name, v.Nodes)
	}

	w, err := cr.Watch(registry.WatchService("go.micro.api"))
	if err != nil {
		panic(err)
	}

	for {
		r, err := w.Next()
		if err != nil {
			panic(err)
		}
		if r.Action == "update" {
			p(r.Service.Name, r.Service.Nodes)
		} else if r.Action == "delete" {
			p(r.Service.Name, r.Service.Nodes)
		} else if r.Action == "create" {
			p(r.Service.Name, r.Service.Nodes)
		}
	}

}

func p(name string, nodes []*registry.Node) {
	d := apps[name]
	if d == nil {
		apps[name] = []string{}
	}
	var addrs []string
	for _, v := range nodes {
		address := v.Address
		port := v.Metadata["xBasisRequestPort"]
		addrs = append(addrs, address+port)
	}

	apps[name] = addrs

	spew.Dump(apps[name])
}
