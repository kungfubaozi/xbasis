package handlers

import (
	"context"
	"konekko.me/gosion/workflow/modules"
	"konekko.me/gosion/workflow/pb"
)

type formService struct {
	modules modules.Modules
}

func (svc *formService) CreatePlaceholder(context.Context, *gosionsvc_external_workflow.CreatePlaceholderRequest, *gosionsvc_external_workflow.CreatePlaceholderResponse) error {
	panic("implement me")
}

func (svc *formService) DeletePlaceholder(context.Context, *gosionsvc_external_workflow.DeletePlaceholderRequest, *gosionsvc_external_workflow.DeletePlaceholderResponse) error {
	panic("implement me")
}

func (svc *formService) UpdatePlaceholder(context.Context, *gosionsvc_external_workflow.UpdatePlaceholderRequest, *gosionsvc_external_workflow.UpdatePlaceholderResponse) error {
	panic("implement me")
}

func (svc *formService) AddField(context.Context, *gosionsvc_external_workflow.AddFieldRequest, *gosionsvc_external_workflow.AddFieldResponse) error {
	panic("implement me")
}

func (svc *formService) RemoveField(context.Context, *gosionsvc_external_workflow.RemoveFieldRequest, *gosionsvc_external_workflow.RemoveFieldResponse) error {
	panic("implement me")
}

func (svc *formService) UpdateFieldProps(context.Context, *gosionsvc_external_workflow.UpdateFieldPropsRequest, *gosionsvc_external_workflow.UpdateFieldPropsResponse) error {
	panic("implement me")
}

func (svc *formService) GetAllTypeFields(context.Context, *gosionsvc_external_workflow.GetAllTypeFieldsRequest, *gosionsvc_external_workflow.GetAllTypeFieldsResponse) error {
	panic("implement me")
}

func (svc *formService) CheckFiledValue(context.Context, *gosionsvc_external_workflow.CheckFiledValueRequest, *gosionsvc_external_workflow.CheckFieldValueResponse) error {
	panic("implement me")
}

func (svc *formService) Search(context.Context, *gosionsvc_external_workflow.SearchRequest, *gosionsvc_external_workflow.SearchResponse) error {
	panic("implement me")
}

func NewFormService() gosionsvc_external_workflow.FormHandler {
	return &formService{}
}
