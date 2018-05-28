package handlers

import (
	"errors"

	"github.com/mannanmcc/rest-api/models"
	"golang.org/x/net/context"
	elastic "gopkg.in/olivere/elastic.v5"
)

const (
	indexName    = "thirdbridge"
	docType      = "company"
	appName      = "companyManager"
	indexMapping = `{
						"mappings" : {
							"company" : {
								"properties" : {
									"name" : { "type" : "text" },
									"status" : { "type" : "text" }
								}
							}
						}
					}`
)

func indexCompany(company models.Company) {
	client, err := elastic.NewClient(elastic.SetURL("http://localhost:9200"))
	if err != nil {
		panic(err)
	}

	err = createCompanyIndexIfDoesNotExist(client)
	if err != nil {
		panic(err)
	}

	addCompanyToIndex(client, company)
}

func createCompanyIndexIfDoesNotExist(client *elastic.Client) error {
	exists, err := client.IndexExists(indexName).Do(context.TODO())
	if err != nil {
		return err
	}

	if exists {
		return nil
	}

	res, err := client.CreateIndex(indexName).
		Body(indexMapping).
		Do(context.TODO())

	if err != nil {
		return err
	}

	if !res.Acknowledged {
		return errors.New("Oops could not create the index")
	}

	return nil
}

func addCompanyToIndex(client *elastic.Client, company models.Company) error {
	_, err := client.Index().
		Index(indexName).
		Type(docType).
		BodyJson(company).
		Do(context.TODO())

	if err != nil {
		return err
	}

	return nil
}
