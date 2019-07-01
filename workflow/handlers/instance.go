package workflowhandlers

import (
	"context"
	commons "konekko.me/xbasis/commons/dto"
	wrapper "konekko.me/xbasis/commons/wrapper"
	workflow "konekko.me/xbasis/workflow/modules"
	pb "konekko.me/xbasis/workflow/pb"
)

type instanceService struct {
	modules workflow.Modules
}

func (svc *instanceService) GetMyLaunchInstances(context.Context, *pb.GetMyLaunchInstancesRequest, *pb.GetMyLaunchInstancesResponse) error {
	panic("implement me")
}

func (svc *instanceService) GetAllInstances(context.Context, *pb.GetAllInstancesRequest, *pb.GetAllInstancesResponse) error {
	panic("implement me")
}

func (svc *instanceService) Stop(context.Context, *pb.StopRequest, *pb.StopResponse) error {
	panic("implement me")
}

func (svc *instanceService) Restart(context.Context, *pb.RestartRequest, *pb.RestartResponse) error {
	panic("implement me")
}

func (svc *instanceService) Continue(context.Context, *pb.ContinueRequest, *pb.ContinueResponse) error {
	panic("implement me")
}

func (svc *instanceService) Submit(ctx context.Context, in *pb.SubmitRequest, out *pb.SubmitResponse) error {
	return wrapper.ContextToAuthorize(ctx, out, func(auth *wrapper.WrapperUser) *commons.State {

		return nil
	})
}

func (svc *instanceService) Search(context.Context, *pb.SearchInstanceRequest, *pb.SearchInstanceResponse) error {
	panic("implement me")
}

func NewInstanceService() pb.InstanceHandler {
	return &instanceService{}
}
