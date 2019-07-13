package analysishandlers

import (
	"context"
	"encoding/json"
	"github.com/olivere/elastic"
	"konekko.me/xbasis/analysis/client"
	analysispb "konekko.me/xbasis/analysis/pb"
	"konekko.me/xbasis/commons/constants"
	"konekko.me/xbasis/commons/dto"
	"konekko.me/xbasis/commons/errstate"
	"konekko.me/xbasis/commons/indexutils"
	wrapper "konekko.me/xbasis/commons/wrapper"
	"strings"
)

type analysisService struct {
	client *indexutils.Client
	log    analysisclient.LogClient
}

func (svc *analysisService) SearchFunctions(ctx context.Context, in *analysispb.SearchFunctionRequest, out *analysispb.SearchFunctionResponse) error {
	return wrapper.ContextToAuthorize(ctx, out, func(auth *wrapper.WrapperUser) *xbasis_commons_dto.State {

		if len(in.AppId) == 0 {
			return nil
		}

		query := elastic.NewBoolQuery()

		v1 := in.Keyword

		e := svc.client.GetElasticClient().Search("xbs-state-functions")

		if len(v1) > 0 {
			q := elastic.NewQueryStringQuery("*" + v1 + "*")
			q.Field("path")
			q.Field("name")
			query.Must(q)
		}
		query.Must(elastic.NewMatchPhraseQuery("appId", in.AppId))

		v, err := e.Type("_doc").Query(query).From(int(in.Size*in.Page)).Size(int(in.Size)).Sort("total", false).Do(context.Background())
		if err != nil {
			return nil
		}

		var datas []*analysispb.StateFunction

		if v.Hits.TotalHits > 0 {
			for _, s := range v.Hits.Hits {
				t := &analysispb.StateFunction{}
				err := json.Unmarshal(*s.Source, t)
				if err == nil {
					//d := &analysispb.StateFunction{
					//	Path:             t.path,
					//	FunctionId:       t.functionId,
					//	FunctionName:     t.functionName,
					//	Total:            t.total,
					//	TodayTotal:       t.todayTotal,
					//	LastDayTotal:     t.lastDayTotal,
					//	Error:            t.error,
					//	TodayError:       t.todayError,
					//	LastDayError:     t.lastDayError,
					//	AvgTiming:        t.avgTiming,
					//	MinTiming:        t.minTiming,
					//	MaxTiming:        t.maxTiming,
					//	Timing:           t.timing,
					//	LastDayUserVisit: t.lastDayUserVisit,
					//	TodayUserVisit:   t.todayUserVisit,
					//	AppId:            t.appId,
					//}
					datas = append(datas, t)
				}
			}
		}

		out.Data = datas

		return errstate.Success
	})
}

func (svc *analysisService) GetFunctionDetail(ctx context.Context, in *analysispb.GetFunctionDetailRequest, out *analysispb.GetFunctionDetailResponse) error {
	return wrapper.ContextToAuthorize(ctx, out, func(auth *wrapper.WrapperUser) *xbasis_commons_dto.State {
		return nil
	})
}

