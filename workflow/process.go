package workflow

import (
	"github.com/pkg/errors"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

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
	//version control
	Version int64 `bson:"version" json:"version"`
}

//流程实例
//用来控制实例的走向等
type processes struct {
	processes map[string]*process
	session   *mgo.Session
}

func (pi *processes) getProcess(id string) (*process, error) {
	if len(pi.processes) > 0 {
		p := pi.processes[id]
		if p == nil {
			var process *process
			err := pi.session.DB("gs_workflow").C("processes").Find(bson.M{"_id": id}).One(process)
			if err != nil {
				return nil, err
			}
			pi.processes[id] = process
			p = process
		}
		return p, nil
	}
	return nil, errors.New("not found")
}

func (pi *processes) next(i *instance) {

}

func (pi *processes) update(id string) {

}

func (pi *processes) close() {
	pi.session.Close()
}
