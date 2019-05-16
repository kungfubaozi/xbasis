package models

type Holder struct {
	Key    map[string]interface{} `json:"key"`
	Status int64                  `json:"status"`
}
