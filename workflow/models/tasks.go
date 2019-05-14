package models

import "konekko.me/gosion/workflow/types"

//提供对外的API
type ApiTask struct {
	*Info
	RequestMethod types.HttpRequestMethod `bson:"request_method" json:"request_method"`
}

//决策
type DecisionTask struct {
	*Info
}

//向外HTTP请求
type HttpTask struct {
	*Info
	RequestMethod types.HttpRequestMethod `bson:"request_method" json:"request_method"`
}

type MailTask struct {
	*Info
}

type StorageTask struct {
	*Info
	Collection string `bson:"collection" json:"collection"`
}

type UserTask struct {
	*Info
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
}
