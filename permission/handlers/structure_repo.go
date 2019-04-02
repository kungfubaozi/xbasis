package permissionhandlers

import (
	"github.com/garyburd/redigo/redis"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"konekko.me/gosion/permission/utils"
)

type structureRepo struct {
	session *mgo.Session
	conn    redis.Conn
}

func (repo *structureRepo) GetCurrent() (string, error) {

}

func (repo *structureRepo) Add(s *structure) error {

}

func (repo *structureRepo) FindCountByNameAndType(name string, t int64) (int, error) {
	return repo.collection().Find(bson.M{"name": name, "type": t}).Count()
}

func (repo *structureRepo) FindById(id string) (*structure, error) {
	var s *structure
	err := repo.collection().Find(bson.M{"_id": id}).One(&s)
	return s, err
}

func (repo *structureRepo) Opening(id string, t int64, opening bool) error {
	if opening {
		//close all
		err := repo.collection().Update(bson.M{"type": t}, bson.M{"$set": bson.M{"opening": false}})
		if err != nil {
			return err
		}
	}
	return repo.collection().Update(bson.M{"_id": id}, bson.M{"$set": bson.M{"opening": opening}})
}

func (repo *structureRepo) OpeningCache(id, appId string, t int64) error {
	_, err := repo.conn.Do("set", permissionutils.GetTypeCurrentStructureKey(appId, t), id)
	return err
}

func (repo *structureRepo) collection() *mgo.Collection {
	return repo.session.DB("gosion").C("structure")
}

func (repo *structureRepo) Close() {
	repo.session.Close()
	repo.conn.Close()
}
