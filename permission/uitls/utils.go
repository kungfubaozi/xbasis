package permissionuitls

import "fmt"

func GetStructureRoleKey(appId string) string {
	return fmt.Sprintf("s/r.%s", appId)
}

func GetStructureUserRoleKey(appId, userId string) string {
	return fmt.Sprintf("s/u/r.%s-%s", appId, userId)
}

func GetStructureFunctionRoleKey(appId, functionId string) string {
	return fmt.Sprintf("s/f/r.%s-%s", appId, functionId)
}

func GetCurrentStructureIdKey(appId string) string {
	return fmt.Sprintf("a/s.%s", appId)
}
