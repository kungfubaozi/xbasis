package event

//触发器
type TriggerStartEvent struct {
	typeEvent
	Trigger string `bson:"trigger" json:"trigger"`
}
