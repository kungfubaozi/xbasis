package applicationhanderls

import (
	"context"
	"fmt"
	"gopkg.in/mgo.v2"
	"konekko.me/gosion/analysis/client"
	external "konekko.me/gosion/application/pb"
	"konekko.me/gosion/commons/config/call"
	"konekko.me/gosion/commons/constants"
	"konekko.me/gosion/commons/dto"
	"konekko.me/gosion/commons/encrypt"
	"konekko.me/gosion/commons/errstate"
	"konekko.me/gosion/commons/generator"
	"konekko.me/gosion/commons/indexutils"
	"konekko.me/gosion/commons/wrapper"
	"time"
)

type applicationService struct {
	session *mgo.Session
	*indexutils.Client
	log analysisclient.LogClient
}

func (svc *applicationService) GetRepo() *applicationRepo {
	return &applicationRepo{session: svc.session.Clone(), Client: svc.Client}
}

//create new application if not exists(name)
func (svc *applicationService) Create(ctx context.Context, in *external.CreateRequest, out *gs_commons_dto.Status) error {
	return gs_commons_wrapper.ContextToAuthorize(ctx, out, func(auth *gs_commons_wrapper.WrapperUser) *gs_commons_dto.State {

		headers := &analysisclient.LogHeaders{
			TraceId:     auth.TraceId,
			ServiceName: gs_commons_constants.ApplicationService,
			ModuleName:  "Application",
		}

		repo := svc.GetRepo()
		defer repo.Close()

		if len(in.Name) == 0 {
			return errstate.ErrRequest
		}

		if repo.ApplicationExists(in.Name) {
			return errstate.ErrApplicationAlreadyExists
		}

		if len(in.Url) > 0 && repo.RedirectUrlExists(in.Url) {
			return errstate.ErrApplicationRedirectUrl
		}

		if in.MustSync && len(in.SyncUrl) <= 10 {
			return errstate.ErrApplicationSyncUrl
		}

		id := gs_commons_generator.NewIDG()

		now := time.Now().UnixNano()

		appId := id.Short()

		enabled := gs_commons_constants.Closed
		if in.Open {
			enabled = gs_commons_constants.Enabled
		}

		//生成加密字符
		secretKey := id.Get()
		c := serviceconfiguration.Get()

		s, err := encrypt.AESEncrypt([]byte(secretKey), []byte(c.CurrencySecretKey))
		if err != nil {
			return errstate.ErrRequest
		}

		info := &appInfo{
			Name:         in.Name,
			CreateAt:     now,
			Id:           appId,
			CreateUserId: auth.User,
			SecretKey:    string(s),
			Settings: &appSetting{
				Enabled:     enabled,
				MustSync:    in.MustSync,
				RedirectURL: in.Url,
				SyncUserURL: in.SyncUrl,
			},
			Clients: []*appClient{
				{
					Id:       id.Short(),
					Platform: gs_commons_constants.PlatformOfAndroid,
					Enabled:  gs_commons_constants.Closed,
				},
				{
					Id:       id.Short(),
					Platform: gs_commons_constants.PlatformOfIOS,
					Enabled:  gs_commons_constants.Closed,
				},
				{
					Id:       id.Short(),
					Platform: gs_commons_constants.PlatformOfWeb,
					Enabled:  gs_commons_constants.Closed,
				},
				{
					Id:       id.Short(),
					Platform: gs_commons_constants.PlatformOfWindows,
					Enabled:  gs_commons_constants.Closed,
				},
				{
					Id:       id.Short(),
					Platform: gs_commons_constants.PlatfromOfMacOS,
					Enabled:  gs_commons_constants.Closed,
				},
				{
					Id:       id.Short(),
					Platform: gs_commons_constants.PlatformOfLinux,
					Enabled:  gs_commons_constants.Closed,
				},
				{
					Id:       id.Short(),
					Platform: gs_commons_constants.PlatformOfFuchsia,
					Enabled:  gs_commons_constants.Closed,
				},
			},
		}

		err = repo.Add(info)
		if err == nil {
			return errstate.Success
		}

		svc.log.Info(&analysisclient.LogContent{
			Headers: headers,
			Action:  "CreateApplication",
		})

		return nil
	})
}

func (svc *applicationService) Remove(ctx context.Context, in *external.RemoveRequest, out *gs_commons_dto.Status) error {
	return gs_commons_wrapper.ContextToAuthorize(ctx, out, func(auth *gs_commons_wrapper.WrapperUser) *gs_commons_dto.State {
		return errstate.Success
	})
}

func (svc *applicationService) ChangeName(ctx context.Context, in *external.ChangeNameRequest, out *gs_commons_dto.Status) error {
	return gs_commons_wrapper.ContextToAuthorize(ctx, out, func(auth *gs_commons_wrapper.WrapperUser) *gs_commons_dto.State {
		return nil
	})
}

func (svc *applicationService) FindByAppId(ctx context.Context, in *external.FindRequest, out *external.SimpleApplicationResponse) error {
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

		out.Info = &external.AppInfo{
			Name:  info.Name,
			AppId: info.Id,
			Settings: &gs_commons_dto.AppSettings{
				Enabled: info.Settings.Enabled,
			},
		}

		var ar []*external.AppClientInfo
		for _, k := range info.Clients {
			c := &external.AppClientInfo{
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

func (svc *applicationService) FindByClientId(context.Context, *external.FindRequest, *external.SimpleApplicationResponse) error {
	panic("implement me")
}

func (svc *applicationService) List(ctx context.Context, in *external.FindRequest, out *external.ListResponse) error {
	return gs_commons_wrapper.ContextToAuthorize(ctx, out, func(auth *gs_commons_wrapper.WrapperUser) *gs_commons_dto.State {

		repo := svc.GetRepo()
		defer repo.Close()

		list, err := repo.FindAll()
		if err != nil {
			fmt.Println("err", err)
			return nil
		}

		var l []*external.AppInfo
		for _, v := range list {
			t := "user"
			switch v.Type {
			case gs_commons_constants.AppTypeManage, gs_commons_constants.AppTypeRoute, gs_commons_constants.AppTypeSafe, gs_commons_constants.AppTypeUser:
				t = "sys"
			}

			l = append(l, &external.AppInfo{
				AppId:    v.Id,
				Enabled:  v.Settings.Enabled,
				CreateAt: v.CreateAt,
				Name:     v.Name,
				Desc:     v.Desc,
				Type:     t,
			})
		}

		if l != nil {
			out.Info = l
			out.State = errstate.Success
		}

		return nil
	})
}

func (svc *applicationService) Switch(context.Context, *external.SwitchRequest, *gs_commons_dto.Status) error {
	panic("implement me")
}

func NewApplicationService(session *mgo.Session, client *indexutils.Client, log analysisclient.LogClient) external.ApplicationHandler {
	return &applicationService{session: session, Client: client, log: log}
}
