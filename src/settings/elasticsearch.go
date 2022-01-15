package settings

import (
	"os"

	"github.com/elastic/go-elasticsearch/v8"
)

func SetupElasticSearch() *elasticsearch.Client {
	es, err := elasticsearch.NewDefaultClient()
	if err != nil {
		panic(err)
	}

	env := os.Getenv("ENV")
	if env != "production" {
		es.Info()
	}
	return es
}
