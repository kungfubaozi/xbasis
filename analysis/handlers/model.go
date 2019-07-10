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
}
