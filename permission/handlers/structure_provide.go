package permissionhandlers

import (
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

func isStructureExists(session *mgo.Session, id string) int {
	c, err := session.DB(dbName).C(structureCollection).Find(bson.M{"_id": id}).Count()
	if err != nil {
		return 0
	}
	return c
}
