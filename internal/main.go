package main

import (
	"go-api-sample/internal/graphql/resolver"
	"go-api-sample/internal/settings"
	"sync"
)

func main() {
	settings.LoadEnv()
	logger := settings.SetupLogger()
	es := settings.SetupElasticSearch()
	db := settings.SetupDB()
	defer db.Close()

	resolver := &resolver.Resolver{
		DbCon:  db,
		Logger: logger,
		Es:     es,
		Wg:     &sync.WaitGroup{},
	}

	settings.SetupServer(resolver)
}
