package permissioncli

import (
	"github.com/micro/go-micro/client"
	"konekko.me/gosion/commons/constants"
	"konekko.me/gosion/permission/pb/ext"
)

func NewVerificationClient(client client.Client) gs_ext_service_permission.VerificationService {
	return gs_ext_service_permission.NewVerificationService(gs_commons_constants.ExtPermissionVerification, client)
}

func NewAccessibleClient(client client.Client) gs_ext_service_permission.AccessibleService {
	return gs_ext_service_permission.NewAccessibleService(gs_commons_constants.ExtAccessibleVerification, client)
}
