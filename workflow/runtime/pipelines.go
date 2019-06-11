package runtime

import (
	"github.com/coocood/freecache"
	"github.com/garyburd/redigo/redis"
	"github.com/vmihailenco/msgpack"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"konekko.me/gosion/analysis/client"
	"konekko.me/gosion/workflow/modules"
	"runtime/debug"
)

type pipelines struct {
	processes processes
	session   *mgo.Session
	pool      *redis.Pool
	log       analysisclient.LogClient
	c         *freecache.Cache
}

func (p *pipelines) Update(processId string) error {
	panic("implement me")
}

func (p *pipelines) Get(processId string) (modules.Pipeline, error) {
	v, err := p.c.Get([]byte(processId))
	var pipe *pipeline
	conn := p.pool.Get()
	version, err := redis.Int64(conn.Do("hget", "pro_vers", processId))
	if err != nil {
		return nil, err
	}

	getNewVersion := func() error {
		err = p.session.DB("gs-flow").C("pipelines").Find(bson.M{"_id": processId}).One(&pipe)
		if err != nil {
			return err
		}
		//set cache
		b, err := msgpack.Marshal(pipe)
		if err != nil {
			return err
		}
		err = p.c.Set([]byte(processId), b, 60*60*24) //24小时
		if err != nil {
			return err
		}
		v = b
		return nil
	}

	if err != nil && err == freecache.ErrNotFound {
		//load in database
		err = getNewVersion()
		if err != nil {
			return nil, err
		}
	}

	if err != nil {
		return nil, err
	}

	//from cache
	if pipe == nil {
		err = msgpack.Unmarshal(v, pipe)
		if err != nil {
			return nil, err
		}
		if version != pipe.Version {
			err = getNewVersion()
			if err != nil {
				return nil, err
			}
		}
	}
	return pipe, err
}

func newPipelines(session *mgo.Session,
	log analysisclient.LogClient, pool *redis.Pool) modules.Pipelines {
	cacheSize := 100 * 1024 * 1024
	cache := freecache.NewCache(cacheSize)
	debug.SetGCPercent(20)

	return &pipelines{c: cache, session: session, pool: pool, log: log}
}
