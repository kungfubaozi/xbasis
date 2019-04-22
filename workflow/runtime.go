package workflow

import "konekko.me/gosion/commons/generator"

type runtime struct {
	script    *luaScript
	instances map[string]*instance
	count     int64
	timer     *Timer
	id        gs_commons_generator.IDGenerator
}

func NewRuntime() *runtime {
	return &runtime{
		script:    newScript(),
		instances: make(map[string]*instance),
		count:     0,
		timer:     newTimer(),
		id:        gs_commons_generator.NewIDG(),
	}
}

func (r *runtime) Initialize() {
	r.loadProcess()
	r.buildInstance()
	r.start()
}

//load form db
func (r *runtime) loadProcess() {

}

func (r *runtime) buildInstance() {

}

func (r *runtime) checkProcess() {

}

func (r *runtime) start() {

}

func (r *runtime) run(i *instance) {
	i.id = r.id.Get()
	i.script = r.script
	r.instances[i.id] = i
	go i.Start()
}
