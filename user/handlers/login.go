package userhandlers

import (
	"context"
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"gopkg.in/mgo.v2"
	"konekko.me/xbasis/analysis/client"
	"konekko.me/xbasis/application/pb/inner"
	"konekko.me/xbasis/authentication/pb/inner"
	constants "konekko.me/xbasis/commons/constants"
	commons "konekko.me/xbasis/commons/dto"
	"konekko.me/xbasis/commons/errstate"
	"konekko.me/xbasis/commons/indexutils"
	regx "konekko.me/xbasis/commons/regx"
	"konekko.me/xbasis/commons/wrapper"
	"konekko.me/xbasis/safety/pb/inner"
	external "konekko.me/xbasis/user/pb"
)

type loginService struct {
	session              *mgo.Session
	innerSecurityService xbasissvc_internal_safety.SecurityService
	innerTokenService    xbasissvc_internal_authentication.TokenService
	innerSyncCheck       xbasissvc_internal_application.UserSyncService
	*indexutils.Client
	log analysisclient.LogClient
}

func (svc *loginService) GetRepo() *userRepo {
	return &userRepo{session: svc.session.Clone(), Client: svc.Client}
}

/**
1）登录生成的token只适用于当前登录的项目(x-client-id)
*/
//web client just support the root project, you need the login to root project and then route to the target client
func (svc *loginService) WithAccount(ctx context.Context, in *external.EntryRequest, out *external.EntryWithAccountResponse) error {
	fmt.Println("entry", in)
	return xbasiswrapper.ContextToAuthorize(ctx, out, func(auth *xbasiswrapper.WrapperUser) *commons.State {
		if len(in.Account) > 0 && len(in.Content) > 0 {
			repo := svc.GetRepo()
			defer repo.Close()

			headers := &analysisclient.LogHeaders{
				TraceId:     auth.TraceId,
				ServiceName: constants.UserService,
				ModuleName:  "Login",
			}

			var info *userModel

			eiup := func() *commons.State {
				return errstate.ErrInvalidUsernameOrPassword
			}

			var id string
			var err error

			if regx.Phone(in.Account) {
				id, err = repo.FindIndexTable("phone", in.Account)
			} else if regx.Email(in.Account) {
				id, err = repo.FindIndexTable("email", in.Account)
			} else {
				id, err = repo.FindIndexTable("account", in.Account)
			}

			if err != nil || len(id) == 0 {
				fmt.Println("err", err)
				return nil
			}

			info, err = repo.FindById(id)
			if err != nil {
				svc.log.Info(&analysisclient.LogContent{
					Headers: headers,
					Action:  "LoginFailedError",
					Message: err.Error(),
				})
				return eiup()
			}

			if len(info.Id) == 0 {
				svc.log.Info(&analysisclient.LogContent{
					Headers: headers,
					Action:  "LoginFailedIdNil",
					Message: "userId nil",
				})
				return nil
			}

			if info != nil && len(info.Id) > 0 {
				//check state

				s, err := svc.innerSecurityService.Get(ctx, &xbasissvc_internal_safety.GetRequest{
					UserId: info.Id,
				})

				if err != nil {
					return nil
				}

				if !s.State.Ok {
					return s.State
				}

				if s.Current != constants.UserStateOfClear {
					svc.log.Info(&analysisclient.LogContent{
						Headers: headers,
						Action:  "UserSecurityCheck",
						Message: "user state not clear",
						Fields: &analysisclient.LogFields{
							"userId": info.Id,
							"ip":     auth.IP,
						},
					})
					return errstate.ErrLoginFailed
				}

				//check if the application needs to synchronize user

				//check password
				err = bcrypt.CompareHashAndPassword([]byte(info.Password), []byte(in.Content))
				if err != nil {
					svc.log.Info(&analysisclient.LogContent{
						Headers: headers,
						Action:  "PasswordError",
						Message: "password error",
						Fields: &analysisclient.LogFields{
							"userId": info.Id,
							"ip":     auth.IP,
						},
					})
					return eiup()
				}

				s2 := &commons.Authorize{
					ClientId:  auth.FromClientId,
					UserId:    info.Id,
					Ip:        auth.IP,
					Device:    auth.UserDevice,
					UserAgent: auth.UserAgent,
					AppId:     auth.AppId,
					Platform:  auth.Platform,
				}

				//generate token
				s1, err := svc.innerTokenService.Generate(ctx, &xbasissvc_internal_authentication.GenerateRequest{
					Auth:       s2,
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
			svc.log.Info(&analysisclient.LogContent{
				Headers: headers,
				Action:  "UserInfoNotFound",
				Message: "cannot find user info",
			})
			return eiup()
		}
		return nil
	})
}

//web client just support the root project, you need the login to root project and then route to the target client
func (svc *loginService) WithValidateCode(ctx context.Context, in *external.EntryRequest, out *external.EntryWithQRCodeResponse) error {
	return xbasiswrapper.ContextToAuthorize(ctx, out, func(auth *xbasiswrapper.WrapperUser) *commons.State {
		return nil
	})
}

//web client just support the root project, you need the login to root project and then route to the target client
func (svc *loginService) WithQRCode(ctx context.Context, in *external.EntryRequest, out *external.EntryWithQRCodeResponse) error {
	return xbasiswrapper.ContextToAuthorize(ctx, out, func(auth *xbasiswrapper.WrapperUser) *commons.State {
		return nil
	})
}

func NewLoginService(session *mgo.Session, securityService xbasissvc_internal_safety.SecurityService,
	tokenService xbasissvc_internal_authentication.TokenService, client *indexutils.Client, log analysisclient.LogClient) external.LoginHandler {
	return &loginService{session: session, innerSecurityService: securityService, innerTokenService: tokenService, Client: client, log: log}
}
