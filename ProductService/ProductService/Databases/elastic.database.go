package Databases

import (
	"ProductService/ProductService/Config"
	"fmt"
	elasticSearch "github.com/elastic/go-elasticsearch/v7"
	"log"
)

var (
	ElastiClient *elasticSearch.Client = elasticCon()
)

func elasticCon() *elasticSearch.Client {
	envGet := Config.GetEnv
	elasticURL := envGet("ELASTICSEARCH_URL")
	elasticUsername := envGet("ELASTICSEARCH_USERNAME")
	elasticPassword := envGet("ELASTICSEARCH_PASSWORD")
	cfg := elasticSearch.Config{
		Addresses: []string{elasticURL},
		Username:  elasticUsername,
		Password:  elasticPassword,
	}

	es, err := elasticSearch.NewClient(cfg)

	if err != nil {
		log.Printf("Error creating the client: %s", err)
	}

	_, err = es.Info()
	if err != nil {
		log.Fatalf("Error getting response: %s", err)
	}
	fmt.Println("Elastic Connected")
	return es
}
