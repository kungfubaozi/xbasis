package permissionhandlers

import (
	"context"
	"github.com/olivere/elastic"
	"gopkg.in/mgo.v2"
	"konekko.me/xbasis/analysis/client"
	"konekko.me/xbasis/commons/actions"
	constants "konekko.me/xbasis/commons/constants"
	commons "konekko.me/xbasis/commons/dto"
	"konekko.me/xbasis/commons/errstate"
	generator "konekko.me/xbasis/commons/generator"
	"konekko.me/xbasis/commons/indexutils"
	"konekko.me/xbasis/commons/transport"
	wrapper "konekko.me/xbasis/commons/wrapper"
	"konekko.me/xbasis/gateway/client"
	external "konekko.me/xbasis/permission/pb"
	"time"
)

type functionService struct {
	*indexutils.Client
	session *mgo.Session
	id      generator.IDGenerator
	log     analysisclient.LogClient
	gateway xbsgatewayclient.GatewayClient
}

func (svc *functionService) ModifySettings(ctx context.Context, in *external.ModifySettingsRequest, out *commons.Status) error {
	return wrapper.ContextToAuthorize(ctx, out, func(auth *wrapper.WrapperUser) *commons.State {

		return nil
	})
}

func (svc *functionService) Search(ctx context.Context, in *external.FunctionSearchRequest, out *external.FunctionSearchResponse) error {
	return wrapper.ContextToAuthorize(ctx, out, func(auth *wrapper.WrapperUser) *commons.State {

		return nil
	})
}

