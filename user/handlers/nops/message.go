package user_nops_handlers

import (
	"context"
	"konekko.me/gosion/commons/dto"
	"konekko.me/gosion/user/pb/nops"
)

type messageService struct {
}

func (svc *messageService) SendVerificationCode(context.Context, *gs_nops_service_message.SendRequest, *gs_commons_dto.Status) error {
	panic("implement me")
}

func NewMessageService() gs_nops_service_message.MessageHandler {
	return &messageService{}
}
