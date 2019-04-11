package safetyclient

import (
	"github.com/micro/go-micro/client"
	"konekko.me/gosion/commons/constants"
	"konekko.me/gosion/safety/pb"
	"konekko.me/gosion/safety/pb/ext"
)

func NewSecurityClient(client client.Client) gs_ext_service_safety.SecurityService {
	return gs_ext_service_safety.NewSecurityService(gs_commons_constants.ExtSafetyService, client)
}

func NewBlacklistClient(client client.Client) gs_service_safety.BlacklistService {
	return gs_service_safety.NewBlacklistService(gs_commons_constants.SafetyService, client)
}
