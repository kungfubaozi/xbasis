package main

import (
	"context"
	"fmt"
	"github.com/olivere/elastic"
	"konekko.me/gosion/permission/utils"
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

	ab := elastic.NewBoolQuery()
	ab.Must(elastic.NewTermQuery("app_id", "5a6d5a6d5979"), elastic.NewTermQuery("opening", true),
		elastic.NewTermsQuery("type", permissionutils.TypeFunctionStructure, permissionutils.TypeUserStructure))
	s := elastic.NewConstantScoreQuery(ab)

	r, err := client.Search("gs_permission_structure").Type("v").Query(s).StoredField("id").Do(context.Background())
	if err == nil && r.Hits.TotalHits == 2 {

		fmt.Println("1", r.Hits.Hits[0].Fields["id"])
		fmt.Println("2", r.Hits.Hits[1])
	} else {
		fmt.Println(err)
	}

}
