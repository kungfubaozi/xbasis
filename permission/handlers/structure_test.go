package permissionhandlers

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/davecgh/go-spew/spew"
	"github.com/olivere/elastic"
	"konekko.me/gosion/permission/utils"
	"testing"
)

func TestStructure(t *testing.T) {
	client, err := elastic.NewClient(elastic.SetURL("http://192.168.2.62:9200/"))
	if err != nil {
		panic(err)
	}

	ab := elastic.NewBoolQuery()
	ab.Must(elastic.NewTermQuery("app_id", "5a6d5a6d5979"))
	ab.Must(elastic.NewTermQuery("opening", true))
	ab.Filter(elastic.NewTermsQuery("type", permissionutils.TypeFunctionStructure))

	src, err := ab.Source()
	if err != nil {
		panic(err)
	}
	data, err := json.MarshalIndent(src, "", "  ")
	if err != nil {
		panic(err)
	}
	fmt.Println(string(data))

	r, err := client.Search("gs_permission_structure").Type("v").Query(ab).FetchSourceContext(elastic.NewFetchSourceContext(true).Include("id")).Do(context.Background())
	if err == nil && r.Hits.TotalHits == 1 {

		spew.Dump(r.Hits.Hits[0])

		var j *structure
		err := json.Unmarshal(*r.Hits.Hits[0].Source, &j)
		if err == nil {
			spew.Dump(j)
		}
		fmt.Println("err", err)

	} else {
		fmt.Println(err)
	}
}
