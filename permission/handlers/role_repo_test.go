package permissionhandlers

import (
	"fmt"
	"gopkg.in/mgo.v2/bson"
	"konekko.me/gosion/commons/dao"
	"sync"
	"testing"
)

func TestRoleRepo_Remove(t *testing.T) {

	session, err := gs_commons_dao.CreateSession("192.168.2.60:27017")
	if err != nil {
		panic(err)
	}
	defer session.Close()

	names, err := session.DB(dbName).CollectionNames()
	if err != nil {
		panic(err)
	}
	var relations []string
	for _, v := range names {
		if len(v) > 20 && v[:20] == "user_roles_relation_" {
			relations = append(relations, v)
		}
	}
	fmt.Println("relation", relations)
	var wg sync.WaitGroup
	wg.Add(len(relations))
	b := bson.M{"structure_id": "61323066-3030-6365-3232-393037373964"}
	u := bson.M{"$pull": bson.M{"roles": "32373134-3337-3431-3433-643265643835"}}

	resp := func(e error) {
		if err == nil {
			err = e
		}
	}

	for _, v := range relations {
		go func() {
			defer wg.Done()
			e := session.DB(dbName).C(v).Update(b, u)
			resp(e)
		}()
	}
	wg.Wait()

	if err != nil {
		panic(err)
	}
}
