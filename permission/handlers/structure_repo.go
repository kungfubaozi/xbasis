package permissionhandlers

import (
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
	id, err := repo.AddData("gs_permission_structure", s)
	if err != nil {
		return err
	}
	if len(id) > 0 {
		s.SID = id
		return repo.collection().Insert(s)
	}
	return indexutils.ErrNotFound
}

func (repo *structureRepo) FindCountByNameAndType(name string, t int64) (int, error) {
	return repo.collection().Find(bson.M{"name": name, "type": t}).Count()
}

func (repo *structureRepo) FindById(id string) (*structure, error) {
	var s *structure
	err := repo.collection().Find(bson.M{"_id": id}).One(&s)
	return s, err
}

func (repo *structureRepo) Opening(appId, structureId, sid string, opening bool) error {
	ok, err := repo.Update("gs_permission_structure", sid, map[string]interface{}{"opening": opening})
	if err != nil {
		return err
	}
	if ok {
		return repo.collection().Update(bson.M{"_id": structureId}, bson.M{"$set": bson.M{"opening": opening}})
	}
	return indexutils.ErrNotFound
}

func (repo *structureRepo) setUserStructureConfig(appId, structureId string) {

}

func (repo *structureRepo) setFunctionStructureConfig(appId, structureId string) {

}

func (repo *structureRepo) collection() *mgo.Collection {
	return repo.session.DB(dbName).C(structureCollection)
}

func (repo *structureRepo) Close() {
	if repo.session != nil {
		repo.session.Close()
	}
}
