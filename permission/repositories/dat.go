package permission_repositories

type DurationAccess struct {
	UserId    string
	Path      string
	ExpiredAt int64
	ClientId  string
	CreateAt  int64
}
