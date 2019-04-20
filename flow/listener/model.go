package listener

type TaskListener struct {
	Id       string `bson:"id" json:"id"`
	Name     string `bson:"name" json:"name"`
	Desc     string `bson:"desc" json:"desc"`
	CreateAt int64  `bson:"create_at" json:"create_at"`
	Listener string `bson:"listener" json:"listener"`
}

type ExecutionListener struct {
	Id       string `bson:"id" json:"id"`
	Name     string `bson:"name" json:"name"`
	Desc     string `bson:"desc" json:"desc"`
	CreateAt int64  `bson:"create_at" json:"create_at"`
	Listener string `bson:"listener" json:"listener"`
}
