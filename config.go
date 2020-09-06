package main

import (
	"fmt"
	"github.com/joho/godotenv"
	"os"
)

type Config struct {
	// The port to bind the web application server to
	Port int

	// Path to database file
	DatabaseURI string

	// APIKey for jwt token generation
	APIKey string
}

func InitConfig(dbURI string, port int, envPath string) (*Config, error) {
	if err := godotenv.Load(envPath); err != nil {
		return nil, fmt.Errorf("no %s file found", envPath)
	}

	config := &Config{
		Port:        port,
		DatabaseURI: dbURI,
		APIKey:      getEnv("API_KEY", "default_api_key"),
	}
	return config, nil
}

func getEnv(key string, defaultVal string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return defaultVal
}
