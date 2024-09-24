package elastic

import (
	"context"
	"fmt"

	elastic "github.com/olivere/elastic/v7"
)

type ElasticClient struct {
	client *elastic.Client
}

func NewElasticClient(url string) *ElasticClient {
	client, err := elastic.NewClient(elastic.SetURL(url))
	if err != nil {
		panic(err)
	}
	elasticClient := &ElasticClient{client: client}
	if err := elasticClient.createIndexIfNotExist("articles"); err != nil {
		fmt.Println("Error is occurred on create index: ", err)
	}
	return elasticClient
}

func (c *ElasticClient) createIndexIfNotExist(index string) error {
	exists, err := c.client.IndexExists(index).Do(context.Background())
	if err != nil {
		return fmt.Errorf("error checking index existence: %w", err)
	}
	if !exists {
		createIndex, err := c.client.CreateIndex(index).Do(context.Background())
		if err != nil {
			return fmt.Errorf("error creating index: %w", err)
		}
		if !createIndex.Acknowledged {
			return fmt.Errorf("index creation not acknowledged")
		}
	}
	return nil
}

func (c *ElasticClient) Search(index string, query elastic.Query) (*elastic.SearchResult, error) {
	return c.client.Search().Index(index).Query(query).Do(context.Background())
}

func (c *ElasticClient) Index(index string, body interface{}) (*elastic.IndexResponse, error) {
	return c.client.Index().Index(index).BodyJson(body).Do(context.Background())
}

func (c *ElasticClient) Delete(index string, id string) (*elastic.DeleteResponse, error) {
	return c.client.Delete().Index(index).Id(id).Do(context.Background())
}
