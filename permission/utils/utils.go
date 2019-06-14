package permissionutils

import (
	"fmt"
)

const (
	// user structure
	TypeUserStructure = 2 << 5

	//function structure
	TypeFunctionStructure = 4 << 6
)

func GetAppRoleKey(appId string) string {
	return fmt.Sprintf("s/r.%s", appId)
}

func GetAppUserRoleKey(appId, userId string) string {
	return fmt.Sprintf("s/u/r.%s-%s", appId, userId)
}

func GetAppFunctionRoleKey(appId, functionId string) string {
	return fmt.Sprintf("s/f/r.%s-%s", appId, functionId)
}

func GetTypeAppKey(structureId string, t int64) string {
	return fmt.Sprintf("a/s/t.%s-%d", structureId, t)
}

func GetTypeCurrentAppKey(structureId string, t int64) string {
	return fmt.Sprintf("a/c/s/u.%s-%d", structureId, t)
}
