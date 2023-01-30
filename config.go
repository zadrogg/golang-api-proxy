package main

import (
	"os"
	"strconv"
)

type Config struct {
	Dadata dadata
	Redis  redis
}

type redis struct {
	host     string
	port     int
	password string
}

type dadata struct {
	dadataKey       string
	dadataSecretKey string
}

func getConfig() *Config {
	return &Config{
		Dadata: dadata{
			dadataKey:       getEnv("DADATA_KEY", ""),
			dadataSecretKey: getEnv("DADATA_SECRET_KEY", ""),
		},
		Redis: redis{
			host:     getEnv("REDIS_HOST", "localhost"),
			port:     getEnvInt("REDIS_PORT", 6379),
			password: getEnv("REDIS_PASSWORD", ""),
		},
	}
}

func getEnv(key string, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}

	return defaultValue
}

func getEnvInt(key string, defaultValue int) int {
	valueStr := getEnv(key, "")
	if value, err := strconv.Atoi(valueStr); err == nil {
		return value
	}

	return defaultValue
}
