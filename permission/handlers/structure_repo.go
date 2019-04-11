package permissionhandlers

import (
	"errors"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"konekko.me/gosion/commons/indexutils"
)

type structureRepo struct {
	session *mgo.Session
	*indexutils.Client
}

func (repo *structureRepo) GetCurrent() (string, error) {
	return "", nil
}

func (repo *structureRepo) Add(s *structure) error {
	return repo.collection().Insert(s)
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
	//_, err := repo.conn.Do("set", permissionutils.GetTypeCurrentStructureKey(appId, t), id)
	//return err
	panic(errors.New(""))
}

func (repo *structureRepo) collection() *mgo.Collection {
	return repo.session.DB(dbName).C(structureCollection)
}

func (repo *structureRepo) Close() {
	if repo.session != nil {
		repo.session.Close()
	}
}
