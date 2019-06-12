package permissioncli

import (
	"github.com/micro/go-micro/client"
	"konekko.me/gosion/commons/constants"
	"konekko.me/gosion/permission/pb/inner"
)

func NewVerificationClient(client client.Client) gosionsvc_internal_permission.VerificationService {
	return gosionsvc_internal_permission.NewVerificationService(gs_commons_constants.InternalPermission, client)
}

func NewAccessibleClient(client client.Client) gosionsvc_internal_permission.AccessibleService {
	return gosionsvc_internal_permission.NewAccessibleService(gs_commons_constants.InternalPermission, client)
}
