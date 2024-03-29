package models

//决策
type DecisionTask struct {
	*Info
}

//邮件服务
type MailTask struct {
	*Info
	Headers map[string]string `bson:"headers" json:"headers"`
	To      string            `bson:"to" json:"to"`
	From    string            `bson:"from" json:"from"`
	Subject string            `bson:"subject" json:"subject"`
	Html    string            `bson:"html" json:"html"`
}

type UserTask struct {
	*Info
	UserIds    []string `bson:"user_ids" json:"user_ids"`
	UserGroups []string `bson:"user_groups" json:"user_groups"`
	UserRoles  []string `bson:"user_roles" json:"user_roles"`
	FormRef    string   `bson:"form_ref" json:"form_ref"`
	NextSize   int64    `bson:"next_size" json:"next_size"` //需要多少人签批可以继续下一步
}

//协助任务
//虽然usertask差不多，但是协助人并不参与到整体流程中，只是提供一些意见给其他人
type AssistTask struct {
	UserIds    []string `bson:"user_ids" json:"user_ids"`
	UserGroups []string `bson:"user_groups" json:"user_groups"`
	UserRoles  []string `bson:"user_roles" json:"user_roles"`
	FormRef    string   `bson:"form_ref" json:"form_ref"`
	NextSize   int64    `bson:"next_size" json:"next_size"` //需要多少人签批可以继续下一步
}

//通知其他不参与此项流程的用户
//被通知的用户可以查看当前流程，
type NotifyTask struct {
	*Info
	UserIds    []string `bson:"user_ids" json:"user_ids"`
	UserGroups []string `bson:"user_groups" json:"user_groups"`
	UserRoles  []string `bson:"user_roles" json:"user_roles"`
}

//定时任务
type TimerTask struct {
	*Info
}

//事件任务
type EventTask struct {
	*Info
	Event string `bson:"event" json:"event"`
}
