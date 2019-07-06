package main

import (
	"github.com/Sirupsen/logrus"
	"github.com/micro/go-micro/registry"
	"konekko.me/xbasis/commons/errstate"
	"math/rand"
)

var apps map[string][]string
var watchs []string

var cr registry.Registry

func (r *request) address() bool {

	addresses := apps[r.serviceName]
	if addresses == nil {

		//get service
		s, err := cr.GetService(r.serviceName)
		if err != nil {
			r.services._log.WithFields(logrus.Fields{
				"serviceName": r.serviceName,
				"err":         err.Error(),
			}).Error("get service name")
			r.json(errstate.ErrSystem)
			return false
		}

		for _, v := range s {
			r.nodes(v.Name, v.Nodes)
		}

		w := true
		for _, v := range watchs {
			if v == r.serviceName {
				w = false
				break
			}
		}

		//start watch
		if w {
			watchs = append(watchs, r.serviceName)
			r.services._log.WithFields(logrus.Fields{
				"service": r.serviceName,
			}).Warn("start watch")
			go func(name string) {
				w, err := cr.Watch(registry.WatchService(name))
				if err != nil {
					panic(err)
				}

				for {
					n, err := w.Next()
					if err != nil {
						panic(err)
					}
					r.nodes(n.Service.Name, n.Service.Nodes)
				}
			}(r.serviceName)
		}

		addresses = apps[r.serviceName]
		if addresses == nil || len(addresses) == 0 {
			r.json(errstate.ErrInvalidServiceNode)
			return false
		}

	}

	r.path = addresses[rand.Intn(len(addresses))]
	return true
}

func (r *request) nodes(name string, nodes []*registry.Node) {
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

}
