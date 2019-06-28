package safetyhanders

import (
	"github.com/garyburd/redigo/redis"
	"github.com/vmihailenco/msgpack"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"sync"
)

type lockingRepo struct {
	session *mgo.Session
	conn    redis.Conn
}

func (repo *lockingRepo) Add(l *lockingUser) error {
	var wg sync.WaitGroup
	wg.Add(2)
	var err error

	resp := func(e error) {
		if err == nil {
			err = e
		}
	}

	go func() {
		defer wg.Done()
		b, e := msgpack.Marshal(l)
		if e != nil {
			resp(e)
			return
		}
		_, e = repo.conn.Do("hset", "locking", l.UserId, b)
		resp(e)
	}()

	go func() {
		defer wg.Done()
		resp(repo.collection().Insert(l))
	}()

	wg.Wait()

	return err
}

func (repo *lockingRepo) IsExists(userId string) (bool, error) {
	_, err := redis.Bytes(repo.conn.Do("hget", "locking", userId))
	if err != nil && err == redis.ErrNil {
		return false, nil
	}
	if err != nil {
		return false, err
	}
	return true, nil
}

func (repo *lockingRepo) collection() *mgo.Collection {
	return repo.session.DB(dbName).C("locking")
}

func (repo *lockingRepo) Remove(userId string) error {
	err := repo.collection().Remove(bson.M{"user_id": userId})
	if err != nil {
		return err
	}
	_, err = repo.conn.Do("del", "locking", userId)
	if err != nil && err == redis.ErrNil {
		return nil
	}
	return err
}

func (repo *lockingRepo) Close() {
	repo.session.Close()
	repo.conn.Close()
}
