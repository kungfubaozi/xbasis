package safetyhanders

import (
	"context"
	"errors"
	"fmt"
	"github.com/garyburd/redigo/redis"
	"github.com/olivere/elastic"
	"gopkg.in/mgo.v2"
	"strconv"
	"time"
)

type blacklistRepo struct {
	session       *mgo.Session
	elasticClient *elastic.Client
}

func (repo *blacklistRepo) Save(bt int64, content, userId string) error {
	b := &blacklist{
		Type:         bt,
		Content:      content,
		CreateAt:     time.Now().UnixNano(),
		CreateUserId: userId,
	}

	v, err := repo.elasticClient.Index().Index("gs_safety_blacklist").Type("v").BodyJson(b).Do(context.Background())
	if err != nil || v.Status == 0 {
		return err
	}

	return nil
}

func (repo *blacklistRepo) Remove(id string) error {
	q := elastic.NewMatchQuery("content", id)
	v, err := repo.elasticClient.Search("gs_safety_blacklist").Type("v").Query(q).Do(context.Background())
	if err != nil {
		return err
	}
	if v.Hits.TotalHits > 0 {
		d := v.Hits.Hits[0]
		r, err := repo.elasticClient.Delete().Index("gs_safety_blacklist").Type("v").Id(d.Id).Do(context.Background())
		if err != nil {
			return err
		}
		if r.Status != 0 {
			return errors.New("")
		}
		return nil
	}
	return errors.New("not found")
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
	if repo.session != nil {
		repo.session.Close()
	}
}

func (repo *blacklistRepo) collection() *mgo.Collection {
	return repo.session.DB(dbName).C(blacklistCollection)
}
