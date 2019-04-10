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

func GetStructureRoleKey(appId string) string {
	return fmt.Sprintf("s/r.%s", appId)
}

func GetStructureUserRoleKey(appId, userId string) string {
	return fmt.Sprintf("s/u/r.%s-%s", appId, userId)
}

func GetStructureFunctionRoleKey(appId, functionId string) string {
	return fmt.Sprintf("s/f/r.%s-%s", appId, functionId)
}

func GetTypeStructureKey(structureId string, t int64) string {
	return fmt.Sprintf("a/s/t.%s-%d", structureId, t)
}

func GetTypeCurrentStructureKey(structureId string, t int64) string {
	return fmt.Sprintf("a/c/s/u.%s-%d", structureId, t)
}
