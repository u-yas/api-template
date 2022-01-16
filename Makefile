dev:
	go run ./src/main.go

build:
	go build ./src/main.go

compose:
	cd build
	docker compose up

generate-init:
	go run github.com/99designs/gqlgen init

generate-schema:
	go run github.com/99designs/gqlgen generate

migrate-up:
	goose postgres -dir ./migrations/v0 ${POSTGRESQL_URL} up

migrate-down:
	goose postgres -dir ./migrations/v0 ${POSTGRESQL_URL} down


sqlboiler-prd:
	sqlboiler psql -c ./sqlboiler-prd.toml

settup-env:
	cp .env.local.sample .env.local
	export POSTGRESQL_URL=postgres://user:password@localhost:5432/sample_db?sslmode=disable
	export GOOSE_DRIVER=postgres
	export GOOSE_DBSTRING=$POSTGRESQL_URL


