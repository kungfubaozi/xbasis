package permissionhandlers

import (
	"github.com/garyburd/redigo/redis"
	"gopkg.in/mgo.v2"
	"konekko.me/gosion/commons/generator"
	"konekko.me/gosion/permission/utils"
)

type bindingRepo struct {
	session *mgo.Session
	conn    redis.Conn
	id      gs_commons_generator.IDGenerator
}

func (repo *bindingRepo) GetFunctionRoleMembers(structureId, functionId string) ([]interface{}, error) {
	return redis.Values(repo.conn.Do("SMEMBERS", permissionutils.GetStructureFunctionRoleKey(structureId, functionId)))
}

func (repo *bindingRepo) GetUserRoleMembers(structureId, userId string) ([]interface{}, error) {
	return redis.Values(repo.conn.Do("SMEMBERS", permissionutils.GetStructureUserRoleKey(structureId, userId)))
}

func (repo *bindingRepo) SetUserRoleMembersInCache(userId, structureId string, roles []string) error {
	_, err := repo.conn.Do("sadd", permissionutils.GetStructureUserRoleKey(structureId, userId), roles)
	return err
}

func (repo *bindingRepo) SetUserRoleMembersInDB(userId, structureId string, roles []string) {

}

func (repo *bindingRepo) Exists(structureId, roleId string) (bool, error) {
	return redis.Bool(repo.conn.Do("hexists", permissionutils.GetStructureRoleKey(structureId), roleId))
}

func (repo *bindingRepo) Close() {
	repo.conn.Close()
	repo.session.Close()
}
