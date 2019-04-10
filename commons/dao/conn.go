package gs_commons_dao

import (
	"fmt"
	"github.com/garyburd/redigo/redis"
	"gopkg.in/mgo.v2"
	"time"
)

func CreatePool(addr string) (*redis.Pool, error) {

	pool := &redis.Pool{
		MaxIdle:     10,
		MaxActive:   20,
		IdleTimeout: 360 * time.Second,
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", addr)

			if err != nil {
				return nil, err
			}

			return c, err
		},
	}

	conn, err := pool.Dial()
	if err != nil {
		return nil, err
	}
	conn.Close()

	return pool, nil
}

// 建立与mongo的连接
func CreateSession(addr string) (*mgo.Session, error) {

	s, err := mgo.Dial(addr)

	if err != nil {
		fmt.Println("connect to mongo error:", err)
		return nil, err
	}

	s.SetMode(mgo.Monotonic, true)

	return s, err

}