func (svc *functionService) GetFunctionItems(ctx context.Context, in *external.GetFunctionItemsRequest, out *external.GetFunctionItemsResponse) error {
	return wrapper.ContextToAuthorize(ctx, out, func(auth *wrapper.WrapperUser) *commons.State {
		if len(in.AppId) == 0 {
			return nil
		}

		repo := svc.GetRepo()
		defer repo.Close()

		groups, err := repo.FindChildGroups(in.AppId, in.Id)
		if err != nil {
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
	return wrapper.ContextToAuthorize(ctx, out, func(auth *wrapper.WrapperUser) *commons.State {
		if len(in.Id) == 0 || len(in.AppId) == 0 {
			return nil
		}
		repo := svc.GetRepo()
		defer repo.Close()

		headers := &analysisclient.LogHeaders{
			TraceId:     auth.TraceId,
			ModuleName:  "Function",
			ServiceName: constants.PermissionService,
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
				Type:    constants.AuthTypeOfValcode,
				Enabled: isAuthEnabled(constants.AuthTypeOfValcode),
			},
			{
				Name:    "Token",
				Type:    constants.AuthTypeOfToken,
				Enabled: isAuthEnabled(constants.AuthTypeOfToken),
			},
			{
				Name:    "MobileConfirm",
				Type:    constants.AuthTypeOfMobileConfirm,
				Enabled: isAuthEnabled(constants.AuthTypeOfMobileConfirm),
			},
			{
				Name:    "Face",
				Type:    constants.AuthTypeOfFace,
				Enabled: isAuthEnabled(constants.AuthTypeOfFace),
			},
			{
				Name:    "Gosion Mini Program",
				Type:    constants.AuthTypeOfMiniProgramCodeConfirm,
				Enabled: isAuthEnabled(constants.AuthTypeOfMiniProgramCodeConfirm),
			},
		}

		isGrant := func(t int64) bool {
			for _, v := range f.NoGrantPlatforms {
				if v == t {
					return false
				}
			}
			return true
		}

		function.Platforms = []*external.FunctionNoGrantPlatforms{
			{
				Name:    "Web",
				Type:    constants.PlatformOfWeb,
				Enabled: isGrant(constants.PlatformOfWeb),
			},
			{
				Name:    "Android",
				Type:    constants.PlatformOfAndroid,
				Enabled: isGrant(constants.PlatformOfAndroid),
			},
			{
				Name:    "Fuchsia",
				Type:    constants.PlatformOfFuchsia,
				Enabled: isGrant(constants.PlatformOfFuchsia),
			},
			{
				Name:    "iOS",
				Type:    constants.PlatformOfIOS,
				Enabled: isGrant(constants.PlatformOfIOS),
			},
			{
				Name:    "Linux",
				Type:    constants.PlatformOfLinux,
				Enabled: isGrant(constants.PlatformOfLinux),
			},
			{
				Name:    "Windows",
				Type:    constants.PlatformOfWindows,
				Enabled: isGrant(constants.PlatformOfWindows),
			},
		}

		svc.log.Info(&analysisclient.LogContent{
			Headers: headers,
			Action:  loggeractions.FindFunction,
			Message: "found",
		})

		//now := xbasisdate.FormatDate(time.Now(), xbasisdate.YYYY_I_MM_I_DD)
		//q := elastic.NewBoolQuery()
		//q.Must(elastic.NewMatchPhraseQuery("funcId", f.Id), elastic.NewMatchPhraseQuery("action", loggeractions.UserRequestApi))
		//
		//var wg sync.WaitGroup
		//wg.Add(2)
		//
		////today
		//go func() {
		//	defer wg.Done()
		//	total, user, err := svc.find(now, q)
		//	if err != nil {
		//		return
		//	}
		//	function.TodayVisit = total / 4
		//	function.TodayUserVisit = user
		//}()
		//
		////total
		//go func() {
		//	defer wg.Done()
		//	total, user, err := svc.find("*", q)
		//	if err != nil {
		//		return
		//	}
		//	function.TotalVisit = total / 4
		//	function.TotalUserVisit = user
		//}()
		//
		//wg.Wait()

		out.Data = function

		return errstate.Success
	})
}

func (svc *functionService) find(now string, q elastic.Query) (int64, int64, error) {
	v, err := svc.GetElasticClient().Search(statementIndex+now).Type("_doc").Query(q).
		Aggregation("user", elastic.NewCardinalityAggregation().Field("who.keyword")).
		Aggregation("total", elastic.NewSumAggregation().Field("total")).
		Do(context.Background())
	if err != nil {
		return 0, 0, err
	}
	var total int64 = 0
	var user int64 = 0
	if v.Aggregations != nil {
		c, ok := v.Aggregations.Cardinality("total")
		if ok {
			total = int64(*c.Value)
		}
		c, ok = v.Aggregations.Sum("user")
		if ok {
			user = int64(*c.Value)
		}
	}
	return total, user, nil
}

func (svc *functionService) GetRepo() *functionRepo {
	return &functionRepo{session: svc.session.Clone(), Client: svc.Client}
}

func (svc *functionService) GetRoleRepo() *roleRepo {
	return &roleRepo{session: svc.session.Clone()}
}

func (svc *functionService) Add(ctx context.Context, in *external.FunctionRequest, out *commons.Status) error {
	return wrapper.ContextToAuthorize(ctx, out, func(auth *wrapper.WrapperUser) *commons.State {

		repo := svc.GetRepo()
		defer repo.Close()

		if len(in.AppId) == 0 || len(in.BindGroupId) == 0 || len(in.Name) == 0 {
			return errstate.ErrRequest
		}

		//find group exists
		if !repo.FindGroupExists(in.BindGroupId, in.AppId) {
			return errstate.ErrFunctionBindGroupId
		}

		header := &analysisclient.LogHeaders{
			TraceId:     auth.TraceId,
			ServiceName: constants.PermissionService,
			ModuleName:  "AddFunction",
		}

		//check authTypes
		if in.AuthTypes != nil && len(in.AuthTypes) > 0 {
			for _, v := range in.AuthTypes {
				switch v {
				case constants.AuthTypeOfFace:
				case constants.AuthTypeOfToken:
				case constants.AuthTypeOfValcode:
				case constants.AuthTypeOfMobileConfirm:
				case constants.AuthTypeOfMiniProgramUserConfirm:
				case constants.AuthTypeOfMiniProgramCodeConfirm:
				default:
					return errstate.ErrFunctionAuthType
				}
			}
		}

		_, err := repo.FindApi(in.AppId, in.Api)
		if err != nil && err == mgo.ErrNotFound {

			f := &function{
				Id:               svc.id.Get(),
				Name:             in.Name,
				Type:             in.Type,
				CreateUserId:     auth.Token.UserId,
				CreateAt:         time.Now().UnixNano(),
				BindGroupId:      in.BindGroupId,
				AppId:            in.AppId,
				Api:              in.Api,
				AuthTypes:        in.AuthTypes,
				NoGrantPlatforms: []int64{},
			}

			err := repo.AddFunction(f)

			if err == nil {

				svc.gateway.SendFunctionChanged(&xbasistransport.AppFunction{
					Id:               f.Id,
					Name:             f.Name,
					NoGrantPlatforms: f.NoGrantPlatforms,
					AppId:            f.AppId,
					Path:             f.Api,
					AuthTypes:        f.AuthTypes,
					ValTokenTimes:    f.ValTokenTimes,
					Share:            f.Share,
				})

				svc.log.Info(&analysisclient.LogContent{
					Headers: header,
					Action:  "AddFunction",
					Fields: &analysisclient.LogFields{
						"id":     f.Id,
						"name":   f.Name,
						"app_id": f.AppId,
					},
				})

				return errstate.Success
			}

		}

		return nil
	})
}

func (svc *functionService) Rename(ctx context.Context, in *external.FunctionRequest, out *commons.Status) error {
	return wrapper.ContextToAuthorize(ctx, out, func(auth *wrapper.WrapperUser) *commons.State {
		return nil
	})
}

func (svc *functionService) Move(ctx context.Context, in *external.FunctionRequest, out *commons.Status) error {
	return wrapper.ContextToAuthorize(ctx, out, func(auth *wrapper.WrapperUser) *commons.State {
		return nil
	})
}

//one application one root group, bindGroupId = appId
func (svc *functionService) AddGroup(ctx context.Context, in *external.FunctionGroupRequest, out *commons.Status) error {
	return wrapper.ContextToAuthorize(ctx, out, func(auth *wrapper.WrapperUser) *commons.State {

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

func (svc *functionService) MoveGroup(ctx context.Context, in *external.FunctionGroupRequest, out *commons.Status) error {
	return wrapper.ContextToAuthorize(ctx, out, func(auth *wrapper.WrapperUser) *commons.State {
		return nil
	})
}

func (svc *functionService) RenameGroup(ctx context.Context, in *external.FunctionGroupRequest, out *commons.Status) error {
	return wrapper.ContextToAuthorize(ctx, out, func(auth *wrapper.WrapperUser) *commons.State {
		return nil
	})
}

func NewFunctionService(client *indexutils.Client, session *mgo.Session, log analysisclient.LogClient, gateway xbsgatewayclient.GatewayClient) external.FunctionHandler {
	return &functionService{Client: client, session: session, id: generator.NewIDG(), log: log, gateway: gateway}
}
