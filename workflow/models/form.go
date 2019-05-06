package models

import (
	"konekko.me/gosion/workflow/types"
)

type TypeForm struct {
	*Info
	CreateUserId string       `json:"create_user_id" bson:"create_user_id"`
	Require      bool         `bson:"require" json:"require"`
	Fields       []*TypeField `bson:"fields" json:"fields"`
}

type TypeField struct {
	Id         string          `bson:"id" json:"id"`
	Name       string          `bson:"name" json:"name"`
	Type       types.FiledType `json:"type" bson:"type"`
	FieldText  string          `bson:"field_text" json:"field_text"`
	Properties interface{}     `bson:"properties" json:"properties"`
	Require    bool            `bson:"require" json:"require"`
	Label      string          `bson:"label" json:"label"`
}

type TimeViewProperties struct {
	DefaultNow bool  `bson:"default_now" json:"default_now"`
	MaxTime    int64 `bson:"max_time" json:"max_time"`
	MinTime    int64 `bson:"min_time" json:"min_time"`
}

type SelectViewProperties struct {
	MaxSelect     int64         `bson:"max_select" json:"max_select"`
	MinSelect     int64         `bson:"min_select" json:"min_select"`
	DefaultValues []interface{} `bson:"default_values" json:"default_values"`
}

type EditViewProperties struct {
	MaxLength    int64       `bson:"max_length" json:"max_length"`
	MinLength    int64       `bson:"min_length" json:"min_length"`
	DefaultValue interface{} `bson:"default_value" json:"default_value"`
	Regx         string      `bson:"regx" json:"regx"` //正则匹配
}

type FileViewProperties struct {
	MaxSize int64    `bson:"max_size" json:"max_size"`
	MinSize int64    `bson:"min_size" json:"min_size"`
	Format  []string `bson:"format" json:"format"`
}

type SpaceViewProperties struct {
	Left   int64 `bson:"left" json:"left"`
	Top    int64 `bson:"top" json:"top"`
	Bottom int64 `bson:"bottom" json:"bottom"`
	Right  int64 `bson:"right" json:"right"`
}

type SubmitForm struct {
	*Info
	FormId         string      `bson:"form_id" json:"form_id"`
	SubmitByUserId string      `bson:"submit_by_user_id" json:"submit_by_user_id"`
	SubmitAtFlow   string      `bson:"submit_at_flow" json:"submit_at_flow"`
	Data           interface{} `bson:"data" json:"data"`
}

//sha1(userId+formId+propId)
type ValidateCodeViewProperties struct {
	Type int64  `bson:"type" json:"type"`
	Id   string `bson:"id" json:"id"` //发送验证码时需要(针对性发送)
}

type TypeFieldValue struct {
	Key    string
	Values []interface{}
}
