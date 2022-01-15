package settings

import (
	"os"

	"go.uber.org/zap"
)

func SetupLogger() *zap.Logger {
	var logger *zap.Logger
	env := os.Getenv("ENV")

	switch env {
	case "production":
		logger, _ = zap.NewProduction()
	case "development":
		logger, _ = zap.NewDevelopment()
	default:
		logger, _ = zap.NewDevelopment()
	}

	defer logger.Sync()

	return logger
}
