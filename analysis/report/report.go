package analysisreport

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/micro/go-micro/metadata"
	"konekko.me/xbasis/analysis/client"
	"konekko.me/xbasis/safety/pb"
)

type Report interface {
	Action(v []byte)
}

type report struct {
	log              analysisclient.LogClient
	lockingService   xbasissvc_external_safety.LockingService
	blacklistService xbasissvc_external_safety.BlacklistService
}

func (svc *report) Action(v []byte) {
	var data map[string]interface{}
	err := json.Unmarshal(v, &data)
	if err != nil {
		fmt.Println("err format")
		return
	}
	traceId, ok := data["traceId"].(string)
	if !ok {
		fmt.Println("err format traceId")
		return
	}
	action, ok := data["action"].(string)
	if !ok {
		fmt.Println("err format action")
		return
	}
	ctx := metadata.NewContext(context.Background(), map[string]string{"trace_id": traceId})
	switch action {
	case "PasswordError":
		userId, ok := data["userId"].(string)
		if !ok {
			fmt.Println("err format userId")
			return
		}
		currentErrorCount, ok := data["currentErrorCount"].(int64)
		if !ok {
			fmt.Println("err format currentErrorCount")
			return
		}
		hisErrorStatCount, ok := data["hisErrorStatCount"].(int64)
		if !ok {
			fmt.Println("err format hisErrorStatCount")
			return
		}
		go svc.passwordError(ctx, userId, currentErrorCount, hisErrorStatCount)
		break
	}
}

func (svc *report) passwordError(ctx context.Context, userId string, currentErrorCount, hisErrorStatCount int64) {
	var lockTime int64
	if hisErrorStatCount == 1 {
		lockTime = 5 * 60
	} else if hisErrorStatCount == 2 {
		lockTime = 2 * 60 * 60
	} else if hisErrorStatCount == 3 {
		lockTime = 12 * 60 * 60
	}
	if lockTime > 0 {
		s, err := svc.lockingService.Lock(ctx, &xbasissvc_external_safety.LockRequest{
			UserId: userId,
			Time:   lockTime,
		})
		if err != nil {
			fmt.Println("lock user error", err)
			return
		}
		fmt.Println("lock resp", s)
	}
}

func NewAction() Report {
	return &report{}
}
