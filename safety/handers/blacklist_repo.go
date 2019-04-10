package safetyhanders

import (
	"fmt"
	"github.com/garyburd/redigo/redis"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"strconv"
	"time"
)

type blacklistRepo struct {
	session *mgo.Session
	conn    redis.Conn
}

func (repo *blacklistRepo) Save(bt int64, content, userId string) error {
	_, err := repo.conn.Do("hmset", "blacklist-"+strconv.FormatInt(bt, 10), content, "denied")
	if err != nil {
		return err
	}

	b := &blacklist{
		Type:         bt,
		Content:      content,
		CreateAt:     time.Now().UnixNano(),
		CreateUserId: userId,
	}

	return repo.collection().Insert(b)
}

func (repo *blacklistRepo) Remove(id string) error {

	var bl blacklist
	err := repo.collection().Find(bson.M{"_id": id}).One(&bl)
	if err != nil {
		return err
	}

	_, err = repo.conn.Do("hdel", "blacklist-"+strconv.FormatInt(bl.Type, 10), bl.Content)

	if err != nil && err == redis.ErrNil {
		err = nil
	}
	if err != nil {
		return err
	}
	return repo.collection().Remove(bson.M{"_id": id})
}

func (repo *blacklistRepo) CacheExists(bt int64, content string) bool {
	e, err := redis.Bool(repo.conn.Do("hexists", "blacklist-"+strconv.FormatInt(bt, 10), content))
	if err != nil {
		fmt.Println("Err", err)
		return true
	}
	return e
}

func (repo *blacklistRepo) Close() {
	repo.conn.Close()
	repo.session.Close()
}

func (repo *blacklistRepo) collection() *mgo.Collection {
	return repo.session.DB(dbName).C(blacklistCollection)
}
