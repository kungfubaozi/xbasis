package workflow

import (
	"context"
	"konekko.me/gosion/commons/dto"
	"konekko.me/gosion/commons/wrapper"
	"konekko.me/gosion/workflow/pb"
)

type formService struct {
}

func (svc *formService) CreatePlaceholder(ctx context.Context, in *gs_service_workflow.CreatePlaceholderRequest, out *gs_service_workflow.CreatePlaceholderResponse) error {
	return gs_commons_wrapper.ContextToAuthorize(ctx, out, func(auth *gs_commons_wrapper.WrapperUser) *gs_commons_dto.State {
		return nil
	})
}

func (svc *formService) DeletePlaceholder(ctx context.Context, in *gs_service_workflow.DeletePlaceholderRequest, out *gs_service_workflow.DeletePlaceholderResponse) error {
	return gs_commons_wrapper.ContextToAuthorize(ctx, out, func(auth *gs_commons_wrapper.WrapperUser) *gs_commons_dto.State {
		return nil
	})
}

func (svc *formService) UpdatePlaceholder(ctx context.Context, in *gs_service_workflow.UpdatePlaceholderRequest, out *gs_service_workflow.UpdatePlaceholderResponse) error {
	return gs_commons_wrapper.ContextToAuthorize(ctx, out, func(auth *gs_commons_wrapper.WrapperUser) *gs_commons_dto.State {
		return nil
	})
}

func (svc *formService) AddField(ctx context.Context, in *gs_service_workflow.AddFieldRequest, out *gs_service_workflow.AddFieldResponse) error {
	return gs_commons_wrapper.ContextToAuthorize(ctx, out, func(auth *gs_commons_wrapper.WrapperUser) *gs_commons_dto.State {
		return nil
	})
}

func (svc *formService) RemoveField(ctx context.Context, in *gs_service_workflow.RemoveFieldRequest, out *gs_service_workflow.RemoveFieldResponse) error {
	return gs_commons_wrapper.ContextToAuthorize(ctx, out, func(auth *gs_commons_wrapper.WrapperUser) *gs_commons_dto.State {
		return nil
	})
}

func (svc *formService) UpdateFieldProps(ctx context.Context, in *gs_service_workflow.UpdateFieldPropsRequest, out *gs_service_workflow.UpdateFieldPropsResponse) error {
	return gs_commons_wrapper.ContextToAuthorize(ctx, out, func(auth *gs_commons_wrapper.WrapperUser) *gs_commons_dto.State {
		return nil
	})
}

func (svc *formService) GetAllTypeFields(ctx context.Context, in *gs_service_workflow.GetAllTypeFieldsRequest, out *gs_service_workflow.GetAllTypeFieldsResponse) error {
	return gs_commons_wrapper.ContextToAuthorize(ctx, out, func(auth *gs_commons_wrapper.WrapperUser) *gs_commons_dto.State {
		return nil
	})
}

func newFormService() gs_service_workflow.FormHandler {
	return &formService{}
}
