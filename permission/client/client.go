package permissioncli

import (
	"github.com/micro/go-micro/client"
	"konekko.me/gosion/commons/constants"
	"konekko.me/gosion/permission/pb"
)

func NewVerificationClient(client client.Client) gs_service_permission.VerificationService {
	return gs_service_permission.NewVerificationService(gs_commons_constants.ExtPermissionVerification, client)
}
