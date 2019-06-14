package permissionhandlers

import (
	"fmt"
	"gopkg.in/mgo.v2"
	"konekko.me/gosion/commons/hashcode"
)

func getURFIndex(fix string) string {
	f := "*"
	if fix != "*" {
		fix = fmt.Sprintf("%d", hashcode.Get(fix)%5)
	}
	return fmt.Sprintf("gosion_urf_relations.%s", f)
}

func mgoignore(err error) bool {
	if err != nil {
		if err == mgo.ErrNotFound {
			return true
		}
	}
	return true
}
