package permissioncli

import (
	"github.com/micro/go-micro/client"
	"konekko.me/gosion/commons/constants"
	"konekko.me/gosion/permission/pb"
	"konekko.me/gosion/permission/pb/inner"
)

func NewVerificationClient(client client.Client) gosionsvc_internal_permission.VerificationService {
	return gosionsvc_internal_permission.NewVerificationService(gs_commons_constants.InternalPermission, client)
}

func NewAccessibleClient(client client.Client) gosionsvc_internal_permission.AccessibleService {
	return gosionsvc_internal_permission.NewAccessibleService(gs_commons_constants.InternalPermission, client)
}

func NewBindingClient(client client.Client) gosionsvc_external_permission.BindingService {
	return gosionsvc_external_permission.NewBindingService(gs_commons_constants.PermissionService, client)
}

func NewGroupClient(client client.Client) gosionsvc_external_permission.UserGroupService {
	return gosionsvc_external_permission.NewUserGroupService(gs_commons_constants.PermissionService, client)
}
