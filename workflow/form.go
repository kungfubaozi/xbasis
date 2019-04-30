package workflow

type FiledType int64

const (
	VTTitleView FiledType = iota

	//int32 int64 float32 float64
	VTNumberView

	//string
	VTEditView

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
	VTFileView

	//a href
	VTALinkView

	//line
	VTLineView

	//top left right bottom space(margin)
	VTSpaceView

	//face check(self)
	VTSelfFaceView

	//user sign
	VTDrawView

	//validate code(email, phone)
	VTValidateCodeView
)

type typeForm struct {
	Basic        *basicModel  `bson:"basic" json:"basic"`
	CreateUserId string       `json:"create_user_id" bson:"create_user_id"`
	Require      bool         `bson:"require" json:"require"`
	Fields       []*typeField `bson:"fields" json:"fields"`
}

type typeField struct {
	Id         string      `bson:"id" json:"id"`
	Name       string      `bson:"name" json:"name"`
	Type       FiledType   `json:"type" bson:"type"`
	FieldText  string      `bson:"field_text" json:"field_text"`
	Properties interface{} `bson:"properties" json:"properties"`
	Require    bool        `bson:"require" json:"require"`
	Label      string      `bson:"label" json:"label"`
}

type timeViewProperties struct {
	DefaultNow bool  `bson:"default_now" json:"default_now"`
	MaxTime    int64 `bson:"max_time" json:"max_time"`
	MinTime    int64 `bson:"min_time" json:"min_time"`
}

type selectViewProperties struct {
	MaxSelect     int64         `bson:"max_select" json:"max_select"`
	MinSelect     int64         `bson:"min_select" json:"min_select"`
	DefaultValues []interface{} `bson:"default_values" json:"default_values"`
}

type editViewProperties struct {
	MaxLength    int64       `bson:"max_length" json:"max_length"`
	MinLength    int64       `bson:"min_length" json:"min_length"`
	DefaultValue interface{} `bson:"default_value" json:"default_value"`
	Regx         string      `bson:"regx" json:"regx"` //正则匹配
}

type fileViewProperties struct {
	MaxSize int64    `bson:"max_size" json:"max_size"`
	MinSize int64    `bson:"min_size" json:"min_size"`
	Format  []string `bson:"format" json:"format"`
}

type spaceViewProperties struct {
	Left   int64 `bson:"left" json:"left"`
	Top    int64 `bson:"top" json:"top"`
	Bottom int64 `bson:"bottom" json:"bottom"`
	Right  int64 `bson:"right" json:"right"`
}

type submitForm struct {
	Basic          *basicModel `bson:"basic" json:"basic"`
	FormId         string      `bson:"form_id" json:"form_id"`
	SubmitByUserId string      `bson:"submit_by_user_id" json:"submit_by_user_id"`
	SubmitAtFlow   string      `bson:"submit_at_flow" json:"submit_at_flow"`
	Data           interface{} `bson:"data" json:"data"`
}

//sha1(userId+formId+propId)
type validateCodeViewProperties struct {
	Type int64  `bson:"type" json:"type"`
	Id   string `bson:"id" json:"id"` //发送验证码时需要(针对性发送)
}

type typeFieldValue struct {
	Key    string
	Values []interface{}
}

type form struct {
}
