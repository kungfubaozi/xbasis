package safetyclient

import (
	"github.com/micro/go-micro/client"
	"konekko.me/gosion/commons/constants"
	"konekko.me/gosion/safety/pb"
	"konekko.me/gosion/safety/pb/inner"
)

func NewSecurityClient(client client.Client) gosionsvc_internal_safety.SecurityService {
	return gosionsvc_internal_safety.NewSecurityService(gs_commons_constants.InternalSafetyService, client)
}

func NewBlacklistClient(client client.Client) gosionsvc_external_safety.BlacklistService {
	return gosionsvc_external_safety.NewBlacklistService(gs_commons_constants.SafetyService, client)
}
