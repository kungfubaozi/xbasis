package task

type Activity interface {
	Do(values interface{})

	Name()

	String()

	Id()
}
