package authenticationhandlers

type simpleUserToken struct {
	UserId    string
	AppId     string
	ClientId  string
	Relation  string
	Type      int64
	Structure string
}

type userAuthorizeInfo struct {
	Relation  string
	UserId    string
	Ip        string
	Device    string
	Platform  int64
	UserAgent string
	AppId     string
	ClientId  string
}