package userhandlers

import (
	"context"
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"gopkg.in/mgo.v2"
	"konekko.me/xbasis/analysis/client"
	"konekko.me/xbasis/application/pb/inner"
	"konekko.me/xbasis/commons/config/call"
	constants "konekko.me/xbasis/commons/constants"
	commmons "konekko.me/xbasis/commons/dto"
	"konekko.me/xbasis/commons/errstate"
	generator "konekko.me/xbasis/commons/generator"
	"konekko.me/xbasis/commons/indexutils"
	regx "konekko.me/xbasis/commons/regx"
	wrapper "konekko.me/xbasis/commons/wrapper"
	"konekko.me/xbasis/permission/pb"
	external "konekko.me/xbasis/user/pb"
	"time"
)

type registerService struct {
	session                  *mgo.Session
	inviteService            external.InviteService
	client                   *indexutils.Client
	bindingService           xbasissvc_external_permission.BindingService
	groupService             xbasissvc_external_permission.UserGroupService
	id                       generator.IDGenerator
	applicationStatusService xbasissvc_internal_application.ApplicationStatusService
	log                      analysisclient.LogClient
}

func (svc *registerService) GetRepo() *userRepo {
	return &userRepo{session: svc.session.Clone(), Client: svc.client}
}

//自注册的用户只能有访问当前项目的权限
//管理员invite可以选择可以访问哪些项目
func (svc *registerService) New(ctx context.Context, in *external.NewRequest, out *commmons.Status) error {
	return wrapper.ContextToAuthorize(ctx, out, func(auth *wrapper.WrapperUser) *commmons.State {
		configuration := serviceconfiguration.Get()

		header := &analysisclient.LogHeaders{
			TraceId:     auth.TraceId,
			ServiceName: constants.UserService,
			ModuleName:  "Register",
		}

		status, err := svc.applicationStatusService.GetAppClientStatus(ctx, &xbasissvc_internal_application.GetAppClientStatusRequest{
			ClientId: in.ClientId,
		})
		if err != nil {
			return errstate.ErrRequest
		}

		if !status.State.Ok {
			return status.State
		}

		//只允许注册到route项目中
		if status.Type != constants.AppTypeRoute {
			return errstate.ErrRequest
		}

		key := ""
		value := in.Contract
		user := &userModel{
			RegisterAt: auth.FromClientId,
			CreateAt:   time.Now().UnixNano(),
		}
		if configuration.RegisterType == 1001 { //phone
			if len(in.Contract) <= 8 {
				return errstate.ErrRequest
			}
			if !regx.Phone(in.Contract) {
				return errstate.ErrFormatPhone
			}
			key = "phone"
			user.Phone = in.Contract
		} else if configuration.RegisterType == 1002 { //email
			if len(in.Contract) <= 8 {
				return errstate.ErrRequest
			}
			if !regx.Email(in.Contract) {
				return errstate.ErrFormatEmail
			}
			key = "email"
			user.Email = in.Contract
		} else {
			return errstate.ErrRequest
		}

		if auth.Access == nil {
			return errstate.ErrSystem
		}

		//验证验证码发送对象是否当前contract
		if auth.Access.To != in.Contract {
			return errstate.ErrValidationCode
		}

		//查找是否被注册过
		repo := svc.GetRepo()
		defer repo.Close()

		u, err := repo.FindIndexTable(key, value)
		if err != nil && err != indexutils.ErrNotFound {
			return nil
		}

		if len(u) != 0 {
			fmt.Println("err1", "8")
			return errstate.ErrUserAlreadyRegister
		}

		//查找是否为邀请用户
		s, err := svc.inviteService.HasInvited(ctx, &external.HasInvitedRequest{
			Phone: in.Contract,
			Email: in.Contract,
		})

		if err != nil {
			fmt.Println("err1", "9", err)
			return nil
		}

		if !s.State.Ok {
			fmt.Println("err1", "10")
			return s.State
		}

		invited := len(s.UserId) != 0

		if len(in.Password) < 6 {
			return errstate.ErrPasswordLength
		}

		p, err := bcrypt.GenerateFromPassword([]byte(in.Password), bcrypt.DefaultCost)
		if err != nil {
			return errstate.ErrSystem
		}
		userId := svc.id.Get()
		if invited {
			userId = s.UserId
		}
		user.Id = userId
		user.Password = string(p)

		info := &userInfo{
			UserId:   userId,
			Username: in.Username,
		}

		if invited {
			v, err := svc.inviteService.GetDetail(ctx, &external.HasInvitedRequest{
				UserId: s.UserId,
			})
			if err != nil {
				return nil
			}
			if !v.State.Ok {
				return v.State
			}
			info.RealName = v.RealName
			if len(v.Email) > 0 {
				user.Email = v.Email
			}
			if len(v.Phone) > 0 {
				user.Phone = v.Phone
			}
			if len(v.Username) > 0 {
				info.Username = v.Username
			}
		}

		//需要设置用户名
		if len(info.Username) == 0 {
			return errstate.ErrUserNeedSetUsername
		}

		err = repo.AddUser(user)
		if err != nil {
			return errstate.ErrSystem
		}

		err = repo.AddUserInfo(info)
		if err != nil {
			return errstate.ErrRequest
		}

		if invited {
			s, err := svc.inviteService.SetState(ctx, &external.SetStateRequest{
				UserId: userId,
				State:  constants.InviteStateOfRegister,
			})
			if err != nil {
				return nil
			}
			if !s.State.Ok {
				return s.State
			}
		}

		svc.log.Info(&analysisclient.LogContent{
			Headers: header,
			Action:  "NewUserRegister",
			Fields: &analysisclient.LogFields{
				"username":  info.Username,
				"user_id":   info.UserId,
				"timestamp": time.Now().Unix(),
				"client_id": in.ClientId,
				"app_id":    auth.AppId,
			},
			Index: &analysisclient.LogIndex{
				Name: "users",
				Id:   userId,
				Fields: &analysisclient.LogFields{
					"username":  info.Username,
					"real_name": info.RealName,
					"phone":     user.Phone,
					"email":     user.Email,
					"user_id":   user.Id,
					"invite":    false,
				},
			},
		})

		return errstate.Success
	})
}

func NewRegisterService(log analysisclient.LogClient, session *mgo.Session, inviteService external.InviteService,
	client *indexutils.Client,
	bindingService xbasissvc_external_permission.BindingService,
	groupService xbasissvc_external_permission.UserGroupService,
	applicationStatusService xbasissvc_internal_application.ApplicationStatusService) external.RegisterHandler {
	return &registerService{log: log, session: session, inviteService: inviteService, client: client, bindingService: bindingService,
		groupService: groupService, id: generator.NewIDG(), applicationStatusService: applicationStatusService}
}
