package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/davecgh/go-spew/spew"
	"github.com/olivere/elastic"
	"konekko.me/gosion/analysis/logger/pb"
	"konekko.me/gosion/analysis/pb"
)

type blacklist struct {
	Type         int64  `json:"type"`
	Content      string `json:"content"`
	CreateAt     int64  `json:"create_at"`
	CreateUserId string `json:"create_user_id"`
}

func main() {

	client, err := elastic.NewClient(elastic.SetURL("http://192.168.2.62:9200/"))
	if err != nil {
		panic(err)
	}

	//b := &blacklist{
	//	Type:         1,
	//	Content:      "123456",
	//	CreateAt:     time.Now().UnixNano(),
	//	CreateUserId: "unid",
	//}
	//
	//rs, err := client.Index().Index("gs_safety_blacklist").Type("v").BodyJson(b).Do(context.Background())
	//if err != nil {
	//	panic(err)
	//}
	//fmt.Println(rs.Result)
	//fmt.Println(rs.Status)

	//q := elastic.NewMatchQuery("content", "123456")
	//v, err := client.Search("gs_safety_blacklist").Type("v").Query(q).Do(context.Background())
	//if err != nil {
	//	panic(err)
	//}
	//if v.Hits.TotalHits > 0 {
	//	d := v.Hits.Hits[0]
	//	r, err := client.Delete().Index("gs_safety_blacklist").Type("v").Id(d.Id).Do(context.Background())
	//	if err != nil {
	//		panic(err)
	//	}
	//	fmt.Println(r.Result)
	//	fmt.Println(r.Status)
	//}

	//ab := elastic.NewBoolQuery()
	//ab.Must(elastic.NewTermQuery("app_id", "5268597a5534"))
	//
	//r, err := client.Search("gs_structures").Type("v").Query(ab).FetchSourceContext(elastic.NewFetchSourceContext(true).Include("id")).Do(context.Background())
	//if err == nil && r.Hits.TotalHits == 2 {
	//	var m map[string]interface{}
	//	err := json.Unmarshal(*r.Hits.Hits[0].Source, &m)
	//	if err != nil {
	//		panic(err)
	//	}
	//	spew.Dump(m)
	//	fmt.Println("id", m["id"])
	//} else {
	//	fmt.Println(err)
	//}

	dataMap := map[string]interface{}{
		"x-action":       "UserRequestApi",
		"x-fields-appId": "526d59544a6a",
		"y-timestamp-1":  0,
		"y-timestamp-2":  1559730148069,
	}

	m, _ := json.Marshal(dataMap)

	source(client, &gs_service_analysis.GetDataRequest{
		Map: string(m),
		XAxis: &gs_service_analysis.XAxisRequest{
			Name: "test-x",
			Factors: []*gs_service_analysis.XAxisFactor{
				{
					Field:     "action",
					Operation: "equals",
					Value:     "x-action",
				},
				{
					Field:     "fields.appId",
					Operation: "equals",
					Value:     "x-fields-appId",
				},
			},
		},
		YAxis: &gs_service_analysis.YAxisRequest{
			Size: 3,
			Name: "test-y",
			Factors: []*gs_service_analysis.YAxisFactor{
				{
					Field:     "fields.id.keyword",
					Operation: "terms",
					Name:      "count",
					Factor: &gs_service_analysis.YAxisFactor{
						Field:     "timestamp",
						Operation: "range",
						Name:      "today",
						Value:     "y-timestamp",
					},
				},
				{
					Field:     "fields.appId.keyword",
					Operation: "cardinality",
					Name:      "appId",
				},
			},
		},
	})

}

func source(client *elastic.Client, in *gs_service_analysis.GetDataRequest) {
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

	showSource := func(source interface{}, err error) {
		b, _ := json.Marshal(source)
		fmt.Println(string(b))
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

	showSource(query.Source())

	search := client.Search("gosion-logger.*")

	source := elastic.NewSearchSource().Query(query)

	loopYAxis(source, in.YAxis.Factors, getValue)

	search.SearchSource(source)
	//showSource(source.Source())

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
			println(string(s))
		}
	}
	spew.Dump(result)
}

func loopYAxis(search *elastic.SearchSource, factors []*gs_service_analysis.YAxisFactor, getValue func(j string) interface{}) {
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

func getAgg(v *gs_service_analysis.YAxisFactor, getValue func(j string) interface{}) (string, elastic.Aggregation) {
	if v != nil {
		if v.Operation == "range" {
			agg1 := elastic.NewRangeAggregation().Field(v.Field).AddRange(getValue(v.Value+"-1"), getValue(v.Value+"-2"))
			return v.Name, agg1
		}
	}
	return "", nil
}
