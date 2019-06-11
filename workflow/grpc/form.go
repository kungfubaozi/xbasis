package grpc

import (
	"context"
	"konekko.me/gosion/workflow/modules"
	"konekko.me/gosion/workflow/pb"
)

type formService struct {
	modules modules.Modules
}

func (svc *formService) CheckFiledValue(context.Context, *gs_service_workflow.CheckFiledValueRequest, *gs_service_workflow.CheckFieldValueResponse) error {
	panic("implement me")
}

func (svc *formService) CreatePlaceholder(context.Context, *gs_service_workflow.CreatePlaceholderRequest, *gs_service_workflow.CreatePlaceholderResponse) error {
	panic("implement me")
}

func (svc *formService) DeletePlaceholder(context.Context, *gs_service_workflow.DeletePlaceholderRequest, *gs_service_workflow.DeletePlaceholderResponse) error {
	panic("implement me")
}

func (svc *formService) UpdatePlaceholder(context.Context, *gs_service_workflow.UpdatePlaceholderRequest, *gs_service_workflow.UpdatePlaceholderResponse) error {
	panic("implement me")
}

func (svc *formService) AddField(context.Context, *gs_service_workflow.AddFieldRequest, *gs_service_workflow.AddFieldResponse) error {
	panic("implement me")
}

func (svc *formService) RemoveField(context.Context, *gs_service_workflow.RemoveFieldRequest, *gs_service_workflow.RemoveFieldResponse) error {
	panic("implement me")
}

func (svc *formService) UpdateFieldProps(context.Context, *gs_service_workflow.UpdateFieldPropsRequest, *gs_service_workflow.UpdateFieldPropsResponse) error {
	panic("implement me")
}

func (svc *formService) GetAllTypeFields(context.Context, *gs_service_workflow.GetAllTypeFieldsRequest, *gs_service_workflow.GetAllTypeFieldsResponse) error {
	panic("implement me")
}

func NewFormService(modules modules.Modules) gs_service_workflow.FormHandler {
	return &formService{modules: modules}
}
