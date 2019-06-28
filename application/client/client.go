package applicationclient

import (
	"github.com/micro/go-micro/client"
	_ "github.com/micro/go-micro/registry/consul"
	"konekko.me/xbasis/application/pb"
	"konekko.me/xbasis/application/pb/inner"
	"konekko.me/xbasis/commons/constants"
)

func NewClient(client client.Client) xbasissvc_external_application.ApplicationService {
	return xbasissvc_external_application.NewApplicationService(xbasisconstants.ApplicationService, client)
}

func NewStatusClient(client client.Client) xbasissvc_internal_application.ApplicationStatusService {
	return xbasissvc_internal_application.NewApplicationStatusService(xbasisconstants.InternalApplicationService, client)
}

func NewSyncClient(client client.Client) xbasissvc_internal_application.UserSyncService {
	return xbasissvc_internal_application.NewUserSyncService(xbasisconstants.InternalApplicationService, client)
}
