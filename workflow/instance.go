package workflow

//single
type instance struct {
	scripts map[string]string
	start   chan bool
	next    chan int
}

var instances []*instance

func putInstance(i *instance) {
	instances = append(instances, i)
	go i.Start()
}

func (i *instance) Start() {
	for {
		select {
		case s := <-i.start:
			if !s {
				return
			}

			break
		case next := <-i.next:
			break
		}

	}
}
