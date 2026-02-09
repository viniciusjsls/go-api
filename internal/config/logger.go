package config

import (
	"os"

	"go.uber.org/zap"
)

func InitLogger() *zap.Logger {
	env := os.Getenv("APP_ENV")
	if env == "" {
		env = "development"
	}

	var logger *zap.Logger
	var err error

	if env == "production" {
		logger, err = zap.NewProduction()
	} else {
		logger, err = zap.NewDevelopment()
	}

	if err != nil {
		panic(err)
	}

	return logger
}
