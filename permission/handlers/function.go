package permissionhandlers

import (
	"context"
	"fmt"
	"gopkg.in/mgo.v2"
	"konekko.me/gosion/commons/constants"
	"konekko.me/gosion/commons/dto"
	"konekko.me/gosion/commons/errstate"
	"konekko.me/gosion/commons/generator"
	"konekko.me/gosion/commons/indexutils"
	"konekko.me/gosion/commons/wrapper"
	"konekko.me/gosion/permission/pb"
	"sync"
	"time"
)

type functionService struct {
	*indexutils.Client
	session *mgo.Session
	id      gs_commons_generator.IDGenerator
}

func (svc *functionService) GetFunctionItems(ctx context.Context, in *gs_service_permission.GetFunctionItemsRequest, out *gs_service_permission.GetFunctionItemsResponse) error {
	return gs_commons_wrapper.ContextToAuthorize(ctx, out, func(auth *gs_commons_wrapper.WrapperUser) *gs_commons_dto.State {
		if len(in.StructureId) == 0 {
			return nil
		}

		repo := svc.GetRepo()
		defer repo.Close()

		groups, err := repo.FindChildGroups(in.StructureId, in.Id)
		if err != nil {
			return nil
		}

		var data []*gs_service_permission.FindItemResponse
		for _, v := range groups {
			data = append(data, &gs_service_permission.FindItemResponse{
				Function: false,
				Id:       v.Id,
				Name:     v.Name,
			})
		}

		//find group and function
		if len(in.Id) > 0 {
			functions, err := repo.FindChildFunctions(in.StructureId, in.Id)
			if err != nil {
				return nil
			}
			for _, v := range functions {
				data = append(data, &gs_service_permission.FindItemResponse{
					Function: true,
					Id:       v.Id,
					Name:     v.Name,
				})
			}
		}

		out.Data = data
		return errstate.Success
	})
}

