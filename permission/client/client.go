package permissioncli

import (
	"github.com/micro/go-micro/client"
	"konekko.me/xbasis/commons/constants"
	"konekko.me/xbasis/permission/pb"
	"konekko.me/xbasis/permission/pb/inner"
)

func NewVerificationClient(client client.Client) xbasissvc_internal_permission.VerificationService {
	return xbasissvc_internal_permission.NewVerificationService(xbasisconstants.InternalPermission, client)
}

func NewAccessibleClient(client client.Client) xbasissvc_internal_permission.AccessibleService {
	return xbasissvc_internal_permission.NewAccessibleService(xbasisconstants.InternalPermission, client)
}

func NewBindingClient(client client.Client) xbasissvc_external_permission.BindingService {
	return xbasissvc_external_permission.NewBindingService(xbasisconstants.PermissionService, client)
}

func NewGroupClient(client client.Client) xbasissvc_external_permission.UserGroupService {
	return xbasissvc_external_permission.NewUserGroupService(xbasisconstants.PermissionService, client)
}
