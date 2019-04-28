package workflow

import (
	"context"
	"konekko.me/gosion/commons/dto"
	"konekko.me/gosion/commons/wrapper"
	"konekko.me/gosion/workflow/pb"
)

type instanceService struct {
	runtime *runtime
}

func (svc *instanceService) Submit(ctx context.Context, in *gs_service_workflow.SubmitRequest, out *gs_service_workflow.SubmitResponse) error {
	return gs_commons_wrapper.ContextToAuthorize(ctx, out, func(auth *gs_commons_wrapper.WrapperUser) *gs_commons_dto.State {
		if len(in.InstanceId) > 0 && in.Data != nil && len(in.Data) > 0 {
			//check instance if exists and get stepper task

			//get map data

		}
		return nil
	})
}

func (svc *instanceService) GetMyLaunchInstances(ctx context.Context, in *gs_service_workflow.GetMyLaunchInstancesRequest, out *gs_service_workflow.GetMyLaunchInstancesResponse) error {
	return gs_commons_wrapper.ContextToAuthorize(ctx, out, func(auth *gs_commons_wrapper.WrapperUser) *gs_commons_dto.State {
		return nil
	})
}

func (svc *instanceService) GetAllInstances(ctx context.Context, in *gs_service_workflow.GetAllInstancesRequest, out *gs_service_workflow.GetAllInstancesResponse) error {
	return gs_commons_wrapper.ContextToAuthorize(ctx, out, func(auth *gs_commons_wrapper.WrapperUser) *gs_commons_dto.State {
		return nil
	})
}

func (svc *instanceService) Stop(ctx context.Context, in *gs_service_workflow.StopRequest, out *gs_service_workflow.StopResponse) error {
	return gs_commons_wrapper.ContextToAuthorize(ctx, out, func(auth *gs_commons_wrapper.WrapperUser) *gs_commons_dto.State {
		return nil
	})
}

func (svc *instanceService) Restart(ctx context.Context, in *gs_service_workflow.RestartRequest, out *gs_service_workflow.RestartResponse) error {
	return gs_commons_wrapper.ContextToAuthorize(ctx, out, func(auth *gs_commons_wrapper.WrapperUser) *gs_commons_dto.State {
		return nil
	})
}

func (svc *instanceService) Continue(ctx context.Context, in *gs_service_workflow.ContinueRequest, out *gs_service_workflow.ContinueResponse) error {
	return gs_commons_wrapper.ContextToAuthorize(ctx, out, func(auth *gs_commons_wrapper.WrapperUser) *gs_commons_dto.State {
		return nil
	})
}

func newInstanceService() gs_service_workflow.InstanceHandler {
	return &instanceService{}
}
