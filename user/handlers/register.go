package userhandlers

import (
	"context"
	"github.com/samuel/go-zookeeper/zk"
	"github.com/vmihailenco/msgpack"
	"golang.org/x/crypto/bcrypt"
	"gopkg.in/mgo.v2"
	"konekko.me/gosion/commons/config"
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
	session        *mgo.Session
	inviteService  external.InviteService
	client         *indexutils.Client
	conn           *zk.Conn
	bindingService gosionsvc_external_permission.BindingService
	groupService   gosionsvc_external_permission.UserGroupService
	id             gs_commons_generator.IDGenerator
}

func (svc *registerService) GetRepo() *userRepo {
	return &userRepo{session: svc.session.Clone(), Client: svc.client}
}

//自注册的用户只能有访问当前项目的权限
//管理员invite可以选择可以访问哪些项目
func (svc *registerService) New(ctx context.Context, in *external.NewRequest, out *gs_commons_dto.Status) error {
	return gs_commons_wrapper.ContextToAuthorize(ctx, out, func(auth *gs_commons_wrapper.WrapperUser) *gs_commons_dto.State {
		configuration := serviceconfiguration.Get()
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

		invited := len(s.Content) != 0

		if invited {

		}

		if len(in.Password) < 6 {
			return errstate.ErrPasswordLength
		}

		p, err := bcrypt.GenerateFromPassword([]byte(in.Password), bcrypt.DefaultCost)
		if err != nil {
			return errstate.ErrSystem
		}

		userId := svc.id.Get()
		user.Id = userId
		user.Password = string(p)

		err = repo.AddUser(user)
		if err != nil {
			return errstate.ErrSystem
		}

		//注册时并不会直接设置账户的角色等，只会赋予RouteApp中的基本角色，其他的应用会在进入时同步该用户
		//只有用户同步时才会设置对应项目的基本角色
		b, _, err := svc.conn.Get(gs_commons_constants.ZKAutonomyRegister)
		app := &gs_commons_config.AutonomyRouteConfig{}
		err = msgpack.Unmarshal(b, app)
		if err != nil {
			return errstate.ErrRequest
		}
		s, err = svc.bindingService.UserRole(ctx, &gosionsvc_external_permission.BindingRoleRequest{
			Id:     userId,
			RoleId: app.RoleId,
			AppId:  app.AppId,
		})
		if err != nil {
			return errstate.ErrRequest
		}
		if !s.State.Ok {
			return s.State
		}
		if len(app.BindGroupId) > 0 {
			s, err = svc.groupService.AddUser(ctx, &gosionsvc_external_permission.SimpleUserNode{
				GroupId: app.BindGroupId,
				UserId:  userId,
				AppId:   app.AppId,
			})
			if err != nil {
				return errstate.ErrRequest
			}
			if !s.State.Ok {
				return s.State
			}
		}

		return nil
	})
}

func NewRegisterService(session *mgo.Session, conn *zk.Conn) external.RegisterHandler {
	return &registerService{session: session}
}
