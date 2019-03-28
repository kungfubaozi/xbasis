package user_handlers

import (
	"context"
	"golang.org/x/crypto/bcrypt"
	"gopkg.in/mgo.v2"
	"konekko.me/gosion/authentication/pb/nops"
	"konekko.me/gosion/commons/constants"
	"konekko.me/gosion/commons/dto"
	"konekko.me/gosion/commons/errstate"
	"konekko.me/gosion/commons/regx"
	"konekko.me/gosion/commons/wrapper"
	"konekko.me/gosion/safety/pb/nops"
	"konekko.me/gosion/user/pb"
	"konekko.me/gosion/user/repositories"
)

type loginService struct {
	session         *mgo.Session
	securityService gs_nops_service_safety.SecurityService
	tokenService    gs_nops_service_authentication.TokenService
}

func (svc *loginService) GetRepo() *user_repositories.UserRepo {
	return &user_repositories.UserRepo{Session: svc.session.Clone()}
}

//web client just support the root project, you need the login to root project and then route to the target client
func (svc *loginService) WithAccount(ctx context.Context, in *gs_service_user.EntryRequest, out *gs_service_user.EntryWithAccountResponse) error {
	return gs_commons_wrapper.ContextToAuthorize(ctx, out, func(auth *gs_commons_wrapper.WrapperUser) *gs_commons_dto.State {
		if len(in.Account) > 0 && len(in.Content) > 0 {
			repo := svc.GetRepo()
			defer repo.Close()

			var info *user_repositories.UserInfo

			eiup := func() *gs_commons_dto.State {
				return errstate.ErrInvalidUsernameOrPassword
			}

			if gs_commons_regx.Phone(in.Account) || gs_commons_regx.Email(in.Account) {
				i, err := repo.FindByContract(in.Account)
				if err != nil {
					return eiup()
				}
				info, err = repo.FindById(i.UserId)
				if err != nil {
					return eiup()
				}
			} else {
				i, err := repo.FindByAccount(in.Account)
				if err != nil {
					return eiup()
				}
				info = i
			}
			if info != nil && len(info.Id) > 0 {
				//check state
				s, err := svc.securityService.Get(ctx, &gs_nops_service_safety.GetRequest{
					UserId: info.Id,
				})

				if err != nil {
					return nil
				}

				if !s.State.Ok {
					return s.State
				}

				if s.Current != gs_commons_constants.UserStateOfClear {
					return errstate.ErrLoginFailed
				}

				//check password
				err = bcrypt.CompareHashAndPassword([]byte(in.Content), []byte(info.Password))
				if err != nil {
					return eiup()
				}

				//generate token
				s1, err := svc.tokenService.Generate(ctx, &gs_nops_service_authentication.GenerateRequest{
					Auth: &gs_commons_dto.Authorize{
						ClientId:  in.ClientId,
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
				out.State = errstate.Success

				return nil
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

func NewLoginService() gs_service_user.LoginHandler {
	return &loginService{}
}
