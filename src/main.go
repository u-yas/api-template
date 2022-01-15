package main

import (
	"api-template/src/graphql/resolver"
	"api-template/src/settings"
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
