package config

import "os"

type Config struct {
	AppEnv string
	Port   string
}

func Load() (*Config, error) {
	return &Config{
		AppEnv: getEnv("APP_ENV", "development"),
		Port:   ":" + getEnv("PORT", "8080"),
	}, nil
}

func getEnv(varName, defaultValue string) string {
	osAppEnv := os.Getenv(varName)

	if osAppEnv == "" {
		return defaultValue
	}

	return osAppEnv
}
