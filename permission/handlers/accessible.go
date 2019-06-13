package permissionhandlers

import (
	"context"
	"konekko.me/gosion/analysis/client"
	"konekko.me/gosion/commons/dto"
	"konekko.me/gosion/commons/errstate"
	"konekko.me/gosion/commons/indexutils"
	"konekko.me/gosion/commons/wrapper"
	inner "konekko.me/gosion/permission/pb/inner"
)

type accessibleService struct {
	*indexutils.Client
	log analysisclient.LogClient
}

func (svc *accessibleService) Check(ctx context.Context, in *inner.CheckRequest, out *gs_commons_dto.Status) error {
	return gs_commons_wrapper.ContextToAuthorize(ctx, out, func(auth *gs_commons_wrapper.WrapperUser) *gs_commons_dto.State {
		//get user require roles
		//var userroles map[string]interface{}
		//
		//headers := &analysisclient.LogHeaders{
		//	TraceId:     auth.TraceId,
		//	ModuleName:  "Accessible",
		//	ServiceName: gs_commons_constants.InternalPermission,
		//}
		//
		//ok, err := svc.Client.QueryFirst("gs-user-roles-relation",
		//	map[string]interface{}{"structure_id": in.StructureId, "user_id": in.UserId}, &userroles, "roles")
		//if err != nil || !ok {
		//
		//	svc.log.Info(&analysisclient.LogContent{
		//		Headers:   headers,
		//		Action:    "FindUserRoleRelation",
		//		Message:   "find roles by userId err.",
		//		StateCode: errstate.ErrRequest.Code,
		//		Fields: &analysisclient.LogFields{
		//			"userId": in.UserId,
		//			"funcS":  in.StructureId,
		//		},
		//	})
		//
		//	return errstate.ErrRequest
		//}
		//
		//userRoles := userroles["roles"].([]interface{})
		//
		//if userRoles != nil && len(userRoles) > 0 && len(in.FunctionRoles) > 0 {
		//	roles := make(map[string]string)
		//	ok := false
		//	for _, v := range userRoles {
		//		roles[v.(string)] = "ok"
		//	}
		//	for _, v := range in.FunctionRoles {
		//		if roles[v] == "ok" {
		//			ok = true
		//			//这里还没有做角色认证, 需要判断角色是否有效
		//
		//			break
		//		}
		//	}
		//
		//	if ok {
		//		svc.log.Info(&analysisclient.LogContent{
		//			Headers:   headers,
		//			Action:    "RoleGrant",
		//			StateCode: errstate.ErrRequest.Code,
		//			Fields: &analysisclient.LogFields{
		//				"userId": in.UserId,
		//				"funcS":  in.StructureId,
		//			},
		//		})
		//		return errstate.Success
		//	}
		//}
		//
		//svc.log.Info(&analysisclient.LogContent{
		//	Headers:   headers,
		//	Action:    "RoleUnauthorized",
		//	StateCode: errstate.ErrRequest.Code,
		//	Fields: &analysisclient.LogFields{
		//		"userId": in.UserId,
		//		"funcS":  in.StructureId,
		//	},
		//})

		c, err := svc.Client.Count(getURFIndex(in.UserId), map[string]interface{}{"userId": in.UserId, "functionId": in.FunctionId})
		if err != nil {
			return errstate.ErrSystem
		}
		if c == 1 {
			return errstate.Success
		}
		return errstate.ErrUserPermission
	})
}

func NewAccessibleService(c *indexutils.Client, log analysisclient.LogClient) inner.AccessibleHandler {
	return &accessibleService{Client: c, log: log}
}
