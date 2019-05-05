package task

import "konekko.me/gosion/flow/base"

type ApiTask struct {
	*base.Info
	RequestMethod int64 `bson:"request_method" json:"request_method"`
}
