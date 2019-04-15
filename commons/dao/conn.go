package gs_commons_dao

import (
	"fmt"
	"github.com/garyburd/redigo/redis"
	"gopkg.in/mgo.v2"
)

func CreatePool(addr string) (*redis.Pool, error) {

	pool := &redis.Pool{
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", addr)

			fmt.Println("conn", err)

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