func (svc *functionService) GetFunctionItemDetail(ctx context.Context, in *gs_service_permission.GetFunctionItemRequest, out *gs_service_permission.GetFunctionItemResponse) error {
	return gs_commons_wrapper.ContextToAuthorize(ctx, out, func(auth *gs_commons_wrapper.WrapperUser) *gs_commons_dto.State {
		if len(in.Id) == 0 || len(in.StructureId) == 0 {
			return nil
		}
		repo := svc.GetRepo()
		defer repo.Close()

		f, err := repo.FindApiById(in.StructureId, in.Id)
		if err != nil {
			fmt.Println("find api err", err)
			return nil
		}

		function := &gs_service_permission.FunctionItemDetail{
			Name:         f.Name,
			Id:           f.Id,
			ValTokenLife: f.ValTokenLife,
			Api:          f.Api,
			CreateAt:     f.CreateAt,
			Share:        f.Share,
		}

		if len(f.Roles) > 0 {
			rrepo := svc.GetRoleRepo()
			defer rrepo.Close()
			size := len(f.Roles)
			var wg sync.WaitGroup
			var err error

			setRole := func(r *role) {
				function.Roles = append(function.Roles, &gs_service_permission.FunctionBindRole{
					Name: r.Name,
					Id:   r.Id,
				})
			}

			resp := func(e error) {
				if err == nil {
					err = e
				}
			}

			findRole := func(v string) {
				role, err := rrepo.FindRoleById(v)
				if err != nil {
					resp(err)
					return
				}
				setRole(role)
			}

			if size > 2 {
				wg.Add(2)
				s := size / 2
				go func() {
					defer wg.Done()
					a := f.Roles[:s]
					for _, v := range a {
						findRole(v)
					}
				}()
				go func() {
					defer wg.Done()
					a := f.Roles[s:]
					for _, v := range a {
						findRole(v)
					}
				}()
			} else {
				wg.Add(len(f.Roles))
				for _, v := range f.Roles {
					go func() {
						defer wg.Done()
						findRole(v)
					}()
				}
			}
			wg.Wait()
			if err != nil {
				return nil
			}
		}

		isAuthEnabled := func(t int64) bool {
			for _, v := range f.AuthTypes {
				if v == t {
					return true
				}
			}
			return false
		}

		function.AuthTypes = []*gs_service_permission.FunctionAuthTypes{
			{
				Name:    "ValCode",
				Type:    gs_commons_constants.AuthTypeOfValcode,
				Enabled: isAuthEnabled(gs_commons_constants.AuthTypeOfValcode),
			},
			{
				Name:    "Token",
				Type:    gs_commons_constants.AuthTypeOfToken,
				Enabled: isAuthEnabled(gs_commons_constants.AuthTypeOfToken),
			},
			{
				Name:    "MobileConfirm",
				Type:    gs_commons_constants.AuthTypeOfMobileConfirm,
				Enabled: isAuthEnabled(gs_commons_constants.AuthTypeOfMobileConfirm),
			},
			{
				Name:    "Face",
				Type:    gs_commons_constants.AuthTypeOfFace,
				Enabled: isAuthEnabled(gs_commons_constants.AuthTypeOfFace),
			},
			{
				Name:    "Gosion Mini Program",
				Type:    gs_commons_constants.AuthTypeOfMiniProgramCodeConfirm,
				Enabled: isAuthEnabled(gs_commons_constants.AuthTypeOfMiniProgramCodeConfirm),
			},
		}

		isGrant := func(t int64) bool {
			for _, v := range f.GrantPlatforms {
				if v == t {
					return true
				}
			}
			return false
		}

		function.Platforms = []*gs_service_permission.FunctionGrantPlatforms{
			{
				Name:    "Web",
				Type:    gs_commons_constants.PlatformOfWeb,
				Enabled: isGrant(gs_commons_constants.PlatformOfWeb),
			},
			{
				Name:    "Android",
				Type:    gs_commons_constants.PlatformOfAndroid,
				Enabled: isGrant(gs_commons_constants.PlatformOfAndroid),
			},
			{
				Name:    "Fuchsia",
				Type:    gs_commons_constants.PlatformOfFuchsia,
				Enabled: isGrant(gs_commons_constants.PlatformOfFuchsia),
			},
			{
				Name:    "iOS",
				Type:    gs_commons_constants.PlatformOfIOS,
				Enabled: isGrant(gs_commons_constants.PlatformOfIOS),
			},
			{
				Name:    "Linux",
				Type:    gs_commons_constants.PlatformOfLinux,
				Enabled: isGrant(gs_commons_constants.PlatformOfLinux),
			},
			{
				Name:    "Windows",
				Type:    gs_commons_constants.PlatformOfWindows,
				Enabled: isGrant(gs_commons_constants.PlatformOfWindows),
			},
		}

		out.Data = function

		return errstate.Success
	})
}

func (svc *functionService) GetRepo() *functionRepo {
	return &functionRepo{session: svc.session.Clone(), Client: svc.Client}
}

func (svc *functionService) GetRoleRepo() *roleRepo {
	return &roleRepo{session: svc.session.Clone()}
}

