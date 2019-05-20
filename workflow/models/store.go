package models

type Holder struct {
	Status        int      `json:"status"`
	RelationNodes []string `json:"relation_nodes"`
	NodeId        string   `json:"node_id"`     //当前节点
	InstanceId    string   `json:"instance_id"` //实例ID
}
