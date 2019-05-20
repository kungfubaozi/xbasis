package runtime

import (
	"konekko.me/gosion/commons/generator"
	"konekko.me/gosion/workflow/distribute"
	"konekko.me/gosion/workflow/models"
	"konekko.me/gosion/workflow/types"
	"testing"
)

func TestProcesses_AddProcess(t *testing.T) {

	id := gs_commons_generator.NewIDG()

	p := &processes{
		id:       id,
		relation: distribute.NewRelation(),
	}

	startEventId := id.Get()
	endEventId := id.Get()

	ut2 := id.Get()
	ut3 := id.Get()
	ut4 := id.Get()
	ut5 := id.Get()
	ut6 := id.Get()

	ig1 := id.Get()
	ig2 := id.Get()
	ig3 := id.Get()
	ig4 := id.Get()

	p.AddProcess(&models.Process{
		Id:   id.Get(),
		Name: "TestProcess",
		Gateways: &models.Gateways{
			Inclusive: []*models.InclusiveGateway{
				{
					Info: &models.Info{
						Id:   ig1,
						Key:  "ig1",
						Name: "ig1",
					},
				},
				{
					Info: &models.Info{
						Id:   ig2,
						Key:  "ig2",
						Name: "ig2",
					},
				},
				{
					Info: &models.Info{
						Id:   ig3,
						Key:  "ig3",
						Name: "ig3",
					},
				},
				{
					Info: &models.Info{
						Id:   ig4,
						Key:  "ig4",
						Name: "ig4",
					},
				},
			},
		},
		Tasks: &models.Tasks{
			UserTasks: []*models.UserTask{
				{
					Info: &models.Info{
						Id:   ut2,
						Key:  "ut2",
						Name: "ut2",
					},
				},
				{
					Info: &models.Info{
						Id:   ut3,
						Key:  "ut3",
						Name: "ut3",
					},
				},
				{
					Info: &models.Info{
						Id:   ut4,
						Key:  "ut4",
						Name: "ut4",
					},
				},
				{
					Info: &models.Info{
						Id:   ut5,
						Key:  "ut5",
						Name: "ut5",
					},
				},
				{
					Info: &models.Info{
						Id:   ut6,
						Key:  "ut6",
						Name: "ut6",
					},
				},
			},
		},
		StartEvent: &models.TypeEvent{
			Id:   startEventId,
			Key:  "start",
			Type: types.CTStartEvent,
			Event: &models.StartEvent{
				NodeEvent: &models.NodeEvent{
					Id:   startEventId,
					Name: "start",
					Key:  "start",
				},
			},
		},
		EndEvents: []*models.TypeEvent{
			{
				Id:   endEventId,
				Key:  "end",
				Type: types.CTEndEvent,
				Event: &models.EndEvent{
					NodeEvent: &models.NodeEvent{
						Id:   endEventId,
						Key:  "end",
						Name: "end",
					},
				},
			},
		},
		Flows: []*models.SequenceFlow{
			{
				Info: &models.Info{
					Id: id.Get(),
				},
				Start:     startEventId,
				StartType: types.CTStartEvent,
				End:       ig1,
				EndType:   types.CTInclusiveGateway,
			},
			{
				Info: &models.Info{
					Id: id.Get(),
				},
				Start:     ig1,
				StartType: types.CTInclusiveGateway,
				End:       ig2,
				EndType:   types.CTInclusiveGateway,
			},
			{
				Info: &models.Info{
					Id: id.Get(),
				},
				Start:     ig2,
				StartType: types.CTInclusiveGateway,
				End:       ut2,
				EndType:   types.CTUserTask,
			},
			{
				Info: &models.Info{
					Id: id.Get(),
				},
				Start:     ig2,
				StartType: types.CTInclusiveGateway,
				End:       ut3,
				EndType:   types.CTUserTask,
			},
			{
				Info: &models.Info{
					Id: id.Get(),
				},
				Start:     ut2,
				StartType: types.CTUserTask,
				End:       ig3,
				EndType:   types.CTInclusiveGateway,
			},
			{
				Info: &models.Info{
					Id: id.Get(),
				},
				Start:     ut3,
				StartType: types.CTUserTask,
				End:       ig3,
				EndType:   types.CTInclusiveGateway,
			},
			{
				Info: &models.Info{
					Id: id.Get(),
				},
				Start:     ig1,
				StartType: types.CTInclusiveGateway,
				End:       ut4,
				EndType:   types.CTUserTask,
			},
			{
				Info: &models.Info{
					Id: id.Get(),
				},
				Start:     ig1,
				StartType: types.CTInclusiveGateway,
				End:       ut5,
				EndType:   types.CTUserTask,
			},
			{
				Info: &models.Info{
					Id: id.Get(),
				},
				Start:     ut4,
				StartType: types.CTUserTask,
				End:       ig4,
				EndType:   types.CTInclusiveGateway,
			},
			{
				Info: &models.Info{
					Id: id.Get(),
				},
				Start:     ut5,
				StartType: types.CTUserTask,
				End:       ig4,
				EndType:   types.CTInclusiveGateway,
			},
			{
				Info: &models.Info{
					Id: id.Get(),
				},
				Start:     ig3,
				StartType: types.CTInclusiveGateway,
				End:       ig4,
				EndType:   types.CTInclusiveGateway,
			},
			{
				Info: &models.Info{
					Id: id.Get(),
				},
				Start:     ig4,
				StartType: types.CTInclusiveGateway,
				End:       ut6,
				EndType:   types.CTUserTask,
			},
			{
				Info: &models.Info{
					Id: id.Get(),
				},
				Start:     ut6,
				StartType: types.CTUserTask,
				End:       endEventId,
				EndType:   types.CTEndEvent,
			},
		},
	})
}
