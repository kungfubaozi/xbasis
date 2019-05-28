package userhandlers

import (
	"context"
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"gopkg.in/mgo.v2"
	"konekko.me/gosion/application/pb/ext"
	"konekko.me/gosion/authentication/pb/ext"
	"konekko.me/gosion/commons/constants"
	"konekko.me/gosion/commons/dto"
	"konekko.me/gosion/commons/errstate"
	"konekko.me/gosion/commons/indexutils"
	"konekko.me/gosion/commons/regx"
	"konekko.me/gosion/commons/wrapper"
	"konekko.me/gosion/safety/pb/ext"
	"konekko.me/gosion/user/pb"
)

type loginService struct {
	session            *mgo.Session
	extSecurityService gs_ext_service_safety.SecurityService
	extTokenService    gs_ext_service_authentication.TokenService
	extSyncCheck       gs_ext_service_application.UsersyncService
	*indexutils.Client
}

func (svc *loginService) GetRepo() *userRepo {
	return &userRepo{session: svc.session.Clone(), Client: svc.Client}
}

/**
1）登录生成的token只适用于当前登录的项目(x-client-id)
2)
*/
//web client just support the root project, you need the login to root project and then route to the target client
func (svc *loginService) WithAccount(ctx context.Context, in *gs_service_user.EntryRequest, out *gs_service_user.EntryWithAccountResponse) error {
	return gs_commons_wrapper.ContextToAuthorize(ctx, out, func(auth *gs_commons_wrapper.WrapperUser) *gs_commons_dto.State {
		if len(in.Account) > 0 && len(in.Content) > 0 {
			repo := svc.GetRepo()
			defer repo.Close()

			var info *userModel

			eiup := func() *gs_commons_dto.State {
				return errstate.ErrInvalidUsernameOrPassword
			}

			var id string
			var err error

			if gs_commons_regx.Phone(in.Account) {
				id, err = repo.FindIndexTable("phone", in.Account)
			} else if gs_commons_regx.Email(in.Account) {
				id, err = repo.FindIndexTable("email", in.Account)
			} else {
				id, err = repo.FindIndexTable("account", in.Account)
			}

			if err != nil || len(id) == 0 {
				return nil
			}

			info, err = repo.FindById(id)
			if err != nil {
				return eiup()
			}

			if len(info.Id) == 0 {
				return nil
			}

			if info != nil && len(info.Id) > 0 {
				//check state

				fmt.Println("entry")

				s, err := svc.extSecurityService.Get(ctx, &gs_ext_service_safety.GetRequest{
					UserId: info.Id,
				})

				if err != nil {
					fmt.Println("err", err)
					return nil
				}

				if !s.State.Ok {
					return s.State
				}

				if s.Current != gs_commons_constants.UserStateOfClear {
					return errstate.ErrLoginFailed
				}

				//check if the application needs to synchronize user

				//check password
				err = bcrypt.CompareHashAndPassword([]byte(info.Password), []byte(in.Content))
				if err != nil {
					return eiup()
				}

				fmt.Println("login clientId", auth.ClientId)

				//generate token
				s1, err := svc.extTokenService.Generate(ctx, &gs_ext_service_authentication.GenerateRequest{
					Auth: &gs_commons_dto.Authorize{
						ClientId:  auth.ClientId,
						UserId:    info.Id,
						Ip:        auth.IP,
						Device:    auth.UserDevice,
						UserAgent: auth.UserAgent,
						AppId:     auth.AppId,
					},
					Route:      false,
					RelationId: "",
				})
				if err != nil {
					return nil
				}

				if !s1.State.Ok {
					return s1.State
				}

				//response token
				out.AccessToken = s1.AccessToken

				out.RefreshToken = s1.RefreshToken

				return errstate.Success
			}
			return eiup()
		}
		return nil
	})
}

//web client just support the root project, you need the login to root project and then route to the target client
func (svc *loginService) WithValidateCode(ctx context.Context, in *gs_service_user.EntryRequest, out *gs_service_user.EntryWithQRCodeResponse) error {
	return gs_commons_wrapper.ContextToAuthorize(ctx, out, func(auth *gs_commons_wrapper.WrapperUser) *gs_commons_dto.State {
		return nil
	})
}

//web client just support the root project, you need the login to root project and then route to the target client
func (svc *loginService) WithQRCode(ctx context.Context, in *gs_service_user.EntryRequest, out *gs_service_user.EntryWithQRCodeResponse) error {
	return gs_commons_wrapper.ContextToAuthorize(ctx, out, func(auth *gs_commons_wrapper.WrapperUser) *gs_commons_dto.State {
		return nil
	})
}

func NewLoginService(session *mgo.Session, securityService gs_ext_service_safety.SecurityService,
	tokenService gs_ext_service_authentication.TokenService, client *indexutils.Client) gs_service_user.LoginHandler {
	return &loginService{session: session, extSecurityService: securityService, extTokenService: tokenService, Client: client}
}
