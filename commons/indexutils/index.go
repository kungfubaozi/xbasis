package indexutils

import (
	"context"
	"encoding/json"
	"errors"
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
	s, err := cli.client.Index().Index(index).Type("v").BodyJson(v).Do(context.Background())
	if err != nil {
		return "", err
	}
	if s.Status == 0 {
		return s.Id, nil
	}
	return "", nil
}

func (cli *Client) QueryFirst(index string, kvs map[string]interface{}, result interface{}) (bool, error) {
	ok, v, err := cli._queryFirst(index, kvs)
	if err != nil {
		return false, err
	}
	if ok {
		err = json.Unmarshal(*v[0].Source, result)
		if err != nil {
			return false, err
		}
		return true, nil
	}
	return false, nil
}

func (cli *Client) Delete(index string, kvs map[string]interface{}) (bool, error) {
	ok, v, err := cli._queryFirst(index, kvs)
	if err != nil {
		return false, err
	}
	if ok {
		a, err := cli.client.Delete().Index(index).Type("v").Id(v[0].Id).Do(context.Background())
		if err != nil {
			return false, nil
		}
		if a.Status == 0 {
			return true, nil
		}
	}
	return false, nil
}

func (cli *Client) _queryFirst(index string, kvs map[string]interface{}) (bool, []*elastic.SearchHit, error) {
	b := elastic.NewBoolQuery()
	for k, v := range kvs {
		b.Must(elastic.NewMatchQuery(k, v))
	}
	v, err := cli.client.Search(index).Type("v").Query(b).Do(context.Background())
	if err != nil {
		return false, nil, err
	}
	if v.Hits.TotalHits > 0 {
		return true, v.Hits.Hits, nil
	}
	return false, nil, nil
}

func (cli *Client) Update(index, id string, values map[string]interface{}) (bool, error) {
	a, err := cli.client.Update().Index(index).Type("v").Id(id).Upsert(values).Do(context.Background())
	if err != nil {
		return false, err
	}
	if a.Status == 0 {
		return true, nil
	}
	return false, ErrNotFound
}

func NewClient(url string) (*Client, error) {
	client, err := elastic.NewClient(elastic.SetURL(url))
	if err != nil {
		return nil, err
	}
	return &Client{client: client}, nil
}
