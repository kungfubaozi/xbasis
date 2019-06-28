package permissionhandlers

import (
	"context"
	"fmt"
	"gopkg.in/mgo.v2"
	"konekko.me/xbasis/analysis/client"
	commons "konekko.me/xbasis/commons/dto"
	"konekko.me/xbasis/commons/errstate"
	"konekko.me/xbasis/commons/indexutils"
	"konekko.me/xbasis/commons/wrapper"
	inner "konekko.me/xbasis/permission/pb/inner"
)

type accessibleService struct {
	*indexutils.Client
	log     analysisclient.LogClient
	session *mgo.Session
}

func (svc *accessibleService) GetRepo() *bindingRepo {
	return &bindingRepo{session: svc.session.Clone(), Client: svc.Client}
}

func (svc *accessibleService) HasGrant(ctx context.Context, in *inner.HasGrantRequest, out *commons.Status) error {
	return xbasiswrapper.ContextToAuthorize(ctx, out, func(auth *xbasiswrapper.WrapperUser) *commons.State {

		repo := svc.GetRepo()
		defer repo.Close()

		r, err := repo.FindRelationUserById(in.UserId, in.AppId)
		if err != nil {
			return nil
		}

		if len(in.Role) > 10 {
			for _, v := range r.Roles {
				if v == in.Role {
					return errstate.ErrUserAlreadyBindRole
				}
			}
		}

		return errstate.Success
	})
}

func (svc *accessibleService) Check(ctx context.Context, in *inner.CheckRequest, out *commons.Status) error {
	return xbasiswrapper.ContextToAuthorize(ctx, out, func(auth *xbasiswrapper.WrapperUser) *commons.State {
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
			fmt.Println("err 1")
			return errstate.ErrSystem
		}
		if c == 1 {
			fmt.Println("err 2")
			return errstate.Success
		}
		fmt.Println("err 3")
		return errstate.ErrUserPermission
	})
}

func NewAccessibleService(c *indexutils.Client, session *mgo.Session, log analysisclient.LogClient) inner.AccessibleHandler {
	return &accessibleService{Client: c, log: log, session: session}
}
