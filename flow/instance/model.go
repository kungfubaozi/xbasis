package flowinstance

type FlowInstance struct {
	Id       string `bson:"id" json:"id"`
	Name     string `bson:"name" json:"name"`
	Desc     string `bson:"desc" json:"desc"`
	CreateAt int64  `bson:"create_at" json:"create_at"`

	//StartEvent/TimerStartEvent/MessageStartEvent
	StartEvent     interface{} `bson:"start_event" json:"start_event"`
	StartEventType int64       `bson:"start_event_type"`

	//EndEvent
	EndEvent     interface{} `bson:"end_event" json:"end_event"`
	EndEventType int64       `bson:"end_event_type" json:"end_event_type"`
}
