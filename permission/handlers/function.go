package permissionhandlers

import (
	"context"
	"fmt"
	"github.com/olivere/elastic"
	"gopkg.in/mgo.v2"
	"konekko.me/gosion/analysis/client"
	"konekko.me/gosion/commons/actions"
	"konekko.me/gosion/commons/constants"
	"konekko.me/gosion/commons/date"
	"konekko.me/gosion/commons/dto"
	"konekko.me/gosion/commons/errstate"
	"konekko.me/gosion/commons/generator"
	"konekko.me/gosion/commons/indexutils"
	"konekko.me/gosion/commons/wrapper"
	external "konekko.me/gosion/permission/pb"
	"sync"
	"time"
)

type functionService struct {
	*indexutils.Client
	session *mgo.Session
	id      gs_commons_generator.IDGenerator
	log     analysisclient.LogClient
}

func (svc *functionService) Search(context.Context, *external.FunctionSearchRequest, *external.FunctionSearchResponse) error {
	panic("implement me")
}

func (svc *functionService) GetFunctionItems(ctx context.Context, in *external.GetFunctionItemsRequest, out *external.GetFunctionItemsResponse) error {
	return gs_commons_wrapper.ContextToAuthorize(ctx, out, func(auth *gs_commons_wrapper.WrapperUser) *gs_commons_dto.State {
		if len(in.AppId) == 0 {
			return nil
		}

		repo := svc.GetRepo()
		defer repo.Close()

		groups, err := repo.FindChildGroups(in.AppId, in.Id)
		if err != nil {
			fmt.Println("123")
			return nil
		}

		var data []*external.FindItemResponse
		for _, v := range groups {
			data = append(data, &external.FindItemResponse{
				Function: false,
				Id:       v.Id,
				Name:     v.Name,
			})
		}

		//find group and function
		if len(in.Id) > 0 {
			functions, err := repo.FindChildFunctions(in.AppId, in.Id)
			if err != nil {
				return nil
			}
			for _, v := range functions {
				data = append(data, &external.FindItemResponse{
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

func (svc *functionService) GetFunctionItemDetail(ctx context.Context, in *external.GetFunctionItemRequest, out *external.GetFunctionItemResponse) error {
	return gs_commons_wrapper.ContextToAuthorize(ctx, out, func(auth *gs_commons_wrapper.WrapperUser) *gs_commons_dto.State {
		if len(in.Id) == 0 || len(in.AppId) == 0 {
			return nil
		}
		repo := svc.GetRepo()
		defer repo.Close()

		headers := &analysisclient.LogHeaders{
			TraceId:     auth.TraceId,
			ModuleName:  "Function",
			ServiceName: gs_commons_constants.PermissionService,
		}

		f, err := repo.FindApiById(in.AppId, in.Id)
		if err != nil {
			svc.log.Info(&analysisclient.LogContent{
				Headers: headers,
				Action:  loggeractions.InvalidFunction,
				Message: "not found function",
			})
			return nil
		}

		function := &external.FunctionItemDetail{
			Name:          f.Name,
			Id:            f.Id,
			ValTokenTimes: f.ValTokenTimes,
			Api:           f.Api,
			CreateAt:      f.CreateAt,
			Share:         f.Share,
		}

		if len(f.Roles) > 0 {
			rrepo := svc.GetRoleRepo()
			defer rrepo.Close()
			size := len(f.Roles)
			var wg sync.WaitGroup
			var err error

			setRole := func(r *role) {
				function.Roles = append(function.Roles, &external.FunctionBindRole{
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

		function.AuthTypes = []*external.FunctionAuthTypes{
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

		function.Platforms = []*external.FunctionGrantPlatforms{
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

		svc.log.Info(&analysisclient.LogContent{
			Headers: headers,
			Action:  loggeractions.FindFunction,
			Message: "found",
		})

		now := gs_commons_date.FormatDate(time.Now(), gs_commons_date.YYYY_I_MM_I_DD)
		q := elastic.NewBoolQuery()
		q.Must(elastic.NewMatchPhraseQuery("fields.id", f.Id), elastic.NewMatchPhraseQuery("action", "UserRequestApi"))

		var wg sync.WaitGroup
		wg.Add(4)
		//today visit count
		go func() {
			defer wg.Done()
			c, err := svc.GetElasticClient().Count("gosion-logger." + now).Type("_doc").Query(q).Do(context.Background())
			if err != nil {
				return
			}
			function.TodayVisit = c
		}()

		//today visit user count
		go func() {
			defer wg.Done()
			v, err := svc.GetElasticClient().Search("gosion-logger."+now).Type("_doc").Query(q).Aggregation("count", elastic.NewCardinalityAggregation().Field("headers.userId.keyword")).Do(context.Background())
			if err != nil {
				return
			}
			if v.Aggregations != nil {
				c, ok := v.Aggregations.Cardinality("count")
				if ok {
					function.TodayVisitUser = int64(*c.Value)
				}
			}
		}()

		//month
		go func() {
			defer wg.Done()
			t := gs_commons_date.FormatDate(time.Now(), gs_commons_date.YYYY_I_MM)
			c, err := svc.GetElasticClient().Count(fmt.Sprintf("gosion-logger.%s.*", t)).Type("_doc").Query(q).Do(context.Background())
			if err != nil {
				return
			}
			function.MonthVisit = c
		}()

		//total
		go func() {
			defer wg.Done()
			c, err := svc.GetElasticClient().Count("gosion-logger.*").Type("_doc").Query(q).Do(context.Background())
			if err != nil {
				return
			}
			function.TotalVisit = c
		}()

		wg.Wait()

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

func (svc *functionService) Add(ctx context.Context, in *external.FunctionRequest, out *gs_commons_dto.Status) error {
	return gs_commons_wrapper.ContextToAuthorize(ctx, out, func(auth *gs_commons_wrapper.WrapperUser) *gs_commons_dto.State {

		repo := svc.GetRepo()
		defer repo.Close()

		if len(in.AppId) == 0 || len(in.BindGroupId) == 0 || len(in.Name) == 0 {
			return errstate.ErrRequest
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

		_, err := repo.FindApi(in.AppId, in.Api)
		if err != nil && err == mgo.ErrNotFound {

			f := &function{
				Id:           svc.id.String(),
				Name:         in.Name,
				Type:         in.Type,
				CreateUserId: auth.Token.UserId,
				CreateAt:     time.Now().UnixNano(),
				BindGroupId:  in.BindGroupId,
				AppId:        in.AppId,
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

func (svc *functionService) Rename(ctx context.Context, in *external.FunctionRequest, out *gs_commons_dto.Status) error {
	return gs_commons_wrapper.ContextToAuthorize(ctx, out, func(auth *gs_commons_wrapper.WrapperUser) *gs_commons_dto.State {
		return nil
	})
}

func (svc *functionService) Move(ctx context.Context, in *external.FunctionRequest, out *gs_commons_dto.Status) error {
	return gs_commons_wrapper.ContextToAuthorize(ctx, out, func(auth *gs_commons_wrapper.WrapperUser) *gs_commons_dto.State {
		return nil
	})
}

//one application one root group, bindGroupId = appId
func (svc *functionService) AddGroup(ctx context.Context, in *external.FunctionGroupRequest, out *gs_commons_dto.Status) error {
	return gs_commons_wrapper.ContextToAuthorize(ctx, out, func(auth *gs_commons_wrapper.WrapperUser) *gs_commons_dto.State {

		if len(in.Name) == 0 && len(in.AppId) == 0 {
			return nil
		}

		repo := svc.GetRepo()
		defer repo.Close()

		i := svc.id.Get()

		err := repo.AddGroup(&functionGroup{
			Id:           i,
			Name:         in.Name,
			AppId:        in.AppId,
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

func (svc *functionService) MoveGroup(ctx context.Context, in *external.FunctionGroupRequest, out *gs_commons_dto.Status) error {
	return gs_commons_wrapper.ContextToAuthorize(ctx, out, func(auth *gs_commons_wrapper.WrapperUser) *gs_commons_dto.State {
		return nil
	})
}

func (svc *functionService) RenameGroup(ctx context.Context, in *external.FunctionGroupRequest, out *gs_commons_dto.Status) error {
	return gs_commons_wrapper.ContextToAuthorize(ctx, out, func(auth *gs_commons_wrapper.WrapperUser) *gs_commons_dto.State {
		return nil
	})
}

func NewFunctionService(client *indexutils.Client, session *mgo.Session, log analysisclient.LogClient) external.FunctionHandler {
	return &functionService{Client: client, session: session, id: gs_commons_generator.NewIDG(), log: log}
}
