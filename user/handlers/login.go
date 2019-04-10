package userhandlers

import (
	"context"
	"fmt"
	"github.com/olivere/elastic"
	"golang.org/x/crypto/bcrypt"
	"gopkg.in/mgo.v2"
	"konekko.me/gosion/authentication/pb/ext"
	"konekko.me/gosion/commons/constants"
	"konekko.me/gosion/commons/dto"
	"konekko.me/gosion/commons/errstate"
	"konekko.me/gosion/commons/regx"
	"konekko.me/gosion/commons/wrapper"
	"konekko.me/gosion/safety/pb/ext"
	"konekko.me/gosion/user/pb"
)

type loginService struct {
	session         *mgo.Session
	securityService gs_ext_service_safety.SecurityService
	tokenService    gs_ext_service_authentication.TokenService
	client          *elastic.Client
}

func (svc *loginService) GetRepo() *userRepo {
	return &userRepo{session: svc.session.Clone(), elastic: svc.client}
}

//web client just support the root project, you need the login to root project and then route to the target client
func (svc *loginService) WithAccount(ctx context.Context, in *gs_service_user.EntryRequest, out *gs_service_user.EntryWithAccountResponse) error {
	return gs_commons_wrapper.ContextToAuthorize(ctx, out, func(auth *gs_commons_wrapper.WrapperUser) *gs_commons_dto.State {
		if len(in.Account) > 0 && len(in.Content) > 0 {
			repo := svc.GetRepo()
			defer repo.Close()

			var info *userInfo

			eiup := func() *gs_commons_dto.State {
				return errstate.ErrInvalidUsernameOrPassword
			}

			t := accountIndexType

			if gs_commons_regx.Phone(in.Account) {
				t = phoneIndexType
			} else if gs_commons_regx.Email(in.Account) {
				t = emailIndexType
			}

			i, err := repo.FindIndexTable(t, in.Account)
			if err != nil {
				return eiup()
			}
			info, err = repo.FindById(i)
			if err != nil {
				return eiup()
			}

			fmt.Println("ok", info.Id)

			if info != nil && len(info.Id) > 0 {
				//check state
				s, err := svc.securityService.Get(ctx, &gs_ext_service_safety.GetRequest{
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
				s1, err := svc.tokenService.Generate(ctx, &gs_ext_service_authentication.GenerateRequest{
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

func NewLoginService(session *mgo.Session, securityService gs_ext_service_safety.SecurityService,
	tokenService gs_ext_service_authentication.TokenService, client *elastic.Client) gs_service_user.LoginHandler {
	return &loginService{session: session, securityService: securityService, tokenService: tokenService, client: client}
}
