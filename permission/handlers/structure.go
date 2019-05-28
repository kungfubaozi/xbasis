package permissionhandlers

import (
	"context"
	"gopkg.in/mgo.v2"
	"konekko.me/gosion/application/pb"
	"konekko.me/gosion/commons/dto"
	"konekko.me/gosion/commons/errstate"
	"konekko.me/gosion/commons/generator"
	"konekko.me/gosion/commons/indexutils"
	"konekko.me/gosion/commons/wrapper"
	"konekko.me/gosion/permission/pb"
	"konekko.me/gosion/permission/utils"
	"time"
)

type structureService struct {
	session            *mgo.Session
	applicationService gs_service_application.ApplicationService
	*indexutils.Client
}

func (svc *structureService) CreateUserStructure(ctx context.Context, in *gs_service_permission.CreateRequest, out *gs_commons_dto.Status) error {
	return gs_commons_wrapper.ContextToAuthorize(ctx, out, func(auth *gs_commons_wrapper.WrapperUser) *gs_commons_dto.State {
		return svc.Create(ctx, in.Name, auth.User, in.AppId, permissionutils.TypeUserStructure)
	})
}

func (svc *structureService) CreateFunctionStructure(ctx context.Context, in *gs_service_permission.CreateRequest, out *gs_commons_dto.Status) error {
	return gs_commons_wrapper.ContextToAuthorize(ctx, out, func(auth *gs_commons_wrapper.WrapperUser) *gs_commons_dto.State {
		return svc.Create(ctx, in.Name, auth.User, in.AppId, permissionutils.TypeFunctionStructure)
	})
}

func (svc *structureService) GetRepo() *structureRepo {
	return &structureRepo{Client: svc.Client, session: svc.session.Clone()}
}

func (svc *structureService) Create(ctx context.Context, name, user, appId string, t int64) *gs_commons_dto.State {
	if len(name) == 0 {
		return nil
	}

	repo := svc.GetRepo()
	defer repo.Close()

	id := gs_commons_generator.NewIDG()

	c, err := repo.FindCountByNameAndType(name, t)
	if err != nil && err == mgo.ErrNotFound || c == 0 { //not found
		err = repo.Add(&structure{
			Id:           id.Get(),
			Name:         name,
			CreateAt:     time.Now().UnixNano(),
			CreateUserId: user,
			AppId:        appId,
			Type:         t,
		})
		if err == nil {
			return errstate.Success
		}
	}

	return nil
}

func (svc *structureService) GetList(ctx context.Context, in *gs_service_permission.GetListRequest, out *gs_service_permission.GetListResponse) error {
	return gs_commons_wrapper.ContextToAuthorize(ctx, out, func(auth *gs_commons_wrapper.WrapperUser) *gs_commons_dto.State {

		if len(in.AppId) == 0 {
			return nil
		}

		repo := svc.GetRepo()
		defer repo.Close()

		app, err := svc.applicationService.FindByAppId(ctx, &gs_service_application.FindRequest{
			Content: in.AppId,
		})
		if err != nil {
			return nil
		}

		if !app.State.Ok {
			return app.State
		}

		t := permissionutils.TypeFunctionStructure
		if in.User {
			t = permissionutils.TypeUserStructure
		}

		s, err := repo.FindByAppIdAndType(in.AppId, t)
		if err != nil {
			return nil
		}

		var data []*gs_service_permission.SimpleStructure
		for _, v := range s {
			opening := false
			if v.Type == permissionutils.TypeFunctionStructure && v.Id == app.Info.FuncS {
				opening = true
			} else if v.Type == permissionutils.TypeUserStructure && v.Id == app.Info.UserS {
				opening = true
			}
			data = append(data, &gs_service_permission.SimpleStructure{
				Name:     v.Name,
				CreateAt: v.CreateAt,
				Opening:  opening,
				Id:       v.Id,
			})
		}

		if data != nil {
			out.Data = data
			return errstate.Success
		}

		return nil
	})
}

func NewStructureService(session *mgo.Session, client *indexutils.Client) gs_service_permission.StructureHandler {
	return &structureService{session: session, Client: client}
}
