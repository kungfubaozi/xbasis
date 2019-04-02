package permissionhandlers

import (
	"github.com/garyburd/redigo/redis"
	"gopkg.in/mgo.v2"
)

type structureRepo struct {
	session *mgo.Session
	conn    redis.Conn
}

func (repo *structureRepo) GetCurrent() (string, error) {

}
