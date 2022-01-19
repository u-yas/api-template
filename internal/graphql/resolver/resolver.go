package resolver

import (
	"database/sql"
	"sync"

	"github.com/elastic/go-elasticsearch/v8"
	"github.com/labstack/echo"
	"go.uber.org/zap"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	DbCon   *sql.DB
	Logger  *zap.Logger
	HttpCon *echo.Context
	Es      *elasticsearch.Client
	Wg      *sync.WaitGroup
}
