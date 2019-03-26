package safety_handers

import (
	"context"
	"gopkg.in/mgo.v2"
	"konekko.me/gosion/commons/dto"
	"konekko.me/gosion/commons/wrapper"
	"konekko.me/gosion/safety/pb/nops"
)

type securityService struct {
	session *mgo.Session
}

//locking
//different places
func (svc *securityService) Get(ctx context.Context, in *gs_nops_service_safety.GetRequest, out *gs_nops_service_safety.GetResponse) error {
	return gs_commons_wrapper.ContextToAuthorize(ctx, out, func(auth *gs_commons_wrapper.WrapperUser) *gs_commons_dto.State {

		return nil
	})
}

func NewSecurityService(session *mgo.Session) gs_nops_service_safety.SecurityHandler {
	return &securityService{session: session}
}
