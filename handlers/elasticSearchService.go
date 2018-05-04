package handlers

import (
	"encoding/json"
	"errors"
	"log"

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
	client, err := newElasticSearchClient()
	if err != nil {
		panic(err)
	}

	err = createCompanyIndexIfDoesNotExist(client)
	if err != nil {
		panic(err)
	}

	addCompanyToIndex(client, company)
}

func searchCompanyByName(searchTerm string) ([]models.Company, error) {
	var companies []models.Company
	index := "thirdbridge"

	client, err := newElasticSearchClient()
	if err != nil {
		panic(err)
	}
	q := elastic.NewMultiMatchQuery(searchTerm, "Name", "abbreviation").Type("phrase_prefix")
	searchResult, err := client.Search().Index(index).Type("company").Query(q).Do(context.TODO())

	if err != nil {
		return nil, err
	}

	for _, hit := range searchResult.Hits.Hits {
		var company models.Company
		err := json.Unmarshal(*hit.Source, &company)
		log.Printf("result in elasticsearch with :%s found:: %+v\n", searchTerm, company)
		if err != nil {
			return nil, err
		}
		companies = append(companies, company)
	}

	return companies, nil
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
