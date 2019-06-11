package loggeractions

const (
	//访问App
	VisitApplication = "VisitApp"

	//未找到功能
	InvalidFunction = "InvalidFunction"

	//找到功能
	FindFunction = "FindFunction"

	//用户请求API(还不知道用户是谁)
	RequestApi = "RequestApi"

	//用户通过验证请求API（走完全部验证流程，知道用户是谁）
	UserRequestApi = "UserRequestApi"
)
