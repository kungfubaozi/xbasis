package flowgateway

//排他网关
type ExclusiveGateway struct {
	Id       string `bson:"id" json:"id"`
	Name     string `bson:"name" json:"name"`
	Desc     string `bson:"desc" json:"desc"`
	CreateAt int64  `bson:"create_at" json:"create_at"`
}

//并行网关
type ParallelGateway struct {
	Id       string `bson:"id" json:"id"`
	Name     string `bson:"name" json:"name"`
	Desc     string `bson:"desc" json:"desc"`
	CreateAt int64  `bson:"create_at" json:"create_at"`
}

//包容网关
type InclusiveGateway struct {
	Id       string `bson:"id" json:"id"`
	Name     string `bson:"name" json:"name"`
	Desc     string `bson:"desc" json:"desc"`
	CreateAt int64  `bson:"create_at" json:"create_at"`
}

//事件网关
type EventGateway struct {
	Id       string `bson:"id" json:"id"`
	Name     string `bson:"name" json:"name"`
	Desc     string `bson:"desc" json:"desc"`
	CreateAt int64  `bson:"create_at" json:"create_at"`
}
