package authentication_repositories

import "github.com/garyburd/redigo/redis"

type TokenRepo struct {
	Conn redis.Conn
}

func (repo *TokenRepo) Get() {

}

func (repo *TokenRepo) findByUserAndClientId(userId, clientId string) {
	repo.Conn.Do("hmget", "line."+userId)
}
