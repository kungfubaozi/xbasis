package permissionhandlers

import (
	"fmt"
	"gopkg.in/mgo.v2"
	"konekko.me/xbasis/commons/hashcode"
)

func getURFIndex(fix string) string {
	f := "*"
	if fix != "*" {
		fix = fmt.Sprintf("%d", hashcode.Equa(fix))
	}
	return fmt.Sprintf("xbs-urf-relations.%s", f)
}

func mgoignore(err error) bool {
	if err != nil {
		if err == mgo.ErrNotFound {
			return true
		}
	}
	return true
}
