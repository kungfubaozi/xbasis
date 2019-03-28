package authentication_repositories

import (
	"fmt"
	"github.com/garyburd/redigo/redis"
)

type TokenRepo struct {
	Conn redis.Conn
}

func (repo *TokenRepo) fix(userId string) string {
	return fmt.Sprintf("auth-u.%s", userId)
}

func (repo *TokenRepo) Get(userId, key string) ([]byte, error) {
	return redis.Bytes(repo.Conn.Do("hmget", repo.fix(userId), key))
}

func (repo *TokenRepo) Remove(userId, key string) error {
	_, err := repo.Conn.Do("hdel", repo.fix(userId), key)
	return err
}

func (repo *TokenRepo) Add(userId, clientId, relation string, b []byte) error {
	_, err := repo.Conn.Do("hmset", repo.fix(userId), fmt.Sprintf("%s.%s", clientId, relation), b)
	return err
}

func (repo *TokenRepo) Close() {
	repo.Conn.Close()
}

func (repo *TokenRepo) SizeOf(userId string) ([]interface{}, error) {
	v, err := redis.Values(repo.Conn.Do("hkeys", repo.fix(userId)))
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
