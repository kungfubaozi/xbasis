package analysishandlers

import (
	"context"
	"encoding/json"
	"github.com/davecgh/go-spew/spew"
	"github.com/olivere/elastic"
	"konekko.me/gosion/analysis/client"
	external "konekko.me/gosion/analysis/pb"
	"konekko.me/gosion/commons/dto"
	"konekko.me/gosion/commons/errstate"
	"konekko.me/gosion/commons/indexutils"
	"konekko.me/gosion/commons/wrapper"
)

type loggerService struct {
	*indexutils.Client
	log analysisclient.LogClient
}

func (svc *loggerService) TodayVisit(context.Context, *external.TodayVisitRequest, *external.TodayVisitResponse) error {
	panic("implement me")
}

func (svc *loggerService) UsageFunction(context.Context, *external.UsageFunctionRequest, *external.UsageFunctionResponse) error {
	panic("implement me")
}

func (svc *loggerService) GetAxisData(ctx context.Context, in *external.GetDataRequest, out *external.GetDataResponse) error {
	return gs_commons_wrapper.ContextToAuthorize(ctx, out, func(auth *gs_commons_wrapper.WrapperUser) *gs_commons_dto.State {
		if len(in.XAxis.Factors) == 0 || len(in.YAxis.Factors) == 0 {
			return nil
		}

		headers := &analysisclient.LogHeaders{
			TraceId: auth.TraceId,
		}

		query := elastic.NewBoolQuery()

		values := make(map[string]interface{})
		if len(in.Map) > 0 {
			err := json.Unmarshal([]byte(in.Map), &values)
			if err != nil {
				panic(err)
			}
			spew.Dump(values)
		}

		getValue := func(j string) interface{} {
			return values[j]
		}

		var must []elastic.Query
		var mustnot []elastic.Query
		var filters []elastic.Query

		for _, v := range in.XAxis.Factors {
			if v.Operation == "equals" {
				must = append(must, elastic.NewMatchPhraseQuery(v.Field, getValue(v.Value)))
			}
		}

		if len(must) > 0 {
			query.Must(must...)
		}

		if len(mustnot) > 0 {
			query.MustNot(mustnot...)
		}

		if len(filters) > 0 {
			query.Filter(filters...)
		}

		search := svc.GetElasticClient().Search("gosion-logger.*")

		source := elastic.NewSearchSource().Query(query)

		loopYAxis(source, in.YAxis.Factors, getValue)

		search.SearchSource(source)

		result, err := search.Size(0).Do(context.Background())
		if err != nil {
			panic(err)
		}
		if result.Hits.TotalHits > 0 {
			if result.Aggregations != nil {
				s, err := json.Marshal(result.Aggregations)
				if err != nil {
					panic(err)
				}

				svc.log.Info(&analysisclient.LogContent{
					Headers: headers,
					Action:  "GetAxisData",
				})

				out.Data = string(s)
				return errstate.Success
			}
		}
		return nil
	})
}

func loopYAxis(search *elastic.SearchSource, factors []*external.YAxisFactor, getValue func(j string) interface{}) {
	for _, v := range factors {
		if v.Operation == "terms" {
			agg := elastic.NewTermsAggregation().Field(v.Field)
			name, aggs := getAgg(v.Factor, getValue)
			if len(name) > 0 {
				agg.SubAggregation(name, aggs)
			}
			search.Aggregation(v.Name, agg)
		} else if v.Operation == "cardinality" {
			agg := elastic.NewCardinalityAggregation().Field(v.Field)
			name, aggs := getAgg(v.Factor, getValue)
			if len(name) > 0 {
				agg.SubAggregation(name, aggs)
			}
			search.Aggregation(v.Name, agg)
		}
	}
}

func getAgg(v *external.YAxisFactor, getValue func(j string) interface{}) (string, elastic.Aggregation) {
	if v != nil {
		if v.Operation == "range" {
			agg1 := elastic.NewRangeAggregation().Field(v.Field).AddRange(getValue(v.Value+"-1"), getValue(v.Value+"-2"))
			return v.Name, agg1
		}
	}
	return "", nil
}

func NewLoggerService(log analysisclient.LogClient, client *indexutils.Client) external.LoggerHandler {
	return &loggerService{log: log, Client: client}
}