func (svc *analysisService) GetTrackingStageDetail(ctx context.Context, in *analysispb.GetTrackingDetailRequest, out *analysispb.TrackingStageResponse) error {
	return wrapper.ContextToAuthorize(ctx, out, func(auth *wrapper.WrapperUser) *xbasis_commons_dto.State {

		if len(in.TraceId) < 18 {
			return nil
		}

		query := elastic.NewBoolQuery()
		query.Must(elastic.NewMatchPhraseQuery("headers.traceId", in.TraceId))

		v, err := svc.client.GetElasticClient().Search("xbs-logger.*").Query(query).Type("_doc").Sort("timestamp", false).Do(context.Background())
		if err != nil {
			return nil
		}

		var items []*analysispb.TrackingStageItem
		if v.Hits.TotalHits > 0 {
			for _, s := range v.Hits.Hits {
				t := &tracking{}
				err := json.Unmarshal(*s.Source, t)
				if err == nil {
					item := &analysispb.TrackingStageItem{}
					item.Message = t.Message
					item.ServiceName = t.Header.ServiceName
					item.StateCode = t.StateCode
					item.ModuleName = t.Header.ModuleName
					item.Timing = t.Timing
					item.Level = t.Level
					item.Action = t.Action
					if len(item.Level) > 0 {
						item.Level = strings.ToUpper(item.Level)
					}
					if t.Action == "PermissionVerification" {
						out.Timestamp = t.Timestamp
						out.UserAgent = t.Header.UserAgent
						out.Ip = t.Header.Ip
						out.UserDevice = t.Header.Device
						out.RefClientId = t.Header.RefClientId
						out.ClientId = t.Header.FromClientId
						out.HasDurationToken = t.Header.HasDurationToken
						out.HasAccessToken = t.Header.HasAccessToken
						out.Passed = t.Passed
						out.BasicValidation = t.BasicValidation
						out.DeniedApiClient = t.DeniedApiClient
						out.InvalidApi = t.InvalidApi
						out.Path = t.Header.Path
					} else if t.Action == "UserRequestApiFinished" {
						if t.Fields != nil {
							out.Taking = &analysispb.StageTaking{
								Verification: t.Fields.VerificationTiming,
								All:          t.Fields.AllTiming,
								Process:      t.Fields.ProcessTiming,
							}
						}
					} else if t.Action == "UserRequestApi" {
						out.UserId = t.Header.UserId
					} else if t.Action == "StartUserRequestApi" {

					} else if t.Action == "SimplifiedLookupApi" {
						out.FunctionId = t.Fields.FunctionId
						out.FunctionName = t.Fields.FunctionName
						if len(t.Fields.AuthTypes) > 0 {
							for _, v1 := range t.Fields.AuthTypes {
								switch int(v1) {
								case xbasisconstants.AuthTypeOfFace:
									out.AuthTypes = append(out.AuthTypes, "Face")
									break
								case xbasisconstants.AuthTypeOfToken:
									out.AuthTypes = append(out.AuthTypes, "Token")
									break
								case xbasisconstants.AuthTypeOfMobileConfirm:
									out.AuthTypes = append(out.AuthTypes, "MobileConfirm")
									break
								case xbasisconstants.AuthTypeOfValcode:
									out.AuthTypes = append(out.AuthTypes, "DurationAccessToken")
									break
								}
							}
						}
						out.AppId = t.Fields.AppId
					}
					items = append(items, item)
				}
			}
		}

		out.Data = items

		return errstate.Success
	})
}

func (svc *analysisService) SearchTracking(ctx context.Context, in *analysispb.SearchTrackingRequest, out *analysispb.SearchTrackingResponse) error {
	return wrapper.ContextToAuthorize(ctx, out, func(auth *wrapper.WrapperUser) *xbasis_commons_dto.State {

		if len(in.AppId) == 0 {
			return nil
		}

		query := elastic.NewBoolQuery()

		v1 := in.Value

		e := svc.client.GetElasticClient().Search("xbs-logger.*")

		query.Must(elastic.NewMatchPhraseQuery("action", "PermissionVerification"))

		if len(v1) > 0 {
			q := elastic.NewQueryStringQuery("*" + v1 + "*")
			if len(in.Key) > 0 {
				q.Field(in.Key)
			} else {
				q.Field("headers.path")
				q.Field("headers.userAgent")
				q.Field("headers.fromClientId")
				q.Field("headers.refClientId")
				q.Field("function")
				q.Field("headers.ip")
			}
			query.Must(q)
		}
		query.Must(elastic.NewMatchPhraseQuery("app_id", in.AppId))

		v, err := e.Type("_doc").Query(query).From(int(in.Size*in.Page)).Size(int(in.Size)).Sort("timestamp", false).Do(context.Background())
		if err != nil {
			return nil
		}

		var datas []*analysispb.RequestTracking

		if v.Hits.TotalHits > 0 {
			for _, s := range v.Hits.Hits {

				t := &tracking{}
				err := json.Unmarshal(*s.Source, t)

				if err == nil {
					datas = append(datas, &analysispb.RequestTracking{
						TrackId:          t.Header.TraceId,
						UserAgent:        t.Header.UserAgent,
						Path:             t.Header.Path,
						Ip:               t.Header.Ip,
						RouteTo:          t.RouteTo,
						Function:         t.Function,
						Timestamp:        t.Timestamp,
						HasAccessToken:   t.Header.HasAccessToken,
						HasDurationToken: t.Header.HasDurationToken,
						Passed:           t.Passed,
						Taking: &analysispb.StageTaking{
							All:          t.AllTiming,
							Verification: t.VerificationTiming,
							Process:      t.ProcessTiming,
						},
						BasicValidation: t.BasicValidation,
						InvalidApi:      t.InvalidApi,
						DeniedApiClient: t.DeniedApiClient,
					})
				}

			}
		}

		out.Data = datas

		return errstate.Success
	})
}

func NewAnalysisService(client *indexutils.Client, log analysisclient.LogClient) analysispb.AnalysisHandler {
	return &analysisService{client: client, log: log}
}
