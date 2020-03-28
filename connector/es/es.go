package es

import (
	"context"
	"errors"
	"github.com/elastic/go-elasticsearch/v8"
	"github.com/elastic/go-elasticsearch/v8/esapi"
	"github.com/google/uuid"
	"log"
	"strings"
)

type ElasticSearchClient interface {
	Insert(document Document, coercedId bool) error
}
type elasticSearchClient struct {
	es *elasticsearch.Client
}

func Open(urls []string, username, password string) (ElasticSearchClient, error) {
	config := elasticsearch.Config{
		Addresses: urls,
		Username:  username,
		Password:  password,
	}

	es, err := elasticsearch.NewClient(config)
	return &elasticSearchClient{es: es}, err
}

func (ec elasticSearchClient) Insert(document Document, coercedId bool) error {
	if document.Index == "" {
		return errors.New("invalid index")
	}

	if coercedId && document.ID == "" {
		return errors.New("invalid document id")
	}

	if document.ID == "" {
		document.ID = uuid.New().String()
	}

	req := esapi.IndexRequest{
		Index:      document.Index,
		DocumentID: document.ID,
		Body:       strings.NewReader(document.Json),
	}
	res, err := req.Do(context.Background(), ec.es)
	log.Println(res)
	return err
}
