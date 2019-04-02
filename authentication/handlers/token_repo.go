package authenticationhandlers

import (
	"fmt"
	"github.com/garyburd/redigo/redis"
)

type tokenRepo struct {
	conn redis.Conn
}

func (repo *tokenRepo) fix(userId string) string {
	return fmt.Sprintf("auth-u.%s", userId)
}

func (repo *tokenRepo) Get(userId, key string) ([]byte, error) {
	return redis.Bytes(repo.conn.Do("hmget", repo.fix(userId), key))
}

func (repo *tokenRepo) Remove(userId, key string) error {
	_, err := repo.conn.Do("hdel", repo.fix(userId), key)
	return err
}

func (repo *tokenRepo) Add(userId, clientId, relation string, b []byte) error {
	_, err := repo.conn.Do("hmset", repo.fix(userId), fmt.Sprintf("%s.%s", clientId, relation), b)
	return err
}

func (repo *tokenRepo) Close() {
	repo.conn.Close()
}

func (repo *tokenRepo) SizeOf(userId string) ([]interface{}, error) {
	v, err := redis.Values(repo.conn.Do("hkeys", repo.fix(userId)))
	if err == redis.ErrNil {
		err = nil
	}
	if err != nil {
		return nil, err
	}
	if len(v) == 0 {
		return nil, nil
	}
	return v, nil
}
