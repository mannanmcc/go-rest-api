package handlers

import (
	"errors"
	"fmt"
	"log"
	"reflect"

	"github.com/mannanmcc/rest-api/models"
	"golang.org/x/net/context"
	elastic "gopkg.in/olivere/elastic.v5"
)

const (
	indexName    = "thirdbridge"
	docType      = "company"
	appName      = "myApp"
	indexMapping = `{
						"mappings" : {
							"company" : {
								"properties" : {
									"name" : { "type" : "text" }
								}
							}
						}
					}`
)

// // Company entity to index
// type Company struct {
// 	Name string `json:"name"`
// }

func indexCompany(company models.Company) {
	client, err := elastic.NewClient(elastic.SetURL("http://localhost:9200"))
	if err != nil {
		panic(err)
	}

	err = createIndexCompanyIfDoesNotExist(client)
	if err != nil {
		panic(err)
	}

	// company := Company{
	// 	Name: fmt.Sprintf("new Company: test company"),
	// }

	addCompanyToIndex(client, company)

	err = findAndPrintAppLogs(client)
	if err != nil {
		panic(err)
	}
}

func createIndexCompanyIfDoesNotExist(client *elastic.Client) error {
	exists, err := client.IndexExists(indexName).Do(context.TODO())
	if err != nil {
		return err
	}

	if exists {
		return nil
	}
	log.Printf("creating new index: %s", indexName)
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
	log.Printf("adding company details into index.......")
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

func findAndPrintAppLogs(client *elastic.Client) error {
	termQuery := elastic.NewTermQuery("app", appName)

	res, err := client.Search(indexName).
		Index(indexName).
		Query(termQuery).
		Do(context.TODO())

	if err != nil {
		return err
	}

	var l models.Company
	for _, item := range res.Each(reflect.TypeOf(l)) {
		l := item.(models.Company)
		fmt.Printf("name: %s\n", l.Name)
	}

	return nil
}
