package workflowhandlers

import (
	"context"
	"konekko.me/xbasis/workflow/modules"
	pb "konekko.me/xbasis/workflow/pb"
)

type formService struct {
	modules modules.Modules
}

func (svc *formService) CreatePlaceholder(context.Context, *pb.CreatePlaceholderRequest, *pb.CreatePlaceholderResponse) error {
	panic("implement me")
}

func (svc *formService) DeletePlaceholder(context.Context, *pb.DeletePlaceholderRequest, *pb.DeletePlaceholderResponse) error {
	panic("implement me")
}

func (svc *formService) UpdatePlaceholder(context.Context, *pb.UpdatePlaceholderRequest, *pb.UpdatePlaceholderResponse) error {
	panic("implement me")
}

func (svc *formService) AddField(context.Context, *pb.AddFieldRequest, *pb.AddFieldResponse) error {
	panic("implement me")
}

func (svc *formService) RemoveField(context.Context, *pb.RemoveFieldRequest, *pb.RemoveFieldResponse) error {
	panic("implement me")
}

func (svc *formService) UpdateFieldProps(context.Context, *pb.UpdateFieldPropsRequest, *pb.UpdateFieldPropsResponse) error {
	panic("implement me")
}

func (svc *formService) GetAllTypeFields(context.Context, *pb.GetAllTypeFieldsRequest, *pb.GetAllTypeFieldsResponse) error {
	panic("implement me")
}

func (svc *formService) CheckFiledValue(context.Context, *pb.CheckFiledValueRequest, *pb.CheckFieldValueResponse) error {
	panic("implement me")
}

func (svc *formService) Search(context.Context, *pb.SearchFormRequest, *pb.SearchFormResponse) error {
	panic("implement me")
}

func NewFormService() pb.FormHandler {
	return &formService{}
}
