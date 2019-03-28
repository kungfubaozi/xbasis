package authentication_repositories

type SimpleUserToken struct {
	UserId   string
	AppId    string
	ClientId string
	Relation string
	Type     int64
	ExpireAt int64 //expire at time ; nanoTime
}

type UserAuthorizeInfo struct {
	Relation  string
	UserId    string
	Ip        string
	Device    string
	Platform  int64
	UserAgent string
	AppId     string
	ClientId  string
}
