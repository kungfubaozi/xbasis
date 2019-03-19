package safety_repositories

import (
	"github.com/garyburd/redigo/redis"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"strconv"
	"time"
)

type Blacklist struct {
	Type         int64  `bson:"type"`
	Content      string `bson:"content"`
	CreateAt     int64  `bson:"create_at"`
	CreateUserId string `bson:"create_user_id"`
}

type BlacklistRepo struct {
	Session *mgo.Session
	Conn    redis.Conn
}

func (repo *BlacklistRepo) Save(bt int64, content, userId string) error {
	_, err := repo.Conn.Do("hmset", "blacklist-"+strconv.FormatInt(bt, 10), content, "denied")
	if err != nil {
		return err
	}

	b := &Blacklist{
		Type:         bt,
		Content:      content,
		CreateAt:     time.Now().UnixNano(),
		CreateUserId: userId,
	}

	return repo.collection().Insert(b)
}

func (repo *BlacklistRepo) Remove(id string) error {

	var bl Blacklist
	err := repo.collection().Find(bson.M{"_id": id}).One(&bl)
	if err != nil {
		return err
	}

	_, err = repo.Conn.Do("hdel", "blacklist-"+strconv.FormatInt(bl.Type, 10), bl.Content)

	if err != nil && err == redis.ErrNil {
		err = nil
	}
	if err != nil {
		return err
	}
	return repo.collection().Remove(bson.M{"_id": id})
}

func (repo *BlacklistRepo) CacheExists(bt int64, content string) bool {
	_, err := redis.String(repo.Conn.Do("hexists", "blacklist-"+strconv.FormatInt(bt, 10), content))
	if err != nil && err == redis.ErrNil {
		return false
	}
	if err != nil {
		return true
	}
	return true
}

func (repo *BlacklistRepo) Close() {
	repo.Conn.Close()
	repo.Session.Close()
}

func (repo *BlacklistRepo) collection() *mgo.Collection {
	return repo.Session.DB("gosion").C("blacklist")
}
