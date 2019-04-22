package workflow

import "konekko.me/gosion/workflow/activities"

type FiledType int64

const (
	VTTitleView FiledType = iota

	//int32 int64 float32 float64
	VTNumberView

	//string
	VTTextView

	//time
	VTTimeView

	//select user
	VTUserSelectView

	//select user group
	VTUserGroupSelectView

	//value value(number, string, user) select
	VTValueSelectView

	//encrypt string
	VTPasswordView

	//single select(dropdown)
	VTDropdownView

	//upload file
	VTUploadFileView

	//a href
	VTALinkViewbushi

	//line
	VTLineView

	//top left right bottom space(margin)
	VTSpaceView

	//face check(self)
	VTSelfFaceView

	//user sign
	VTDrawView
)

type TypeForm struct {
	Basic        *BasicModel `bson:"basic" json:"basic"`
	CreateUserId string      `json:"create_user_id" bson:"create_user_id"`
	Require      bool        `bson:"require" json:"require"`
	Fields       []TypeField `bson:"fields" json:"fields"`
}

type TypeField struct {
	Id         string      `bson:"id" json:"id"`
	Name       string      `bson:"name" json:"name"`
	Type       FiledType   `json:"type" bson:"type"`
	FieldText  string      `bson:"field_text" json:"field_text"`
	Properties interface{} `bson:"properties" json:"properties"`
}

type TimeViewProperties struct {
	DefaultNow bool  `bson:"default_now" json:"default_now"`
	MaxTime    int64 `bson:"max_time" json:"max_time"`
	MinTime    int64 `bson:"min_time" json:"min_time"`
	Require    bool  `bson:"require" json:"require"`
}

type SelectViewProperties struct {
	Max           float64     `bson:"max" json:"max"`
	Min           float64     `bson:"min" json:"min"`
	Require       bool        `bson:"require" json:"require"`
	DefaultValues interface{} `bson:"default_values" json:"default_values"`
}

type TextViewProperties struct {
	MaxLength    int64       `bson:"max_length" json:"max_length"`
	MinLength    int64       `bson:"min_length" json:"min_length"`
	DefaultValue interface{} `bson:"default_value" json:"default_value"`
	Require      bool        `bson:"require" json:"require"`
}

type UploadFileViewProperties struct {
	MaxSize int64    `bson:"max_size" json:"max_size"`
	MinSize int64    `bson:"min_size" json:"min_size"`
	Format  []string `bson:"format" json:"format"`
}

type SpaceViewProperties struct {
	Left   float64 `bson:"left" json:"left"`
	Top    float64 `bson:"top" json:"top"`
	Bottom float64 `bson:"bottom" json:"bottom"`
	Right  float64 `bson:"right" json:"right"`
}

//获取对应form信息
func getFormById(id string) {

}

func getFormAllField() (map[string]interface{}, error) {

}

func getFormValue(flowId, formId string) (map[string]interface{}, error) {

}
