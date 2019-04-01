package user_nops_handlers

import (
	"context"
	"konekko.me/gosion/commons/dto"
	"konekko.me/gosion/commons/wrapper"
	"konekko.me/gosion/message/cmd/messagecli"
	"konekko.me/gosion/user/pb/ext"
)

type messageService struct {
	message messagecli.MessageClient
}

func (svc *messageService) SendVerificationCode(ctx context.Context, in *gs_ext_service_user.SendRequest, out *gs_commons_dto.Status) error {
	return gs_commons_wrapper.ContextToAuthorize(ctx, in, func(auth *gs_commons_wrapper.WrapperUser) *gs_commons_dto.State {
		return nil
	})
}

func NewMessageService() gs_ext_service_user.MessageHandler {
	return &messageService{}
}