func (svc *functionService) Add(ctx context.Context, in *gs_service_permission.FunctionRequest, out *gs_commons_dto.Status) error {
	return gs_commons_wrapper.ContextToAuthorize(ctx, out, func(auth *gs_commons_wrapper.WrapperUser) *gs_commons_dto.State {

		repo := svc.GetRepo()
		defer repo.Close()

		if len(in.StructureId) == 0 || len(in.BindGroupId) == 0 || len(in.Name) == 0 {
			return errstate.ErrRequest
		}

		if isStructureExists(repo.session, in.StructureId) == 0 {
			return errstate.ErrInvalidStructure
		}

		//find group exists
		if !repo.FindGroupExists(in.BindGroupId) {
			return errstate.ErrFunctionBindGroupId
		}

		//check authTypes
		if in.AuthTypes != nil && len(in.AuthTypes) > 0 {
			for _, v := range in.AuthTypes {
				switch v {
				case gs_commons_constants.AuthTypeOfFace:
				case gs_commons_constants.AuthTypeOfToken:
				case gs_commons_constants.AuthTypeOfValcode:
				case gs_commons_constants.AuthTypeOfMobileConfirm:
				case gs_commons_constants.AuthTypeOfMiniProgramUserConfirm:
				case gs_commons_constants.AuthTypeOfMiniProgramCodeConfirm:
				default:
					return errstate.ErrFunctionAuthType
				}
			}
		}

		_, err := repo.FindApi(in.StructureId, in.Api)
		if err != nil && err == mgo.ErrNotFound {

			f := &function{
				Id:           svc.id.String(),
				Name:         in.Name,
				Type:         in.Type,
				CreateUserId: auth.Token.UserId,
				CreateAt:     time.Now().UnixNano(),
				BindGroupId:  in.BindGroupId,
				StructureId:  in.StructureId,
				Api:          in.Api,
				AuthTypes:    in.AuthTypes,
				GrantPlatforms: []int64{
					gs_commons_constants.PlatformOfWindows,
					gs_commons_constants.PlatformOfLinux,
					gs_commons_constants.PlatformOfIOS,
					gs_commons_constants.PlatformOfFuchsia,
					gs_commons_constants.PlatformOfAndroid,
					gs_commons_constants.PlatformOfWeb},
			}

			err := repo.AddFunction(f)

			if err == nil {
				return errstate.Success
			}

		}

		return nil
	})
}

func (svc *functionService) Rename(ctx context.Context, in *gs_service_permission.FunctionRequest, out *gs_commons_dto.Status) error {
	return gs_commons_wrapper.ContextToAuthorize(ctx, out, func(auth *gs_commons_wrapper.WrapperUser) *gs_commons_dto.State {
		return nil
	})
}

func (svc *functionService) Move(ctx context.Context, in *gs_service_permission.FunctionRequest, out *gs_commons_dto.Status) error {
	return gs_commons_wrapper.ContextToAuthorize(ctx, out, func(auth *gs_commons_wrapper.WrapperUser) *gs_commons_dto.State {
		return nil
	})
}

//one application one root group, bindGroupId = appId
func (svc *functionService) AddGroup(ctx context.Context, in *gs_service_permission.FunctionGroupRequest, out *gs_commons_dto.Status) error {
	return gs_commons_wrapper.ContextToAuthorize(ctx, out, func(auth *gs_commons_wrapper.WrapperUser) *gs_commons_dto.State {

		if len(in.Name) == 0 && len(in.StructureId) == 0 {
			return nil
		}

		repo := svc.GetRepo()
		defer repo.Close()

		i := svc.id.Get()

		err := repo.AddGroup(&functionGroup{
			Id:           i,
			Name:         in.Name,
			StructureId:  in.StructureId,
			BindGroupId:  in.BindGroupId,
			CreateAt:     time.Now().UnixNano(),
			CreateUserId: auth.Token.UserId,
		})

		if err != nil {
			return nil
		}

		out.Content = i

		return errstate.Success
	})
}

func (svc *functionService) MoveGroup(ctx context.Context, in *gs_service_permission.FunctionGroupRequest, out *gs_commons_dto.Status) error {
	return gs_commons_wrapper.ContextToAuthorize(ctx, out, func(auth *gs_commons_wrapper.WrapperUser) *gs_commons_dto.State {
		return nil
	})
}

func (svc *functionService) RenameGroup(ctx context.Context, in *gs_service_permission.FunctionGroupRequest, out *gs_commons_dto.Status) error {
	return gs_commons_wrapper.ContextToAuthorize(ctx, out, func(auth *gs_commons_wrapper.WrapperUser) *gs_commons_dto.State {
		return nil
	})
}

func NewFunctionService(client *indexutils.Client, session *mgo.Session) gs_service_permission.FunctionHandler {
	return &functionService{Client: client, session: session, id: gs_commons_generator.NewIDG()}
}
