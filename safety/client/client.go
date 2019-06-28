package safetyclient

import (
	"github.com/micro/go-micro/client"
	"konekko.me/xbasis/commons/constants"
	"konekko.me/xbasis/safety/pb"
	"konekko.me/xbasis/safety/pb/inner"
)

func NewSecurityClient(client client.Client) xbasissvc_internal_safety.SecurityService {
	return xbasissvc_internal_safety.NewSecurityService(xbasisconstants.InternalSafetyService, client)
}

func NewBlacklistClient(client client.Client) xbasissvc_external_safety.BlacklistService {
	return xbasissvc_external_safety.NewBlacklistService(xbasisconstants.SafetyService, client)
}
