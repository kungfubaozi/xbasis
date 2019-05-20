package models

type Holder struct {
	Status        int      `json:"status"`
	RelationNodes []string `json:"relation_nodes"`
	NodeId        string   `json:"node_id"`     //当前节点
	InstanceId    string   `json:"instance_id"` //实例ID
}

//在进行流程走向判读时，比如等待，需要查找当前实例中所有的忽略节点，不需要加入Gateway
type NodeIgnore struct {
	InstanceId string   `bson:"instance_id" json:"instance_id"`
	GatewayId  string   `bson:"gateway_id" json:"gateway_id"`
	Ignores    []string `bson:"ignores" json:"ignores"` //忽略的节点
}
