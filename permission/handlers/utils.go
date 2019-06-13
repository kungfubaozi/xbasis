package permissionhandlers

import (
	"fmt"
	"konekko.me/gosion/commons/hashcode"
)

func getURFIndex(fix string) string {
	f := "*"
	if fix != "*" {
		fix = fmt.Sprintf("%d", hashcode.Get(fix)%5)
	}
	return fmt.Sprintf("gosion_urf_relations.%s", f)
}
