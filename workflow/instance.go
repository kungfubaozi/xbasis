package workflow

//single
type instance struct {
	script         *luaScript
	id             string
	scripts        map[string]string
	forms          map[string]*typeForm
	flowConditions map[string]string
	running        bool
	start          chan bool
	nextTask       chan int
	indexTask      int //当前位置
}

type users struct {
}

func (i *instance) Start() {
	for {
		select {
		case s := <-i.start:
			if !s {
				i.running = false
				return
			}

			break
		case next := <-i.nextTask:
			i.indexTask = next
			break
		}

	}
}

func (i *instance) SaveFlowStatus() {

}
