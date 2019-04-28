package workflow

import "gopkg.in/mgo.v2"

type runtime struct {
	processes *processes
}

func NewRuntime(session *mgo.Session) *runtime {
	r := &runtime{
		processes: &processes{processes: make(map[string]*process), session: session.Clone()},
	}
	return r
}

func (rt *runtime) Add() {

}

func (rt *runtime) GetForm() {

}
