package workflowhandlers

import (
	"context"
	"encoding/json"
	"gopkg.in/mgo.v2"
	"konekko.me/xbasis/analysis/client"
	date "konekko.me/xbasis/commons/date"
	commons "konekko.me/xbasis/commons/dto"
	"konekko.me/xbasis/commons/errstate"
	generator "konekko.me/xbasis/commons/generator"
	wrapper "konekko.me/xbasis/commons/wrapper"
	"konekko.me/xbasis/workflow/models"
	"konekko.me/xbasis/workflow/modules"
	pb "konekko.me/xbasis/workflow/pb"
	"time"
)

type processService struct {
	modules modules.Modules
	id      generator.IDGenerator
	session *mgo.Session
	log     analysisclient.LogClient
}

func (svc *processService) Launch(context.Context, *pb.LaunchRequest, *pb.LaunchResponse) error {
	panic("implement me")
}

func (svc *processService) Build(context.Context, *pb.BuildRequest, *pb.BuildResponse) error {
	panic("implement me")
}

func (svc *processService) Delete(ctx context.Context, in *pb.DeleteProcessRequest, out *pb.DeleteProcessResponse) error {
	panic("implement me")
}

//update只是更新当前的流程设置，并不直接关系到具体流程
func (svc *processService) Update(ctx context.Context, in *pb.UpdateProcessRequest, out *pb.UpdateProcessResponse) error {
	return wrapper.ContextToAuthorize(ctx, out, func(auth *wrapper.WrapperUser) *commons.State {
		if len(in.AppId) < 8 {
			return errstate.ErrRequest
		}
		processId := in.ProcessId
		if len(in.ProcessId) < 16 {
			processId = svc.id.Get()
		}
		if len(in.NodeDataArray) > 10 {

			header := &analysisclient.LogHeaders{
				TraceId: svc.id.Get(),
			}

			var nodeDataArray []interface{}
			var linkDataArray []interface{}
			err := json.Unmarshal([]byte(in.NodeDataArray), &nodeDataArray)
			if err != nil {
				return errstate.ErrRequest
			}
			err = json.Unmarshal([]byte(in.LinkDataArray), &linkDataArray)
			if err != nil {
				return errstate.ErrRequest
			}

			p := len(in.ProcessId) >= 16

			var array *models.FlowDataArray

			if p {
				a, err := svc.modules.Process().GetFlowDataArray(processId)
				if err != nil {
					return nil
				}
				array = a
				if a.AppId != in.AppId {
					return errstate.ErrRequest
				}
			}

			if array == nil {
				array = &models.FlowDataArray{
					ProcessId: processId,
					Name:      in.Name,
					Desc:      in.Desc,
					Image:     in.Image,
					AppId:     in.AppId,
					CreateAt:  time.Now().Unix(),
				}
			}

			if len(in.Name) > 10 {
				array.Name = in.Name
			}

			if len(in.Desc) > 0 {
				array.Desc = in.Desc
			}

			if len(in.Image) > 100 {
				err = svc.modules.Process().SaveImage(processId, in.Image)
				if err != nil {
					return nil
				}
			}

			array.LinkDataArray = linkDataArray
			array.NodeDataArray = nodeDataArray

			err = svc.modules.Process().SaveFlowDataArrays(array)
			if err != nil {
				return errstate.ErrRequest
			}

			//create index
			svc.log.Info(&analysisclient.LogContent{
				Headers: header,
				Action:  "SaveFlowDataArrays",
				Index: &analysisclient.LogIndex{
					Name: "flowarrays",
					Id:   processId,
					Fields: &analysisclient.LogFields{
						"process_id": processId,
						"app_id":     array.AppId,
						"name":       array.Name,
						"desc":       array.Desc,
						"update_at":  date.FormatDate(time.Now(), date.YYYY_MM_DD_HH_MM_SS),
					},
				},
			})

			out.ProcessId = processId

			return errstate.Success
		}
		return errstate.ErrRequest
	})
}

func (svc *processService) Open(context.Context, *pb.OpenRequest, *pb.OpenResponse) error {
	panic("implement me")
}

func (svc *processService) Detail(ctx context.Context, in *pb.DetailProcessRequest, out *pb.DetailProcessResponse) error {
	return wrapper.ContextToAuthorize(ctx, out, func(auth *wrapper.WrapperUser) *commons.State {

		f, err := svc.modules.Process().GetFlowDataArray(in.ProcessId)
		if err != nil {
			return errstate.ErrRequest
		}

		nda, err := json.Marshal(f.NodeDataArray)
		if err != nil {
			return errstate.ErrRequest
		}

		lda, err := json.Marshal(f.LinkDataArray)
		if err != nil {
			return errstate.ErrRequest
		}

		out.NodeDataArray = string(nda)
		out.LinkDataArray = string(lda)
		out.Name = f.Name
		out.Desc = f.Desc

		return errstate.Success
	})
}

func (svc *processService) GetImage(ctx context.Context, in *pb.GetImageRequest, out *pb.GetImageResponse) error {
	return wrapper.ContextToAuthorize(ctx, out, func(auth *wrapper.WrapperUser) *commons.State {

		a, err := svc.modules.Process().GetImage(in.ProcessId)
		if err != nil {
			return nil
		}

		out.Image = a

		return errstate.Success
	})
}

func (svc *processService) Search(ctx context.Context, in *pb.SearchProcessRequest, out *pb.SearchProcessResponse) error {
	return wrapper.ContextToAuthorize(ctx, out, func(auth *wrapper.WrapperUser) *commons.State {

		items, err := svc.modules.Process().Search(in.AppId, in.Name, in.Page, in.Size)
		if err != nil {
			return errstate.ErrRequest
		}

		var data []*pb.SearchProcessItem
		for _, v := range items {
			data = append(data, &pb.SearchProcessItem{
				AppId:     v.AppId,
				Name:      v.Name,
				Desc:      v.Desc,
				ProcessId: v.ProcessId,
				UpdateAt:  v.UpdateAt,
			})
		}

		out.Data = data

		return errstate.Success
	})
}

func NewProcessService(modules modules.Modules,
	id generator.IDGenerator, log analysisclient.LogClient) pb.ProcessHandler {
	return &processService{modules: modules, id: id, log: log}
}
