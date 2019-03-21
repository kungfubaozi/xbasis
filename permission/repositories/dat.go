package permission_repositories

type DurationAccess struct {
	User          string
	Path          string
	Life          int64
	ClientId      string
	CreateAt      int64
	Stat          string
	Code          int64
	CodeExpiredAt int64
	Key           string
}
