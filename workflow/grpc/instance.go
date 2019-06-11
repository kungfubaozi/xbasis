package grpc

import (
	"context"
	"konekko.me/gosion/workflow/modules"
	"konekko.me/gosion/workflow/pb"
)

type instanceService struct {
	modules modules.Modules
}

//获取我新建的流程
func (svc *instanceService) GetMyLaunchInstances(context.Context, *gs_service_workflow.GetMyLaunchInstancesRequest, *gs_service_workflow.GetMyLaunchInstancesResponse) error {
	panic("implement me")
}

//获取所有的流程
func (svc *instanceService) GetAllInstances(context.Context, *gs_service_workflow.GetAllInstancesRequest, *gs_service_workflow.GetAllInstancesResponse) error {
	panic("implement me")
}

//停止流程
func (svc *instanceService) Stop(context.Context, *gs_service_workflow.StopRequest, *gs_service_workflow.StopResponse) error {
	panic("implement me")
}

//重新启动
func (svc *instanceService) Restart(context.Context, *gs_service_workflow.RestartRequest, *gs_service_workflow.RestartResponse) error {
	panic("implement me")
}

//继续
func (svc *instanceService) Continue(context.Context, *gs_service_workflow.ContinueRequest, *gs_service_workflow.ContinueResponse) error {
	panic("implement me")
}

//提交
func (svc *instanceService) Submit(context.Context, *gs_service_workflow.SubmitRequest, *gs_service_workflow.SubmitResponse) error {
	panic("implement me")
}

func NewInstanceService(modules modules.Modules) gs_service_workflow.InstanceHandler {
	return &instanceService{modules: modules}
}
