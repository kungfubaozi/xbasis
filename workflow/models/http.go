package models

import "konekko.me/gosion/workflow/types"

//异步任务
//当执行异步任务时，会默认挂载一个event，当目标任务执行完成后发送一个对应的event通知，流程继续
type AsyncCallTask struct {
	*Info
	Target  string `bson:"target" json:"target"`
	Trigger string `bson:"trigger" json:"trigger"` //触发此task完成
}

//向外HTTP请求
type HttpTask struct {
	*Info
	RequestMethod types.HttpRequestMethod `bson:"request_method" json:"request_method"`
}

type StorageTask struct {
	*Info
	Collection string `bson:"collection" json:"collection"`
}

//提供对外的API
//是个单实例服务
type ApiTask struct {
	*Info
	RequestMethod types.HttpRequestMethod `bson:"request_method" json:"request_method"`
}
