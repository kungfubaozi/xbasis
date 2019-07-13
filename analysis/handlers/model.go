package analysishandlers

import "konekko.me/xbasis/analysis/client"

type tracking struct {
	Header             *analysisclient.LogHeaders `json:"headers"`
	Timestamp          int64                      `json:"timestamp"`
	Message            string                     `json:"message"`
	StateCode          int64                      `json:"stateCode"`
	Function           string                     `json:"function"`
	Passed             bool                       `json:"passed"`
	RouteTo            string                     `json:"route_to"`
	AllTiming          float64                    `json:"all"`
	ProcessTiming      float64                    `json:"processing"`
	VerificationTiming float64                    `json:"verification"`
	BasicValidation    bool                       `json:"basic_validation"`
	InvalidApi         bool                       `json:"invalid_api"`
	DeniedApiClient    bool                       `json:"denied_api_client"`
	Fields             *trackingFields            `json:"fields"`
	Action             string                     `json:"action"`
	Timing             int64                      `json:"timing"` //距离上一个处理耗时
	Level              string                     `json:"level"`
}

type trackingFields struct {
	AppId              string    `json:"app_id"`
	FunctionId         string    `json:"function_id"`
	FunctionName       string    `json:"function_name"`
	AuthTypes          []float64 `json:"auth_types"`
	UserId             string    `json:"user_id"`
	Platform           float64   `json:"platform"`
	AllTiming          float64   `json:"all"`
	ProcessTiming      float64   `json:"processing"`
	VerificationTiming float64   `json:"verification"`
}

type stateFunction struct {
	functionId       string
	total            int64
	todayTotal       int64
	lastDayTotal     int64
	error            int64
	todayError       int64
	lastDayError     int64
	avgTiming        int64
	minTiming        int64
	maxTiming        int64
	timing           int64
	lastDayUserVisit int64
	todayUserVisit   int64
	functionName     string
	path             string
	appId            string
	mnovps           int64
	mnovpm           int64
	mnovph           int64
}
