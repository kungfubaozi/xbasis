package permissionhandlers

import (
	"github.com/garyburd/redigo/redis"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"konekko.me/gosion/commons/generator"
	"konekko.me/gosion/permission/utils"
	"time"
)

type roleRepo struct {
	session *mgo.Session
	id      gs_commons_generator.IDGenerator
	conn    redis.Conn
}

func (repo *roleRepo) FindByName(structureId, name string) (*role, error) {
	var r role
	err := repo.collection().Find(bson.M{"name": name, "structure_id": structureId}).One(&r)
	return &r, err
}

func (repo *roleRepo) Remove(structureId, roleId string) error {
	_, err := repo.conn.Do("hdel", permissionutils.GetStructureRoleKey(structureId), roleId)
	if err != nil && err == redis.ErrNil {
		err = nil
	}
	if err != nil {
		return err
	}
	return repo.collection().Remove(bson.M{"structure_id": structureId, "_id": roleId})
}

func (repo *roleRepo) Save(name, structureId, userId string) error {
	r := &role{
		Id:           repo.id.Get(),
		CreateAt:     time.Now().UnixNano(),
		Name:         name,
		StructureId:  structureId,
		CreateUserId: userId,
	}
	err := repo.collection().Insert(r)
	if err != nil {
		return err
	}
	_, err = repo.conn.Do("hmset", permissionutils.GetStructureRoleKey(structureId), r.Id, r.Name)
	return err
}

func (repo *roleRepo) FindRoleById(roleId string) (*role, error) {
	var role *role
	err := repo.collection().Find(bson.M{"_id": roleId}).One(&role)
	return role, err
}

func (repo *roleRepo) FindRolesByStructure(structureId string, page, size int64) ([]*role, error) {
	var roles []*role
	err := repo.collection().Find(bson.M{"structure_id": structureId}).Limit(int(size)).Skip(int(page * size)).All(&roles)
	return roles, err
}

func (repo *roleRepo) collection() *mgo.Collection {
	return repo.session.DB(dbName).C(roleCollection)
}

func (repo *roleRepo) Close() {
	repo.session.Close()
	if repo.conn != nil {
		repo.conn.Close()
	}
}
