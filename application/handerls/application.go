package applicationhanderls

import (
	"context"
	"fmt"
	"gopkg.in/mgo.v2"
	"konekko.me/xbasis/analysis/client"
	external "konekko.me/xbasis/application/pb"
	"konekko.me/xbasis/commons/config/call"
	constants "konekko.me/xbasis/commons/constants"
	commons "konekko.me/xbasis/commons/dto"
	"konekko.me/xbasis/commons/encrypt"
	"konekko.me/xbasis/commons/errstate"
	generator "konekko.me/xbasis/commons/generator"
	"konekko.me/xbasis/commons/indexutils"
	wrapper "konekko.me/xbasis/commons/wrapper"
	"time"
)

type applicationService struct {
	session *mgo.Session
	*indexutils.Client
	id  generator.IDGenerator
	log analysisclient.LogClient
}

func (svc *applicationService) GetRepo() *applicationRepo {
	return &applicationRepo{session: svc.session.Clone(), Client: svc.Client}
}

//create new application if not exists(name)
func (svc *applicationService) Create(ctx context.Context, in *external.CreateRequest, out *commons.Status) error {
	return wrapper.ContextToAuthorize(ctx, out, func(auth *wrapper.WrapperUser) *commons.State {

		headers := &analysisclient.LogHeaders{
			TraceId:     auth.TraceId,
			ServiceName: constants.ApplicationService,
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

		id := svc.id

		now := time.Now().UnixNano()

		appId := id.Short()

		enabled := constants.Closed
		if in.Open {
			enabled = constants.Enabled
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
				RedirectURL: in.Url,
				SyncUserURL: in.SyncUrl,
			},
			Clients: []*appClient{
				{
					Id:       id.Short(),
					Platform: constants.PlatformOfAndroid,
					Enabled:  constants.Closed,
				},
				{
					Id:       id.Short(),
					Platform: constants.PlatformOfIOS,
					Enabled:  constants.Closed,
				},
				{
					Id:       id.Short(),
					Platform: constants.PlatformOfWeb,
					Enabled:  constants.Closed,
				},
				{
					Id:       id.Short(),
					Platform: constants.PlatformOfWindows,
					Enabled:  constants.Closed,
				},
				{
					Id:       id.Short(),
					Platform: constants.PlatfromOfMacOS,
					Enabled:  constants.Closed,
				},
				{
					Id:       id.Short(),
					Platform: constants.PlatformOfLinux,
					Enabled:  constants.Closed,
				},
				{
					Id:       id.Short(),
					Platform: constants.PlatformOfFuchsia,
					Enabled:  constants.Closed,
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

func (svc *applicationService) Remove(ctx context.Context, in *external.RemoveRequest, out *commons.Status) error {
	return wrapper.ContextToAuthorize(ctx, out, func(auth *wrapper.WrapperUser) *commons.State {
		return errstate.Success
	})
}

func (svc *applicationService) ChangeName(ctx context.Context, in *external.ChangeNameRequest, out *commons.Status) error {
	return wrapper.ContextToAuthorize(ctx, out, func(auth *wrapper.WrapperUser) *commons.State {
		return nil
	})
}

func (svc *applicationService) FindByAppId(ctx context.Context, in *external.FindRequest, out *external.SimpleApplicationResponse) error {
	return wrapper.ContextToAuthorize(ctx, out, func(auth *wrapper.WrapperUser) *commons.State {

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
			Settings: &commons.AppSettings{
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
	return wrapper.ContextToAuthorize(ctx, out, func(auth *wrapper.WrapperUser) *commons.State {

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
			case constants.AppTypeManage, constants.AppTypeRoute, constants.AppTypeSafe, constants.AppTypeUser:
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

func (svc *applicationService) Switch(context.Context, *external.SwitchRequest, *commons.Status) error {
	panic("implement me")
}

func NewApplicationService(session *mgo.Session, client *indexutils.Client, log analysisclient.LogClient) external.ApplicationHandler {
	return &applicationService{session: session, Client: client, log: log, id: generator.NewIDG()}
}
