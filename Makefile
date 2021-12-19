build:
	go build ./src/main.go

generate-init:
	go run github.com/99designs/gqlgen init

generate-schema:
	go run github.com/99designs/gqlgen generate

migrate-drop:
	migrate -database ${POSTGRESQL_URL} -path ./migrations drop 

sqlboiler:
	sqlboiler psql

sqlboiler-prd:
	sqlboiler psql -c ./sqlboiler-prd.toml

migrate-up-prd:
	migrate -database ${POSTGRESQL_URL_PRD} -path ./migrations up	
