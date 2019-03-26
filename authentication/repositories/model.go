package authentication_repositories

type SimpleUserToken struct {
	UserId   string
	AppId    string
	ClientId string
	Relation string
	ExpireAt int64 //expire at time ; nanoTime
}
