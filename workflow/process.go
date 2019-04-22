package workflow

type process struct {
	Id           string `bson:"_id" json:"id"`
	Name         string `bson:"name" json:"name"`
	CreateUserId string `bson:"create_user_id" json:"create_user_id"`
	CreateAt     int64  `bson:"create_at" json:"create_at"`
	Desc         string `bson:"desc" json:"desc"`
	//connect flows
	Flows []*sequenceFlow `bson:"flows" json:"flows"`
	//user tasks
	UserTasks []*userTask `bson:"user_tasks" json:"user_tasks"`
	//http tasks
	HttpTasks []*httpTask `bson:"http_tasks" json:"http_tasks"`
	//decision tasks
	DecisionTasks []*decisionTask `bson:"decision_tasks" json:"decision_tasks"`
	//send tasks
	SendTasks []*sendTask `bson:"send_tasks" json:"send_tasks"`
}

func (p *process) ToInstance() (*instance, error) {
	i := &instance{}
}

func newProcess(p *process) {

}
