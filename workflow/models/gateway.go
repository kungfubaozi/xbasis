package models

//包容网关
//包容网关会寻找所有符合条件的流向
//如果没有符合的，那么就去走默认的flow
type InclusiveGateway struct {
	*Info
	Exclusive   bool `bson:"exclusive" json:"exclusive"` //排他
	ScriptFlows int  `bson:"script_flows" json:"script_flows"`
}
