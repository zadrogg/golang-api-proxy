package main

import (
	"os"
	"strconv"
)

type Config struct {
	Dadata dadata
	Redis  redisDB
	Cache  cache
	Server server
}

type redisDB struct {
	host     string
	port     int
	password string
}

type dadata struct {
	dadataKey       string
	dadataSecretKey string
}

type cache struct {
	sign bool
}

type server struct {
	url string
}

func getConfig() *Config {
	return &Config{
		Dadata: dadata{
			dadataKey:       getEnv("DADATA_KEY", ""),
			dadataSecretKey: getEnv("DADATA_SECRET_KEY", ""),
		},
		Redis: redisDB{
			host:     getEnv("REDIS_HOST", "localhost"),
			port:     getEnvInt("REDIS_PORT", 6379),
			password: getEnv("REDIS_PASSWORD", ""),
		},
		Cache: cache{
			sign: getEnvBool("CACHE", false),
		},
		Server: server{
			url: getEnv("SERVER", "localhost:3000"),
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

func getEnvBool(key string, defaultValue bool) bool {
	valueStr := getEnv(key, "")
	if value, err := strconv.ParseBool(valueStr); err == nil {
		return value
	}

	return defaultValue
}
