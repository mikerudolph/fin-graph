package config

import (
	"os"

	"github.com/joho/godotenv"
)

// Config ...
type Config struct {
	IEXKey     string `envconfig:"IEX_KEY"`
	IEXBaseURL string `envconfig:"IEX_BASE_URL"`
}

// GetConfig ...
func GetConfig() Config {
	_ = godotenv.Load()

	config := Config{
		IEXKey:     os.Getenv("IEX_KEY"),
		IEXBaseURL: os.Getenv("IEX_BASE_URL"),
	}

	return config
}
