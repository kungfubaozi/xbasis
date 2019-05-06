package models

import "konekko.me/gosion/workflow/types"

//提供对外的API
type ApiTask struct {
	*Info
	RequestMethod types.HttpRequestMethod `bson:"request_method" json:"request_method"`
}

type DecisionTask struct {
	*Info
}

type GRPCTask struct {
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
}

//通知其他不参与此项流程的用户
//被通知的用户可以查看当前流程，
type NotifyTask struct {
	*Info
}

//流程页面缓存，任一flow的script都可以访问已设定的store
//已结束的实例会清除当前实例的所有store
type StoreTask struct {
	*Info
}
