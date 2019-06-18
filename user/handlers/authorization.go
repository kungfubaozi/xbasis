package userhandlers

import (
	"context"
	"fmt"
	"gopkg.in/mgo.v2"
	"konekko.me/gosion/application/pb/inner"
	"konekko.me/gosion/commons/dto"
	"konekko.me/gosion/commons/errstate"
	"konekko.me/gosion/commons/indexutils"
	"konekko.me/gosion/commons/wrapper"
	external "konekko.me/gosion/user/pb"
)

type authorizationService struct {
	usersyncService          gosionsvc_internal_application.UsersyncService
	applicationStatusService gosionsvc_internal_application.ApplicationStatusService
	session                  *mgo.Session
	client                   *indexutils.Client
}

func (svc *authorizationService) GetRepo() *userRepo {
	return &userRepo{session: svc.session.Clone(), Client: svc.client}
}

func (svc *authorizationService) Sync(ctx context.Context, in *external.SyncRequest, out *gs_commons_dto.Status) error {
	return gs_commons_wrapper.ContextToAuthorize(ctx, out, func(auth *gs_commons_wrapper.WrapperUser) *gs_commons_dto.State {

		if len(in.ClientId) > 8 {

			v, err := svc.applicationStatusService.GetAppClientStatus(ctx, &gosionsvc_internal_application.GetAppClientStatusRequest{
				ClientId: in.ClientId,
			})
			if err != nil {
				fmt.Println("err 1", err)
				return errstate.ErrRequest
			}

			if !v.State.Ok {
				fmt.Println("err 2", v.State)
				return v.State
			}

			c, err := svc.usersyncService.Check(ctx, &gosionsvc_internal_application.CheckRequest{
				UserId: auth.Token.UserId,
				AppId:  v.AppId,
			})
			if err != nil {
				return errstate.ErrSystem
			}
			if c.State.Code == errstate.ErrUserNotSync.Code {
				repo := svc.GetRepo()
				defer repo.Close()

				info, err := repo.FindUserInfo(auth.Token.UserId)
				if err != nil {
					return nil
				}

				s, err := svc.usersyncService.Transport(ctx, &gosionsvc_internal_application.UserInfo{
					GId:      info.UserId,
					Username: info.Username,
					Icon:     info.Icon,
					RealName: info.RealName,
					AppId:    v.AppId,
					AppType:  v.Type,
				})
				if err != nil {
					fmt.Println("err 3", err)
					return errstate.ErrRequest
				}

				fmt.Println("err 4", s.State)

				return s.State
			}
			return c.State
		}

		return nil
	})
}

func NewAuthorizationService(session *mgo.Session, userSyncService gosionsvc_internal_application.UsersyncService,
	applicationStatusService gosionsvc_internal_application.ApplicationStatusService, client *indexutils.Client) external.AuthorizationHandler {
	return &authorizationService{session: session, usersyncService: userSyncService, applicationStatusService: applicationStatusService, client: client}
}
