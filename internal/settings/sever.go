package settings

import (
	"go-api-sample/internal/graphql/generated"
	"go-api-sample/internal/graphql/resolver"
	"log"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"go.uber.org/zap"
)

const defaultPort = "8080"

func SetupServer(resolvers *resolver.Resolver) {
	logger := resolvers.Logger

	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	e := echo.New()
	e.Use(middleware.Recover())

	e.POST("/query", func(c echo.Context) error {
		resolvers.HttpCon = &c
		gqlConfig := generated.Config{
			Resolvers: resolvers,
		}
		graphqlHandler := handler.NewDefaultServer(generated.NewExecutableSchema(gqlConfig))

		graphqlHandler.ServeHTTP(c.Response(), c.Request())
		return nil
	})

	env := os.Getenv("ENV")
	if env == "development" {
		playgroundHandler := playground.Handler("GraphQL", "/query")
		e.GET("/playground", func(c echo.Context) error {
			playgroundHandler.ServeHTTP(c.Response(), c.Request())
			return nil
		})
		log.Printf("connect to http://localhost:%s/playground for GraphQL playground", port)
	}

	err := e.Start(":" + defaultPort)
	if err != nil {
		logger.Fatal("fatal error ", zap.Errors("error", []error{err}))
	}
}
