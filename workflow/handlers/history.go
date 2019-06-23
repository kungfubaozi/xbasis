package workflowhandlers

import (
	"konekko.me/gosion/workflow/modules"
	"konekko.me/gosion/workflow/pb"
)

type historyService struct {
	modules modules.Modules
}

func NewHistoryService() gosionsvc_external_workflow.HistoryHandler {
	return &historyService{}
}
