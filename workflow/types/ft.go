package types


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

