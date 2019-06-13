package permissionhandlers

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/davecgh/go-spew/spew"
	"github.com/olivere/elastic"
	"konekko.me/gosion/commons/indexutils"
	"testing"
)

var index = "gs-urf-relations-test"

func TestBindingService_UserRole(t *testing.T) {
	client, err := indexutils.NewClient("http://192.168.2.62:9200/")
	if err != nil {
		panic(err)
	}

	//addRole(client)
	bindFunction(client)
}

func addRole(client *indexutils.Client) {

	r := &directrelation{
		UserId:     "d3e71ce819d3",
		RoleId:     "c8fe08149210",
		User:       true,
		Function:   false,
		FunctionId: "",
	}

	client.AddData(index, r)
}

func showSource(i interface{}, err error) {
	s, _ := json.Marshal(i)
	fmt.Println(string(s))
}

func bindFunction(client *indexutils.Client) {
	query := elastic.NewBoolQuery()
	query.Must(elastic.NewMatchPhraseQuery("user", true))
	query.Must(elastic.NewMatchPhraseQuery("function", false))
	query.Must(elastic.NewMatchPhraseQuery("role_id", "c8fe08149210"))

	r, err := client.GetElasticClient().UpdateByQuery(index).Query(query).
		Script(elastic.NewScript("ctx._source.function_id = params.id;ctx._source.function = params.t").Params(map[string]interface{}{
			"id": "2d58b482c632",
			"t":  true,
		})).
		Do(context.Background())

	showSource(query.Source())

	if err != nil {
		panic(err)
	}

	spew.Dump(r)
}
