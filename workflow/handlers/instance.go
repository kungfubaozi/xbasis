package workflowhandlers

import (
	"context"
	"konekko.me/xbasis/workflow/modules"
	pb "konekko.me/xbasis/workflow/pb"
)

type instanceService struct {
	modules modules.Modules
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

func (svc *instanceService) Submit(context.Context, *pb.SubmitRequest, *pb.SubmitResponse) error {
	panic("implement me")
}

func (svc *instanceService) Search(context.Context, *pb.SearchInstanceRequest, *pb.SearchInstanceResponse) error {
	panic("implement me")
}

func NewInstanceService() pb.InstanceHandler {
	return &instanceService{}
}
