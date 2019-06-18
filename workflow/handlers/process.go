package handlers

import (
	"context"
	"konekko.me/gosion/workflow/modules"
	"konekko.me/gosion/workflow/pb"
)

type processService struct {
	modules modules.Modules
}

func (svc *processService) Launch(context.Context, *gosionsvc_external_workflow.LaunchRequest, *gosionsvc_external_workflow.LaunchResponse) error {
	panic("implement me")
}

func (svc *processService) Create(context.Context, *gosionsvc_external_workflow.CreateRequest, *gosionsvc_external_workflow.CreateResponse) error {
	panic("implement me")
}

func (svc *processService) Build(context.Context, *gosionsvc_external_workflow.BuildRequest, *gosionsvc_external_workflow.BuildResponse) error {
	panic("implement me")
}

func (svc *processService) Delete(context.Context, *gosionsvc_external_workflow.DeleteRequest, *gosionsvc_external_workflow.DeleteResponse) error {
	panic("implement me")
}

func (svc *processService) Update(context.Context, *gosionsvc_external_workflow.UpdateRequest, *gosionsvc_external_workflow.UpdateResponse) error {
	panic("implement me")
}

func (svc *processService) Open(context.Context, *gosionsvc_external_workflow.OpenRequest, *gosionsvc_external_workflow.OpenResponse) error {
	panic("implement me")
}

func (svc *processService) Detail(context.Context, *gosionsvc_external_workflow.DetailRequest, *gosionsvc_external_workflow.DetailResponse) error {
	panic("implement me")
}

func (svc *processService) GetImage(context.Context, *gosionsvc_external_workflow.GetImageRequest, *gosionsvc_external_workflow.GetImageResponse) error {
	panic("implement me")
}

func (svc *processService) Search(context.Context, *gosionsvc_external_workflow.SearchRequest, *gosionsvc_external_workflow.SearchResponse) error {
	panic("implement me")
}

func NewProcessService() gosionsvc_external_workflow.ProcessHandler {
	return &processService{}
}
