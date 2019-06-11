package userhandlers

import (
	"context"
	"github.com/pkg/errors"
	"gopkg.in/mgo.v2"
	"konekko.me/gosion/commons/dto"
	"konekko.me/gosion/commons/errstate"
	"konekko.me/gosion/commons/wrapper"
	"konekko.me/gosion/message/cmd/messagecli"
	"konekko.me/gosion/user/pb/ext"
)

type messageService struct {
	message messagecli.MessageClient
	session *mgo.Session
}

func (svc *messageService) sendCode(t int64, to, code string) error {
	if t == 1002 { //email

	} else if t == 1001 { //phone
		return svc.message.SendSMS(to, code)
	}
	return errors.New("unsupp")
}

func (svc *messageService) GetRepo() *userRepo {
	return &userRepo{session: svc.session.Clone()}
}

func (svc *messageService) SendVerificationCode(ctx context.Context, in *gs_ext_service_user.SendRequest, out *gs_commons_dto.Status) error {
	return gs_commons_wrapper.ContextToAuthorize(ctx, in, func(auth *gs_commons_wrapper.WrapperUser) *gs_commons_dto.State {
		to := in.To
		if in.Auth {
			repo := svc.GetRepo()
			defer repo.Close()

			model, err := repo.FindById(in.To)
			if err != nil {
				return errstate.ErrRequest
			}

			if in.MessageType == 1002 {
				if len(model.Email) <= 8 {
					return errstate.ErrUnbindEmail
				}
				to = model.Email
			} else if in.MessageType == 1001 {
				if len(model.Phone) <= 8 {
					return errstate.ErrUnbindPhone
				}
				to = model.Phone
			}
			err = svc.sendCode(in.MessageType, to, in.Code)
			if err != nil {
				return errstate.ErrRequest
			}
		}
		err := svc.sendCode(in.MessageType, to, in.Code)
		if err != nil {
			return errstate.ErrRequest
		}
		return errstate.Success
	})
}

func NewMessageService(message messagecli.MessageClient, session *mgo.Session) gs_ext_service_user.MessageHandler {
	return &messageService{message: message, session: session}
}
