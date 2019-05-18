package models

type Holder struct {
	Status      int      `json:"status"`
	ParentNodes []string `json:"parent_nodes"` //所有父节点
	NodeId      string   `json:"node_id"`      //当前节点
	InstanceId  string   `json:"instance_id"`  //实例ID
}
