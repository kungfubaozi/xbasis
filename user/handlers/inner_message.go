package userhandlers

import (
	"context"
	"fmt"
	"github.com/pkg/errors"
	"gopkg.in/mgo.v2"
	commons "konekko.me/xbasis/commons/dto"
	"konekko.me/xbasis/commons/errstate"
	"konekko.me/xbasis/commons/wrapper"
	"konekko.me/xbasis/message/cmd/messagecli"
	inner "konekko.me/xbasis/user/pb/inner"
)

type messageService struct {
	message messagecli.MessageClient
	session *mgo.Session
}

func (svc *messageService) sendCode(t int64, to, code string) error {
	if t == 1002 { //email

	} else if t == 1001 { //phone
		s := fmt.Sprintf("您的验证码为%s，请于10分钟内正确输入，如非本人操作，请忽略此短信。", code)
		return svc.message.SendSMS(to, s)
	}
	return errors.New("unsupp")
}

func (svc *messageService) GetRepo() *userRepo {
	return &userRepo{session: svc.session.Clone()}
}

func (svc *messageService) SendVerificationCode(ctx context.Context, in *inner.SendRequest, out *commons.Status) error {
	return xbasiswrapper.ContextToAuthorize(ctx, out, func(auth *xbasiswrapper.WrapperUser) *commons.State {
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
			fmt.Println("send error", err, in.MessageType)
			return errstate.ErrRequest
		}
		return errstate.Success
	})
}

func NewMessageService(message messagecli.MessageClient, session *mgo.Session) inner.MessageHandler {
	return &messageService{message: message, session: session}
}
