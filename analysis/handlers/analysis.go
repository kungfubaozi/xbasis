package analysishandlers

import (
	"context"
	"encoding/json"
	"github.com/olivere/elastic"
	"konekko.me/xbasis/analysis/client"
	analysispb "konekko.me/xbasis/analysis/pb"
	"konekko.me/xbasis/commons/dto"
	"konekko.me/xbasis/commons/errstate"
	"konekko.me/xbasis/commons/indexutils"
	wrapper "konekko.me/xbasis/commons/wrapper"
)

type analysisService struct {
	client *indexutils.Client
	log    analysisclient.LogClient
}

func (svc *analysisService) GetFunctionDetail(ctx context.Context, in *analysispb.GetFunctionDetailRequest, out *analysispb.GetFunctionDetailResponse) error {
	return wrapper.ContextToAuthorize(ctx, out, func(auth *wrapper.WrapperUser) *xbasis_commons_dto.State {
		return nil
	})
}

func (svc *analysisService) GetTrackingStageDetail(ctx context.Context, in *analysispb.GetTrackingDetailRequest, out *analysispb.TrackingStageResponse) error {
	return wrapper.ContextToAuthorize(ctx, out, func(auth *wrapper.WrapperUser) *xbasis_commons_dto.State {
		return nil
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
