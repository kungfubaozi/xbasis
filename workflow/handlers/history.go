package workflowhandlers

import (
	"konekko.me/xbasis/workflow/modules"
	pb "konekko.me/xbasis/workflow/pb"
)

type historyService struct {
	modules modules.Modules
}

func NewHistoryService() pb.HistoryHandler {
	return &historyService{}
}
