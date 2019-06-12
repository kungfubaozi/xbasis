package userclient

import (
	"github.com/micro/go-micro/client"
	"konekko.me/gosion/commons/constants"
	"konekko.me/gosion/user/pb/inner"
)

func NewExtUserClient(client client.Client) gosionsvc_internal_user.UserService {
	return gosionsvc_internal_user.NewUserService(gs_commons_constants.InternalUserService, client)
}

func NewExtMessageClient(client client.Client) gosionsvc_internal_user.MessageService {
	return gosionsvc_internal_user.NewMessageService(gs_commons_constants.InternalUserService, client)
}
