package config

import (
	"go.uber.org/zap"
)

func InitLogger(cfg *Config) *zap.Logger {
	var logger *zap.Logger
	var err error

	if cfg.AppEnv == "production" {
		logger, err = zap.NewProduction()
	} else {
		logger, err = zap.NewDevelopment()
	}

	if err != nil {
		panic(err)
	}

	defer logger.Sync()

	return logger
}
