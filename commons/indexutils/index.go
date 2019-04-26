package indexutils

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/olivere/elastic"
)

type Client struct {
	client *elastic.Client
}

var ErrNotFound error = errors.New("not")

func (cli *Client) GetElasticClient() *elastic.Client {
	return cli.client
}

func (cli *Client) AddData(index string, v interface{}) (string, error) {
	s, err := cli.client.Index().Index(index).Type("_doc").BodyJson(v).Do(context.Background())

	if err != nil {
		return "", err
	}
	if s.Status == 0 {
		return s.Id, nil
	}
	return "", nil
}

func (cli *Client) QueryFirst(index string, kvs map[string]interface{}, result interface{}, includes ...string) (bool, error) {
	ok, v, err := cli._queryFirst(index, kvs, includes...)
	if err != nil {
		fmt.Println("query first err", err)
		return false, err
	}
	if ok {
		err = json.Unmarshal(*v[0].Source, result)
		if err != nil {
			fmt.Println("json Unmarshal err", err)
			return false, err
		}
		return true, nil
	}
	return false, nil
}

func (cli *Client) Delete(index string, kvs map[string]interface{}) (bool, error) {
	ok, v, err := cli._queryFirst(index, kvs, nil...)
	if err != nil {
		return false, err
	}
	if ok {
		a, err := cli.client.Delete().Index(index).Type("_doc").Id(v[0].Id).Do(context.Background())
		if err != nil {
			return false, nil
		}
		if a.Status == 0 {
			return true, nil
		}
	}
	return false, nil
}

func (cli *Client) _queryFirst(index string, kvs map[string]interface{}, includes ...string) (bool, []*elastic.SearchHit, error) {

	query := cli._buildQuery(kvs)

	e := cli.client.Search(index).Type("_doc").Query(query)
	if includes != nil && len(includes) > 0 {
		e.FetchSourceContext(elastic.NewFetchSourceContext(true).Include(includes...))
	}
	//src, err := query.Source()
	//if err != nil {
	//	panic(err)
	//}
	//data, err := json.MarshalIndent(src, "", "  ")
	//if err != nil {
	//	panic(err)
	//}
	//fmt.Println("query", string(data))

	v, err := e.Do(context.Background())
	if err != nil {
		return false, nil, err
	}

	//fmt.Println("hits", v.Hits.TotalHits)
	//
	//spew.Dump(v.Hits)

	if v.Hits.TotalHits > 0 {
		return true, v.Hits.Hits, nil
	}
	return false, nil, nil
}

func (cli *Client) Update(index, id string, values map[string]interface{}) (bool, error) {
	a, err := cli.client.Update().Index(index).Type("_doc").Id(id).Upsert(values).Do(context.Background())
	if err != nil {
		return false, err
	}
	if a.Status == 0 {
		return true, nil
	}
	return false, ErrNotFound
}

func (cli *Client) Count(index string, kvs map[string]interface{}) (int64, error) {
	a, err := cli.client.Count(index).Type("_doc").Query(cli._buildQuery(kvs)).Do(context.Background())
	if err != nil {
		return 0, err
	}
	return a, nil
}

func (cli *Client) _buildQuery(kvs map[string]interface{}) elastic.Query {
	b := elastic.NewBoolQuery()
	for k, v := range kvs {
		b.Must(elastic.NewMatchPhraseQuery(k, v))
	}
	return b
}

func NewClient(url string) (*Client, error) {
	client, err := elastic.NewClient(elastic.SetURL(url))
	if err != nil {
		return nil, err
	}
	return &Client{client: client}, nil
}
