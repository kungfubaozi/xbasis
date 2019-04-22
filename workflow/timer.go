package workflow

import "time"

type Timer struct {
	running bool
}

func newTimer() *Timer {
	return &Timer{}
}

func (t *Timer) Start() {
	go func() {
		t.running = true
		for {
			time.Sleep(5000)
		}
	}()
}

func (t *Timer) RegisterTime(time string) {
	if !t.running {
		t.Start()
	}
}
