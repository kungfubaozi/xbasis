package workflowhandlers

import (
	"context"
	"konekko.me/gosion/workflow/modules"
	"konekko.me/gosion/workflow/pb"
)

type instanceService struct {
	modules modules.Modules
}

func (svc *instanceService) GetMyLaunchInstances(context.Context, *gosionsvc_external_workflow.GetMyLaunchInstancesRequest, *gosionsvc_external_workflow.GetMyLaunchInstancesResponse) error {
	panic("implement me")
}

func (svc *instanceService) GetAllInstances(context.Context, *gosionsvc_external_workflow.GetAllInstancesRequest, *gosionsvc_external_workflow.GetAllInstancesResponse) error {
	panic("implement me")
}

func (svc *instanceService) Stop(context.Context, *gosionsvc_external_workflow.StopRequest, *gosionsvc_external_workflow.StopResponse) error {
	panic("implement me")
}

func (svc *instanceService) Restart(context.Context, *gosionsvc_external_workflow.RestartRequest, *gosionsvc_external_workflow.RestartResponse) error {
	panic("implement me")
}

func (svc *instanceService) Continue(context.Context, *gosionsvc_external_workflow.ContinueRequest, *gosionsvc_external_workflow.ContinueResponse) error {
	panic("implement me")
}

func (svc *instanceService) Submit(context.Context, *gosionsvc_external_workflow.SubmitRequest, *gosionsvc_external_workflow.SubmitResponse) error {
	panic("implement me")
}

func (svc *instanceService) Search(context.Context, *gosionsvc_external_workflow.SearchInstanceRequest, *gosionsvc_external_workflow.SearchInstanceResponse) error {
	panic("implement me")
}

func NewInstanceService() gosionsvc_external_workflow.InstanceHandler {
	return &instanceService{}
}
