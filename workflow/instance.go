package workflow

//流程实例
//每启动一个流程都会产生一个对应的实例
type instance struct {
	id            string
	current       string
	startByUserId string //流程启动者
	startAt       int64
	processId     string //对应的流程
}

func (i *instance) next() {

}
