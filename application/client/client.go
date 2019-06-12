package applicationclient

import (
	"github.com/micro/go-micro/client"
	_ "github.com/micro/go-micro/registry/consul"
	"konekko.me/gosion/application/pb"
	"konekko.me/gosion/application/pb/inner"
	"konekko.me/gosion/commons/constants"
)

func NewClient(client client.Client) gosionsvc_external_application.ApplicationService {
	return gosionsvc_external_application.NewApplicationService(gs_commons_constants.ApplicationService, client)
}

func NewStatusClient(client client.Client) gosionsvc_internal_application.ApplicationStatusService {
	return gosionsvc_internal_application.NewApplicationStatusService(gs_commons_constants.InternalApplicationService, client)
}

func NewSyncClient(client client.Client) gosionsvc_internal_application.UsersyncService {
	return gosionsvc_internal_application.NewUsersyncService(gs_commons_constants.InternalApplicationService, client)
}
