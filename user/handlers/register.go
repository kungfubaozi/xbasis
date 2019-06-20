package userhandlers

import (
	"context"
	"golang.org/x/crypto/bcrypt"
	"gopkg.in/mgo.v2"
	"konekko.me/gosion/application/pb/inner"
	"konekko.me/gosion/commons/config/call"
	"konekko.me/gosion/commons/constants"
	"konekko.me/gosion/commons/dto"
	"konekko.me/gosion/commons/errstate"
	"konekko.me/gosion/commons/generator"
	"konekko.me/gosion/commons/indexutils"
	"konekko.me/gosion/commons/regx"
	"konekko.me/gosion/commons/wrapper"
	"konekko.me/gosion/permission/pb"
	external "konekko.me/gosion/user/pb"
	"time"
)

type registerService struct {
	session                  *mgo.Session
	inviteService            external.InviteService
	client                   *indexutils.Client
	bindingService           gosionsvc_external_permission.BindingService
	groupService             gosionsvc_external_permission.UserGroupService
	id                       gs_commons_generator.IDGenerator
	applicationStatusService gosionsvc_internal_application.ApplicationStatusService
}

func (svc *registerService) GetRepo() *userRepo {
	return &userRepo{session: svc.session.Clone(), Client: svc.client}
}

//自注册的用户只能有访问当前项目的权限
//管理员invite可以选择可以访问哪些项目
func (svc *registerService) New(ctx context.Context, in *external.NewRequest, out *gs_commons_dto.Status) error {
	return gs_commons_wrapper.ContextToAuthorize(ctx, out, func(auth *gs_commons_wrapper.WrapperUser) *gs_commons_dto.State {
		configuration := serviceconfiguration.Get()

		status, err := svc.applicationStatusService.GetAppClientStatus(ctx, &gosionsvc_internal_application.GetAppClientStatusRequest{
			ClientId: in.ClientId,
		})
		if err != nil {
			return errstate.ErrRequest
		}

		if !status.State.Ok {
			return status.State
		}

		//只允许注册到route项目中
		if status.Type != gs_commons_constants.AppTypeRoute {
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
			if !gs_commons_regx.Phone(in.Contract) {
				return errstate.ErrFormatPhone
			}
			key = "phone"
			user.Phone = in.Contract
		} else if configuration.RegisterType == 1002 { //email
			if len(in.Contract) <= 8 {
				return errstate.ErrRequest
			}
			if !gs_commons_regx.Email(in.Contract) {
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

		if auth.Access.To != in.Contract {
			return errstate.ErrValidationCode
		}

		//查找是否被注册过
		repo := svc.GetRepo()
		defer repo.Close()

		u, err := repo.FindIndexTable(key, value)
		if err != nil {
			return nil
		}

		if len(u) != 0 {
			return errstate.ErrUserAlreadyRegister
		}

		//查找是否为邀请用户
		s, err := svc.inviteService.HasInvited(ctx, &external.HasInvitedRequest{
			Phone: in.Contract,
			Email: in.Contract,
		})

		if err != nil {
			return nil
		}

		if !s.State.Ok {
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

		err = repo.AddUser(user)
		if err != nil {
			return errstate.ErrSystem
		}

		info := &userInfo{
			UserId:   userId,
			Username: in.Username,
		}

		err = repo.AddUserInfo(info)
		if err != nil {
			return errstate.ErrRequest
		}

		index := &userModelIndex{
			Username: in.Username,
			Phone:    user.Phone,
			Email:    user.Email,
		}

		err = repo.AddUserIndex(index)
		if err != nil {
			return errstate.ErrRequest
		}

		if invited {
			s, err := svc.inviteService.SetState(ctx, &external.SetStateRequest{
				UserId: userId,
				State:  gs_commons_constants.InviteStateOfRegister,
			})
			if err != nil {
				return nil
			}
			if !s.State.Ok {
				return s.State
			}
		}

		return errstate.Success
	})
}

func NewRegisterService(session *mgo.Session, inviteService external.InviteService,
	client *indexutils.Client,
	bindingService gosionsvc_external_permission.BindingService,
	groupService gosionsvc_external_permission.UserGroupService,
	applicationStatusService gosionsvc_internal_application.ApplicationStatusService) external.RegisterHandler {
	return &registerService{session: session, inviteService: inviteService, client: client, bindingService: bindingService,
		groupService: groupService, id: gs_commons_generator.NewIDG(), applicationStatusService: applicationStatusService}
}
