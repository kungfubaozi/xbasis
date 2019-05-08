package models

import "konekko.me/gosion/workflow/types"

//设置哪些人可以发起一个流程
//在中间task处理时不会判断此权限，只会在发起时判断权限是否符合
type NodeEvent struct {
	Id                 string   `bson:"id" json:"id"`
	Name               string   `bson:"name" json:"name"`
	Desc               string   `bson:"desc" json:"desc"`
	CreateAt           int64    `bson:"create_at" json:"create_at"`
	ExecutionListeners []string `bson:"execution_listeners" json:"execution_listeners"`
	UserRoles          []string `bson:"user_roles" json:"user_roles"`
	UserGroups         []string `bson:"user_groups" json:"user_groups"`
	UserIds            []string `bson:"user_ids" json:"user_ids"`
}

type StartEvent struct {
	*NodeEvent
	FormRef string `bson:"form_ref" json:"form_ref"`
}

//忽略NodeEvent设置的权限
type EndEvent struct {
	*NodeEvent
}

//忽略NodeEvent设置的权限
type CancelEndEvent struct {
	*NodeEvent
}

//忽略NodeEvent设置的权限
type MessageStartEvent struct {
	*NodeEvent
}

//忽略NodeEvent设置的权限
type TerminateEndEvent struct {
	*NodeEvent
}

type ApiStartEvent struct {
	*NodeEvent
	Method  types.HttpRequestMethod `bson:"method" json:"method"`
	FormRef string                  `bson:"form_ref" json:"form_ref"`
}

//忽略NodeEvent设置的权限
type TimerStartEvent struct {
	*NodeEvent
	Corn string `bson:"corn" json:"corn"`
}

//忽略NodeEvent设置的权限
type TriggerStartEvent struct {
	*NodeEvent
	Trigger string `bson:"trigger" json:"trigger"`
}
