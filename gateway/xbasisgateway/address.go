package xbasisgateway

import (
	"github.com/Sirupsen/logrus"
	"github.com/micro/go-micro/registry"
	"github.com/micro/go-micro/registry/consul"
	"konekko.me/xbasis/commons/errstate"
	"math/rand"
)

var cr = consul.NewRegistry(registry.Addrs("192.168.80.67:8500"),
	registry.Secure(false))

var apps = make(map[string]*appService)
var watchs []string

type appService struct {
	watch       registry.Watcher
	serviceName string
	appId       string
	addresses   []string
	stop        chan bool
}

func (r *request) address() bool {

	svc := apps[r.toAppId]
	if svc == nil {

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
			r.nodes(v)
		}

		w := true
		for _, v := range watchs {
			if v == r.toAppId {
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
			w, err := cr.Watch(registry.WatchService(r.serviceName))
			if err != nil {
				panic(err)
			}
			if apps[r.toAppId].watch == nil {
				apps[r.toAppId].watch = w
			}
			go func() {
				for {
					select {
					case <-svc.stop:
						return
					default:
						n, err := w.Next()
						if err != nil {
							panic(err)
						}
						r.nodes(n.Service)
					}
				}
			}()

		}

		svc = apps[r.toAppId]
		if svc == nil || len(svc.addresses) == 0 {
			r.json(errstate.ErrInvalidServiceNode)
			return false
		}

	} else {
		if svc.serviceName != r.serviceName {
			//cancel watch and re-watch
			svc.stop <- true
			delete(apps, r.toAppId)
			return r.address()
		}
	}

	r.path = svc.addresses[rand.Intn(len(svc.addresses))]
	return true
}

func (r *request) nodes(service *registry.Service) {
	d := apps[r.toAppId]
	if d == nil {
		apps[r.toAppId] = &appService{
			serviceName: r.serviceName,
			appId:       r.toAppId,
			stop:        make(chan bool),
		}
	}
	var addrs []string
	if r.serviceName != service.Name {
		return
	}
	for _, v := range service.Nodes {
		address := v.Address
		port := v.Metadata["xBasisRequestPort"]
		addrs = append(addrs, address+port)
	}

	apps[r.toAppId].addresses = addrs
}
