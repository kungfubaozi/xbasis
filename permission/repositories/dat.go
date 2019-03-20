package permission_repositories

import "github.com/garyburd/redigo/redis"

type DurationAccessRepo struct {
	Conn redis.Conn
}

type DurationAccess struct {
	UserId    string
	Path      string
	ExpiredAt int64
	AppId     string
}
