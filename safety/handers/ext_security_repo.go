package safetyhanders

import (
	"github.com/garyburd/redigo/redis"
	"gopkg.in/mgo.v2"
)

type securityRepo struct {
	session *mgo.Session
	conn    redis.Conn
}
