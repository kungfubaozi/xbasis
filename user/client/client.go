package userclient

import (
	"github.com/micro/go-micro/client"
	"konekko.me/gosion/commons/constants"
	"konekko.me/gosion/user/pb/ext"
)

func NewExtUserClient(client client.Client) gs_ext_service_user.UserService {
	return gs_ext_service_user.NewUserService(gs_commons_constants.ExtUserService, client)
}

func NewExtMessageClient(client client.Client) gs_ext_service_user.MessageService {
	return gs_ext_service_user.NewMessageService(gs_commons_constants.ExtUserService, client)
}
