package task

import "konekko.me/gosion/flow/base"

type apiTask struct {
	*base.Info
	RequestMethod int64 `bson:"request_method" json:"request_method"`
}
