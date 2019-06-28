package workflowhandlers

import (
	"context"
	"gopkg.in/mgo.v2"
	"konekko.me/xbasis/analysis/client"
	commons "konekko.me/xbasis/commons/dto"
	generator "konekko.me/xbasis/commons/generator"
	wrapper "konekko.me/xbasis/commons/wrapper"
	"konekko.me/xbasis/workflow/modules"
	pb "konekko.me/xbasis/workflow/pb"
)

type formService struct {
	modules modules.Modules
	id      generator.IDGenerator
	session *mgo.Session
	log     analysisclient.LogClient
}

func (svc *formService) Update(ctx context.Context, in *pb.UpdateRequest, out *pb.UpdateResponse) error {
	return wrapper.ContextToAuthorize(ctx, out, func(auth *wrapper.WrapperUser) *commons.State {
		return nil
	})
}

func (svc *formService) Delete(ctx context.Context, in *pb.DeleteRequest, out *pb.DeleteResponse) error {
	return wrapper.ContextToAuthorize(ctx, out, func(auth *wrapper.WrapperUser) *commons.State {
		return nil
	})
}

func (svc *formService) GetAllTypeFields(ctx context.Context, in *pb.GetAllTypeFieldsRequest, out *pb.GetAllTypeFieldsResponse) error {
	return wrapper.ContextToAuthorize(ctx, out, func(auth *wrapper.WrapperUser) *commons.State {
		return nil
	})
}

func (svc *formService) CheckFiledValue(ctx context.Context, in *pb.CheckFiledValueRequest, out *pb.CheckFieldValueResponse) error {
	return wrapper.ContextToAuthorize(ctx, out, func(auth *wrapper.WrapperUser) *commons.State {
		return nil
	})
}

func (svc *formService) Search(ctx context.Context, in *pb.SearchFormRequest, out *pb.SearchFormResponse) error {
	return wrapper.ContextToAuthorize(ctx, out, func(auth *wrapper.WrapperUser) *commons.State {
		return nil
	})
}

func (svc *formService) Detail(ctx context.Context, in *pb.DetailRequest, out *pb.DetailResponse) error {
	return wrapper.ContextToAuthorize(ctx, out, func(auth *wrapper.WrapperUser) *commons.State {
		return nil
	})
}

func NewFormService(modules modules.Modules,
	id generator.IDGenerator, log analysisclient.LogClient) pb.FormHandler {
	return &formService{modules: modules, id: id, log: log}
}
