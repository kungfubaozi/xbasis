package permission_uitls

import "fmt"

func GetAppRoleKey(appId string) string {
	return fmt.Sprintf("a/r.%s", appId)
}

func GetAppUserRoleKey(appId, userId string) string {
	return fmt.Sprintf("a/u/r.%s-%s", appId, userId)
}

func GetAppFunctionRoleKey(appId, functionId string) string {
	return fmt.Sprintf("a/f/r.%s-%s", appId, functionId)
}
