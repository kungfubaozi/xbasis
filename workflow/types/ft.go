package types

type FiledType int64

const (
	FTTitleView FiledType = iota

	//int32 int64 float32 float64
	FTNumberView

	//string
	FTEditView

	//time
	FTTimeView

	//select user
	FTUserSelectView

	//select user group
	FTUserGroupSelectView

	//value value(number, string, user) select
	FTValueSelectView

	//encrypt string
	FTPasswordView

	//single select(dropdown)
	FTDropdownView

	//upload file
	FTFileView

	//a href
	FTALinkView

	//line
	FTLineView

	//top left right bottom space(margin)
	FTSpaceView

	//face check(self)
	FTSelfFaceView

	//user sign
	FTDrawView

	//validate code(email, phone)
	FTValidateCodeView

	//radio group view
	FTRadioGroupView
)
