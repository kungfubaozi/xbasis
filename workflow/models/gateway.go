package models

//排他网关
//寻找第一条符合条件的流向
type ExclusiveGateway struct {
	*Info
}

//包容网关
//包容网关会寻找所有符合条件的流向
//如果没有符合的，那么就去走默认的flow
type InclusiveGateway struct {
	*Info
}

//并行网关
//允许将流程分成多条分支
//即使顺序流中定义了条件，也会被忽略
type ParallelGateway struct {
	*Info
}
