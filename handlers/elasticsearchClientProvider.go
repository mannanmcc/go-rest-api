package handlers

import (
	elastic "gopkg.in/olivere/elastic.v5"
)

func newElasticSearchClient() (*elastic.Client, error) {
	client, err := elastic.NewClient(elastic.SetURL("http://localhost:9200"))

	if err != nil {
		return nil, err
	}

	return client, nil
}
