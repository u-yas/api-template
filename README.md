## Specifications

This repository is a sample api server created using echo, gqlgen, sqlboiler, goose, elasticsearch, zap, etc.

## HOW TO USE

1. `go install github.com/pressly/goose/v3/cmd/goose@latest`
2. `go install github.com/volatiletech/sqlboiler/v4@latest`
3. `go mod tidy`
4. `make setup-env`
5. `docker compose up -d`
6. `make dev`