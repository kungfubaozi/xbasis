package userclient

import (
	"github.com/micro/go-micro/client"
	"konekko.me/xbasis/commons/constants"
	"konekko.me/xbasis/user/pb"
	"konekko.me/xbasis/user/pb/inner"
)

func NewExtUserClient(client client.Client) xbasissvc_internal_user.UserService {
	return xbasissvc_internal_user.NewUserService(xbasisconstants.InternalUserService, client)
}

func NewExtMessageClient(client client.Client) xbasissvc_internal_user.MessageService {
	return xbasissvc_internal_user.NewMessageService(xbasisconstants.InternalUserService, client)
}

func NewInviteClient(client client.Client) xbasissvc_external_user.InviteService {
	return xbasissvc_external_user.NewInviteService(xbasisconstants.UserService, client)
}
