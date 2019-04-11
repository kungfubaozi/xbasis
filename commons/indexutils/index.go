package indexutils

import (
	"context"
	"encoding/json"
	"github.com/olivere/elastic"
)

type Client struct {
	client *elastic.Client
}

func (cli *Client) AddData(index string, v interface{}) (bool, error) {
	s, err := cli.client.Index().Index(index).Type("v").BodyJson(v).Do(context.Background())
	if err != nil {
		return false, err
	}
	if s.Status == 0 {
		return true, nil
	}
	return false, nil
}

func (cli *Client) QueryFirst(index string, kvs map[string]interface{}, result interface{}) (bool, error) {
	b := elastic.NewBoolQuery()
	for k, v := range kvs {
		b.Must(elastic.NewMatchQuery(k, v))
	}
	v, err := cli.client.Search(index).Type("v").Query(b).Do(context.Background())
	if err != nil {
		return false, nil
	}
	if v.Hits.TotalHits > 0 {
		err = json.Unmarshal(*v.Hits.Hits[0].Source, result)
		if err != nil {
			return false, nil
		}
		return true, nil
	}
	return false, nil
}

func NewClient(url string) (*Client, error) {
	client, err := elastic.NewClient(elastic.SetURL(url))
	if err != nil {
		return nil, err
	}
	return &Client{client: client}, nil
}
