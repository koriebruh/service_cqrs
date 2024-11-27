package config

import (
	"github.com/elastic/go-elasticsearch/v8"
	"log"
	"net/http"
	"time"
)

func ElasticClient(cfg *Config) *elasticsearch.Client {
	esConfig := elasticsearch.Config{
		Addresses: []string{cfg.Elastic.Host},
		Username:  cfg.Elastic.Username,
		Password:  cfg.Elastic.Password,
		Transport: &http.Transport{
			MaxIdleConns:          cfg.Elastic.MaxIdleConns,
			IdleConnTimeout:       90 * time.Second, // Default timeout for idle connections
			ResponseHeaderTimeout: cfg.Elastic.Timeout,
		},
	}

	client, err := elasticsearch.NewClient(esConfig)
	if err != nil {
		log.Fatalf("Error creating Elasticsearch client: %v", err)
	}

	// Test connection
	res, err := client.Info()
	if err != nil {
		log.Fatalf("Error pinging Elasticsearch: %v", err)
	}
	defer res.Body.Close()

	log.Println("Elasticsearch connection established")
	return client
}
