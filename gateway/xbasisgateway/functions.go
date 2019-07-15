package xbasisgateway

import (
	"github.com/vmihailenco/msgpack"
	"konekko.me/xbasis/commons/constants"
	"konekko.me/xbasis/commons/listen"
	"konekko.me/xbasis/commons/transport"
)

type functions struct {
	//app-path-function
	funcs map[string]AppFunctions
}

//path-function detail
type AppFunctions map[string]*xbasistransport.AppFunction

func (f *functions) find(appId, path string) *xbasistransport.AppFunction {
	v, ok := f.funcs[appId]
	if ok {
		f, ok := v[path]
		if ok {
			return f
		}
	}
	return nil
}

func (f *functions) update(af *xbasistransport.AppFunction) {
	if len(af.AppId) > 0 && len(af.Path) > 0 {
		_, ok := f.funcs[af.AppId]
		if ok {
			d := f.funcs[af.AppId][af.Path]
			if d == nil {
				af.Version = 1
				f.funcs[af.AppId] = map[string]*xbasistransport.AppFunction{}
			} else {
				af.Version = d.Version + 1
			}
			f.funcs[af.AppId][af.Path] = af
		} else {
			f.funcs[af.AppId] = make(map[string]*xbasistransport.AppFunction)
			f.update(af)
		}
	}
}

func (f *functions) listen(addr string) error {
	return listen.CreateKafkaListener(addr, xbasisconstants.SyncFunctionTopic, func(msg []byte) {
		af := &xbasistransport.AppFunction{}
		err := msgpack.Unmarshal(msg, af)
		if err == nil {
			f.update(af)
		}
	})
}
