package workflow

import (
	"context"
	"konekko.me/gosion/commons/dto"
	"konekko.me/gosion/commons/wrapper"
	"konekko.me/gosion/workflow/pb"
)

type processService struct {
}

func (svc *processService) Launch(ctx context.Context, in *gs_service_workflow.LaunchRequest, out *gs_service_workflow.LaunchResponse) error {
	return gs_commons_wrapper.ContextToAuthorize(ctx, out, func(auth *gs_commons_wrapper.WrapperUser) *gs_commons_dto.State {
		return nil
	})
}

func (svc *processService) Create(ctx context.Context, in *gs_service_workflow.CreateRequest, out *gs_service_workflow.CreateResponse) error {
	return gs_commons_wrapper.ContextToAuthorize(ctx, out, func(auth *gs_commons_wrapper.WrapperUser) *gs_commons_dto.State {
		return nil
	})
}

func (svc *processService) Build(ctx context.Context, in *gs_service_workflow.BuildRequest, out *gs_service_workflow.BuildResponse) error {
	return gs_commons_wrapper.ContextToAuthorize(ctx, out, func(auth *gs_commons_wrapper.WrapperUser) *gs_commons_dto.State {
		return nil
	})
}

func (svc *processService) Delete(ctx context.Context, in *gs_service_workflow.DeleteRequest, out *gs_service_workflow.DeleteResponse) error {
	return gs_commons_wrapper.ContextToAuthorize(ctx, out, func(auth *gs_commons_wrapper.WrapperUser) *gs_commons_dto.State {
		return nil
	})
}

func (svc *processService) Update(ctx context.Context, in *gs_service_workflow.UpdateRequest, out *gs_service_workflow.UpdateResponse) error {
	return gs_commons_wrapper.ContextToAuthorize(ctx, out, func(auth *gs_commons_wrapper.WrapperUser) *gs_commons_dto.State {
		return nil
	})
}

func (svc *processService) Open(ctx context.Context, in *gs_service_workflow.OpenRequest, out *gs_service_workflow.OpenResponse) error {
	return gs_commons_wrapper.ContextToAuthorize(ctx, out, func(auth *gs_commons_wrapper.WrapperUser) *gs_commons_dto.State {
		return nil
	})
}

func (svc *processService) Detail(ctx context.Context, in *gs_service_workflow.DetailRequest, out *gs_service_workflow.DetailResponse) error {
	return gs_commons_wrapper.ContextToAuthorize(ctx, out, func(auth *gs_commons_wrapper.WrapperUser) *gs_commons_dto.State {
		return nil
	})
}

func (svc *processService) GetImage(ctx context.Context, in *gs_service_workflow.GetImageRequest, out *gs_service_workflow.GetImageResponse) error {
	return gs_commons_wrapper.ContextToAuthorize(ctx, out, func(auth *gs_commons_wrapper.WrapperUser) *gs_commons_dto.State {
		return nil
	})
}

func newProcessServcie() gs_service_workflow.ProcessHandler {
	return &processService{}
}
