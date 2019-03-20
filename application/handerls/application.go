package application_handerls

import (
	"context"
	"github.com/garyburd/redigo/redis"
	"gopkg.in/mgo.v2"
	"konekko.me/gosion/application/pb"
	"konekko.me/gosion/application/repositories"
	"konekko.me/gosion/commons/constants"
	"konekko.me/gosion/commons/dto"
	"konekko.me/gosion/commons/encrypt"
	"konekko.me/gosion/commons/errstate"
	"konekko.me/gosion/commons/generator"
	"konekko.me/gosion/commons/wrapper"
	"time"
)

type applicationService struct {
	session *mgo.Session
	pool    *redis.Pool
}

func (svc *applicationService) GetRepo() application_repositories.ApplicationRepo {
	return application_repositories.ApplicationRepo{Session: svc.session.Clone(), Conn: svc.pool.Get()}
}

//create new application if not exists(name)
func (svc *applicationService) Create(ctx context.Context, in *gs_service_application.CreateRequest, out *gs_commons_dto.Status) error {
	return gs_commons_wrapper.ContextToAuthorize(ctx, out, func(auth *gs_commons_wrapper.WrapperUser) *gs_commons_dto.State {

		repo := svc.GetRepo()
		defer repo.Close()

		if len(in.Name) == 0 {
			return errstate.ErrRequest
		}

		if repo.ApplicationExists(in.Name) {
			return errstate.ErrApplicationAlreadyExists
		}

		node := gs_commons_generator.ID()

		now := time.Now().UnixNano()

		appId := node.Generate().String()

		info := &application_repositories.AppInfo{
			Name:         in.Name,
			CreateAt:     now,
			Id:           appId,
			CreateUserId: auth.User,
			Settings: &application_repositories.AppSetting{
				Enabled:  gs_commons_constants.Enabled,
				OpenMode: gs_commons_constants.OpenModeOfAllOrganization,
			},
			Clients: []*application_repositories.AppClient{
				{
					Id:       node.Generate().String(),
					Platform: gs_commons_constants.PlatformOfAndroid,
					Secret:   gs_commons_encrypt.SHA256_BASE64(node.Generate().String() + appId),
					Enabled:  gs_commons_constants.Enabled,
				},
				{
					Id:       node.Generate().String(),
					Platform: gs_commons_constants.PlatformOfIOS,
					Secret:   gs_commons_encrypt.SHA256_BASE64(node.Generate().String() + appId),
					Enabled:  gs_commons_constants.Enabled,
				},
				{
					Id:       node.Generate().String(),
					Platform: gs_commons_constants.PlatformOfWeb,
					Secret:   gs_commons_encrypt.SHA256_BASE64(node.Generate().String() + appId),
					Enabled:  gs_commons_constants.Enabled,
				},
				{
					Id:       node.Generate().String(),
					Platform: gs_commons_constants.PlatformOfWindows,
					Secret:   gs_commons_encrypt.SHA256_BASE64(node.Generate().String() + appId),
					Enabled:  gs_commons_constants.Enabled,
				},
				{
					Id:       node.Generate().String(),
					Platform: gs_commons_constants.PlatfromOfMacOS,
					Secret:   gs_commons_encrypt.SHA256_BASE64(node.Generate().String() + appId),
					Enabled:  gs_commons_constants.Enabled,
				},
			},
		}

		err := repo.Upsert(info)
		if err == nil {
			return errstate.Success
		}

		return nil
	})
}

func (svc *applicationService) Remove(ctx context.Context, in *gs_service_application.RemoveRequest, out *gs_commons_dto.Status) error {
	return gs_commons_wrapper.ContextToAuthorize(ctx, out, func(auth *gs_commons_wrapper.WrapperUser) *gs_commons_dto.State {
		return nil
	})
}

func (svc *applicationService) ChangeName(ctx context.Context, in *gs_service_application.ChangeNameRequest, out *gs_commons_dto.Status) error {
	return gs_commons_wrapper.ContextToAuthorize(ctx, out, func(auth *gs_commons_wrapper.WrapperUser) *gs_commons_dto.State {
		return nil
	})
}

func (svc *applicationService) FindByAppId(ctx context.Context, in *gs_service_application.FindRequest, out *gs_service_application.SimpleApplicationResponse) error {
	return gs_commons_wrapper.ContextToAuthorize(ctx, out, func(auth *gs_commons_wrapper.WrapperUser) *gs_commons_dto.State {

		repo := svc.GetRepo()
		defer repo.Close()

		if len(in.Content) == 0 {
			return errstate.ErrRequest
		}

		info, err := repo.FindByApplicationId(in.Content)
		if err != nil {
			return errstate.ErrRequest
		}

		out.Info = &gs_service_application.AppInfo{
			Name:  info.Name,
			AppId: info.Id,
			Settings: &gs_commons_dto.AppSettings{
				Enabled:  info.Settings.Enabled,
				OpenMode: info.Settings.OpenMode,
			},
		}

		var ar []*gs_service_application.AppClientInfo
		for _, k := range info.Clients {
			c := &gs_service_application.AppClientInfo{
				ClientId: k.Id,
				Enabled:  k.Enabled,
				Platform: k.Platform,
			}
			ar = append(ar, c)
		}

		out.Info.Clients = ar

		out.State = errstate.Success

		return nil
	})
}

func (svc *applicationService) Status(ctx context.Context, in *gs_service_application.FindRequest, out *gs_service_application.StatusResponse) error {
	return gs_commons_wrapper.ContextToAuthorize(ctx, out, func(auth *gs_commons_wrapper.WrapperUser) *gs_commons_dto.State {
		repo := svc.GetRepo()
		defer repo.Close()

		if len(in.Content) == 0 {
			return errstate.ErrRequest
		}

		info, err := repo.FindByApplicationId(in.Content)
		if err != nil {
			return errstate.ErrRequest
		}

		if info.Settings.Enabled == gs_commons_constants.Closed {
			return errstate.ErrApplicationClosed
		}

		out.AppId = info.Id
		out.OpenMode = info.Settings.OpenMode
		out.State = errstate.Success

		return nil
	})
}

func (svc *applicationService) FindByClientId(context.Context, *gs_service_application.FindRequest, *gs_service_application.SimpleApplicationResponse) error {
	panic("implement me")
}

func (svc *applicationService) List(ctx context.Context, in *gs_service_application.FindRequest, out *gs_service_application.ListResponse) error {
	return gs_commons_wrapper.ContextToAuthorize(ctx, out, func(auth *gs_commons_wrapper.WrapperUser) *gs_commons_dto.State {

		repo := svc.GetRepo()
		defer repo.Close()

		list, err := repo.FindAll()
		if err != nil {
			return nil
		}

		var l []*gs_service_application.AppInfo
		for _, v := range list {
			l = append(l, &gs_service_application.AppInfo{
				AppId:    v.Id,
				Enabled:  v.Settings.Enabled,
				CreateAt: v.CreateAt,
				Name:     v.Name,
				Desc:     v.Desc,
			})
		}

		if l != nil {
			out.Info = l
			out.State = errstate.Success
		}

		return nil
	})
}

func NewApplicationService(session *mgo.Session, pool *redis.Pool) gs_service_application.ApplicationHandler {
	return &applicationService{session: session, pool: pool}
}
