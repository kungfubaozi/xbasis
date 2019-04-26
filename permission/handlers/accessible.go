package permissionhandlers

import (
	"context"
	"github.com/Sirupsen/logrus"
	"konekko.me/gosion/commons/dto"
	"konekko.me/gosion/commons/errstate"
	"konekko.me/gosion/commons/gslogrus"
	"konekko.me/gosion/commons/indexutils"
	"konekko.me/gosion/commons/wrapper"
	"konekko.me/gosion/permission/pb/ext"
)

type accessibleService struct {
	*indexutils.Client
	*gslogrus.Logger
}

func (svc *accessibleService) Check(ctx context.Context, in *gs_ext_service_permission.CheckRequest, out *gs_commons_dto.Status) error {
	return gs_commons_wrapper.ContextToAuthorize(ctx, out, func(auth *gs_commons_wrapper.WrapperUser) *gs_commons_dto.State {
		//get user require roles
		var userroles map[string]interface{}

		log := svc.WithHeaders(auth.TraceId, auth.ClientId, auth.IP, "", auth.UserAgent, auth.UserDevice)

		ok, err := svc.Client.QueryFirst("gs-user-roles-relation",
			map[string]interface{}{"structure_id": in.StructureId, "user_id": in.UserId}, &userroles, "roles")
		if err != nil || !ok {

			log.WithAction("RelationFind", logrus.Fields{
				"err":               err,
				"state":             errstate.ErrRequest.Code,
				"userId":            in.UserId,
				"functionStructure": in.StructureId,
			}).Error("find roles by userId err.")

			return errstate.ErrRequest
		}

		userRoles := userroles["roles"].([]string)

		if userRoles != nil && len(userRoles) > 0 && len(in.FunctionRoles) > 0 {
			roles := make(map[string]string)
			ok := false
			for _, v := range userRoles {
				roles[v] = "ok"
			}
			for _, v := range in.FunctionRoles {
				if roles[v] == "ok" {
					ok = true
					break
				}
			}

			if ok {
				log.WithAction("RoleInclusion", logrus.Fields{
					"state": errstate.Success.Code,
				}).Info("grant")
				return errstate.Success
			}
		}

		log.WithAction("RoleInclusion", logrus.Fields{
			"state": errstate.ErrUserPermission.Code,
		}).Info("unauthorized")

		return errstate.ErrUserPermission
	})
}

func NewAccessibleService(c *indexutils.Client, log *gslogrus.Logger) gs_ext_service_permission.AccessibleHandler {
	return &accessibleService{Client: c, Logger: log}
}
